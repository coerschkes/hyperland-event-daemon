package event_handlers

import (
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

type FocusedWindowHandler struct {
	pwdRegistry *state.PwdRegistry
}

func NewFocusedWindowHandler(registry *state.PwdRegistry) hyprland.EventHandler {
	return &FocusedWindowHandler{registry}
}

func (h *FocusedWindowHandler) Handle(event string) error {
	if strings.HasPrefix(event, "activewindow>>") {
		currentApplication := h.getCurrentApplication(event)
		currentPwd := h.getCurrentPwd(event)

		h.pwdRegistry.UpdateCurrentApp(currentApplication)
		h.pwdRegistry.UpdateCurrentPwd(currentPwd)
	}

	return nil
}

func (h *FocusedWindowHandler) getCurrentApplication(event string) string {
	parts := strings.Split(event, ">>")
	if len(parts) > 1 {
		sub := parts[1]
		subParts := strings.Split(sub, ",")
		return subParts[0]
	}

	return ""
}

func (h *FocusedWindowHandler) getCurrentPwd(event string) string {
	pwd := strings.Split(event, ":")
	if len(pwd) == 1 {
		return ""
	}

	return pwd[1]
}
