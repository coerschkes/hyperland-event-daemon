package hyprland

import (
	"encoding/json"
	"os/exec"
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type Hyprctl struct{}

func NewHyprctl() *Hyprctl {
	return &Hyprctl{}
}

func (h *Hyprctl) GetMonitors() ([]domain.Monitor, error) {
	cmd := exec.Command("hyprctl", "-j", "monitors")
	output, err := cmd.Output()
	if err != nil {
		return []domain.Monitor{}, err
	}

	var monitors []domain.Monitor
	if err := json.Unmarshal(output, &monitors); err != nil {
		return []domain.Monitor{}, err
	}

	return monitors, nil
}

func (h *Hyprctl) SetMonitorConfiguration(configuration []string) error {
	cmd := exec.Command(
		"hyprctl",
		"keyword",
		"monitor",
		strings.Join(configuration, ","),
	)

	_, err := cmd.CombinedOutput()

	return err
}

func (h *Hyprctl) GetDevices() (domain.Devices, error) {
	cmd := exec.Command("hyprctl", "-j", "devices")
	output, err := cmd.Output()

	if err != nil {
		return domain.Devices{}, err
	}

	var devices domain.Devices
	if err := json.Unmarshal(output, &devices); err != nil {
		return domain.Devices{}, err
	}

	return devices, nil
}

func (h *Hyprctl) SetDeviceConfiguration(name string, deviceConfiguration Configuration) error {
	cmd := exec.Command(
		"hyprctl",
		"-r",
		"--",
		"keyword",
		"device["+name+"]:"+deviceConfiguration.Key,
		deviceConfiguration.Value,
	)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
