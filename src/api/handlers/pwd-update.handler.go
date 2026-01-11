package handlers

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/api"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

const PWD_UPDATE_IDENTIFIER = "pwd-update"
const PWD_UPDATE_MESSAGE_FORMAT = "<identifier>:<pid>:<pwd>"

type PwdUpdateHandler struct {
	pwdRegistry *state.PwdRegistry
}

func NewPwdUpdateHandler(registry state.PwdRegistry) api.MessageHandler {
	return &PwdUpdateHandler{&registry}
}

func (h *PwdUpdateHandler) CanHandle(msg string) bool {
	return strings.HasPrefix(msg, PWD_UPDATE_IDENTIFIER)
}

func (h *PwdUpdateHandler) Handle(msg string) (string, error) {
	fmt.Println("Received msg '" + msg + "' in " + PWD_UPDATE_IDENTIFIER)
	split := strings.Split(msg, ":")

	if len(split) != 3 {
		return "", errors.New("Wrong message format. Unable to handle: '" + msg + "'. Expected: '" + PWD_UPDATE_MESSAGE_FORMAT + "'.")
	}
	pid, err := strconv.Atoi(split[1])
	if err != nil {
		return "", errors.New("Unable to parse pid '" + split[1] + "'.")
	}

	h.pwdRegistry.UpsertEntry(pid, split[2])

	fmt.Println("Set pwd to " + h.pwdRegistry.GetPwd(pid) + " for pid " + strconv.Itoa(pid))

	return "", nil
}
