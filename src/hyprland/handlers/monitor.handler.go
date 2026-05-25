package handlers

import (
	"fmt"
	"os/exec"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

var internalMonitor = "eDP-1"

type MonitorHandler struct {
	monitorProvider hyprland.MonitorProvider
}

func NewMonitorHandler(monitorProvider hyprland.MonitorProvider) *MonitorHandler {
	return &MonitorHandler{monitorProvider}
}

func (h *MonitorHandler) Types() []domain.HyprlandEventType {
	return []domain.HyprlandEventType{domain.HyprlandEventTypeMonitoradded, domain.HyprlandEventTypeMonitorremoved}
}

func (h *MonitorHandler) OnEventReceived(event domain.HyprlandEvent) error {
	monitorName := event.Payload[0]
	if monitorName == "FALLBACK" || monitorName == internalMonitor {
		return nil
	}

	if event.Type == domain.HyprlandEventTypeMonitoradded {
		err := h.handleMonitorAdded(monitorName)
		if err != nil {
			return err
		}
	} else {
		err := h.handleMonitorRemoved()
		if err != nil {
			return err
		}
	}

	go h.restartStatusbar()

	return nil
}

func (h *MonitorHandler) OnStartup() error {
	monitors, err := h.monitorProvider.GetMonitors()
	if err != nil {
		return err
	}

	if len(monitors) == 1 {
		return nil
	}

	for _, monitor := range monitors {
		if monitor.Name == internalMonitor {
			continue
		}

		err = h.handleMonitorAdded(monitor.Name)
		if err != nil {
			return err
		}
	}

	return nil
}

func (h *MonitorHandler) handleMonitorAdded(monitorName string) error {
	err := h.monitorProvider.SetMonitorConfiguration([]string{monitorName, "preferred", "primary"})
	if err != nil {
		return err
	}

	return h.monitorProvider.SetMonitorConfiguration([]string{internalMonitor, "disabled"})
}

func (h *MonitorHandler) handleMonitorRemoved() error {
	monitors, err := h.monitorProvider.GetMonitors()

	if err != nil {
		return err
	}

	if len(monitors) == 1 {
		err := h.monitorProvider.SetMonitorConfiguration([]string{internalMonitor, "preferred", "auto", "1.25"})
		if err != nil {
			return err
		}

		return h.monitorProvider.Reload()
	}

	return nil
}

func (h *MonitorHandler) restartStatusbar() {
	err := exec.Command("systemctl", "--user", "restart", "ags-widgets").Run()
	if err != nil {
		fmt.Println(err)
	}
}
