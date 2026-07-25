package hyprland

import (
	"encoding/json"
	"os/exec"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
)

type Hyprctl struct{}

func NewHyprctl() *Hyprctl {
	return &Hyprctl{}
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

func (h *Hyprctl) Reload() error {
	cmd := exec.Command(
		"hyprctl",
		"reload",
	)

	_, err := cmd.CombinedOutput()
	if err != nil {
		return err
	}

	return nil
}
