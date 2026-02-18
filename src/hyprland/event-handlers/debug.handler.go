package event_handlers

import (
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
)

type DebugHandler struct {
}

func NewDebugHandler(registry *state.PwdRegistry) hyprland.EventHandler {
	return &DebugHandler{}
}

func (h *DebugHandler) Handle(event string) error {
	println("DEBUG: " + event)

	return nil
}
