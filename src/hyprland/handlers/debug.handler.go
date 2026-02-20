package handlers

import (
	"fmt"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type DebugHandler struct{}

func NewDebugHandler() *DebugHandler {
	return &DebugHandler{}
}

func (h *DebugHandler) Types() []domain.HyprlandEventType {
	return domain.AllHyprlandEventTypes
}

func (h *DebugHandler) OnEventReceived(event domain.HyprlandEvent) error {
	fmt.Println("DEBUG: ", event.RawEvent)
	return nil
}
