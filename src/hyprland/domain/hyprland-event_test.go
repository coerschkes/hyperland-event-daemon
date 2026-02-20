package domain

import (
	"slices"
	"testing"

	"github.com/coerschkes/hyprland-event-daemon/src/testutil"
)

func TestNewHyprlandEvent(t *testing.T) {
	type expected struct {
		eventType HyprlandEventType
		payload   []string
	}

	tests := []struct {
		input string
		want  expected
	}{
		{"activewindow>>Alacritty,test@archdev:~/dev", expected{eventType: ActiveWindow, payload: []string{"Alacritty", "test@archdev:~/dev"}}},
		{"activewindowv2>>56127669d000", expected{eventType: ActiveWindowV2, payload: []string{"56127669d000"}}},
		{"activewindow>>", expected{eventType: ActiveWindow, payload: []string{""}}},
		{"activewindow", expected{eventType: ActiveWindow, payload: []string{}}},
		{"foo", expected{eventType: Unknown, payload: []string{}}},
		{"foo>>bar", expected{eventType: Unknown, payload: []string{"bar"}}},
		{"", expected{}},
	}

	for _, tt := range tests {
		got := NewHyprlandEvent(tt.input)
		if !slices.Equal(got.Payload, tt.want.payload) {
			testutil.AssertFail(t, "NewHyprlandEvent", tt.input, got, tt.want, "payload mismatch")
		}

		if got.Type != tt.want.eventType {
			testutil.AssertFail(t, "NewHyprlandEvent", tt.input, got, tt.want, "event type mismatch")
		}
	}
}
