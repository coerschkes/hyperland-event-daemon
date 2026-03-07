package handlers

import (
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type ConfigReloadedHandler struct {
	startupExecutor *hyprland.StartupExecutor
}

func NewConfigReloadedHandler(startupExecutor *hyprland.StartupExecutor) *ConfigReloadedHandler {
	return &ConfigReloadedHandler{startupExecutor}
}

func (h *ConfigReloadedHandler) Types() []domain.HyprlandEventType {
	return []domain.HyprlandEventType{domain.HyprlandEventTypeConfigreloaded}
}

func (h *ConfigReloadedHandler) OnEventReceived(event domain.HyprlandEvent) error {
	h.startupExecutor.Execute()
	return nil
}
