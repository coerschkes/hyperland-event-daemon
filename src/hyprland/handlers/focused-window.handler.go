package handlers

import (
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

type FocusedWindowHandler struct {
	pwdRegistry *state.PwdRegistry
}

func NewFocusedWindowHandler(registry *state.PwdRegistry) *FocusedWindowHandler {
	return &FocusedWindowHandler{registry}
}

func (h *FocusedWindowHandler) Types() []domain.HyprlandEventType {
	return []domain.HyprlandEventType{domain.ActiveWindow}
}

func (h *FocusedWindowHandler) OnEventReceived(event domain.HyprlandEvent) error {
	currentApplication := event.Payload[0]
	currentPwd := h.getCurrentPwd(event.Payload[1])

	h.pwdRegistry.UpdateCurrentApp(currentApplication)
	h.pwdRegistry.UpdateCurrentPwd(currentPwd)

	return nil
}

func (h *FocusedWindowHandler) getCurrentPwd(event string) string {
	pwd := strings.Split(event, ":")
	if len(pwd) == 1 {
		return ""
	}

	return pwd[1]
}
