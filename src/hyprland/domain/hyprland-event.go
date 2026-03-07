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

	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	eventType, err := ParseHyprlandEventType(parts[0])

	if err != nil {
		e.Type = HyprlandEventTypeUnknown
	} else {
		e.Type = eventType
	}

	if len(parts) > 1 {
		sub := parts[1]
		e.Payload = strings.Split(sub, ",")
	}
}
