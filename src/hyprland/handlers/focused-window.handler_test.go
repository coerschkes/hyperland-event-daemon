package handlers

import (
	"testing"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
	"github.com/coerschkes/hyprland-event-daemon/src/state"
	"github.com/coerschkes/hyprland-event-daemon/src/testutil"
)

func TestFocusedWindowHandlerOnEventReceived(t *testing.T) {
	type expected struct {
		pwd string
		app string
	}

	tests := []struct {
		input string
		want  expected
	}{
		{"activewindow>>Alacritty,test@archdev:~/dev/projects", expected{pwd: "~/dev/projects", app: "Alacritty"}},
		{"activewindow>>HelloWorld,test@archdev:~/dev/projects/helloworld", expected{pwd: "~/", app: "HelloWorld"}},
		{"activewindow>>Alacritty,test@archdev:/", expected{pwd: "/", app: "Alacritty"}},
	}

	for _, tt := range tests {
		registry := state.NewPwdRegistry()
		handler := NewFocusedWindowHandler(registry)
		event := domain.NewHyprlandEvent(tt.input)

		handler.OnEventReceived(event)

		if registry.GetCurrentPwd() != tt.want.pwd {
			testutil.AssertFail(t, "OnEventReceived", tt.input, registry.GetCurrentPwd(), tt.want.pwd, "pwd mismatch")
		}

		if registry.GetCurrentApp() != tt.want.app {
			testutil.AssertFail(t, "OnEventReceived", tt.input, registry.GetCurrentApp(), tt.want.app, "app mismatch")
		}
	}
}
