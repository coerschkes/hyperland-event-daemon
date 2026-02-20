package handlers

import (
	"slices"
	"testing"

	"github.com/coerschkes/hyprland-event-daemon/src/hyprland/domain"
	"github.com/coerschkes/hyprland-event-daemon/src/testutil"
)

type MonitorProviderMock struct {
	Mock testutil.Mock
}

func NewMonitorProviderMock(mock testutil.Mock) *MonitorProviderMock {
	return &MonitorProviderMock{mock}
}

type GetMonitorsResponse struct {
	Monitors []domain.Monitor
	Error    error
}

func (m *MonitorProviderMock) GetMonitors() ([]domain.Monitor, error) {
	m.Mock.AddMethodCall("GetMonitors")
	response := m.Mock.MethodResponses["GetMonitors"].(GetMonitorsResponse)
	return response.Monitors, response.Error
}

func (m *MonitorProviderMock) SetMonitorConfiguration(configuration []string) error {
	m.Mock.AddMethodCall("SetMonitorConfiguration", configuration)
	v, ok := m.Mock.MethodResponses["SetMonitorConfiguration"].(error)
	if !ok {
		return nil
	}
	return v
}

// todo: fix
func TestMonitorHandlerOnEventReceived(t *testing.T) {
	tests := []struct {
		input           string
		methodUnderTest string
		mock            *testutil.Mock
		want            []testutil.MethodCall
	}{
		{
			"monitoradded>>HDMI-1",
			"SetMonitorConfiguration",
			testutil.NewMock(make(map[string]any)),
			[]testutil.MethodCall{
				{
					MethodName: "SetMonitorConfiguration",
					CallNumber: 1,
					Params:     []string{"eDP-1", "disabled"},
				},
				{
					MethodName: "SetMonitorConfiguration",
					CallNumber: 2,
					Params:     []string{"HDMI-1", "preferred", "primary"},
				},
			},
		},
	}

	for _, tt := range tests {
		provider := NewMonitorProviderMock(*tt.mock)
		handler := NewMonitorHandler(provider)
		event := domain.NewHyprlandEvent(tt.input)

		err := handler.OnEventReceived(event)

		if err != nil {
			testutil.AssertFail(t, "OnEventReceived", tt.input, err, nil, "unexpected err")
		}

		if !slices.Equal(tt.mock.MethodCalls[tt.methodUnderTest], tt.want) {
			testutil.AssertFail(t, "OnEventReceived", tt.input, tt.mock.MethodCalls[tt.methodUnderTest], tt.want, "method calls mismatch")
		}
	}
}
