package gohypr

type Event interface {
	private()
}

type (
	EventWorkspace struct {
		WorkspaceName string
	}
	EventWorkspaceV2 struct {
		WorkspaceID   string
		WorkspaceName string
	}
	EventFocusedMon struct {
		MonitorName   string
		WorkspaceName string
	}
	EventFocusedMonV2 struct {
		MonitorName string
		WorkspaceID string
	}
	EventActiveWindow struct {
		WindowClass string
		WindowTitle string
	}
	EventActiveWindowV2 struct {
		WindowAddress string
	}
	EventFullscreen struct {
		State bool
	}
	EventMonitorRemoved struct {
		MonitorName string
	}
	EventMonitorRemovedV2 struct {
		MonitorID          string
		MonitorName        string
		MonitorDescription string
	}
	EventMonitorAdded struct {
		MonitorName string
	}
	EventMonitorAddedV2 struct {
		MonitorID          string
		MonitorName        string
		MonitorDescription string
	}
	EventCreateWorkspace struct {
		WorkspaceName string
	}
	EventCreateWorkspaceV2 struct {
		WorkspaceID   string
		WorkspaceName string
	}
	EventDestroyWorkspace struct {
		WorkspaceName string
	}
	EventDestroyWorkspaceV2 struct {
		WorkspaceID   string
		WorkspaceName string
	}
	EventMoveWorkspace struct {
		WorkspaceName string
		MonitorName   string
	}
	EventMoveWorkspaceV2 struct {
		WorkspaceID   string
		WorkspaceName string
		MonitorName   string
	}
	EventRenameWorkspace struct {
		WorkspaceID string
		NewName     string
	}
	EventActiveSpecial struct {
		WorkspaceName string
		MonitorName   string
	}
	EventActiveSpecialV2 struct {
		WorkspaceID   string
		WorkspaceName string
		MonitorName   string
	}
	EventActiveLayout struct {
		KeyboardName string
		LayoutName   string
	}
	EventOpenWindow struct {
		WindowAddress string
		WorkspaceName string
		WindowClass   string
		WindowTitle   string
	}
	EventCloseWindow struct {
		WindowAddress string
	}
	EventKill struct {
		WindowAddress string
	}
	EventMoveWindow struct {
		WindowAddress string
		WorkspaceName string
	}
	EventMoveWindowV2 struct {
		WindowAddress string
		WorkspaceID   string
		WorkspaceName string
	}
	EventOpenLayer struct {
		NameSpace string
	}
	EventCloseLayer struct {
		NameSpace string
	}
	EventSubmap struct {
		SubmapName string
	}
	EventChangeFloatingMode struct {
		WindowAddress string
		Floating      bool
	}
	EventUrgent struct {
		WindowAddress string
	}
	EventScreencast struct {
		State bool
		Owner string
	}
	EventScreencastV2 struct {
		State bool
		Owner string
		Name  string
	}
	EventWindowTitle struct {
		WindowAddress string
	}
	EventWindowTitleV2 struct {
		WindowAddress string
		WindowTitle   string
	}
	EventToggleGroup struct {
		State         bool
		WindowAddress []string
	}
	EventMoveIntoGroup struct {
		WindowAddress string
	}
	EventMoveOutOfGroup struct {
		WindowAddress string
	}
	EventIgnoreGroupLock struct {
		State bool
	}
	EventLockGroups struct {
		State bool
	}
	EventConfigReloaded struct{}
	EventPin            struct {
		WindowAddress string
		PinState      bool
	}
	EventMinimized struct {
		WindowAddress string
		State         bool
	}
	EventBell struct {
		WindowAddress string
	}
)

func (e EventWorkspace) private()          {}
func (e EventWorkspaceV2) private()        {}
func (e EventFocusedMon) private()         {}
func (e EventFocusedMonV2) private()       {}
func (e EventActiveWindow) private()       {}
func (e EventActiveWindowV2) private()     {}
func (e EventFullscreen) private()         {}
func (e EventMonitorRemoved) private()     {}
func (e EventMonitorRemovedV2) private()   {}
func (e EventMonitorAdded) private()       {}
func (e EventMonitorAddedV2) private()     {}
func (e EventCreateWorkspace) private()    {}
func (e EventCreateWorkspaceV2) private()  {}
func (e EventDestroyWorkspace) private()   {}
func (e EventDestroyWorkspaceV2) private() {}
func (e EventMoveWorkspace) private()      {}
func (e EventMoveWorkspaceV2) private()    {}
func (e EventRenameWorkspace) private()    {}
func (e EventActiveSpecial) private()      {}
func (e EventActiveSpecialV2) private()    {}
func (e EventActiveLayout) private()       {}
func (e EventOpenWindow) private()         {}
func (e EventCloseWindow) private()        {}
func (e EventKill) private()               {}
func (e EventMoveWindow) private()         {}
func (e EventMoveWindowV2) private()       {}
func (e EventOpenLayer) private()          {}
func (e EventCloseLayer) private()         {}
func (e EventSubmap) private()             {}
func (e EventChangeFloatingMode) private() {}
func (e EventUrgent) private()             {}
func (e EventScreencast) private()         {}
func (e EventScreencastV2) private()       {}
func (e EventWindowTitle) private()        {}
func (e EventWindowTitleV2) private()      {}
func (e EventToggleGroup) private()        {}
func (e EventMoveIntoGroup) private()      {}
func (e EventMoveOutOfGroup) private()     {}
func (e EventIgnoreGroupLock) private()    {}
func (e EventLockGroups) private()         {}
func (e EventConfigReloaded) private()     {}
func (e EventPin) private()                {}
func (e EventMinimized) private()          {}
func (e EventBell) private()               {}
