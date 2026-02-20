package domain

import (
	"testing"

	"github.com/coerschkes/hyprland-event-daemon/src/testutil"
)

func TestNewHyprlandEventType(t *testing.T) {
	tests := []struct {
		input string
		want  HyprlandEventType
	}{
		{"activewindow", ActiveWindow},
		{"monitorremovedv2", MonitorRemovedV2},
		{"jibberish", Unknown},
		{"", Unknown},
	}

	for _, tt := range tests {
		got := NewHyprlandEventType(tt.input)
		if got != tt.want {
			testutil.AssertFail(t, "NewHyprlandEventType", tt.input, got, tt.want, "result mismatch")
		}
	}
}
