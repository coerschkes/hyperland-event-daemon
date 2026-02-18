package event_handlers

import (
	"os/exec"
	"strings"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland"
)

type Mouse struct {
	Identifier  string
	Sensitivity string
}

var configurations = []Mouse{
	{
		Identifier:  "razer-razer-viper-v3-pro",
		Sensitivity: "-1",
	},
	{
		Identifier:  "razer-razer-viper-v3-pro-1",
		Sensitivity: "-1",
	},
	{
		Identifier:  "razer-razer-viper-v3-pro-mouse",
		Sensitivity: "-1",
	},
}

type MouseSensitivityFixHandler struct{}

func NewMouseSensitivityFixHandler() hyprland.EventHandler {
	return &MouseSensitivityFixHandler{}
}

func (h *MouseSensitivityFixHandler) Handle(event string) error {
	if strings.HasPrefix(event, "activelayout>>razer-razer-viper-v3-pro") {
		for _, configuration := range configurations {
			cmd := exec.Command(
				"hyprctl",
				"-r",
				"--",
				"keyword",
				"device["+configuration.Identifier+"]:sensitivity",
				configuration.Sensitivity,
			)

			_, err := cmd.CombinedOutput()
			if err != nil {
				return err
			}
		}
	}

	return nil
}
