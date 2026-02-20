package domain

import (
	"strings"
)

type HyprlandEvent struct {
	RawEvent string
	Type     HyprlandEventType
	Payload  []string
}

func NewHyprlandEvent(payload string) HyprlandEvent {
	hyprlandEvent := &HyprlandEvent{}

	trimmedPayload := strings.TrimSpace(payload)
	if trimmedPayload == "" {
		return *hyprlandEvent
	}

	hyprlandEvent.RawEvent = trimmedPayload
	hyprlandEvent.parseEvent()

	return *hyprlandEvent
}

func (e *HyprlandEvent) parseEvent() {
	parts := strings.Split(e.RawEvent, ">>")

	e.Type = NewHyprlandEventType(parts[0])

	if len(parts) > 1 {
		sub := parts[1]
		e.Payload = strings.Split(sub, ",")
	}
}
