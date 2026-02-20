package handlers

import (
	"slices"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type MouseHandler struct {
	deviceProvider hyprland.DeviceProvider
}

func NewMouseHandler(deviceProvider hyprland.DeviceProvider) *MouseHandler {
	return &MouseHandler{deviceProvider}
}

func (h *MouseHandler) Types() []domain.HyprlandEventType {
	return []domain.HyprlandEventType{domain.ActiveLayout}
}

func (h *MouseHandler) OnStartup() error {
	devices, err := h.deviceProvider.GetDevices()
	if err != nil {
		return err
	}

	if h.isMouseConnected(devices, domain.RazerViperConfigurations) {
		return h.setMouseSensitivity()
	}

	return nil
}

func (h *MouseHandler) OnEventReceived(event domain.HyprlandEvent) error {
	if event.Payload[0] == "razer-razer-viper-v3-pro" {
		return h.setMouseSensitivity()
	}

	return nil
}

func (h *MouseHandler) setMouseSensitivity() error {
	for _, configuration := range domain.RazerViperConfigurations {
		err := h.deviceProvider.SetDeviceConfiguration(configuration.Name, hyprland.Configuration{Key: "sensitivity", Value: configuration.Sensitivity})
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *MouseHandler) isMouseConnected(foundDevices domain.Devices, mouseConfigurations []domain.Mouse) bool {
	return slices.ContainsFunc(foundDevices.Mice, func(mouse domain.Mouse) bool {
		return slices.ContainsFunc(mouseConfigurations, func(config domain.Mouse) bool {
			return mouse.Name == config.Name
		})
	})
}
