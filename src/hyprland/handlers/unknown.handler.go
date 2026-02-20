package handlers

import (
	"fmt"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type UnknownHandler struct{}

func NewUnknownHandler() *UnknownHandler {
	return &UnknownHandler{}
}

func (h *UnknownHandler) Types() []domain.HyprlandEventType {
	return []domain.HyprlandEventType{domain.Unknown}
}

func (h *UnknownHandler) OnEventReceived(event domain.HyprlandEvent) error {
	fmt.Println("Unknown event received: ", event.RawEvent)
	return nil
}
