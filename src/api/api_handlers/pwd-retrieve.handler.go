package api_handlers

import (
	"errors"
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/api"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

const PWD_RETRIEVE_IDENTIFIER = "pwd-retrieve"
const PWD_RETRIEVE_MESSAGE_FORMAT = "<identifier>"

type PwdRetrieveHandler struct {
	pwdRegistry *state.PwdRegistry
}

func NewPwdRetrieveHandler(registry *state.PwdRegistry) api.MessageHandler {
	return &PwdRetrieveHandler{registry}
}

func (h *PwdRetrieveHandler) CanHandle(msg string) bool {
	return strings.HasPrefix(msg, PWD_RETRIEVE_IDENTIFIER)
}

func (h *PwdRetrieveHandler) Handle(msg string) (string, error) {
	split := strings.Split(msg, ":")

	switch len(split) {
	case 1:
		return h.pwdRegistry.GetCurrentPwd(), nil
	default:
		return "", errors.New("Wrong message format. Unable to handle: '" + msg + "'. Expected: '" + PWD_RETRIEVE_MESSAGE_FORMAT + "'.")
	}
}
