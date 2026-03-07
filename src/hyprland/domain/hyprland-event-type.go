package domain

//go:generate go tool go-enum -f $GOFILE

// ENUM(activewindow, activewindowv2, activelayout, workspace, moveworkspace, moveworkspacev2, monitoradded, monitoraddedv2, monitorremoved, monitorremovedv2, focusedmon, focusedmonv2, openlayer, destroyworkspace, destroyworkspacev2, configreloaded, unknown)
type HyprlandEventType string

func AllHyprlandEventTypes() []HyprlandEventType {
	out := make([]HyprlandEventType, 0, len(_HyprlandEventTypeValue))
	for _, v := range _HyprlandEventTypeValue {
		out = append(out, v)
	}
	return out
}

// const (
// 	ActiveWindow       HyprlandEventType = "activewindow"
// 	ActiveWindowV2     HyprlandEventType = "activewindowv2"
// 	ActiveLayout       HyprlandEventType = "activelayout"
// 	Workspace          HyprlandEventType = "workspace"
// 	MoveWorkspace      HyprlandEventType = "moveworkspace"
// 	MoveWorkspaceV2    HyprlandEventType = "moveworkspacev2"
// 	MonitorAdded       HyprlandEventType = "monitoradded"
// 	MonitorAddedV2     HyprlandEventType = "monitoraddedv2"
// 	MonitorRemoved     HyprlandEventType = "monitorremoved"
// 	MonitorRemovedV2   HyprlandEventType = "monitorremovedv2"
// 	FocusedMon         HyprlandEventType = "focusedmon"
// 	FocusedMonV2       HyprlandEventType = "focusedmonv2"
// 	Openlayer          HyprlandEventType = "openlayer"
// 	DestroyWorkspace   HyprlandEventType = "destroyworkspace"
// 	DestroyWorkspaceV2 HyprlandEventType = "destroyworkspacev2"
// 	Unknown            HyprlandEventType = "unknown"
// )
//
// var AllHyprlandEventTypes = []HyprlandEventType{
// 	ActiveWindow,
// 	ActiveWindowV2,
// 	ActiveLayout,
// 	Workspace,
// 	MoveWorkspace,
// 	MoveWorkspaceV2,
// 	MonitorAdded,
// 	MonitorAddedV2,
// 	MonitorRemoved,
// 	MonitorRemovedV2,
// 	FocusedMon,
// 	FocusedMonV2,
// 	Openlayer,
// 	DestroyWorkspace,
// 	DestroyWorkspaceV2,
// }
//
// func NewHyprlandEventType(s string) HyprlandEventType {
// 	t := HyprlandEventType(s)
//
// 	if slices.Contains(AllHyprlandEventTypes, t) {
// 		return t
// 	}
//
// 	return Unknown
// }
