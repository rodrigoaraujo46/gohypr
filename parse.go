package gohypr

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

func expectFields(fields []string, n int, line string) error {
	if len(fields) != n {
		return fmt.Errorf(
			"%w: expected %d fields, got %d; line: %q",
			ErrUnexpectedFields,
			n,
			len(fields),
			line,
		)
	}

	return nil
}

func parseEvent(line string) (Event, error) {
	line, _ = strings.CutSuffix(line, "\n")
	name, payload, ok := strings.Cut(line, ">>")
	if !ok {
		return nil, errors.New("not a valid event, doesn't contain '>>'")
	}

	fields := strings.FieldsFunc(payload, func(r rune) bool { return r == ',' })

	switch name {
	case "workspace":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventWorkspace{
			WorkspaceName: fields[0],
		}, nil

	case "workspacev2":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventWorkspaceV2{
			WorkspaceID:   fields[0],
			WorkspaceName: fields[1],
		}, nil

	case "focusedmon":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventFocusedMon{
			MonitorName:   fields[0],
			WorkspaceName: fields[1],
		}, nil

	case "focusedmonv2":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventFocusedMonV2{
			MonitorName: fields[0],
			WorkspaceID: fields[1],
		}, nil

	case "activewindow":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventActiveWindow{
			WindowClass: fields[0],
			WindowTitle: fields[1],
		}, nil

	case "activewindowv2":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventActiveWindowV2{
			WindowAddress: fields[0],
		}, nil

	case "fullscreen":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[0])
		if err != nil {
			return nil, err
		}

		return EventFullscreen{
			State: state,
		}, nil

	case "monitorremoved":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventMonitorRemoved{
			MonitorName: fields[0],
		}, nil

	case "monitorremovedv2":
		if err := expectFields(fields, 3, line); err != nil {
			return nil, err
		}

		return EventMonitorRemovedV2{
			MonitorID:          fields[0],
			MonitorName:        fields[1],
			MonitorDescription: fields[2],
		}, nil

	case "monitoradded":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventMonitorAdded{
			MonitorName: fields[0],
		}, nil

	case "monitoraddedv2":
		if err := expectFields(fields, 3, line); err != nil {
			return nil, err
		}

		return EventMonitorAddedV2{
			MonitorID:          fields[0],
			MonitorName:        fields[1],
			MonitorDescription: fields[2],
		}, nil

	case "createworkspace":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventCreateWorkspace{
			WorkspaceName: fields[0],
		}, nil

	case "createworkspacev2":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventCreateWorkspaceV2{
			WorkspaceID:   fields[0],
			WorkspaceName: fields[1],
		}, nil

	case "destroyworkspace":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventDestroyWorkspace{
			WorkspaceName: fields[0],
		}, nil

	case "destroyworkspacev2":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventDestroyWorkspaceV2{
			WorkspaceID:   fields[0],
			WorkspaceName: fields[1],
		}, nil

	case "moveworkspace":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventMoveWorkspace{
			WorkspaceName: fields[0],
			MonitorName:   fields[1],
		}, nil

	case "moveworkspacev2":
		if err := expectFields(fields, 3, line); err != nil {
			return nil, err
		}

		return EventMoveWorkspaceV2{
			WorkspaceID:   fields[0],
			WorkspaceName: fields[1],
			MonitorName:   fields[2],
		}, nil

	case "renameworkspace":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventRenameWorkspace{
			WorkspaceID: fields[0],
			NewName:     fields[1],
		}, nil

	case "activespecial":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventActiveSpecial{
			WorkspaceName: fields[0],
			MonitorName:   fields[1],
		}, nil

	case "activespecialv2":
		if err := expectFields(fields, 3, line); err != nil {
			return nil, err
		}

		return EventActiveSpecialV2{
			WorkspaceID:   fields[0],
			WorkspaceName: fields[1],
			MonitorName:   fields[2],
		}, nil

	case "activelayout":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventActiveLayout{
			KeyboardName: fields[0],
			LayoutName:   fields[1],
		}, nil

	case "openwindow":
		if err := expectFields(fields, 4, line); err != nil {
			return nil, err
		}

		return EventOpenWindow{
			WindowAddress: fields[0],
			WorkspaceName: fields[1],
			WindowClass:   fields[2],
			WindowTitle:   fields[3],
		}, nil

	case "closewindow":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventCloseWindow{
			WindowAddress: fields[0],
		}, nil

	case "kill":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventKill{
			WindowAddress: fields[0],
		}, nil

	case "movewindow":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventMoveWindow{
			WindowAddress: fields[0],
			WorkspaceName: fields[1],
		}, nil

	case "movewindowv2":
		if err := expectFields(fields, 3, line); err != nil {
			return nil, err
		}

		return EventMoveWindowV2{
			WindowAddress: fields[0],
			WorkspaceID:   fields[1],
			WorkspaceName: fields[2],
		}, nil

	case "openlayer":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventOpenLayer{
			NameSpace: fields[0],
		}, nil

	case "closelayer":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventCloseLayer{
			NameSpace: fields[0],
		}, nil

	case "submap":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventSubmap{
			SubmapName: fields[0],
		}, nil

	case "changefloatingmode":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[1])
		if err != nil {
			return nil, err
		}

		return EventChangeFloatingMode{
			WindowAddress: fields[0],
			Floating:      state,
		}, nil

	case "urgent":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventUrgent{
			WindowAddress: fields[0],
		}, nil

	case "screencast":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[0])
		if err != nil {
			return nil, err
		}

		return EventScreencast{
			State: state,
			Owner: fields[1],
		}, nil

	case "screencastv2":
		if err := expectFields(fields, 3, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[0])
		if err != nil {
			return nil, err
		}

		return EventScreencastV2{
			State: state,
			Owner: fields[1],
			Name:  fields[2],
		}, nil

	case "windowtitle":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventWindowTitle{
			WindowAddress: fields[0],
		}, nil

	case "windowtitlev2":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		return EventWindowTitleV2{
			WindowAddress: fields[0],
			WindowTitle:   fields[1],
		}, nil

	case "togglegroup":
		const minFields = 1
		if len(fields) < minFields {
			return nil, fmt.Errorf(
				"%w: expected at least 1 field, got %d; line: %q",
				ErrUnexpectedFields,
				len(fields),
				line,
			)
		}

		state, err := strconv.ParseBool(fields[0])
		if err != nil {
			return nil, err
		}

		return EventToggleGroup{
			State:         state,
			WindowAddress: fields[1:],
		}, nil

	case "moveintogroup":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventMoveIntoGroup{
			WindowAddress: fields[0],
		}, nil

	case "moveoutofgroup":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventMoveOutOfGroup{
			WindowAddress: fields[0],
		}, nil

	case "ignoregrouplock":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[0])
		if err != nil {
			return nil, err
		}

		return EventIgnoreGroupLock{
			State: state,
		}, nil

	case "lockgroups":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[0])
		if err != nil {
			return nil, err
		}

		return EventLockGroups{
			State: state,
		}, nil

	case "configreloaded":
		if err := expectFields(fields, 0, line); err != nil {
			return nil, err
		}

		return EventConfigReloaded{}, nil

	case "pin":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[1])
		if err != nil {
			return nil, err
		}

		return EventPin{
			WindowAddress: fields[0],
			PinState:      state,
		}, nil

	case "minimized":
		if err := expectFields(fields, 2, line); err != nil {
			return nil, err
		}

		state, err := strconv.ParseBool(fields[1])
		if err != nil {
			return nil, err
		}

		return EventMinimized{
			WindowAddress: fields[0],
			State:         state,
		}, nil

	case "bell":
		if err := expectFields(fields, 1, line); err != nil {
			return nil, err
		}

		return EventBell{
			WindowAddress: fields[0],
		}, nil

	default:
		return nil, fmt.Errorf("event %q not expected", name)
	}
}
