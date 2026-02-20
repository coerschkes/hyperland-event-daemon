package domain

import (
	"slices"
)

type HyprlandEventType string

const (
	ActiveWindow       HyprlandEventType = "activewindow"
	ActiveWindowV2     HyprlandEventType = "activewindowv2"
	ActiveLayout       HyprlandEventType = "activelayout"
	Workspace          HyprlandEventType = "workspace"
	MoveWorkspace      HyprlandEventType = "moveworkspace"
	MoveWorkspaceV2    HyprlandEventType = "moveworkspacev2"
	MonitorAdded       HyprlandEventType = "monitoradded"
	MonitorAddedV2     HyprlandEventType = "monitoraddedv2"
	MonitorRemoved     HyprlandEventType = "monitorremoved"
	MonitorRemovedV2   HyprlandEventType = "monitorremovedv2"
	FocusedMon         HyprlandEventType = "focusedmon"
	FocusedMonV2       HyprlandEventType = "focusedmonv2"
	Openlayer          HyprlandEventType = "openlayer"
	DestroyWorkspace   HyprlandEventType = "destroyworkspace"
	DestroyWorkspaceV2 HyprlandEventType = "destroyworkspacev2"
	Unknown            HyprlandEventType = "unknown"
)

var AllHyprlandEventTypes = []HyprlandEventType{
	ActiveWindow,
	ActiveWindowV2,
	Workspace,
	MoveWorkspace,
	MoveWorkspaceV2,
	MonitorAdded,
	MonitorAddedV2,
	MonitorRemoved,
	MonitorRemovedV2,
	FocusedMon,
	FocusedMonV2,
	Openlayer,
	DestroyWorkspace,
	DestroyWorkspaceV2,
}

func NewHyprlandEventType(s string) HyprlandEventType {
	t := HyprlandEventType(s)

	if slices.Contains(AllHyprlandEventTypes, t) {
		return t
	}

	return Unknown
}
