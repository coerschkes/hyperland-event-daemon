package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/api"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

const PWD_RETRIEVE_IDENTIFIER = "pwd-retrieve"
const PWD_RETRIEVE_MESSAGE_FORMAT = "<identifier>(:<pid>)"

type PwdRetrieveHandler struct {
	pwdRegistry *state.PwdRegistry
}

func NewPwdRetrieveHandler(registry state.PwdRegistry) api.MessageHandler {
	return &PwdRetrieveHandler{&registry}
}

func (h *PwdRetrieveHandler) CanHandle(msg string) bool {
	return strings.HasPrefix(msg, PWD_RETRIEVE_IDENTIFIER)
}

func (h *PwdRetrieveHandler) Handle(msg string) (string, error) {
	fmt.Println("RECEIVED: " + msg)
	split := strings.Split(msg, ":")

	switch len(split) {
	case 1:
		return h.pwdRegistry.GetPwd(h.pwdRegistry.GetCurrentPid()), nil
	case 2:
		pid, err := strconv.Atoi(split[1])
		if err != nil {
			return "", errors.New("Unable to parse pid '" + split[1] + "'. Error: " + err.Error())
		}
		return h.pwdRegistry.GetPwd(pid), nil
	default:
		return "", errors.New("Wrong message format. Unable to handle: '" + msg + "'. Expected: '" + PWD_RETRIEVE_MESSAGE_FORMAT + "'.")
	}
}
