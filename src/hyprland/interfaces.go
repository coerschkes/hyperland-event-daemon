package hyprland

import "github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"

type Handler interface {
	Types() []domain.HyprlandEventType
}

type EventHandler interface {
	Handler
	OnEventReceived(event domain.HyprlandEvent) error
}

type StartupHandler interface {
	OnStartup() error
}

type MonitorProvider interface {
	GetMonitors() ([]domain.Monitor, error)
	SetMonitorConfiguration(configuration []string) error
}

type DeviceProvider interface {
	GetDevices() (domain.Devices, error)
	SetDeviceConfiguration(name string, deviceConfiguration Configuration) error
}
