package gohypr

import (
	"bufio"
	"net"
)

type (
	client struct {
		socketFinder func() (string, error)
	}

	EventResult struct {
		Event Event
		Err   error
	}
)

func NewClient(options ...option) *client {
	l := &client{socketFinder: DefaultSocketFinder}
	for _, option := range options {
		option(l)
	}

	return l
}

func (l *client) getConn() (*net.UnixConn, error) {
	path, err := l.socketFinder()
	if err != nil {
		return nil, err
	}

	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: path, Net: "unix"})
	if err != nil {
		return nil, err
	}

	return conn, nil
}

func (l *client) each(yield func(EventResult) bool) {
	path, err := l.socketFinder()
	if err != nil {
		yield(EventResult{Err: err})
		return
	}

	conn, err := net.DialUnix("unix", nil, &net.UnixAddr{Name: path, Net: "unix"})
	if err != nil {
		yield(EventResult{Err: err})
		return
	}
	defer func() { _ = conn.Close() }()

	for r := bufio.NewReader(conn); ; {
		line, err := r.ReadString('\n')
		if err != nil {
			yield(EventResult{Err: err})
			return
		}

		e, err := parseEvent(line)
		if !yield(EventResult{Event: e, Err: err}) {
			return
		}
	}
}

func (l *client) Events() <-chan EventResult {
	events := make(chan EventResult, 32)

	go func() {
		defer close(events)
		l.each(func(r EventResult) bool {
			events <- r
			return true
		})
	}()

	return events
}

func (c *client) OnWorkspace(f func(e EventWorkspace, err error)) {
	c.each(func(r EventResult) bool {
		e, ok := r.Event.(EventWorkspace)
		if ok {
			f(e, r.Err)
		}
		return true
	})
}

func (c *client) OnWorkspaceV2(f func(e EventWorkspaceV2, err error))               {}
func (c *client) OnFocusedMon(f func(e EventFocusedMon, err error))                 {}
func (c *client) OnFocusedMonV2(f func(e EventFocusedMonV2, err error))             {}
func (c *client) OnActiveWindow(f func(e EventActiveWindow, err error))             {}
func (c *client) OnActiveWindowV2(f func(e EventActiveWindowV2, err error))         {}
func (c *client) OnFullscreen(f func(e EventFullscreen, err error))                 {}
func (c *client) OnMonitorRemoved(f func(e EventMonitorRemoved, err error))         {}
func (c *client) OnMonitorRemovedV2(f func(e EventMonitorRemovedV2, err error))     {}
func (c *client) OnMonitorAdded(f func(e EventMonitorAdded, err error))             {}
func (c *client) OnMonitorAddedV2(f func(e EventMonitorAddedV2, err error))         {}
func (c *client) OnCreateWorkspace(f func(e EventCreateWorkspace, err error))       {}
func (c *client) OnCreateWorkspaceV2(f func(e EventCreateWorkspaceV2, err error))   {}
func (c *client) OnDestroyWorkspace(f func(e EventDestroyWorkspace, err error))     {}
func (c *client) OnDestroyWorkspaceV2(f func(e EventDestroyWorkspaceV2, err error)) {}
func (c *client) OnMoveWorkspace(f func(e EventMoveWorkspace, err error))           {}
func (c *client) OnMoveWorkspaceV2(f func(e EventMoveWorkspaceV2, err error))       {}
func (c *client) OnRenameWorkspace(f func(e EventRenameWorkspace, err error))       {}
func (c *client) OnActiveSpecial(f func(e EventActiveSpecial, err error))           {}
func (c *client) OnActiveSpecialV2(f func(e EventActiveSpecialV2, err error))       {}
func (c *client) OnActiveLayout(f func(e EventActiveLayout, err error))             {}
func (c *client) OnOpenWindow(f func(e EventOpenWindow, err error)) {
	c.each(func(r EventResult) bool {
		e, ok := r.Event.(EventOpenWindow)
		if ok {
			f(e, r.Err)
		}
		return true
	})
}
func (c *client) OnCloseWindow(f func(e EventCloseWindow, err error))               {}
func (c *client) OnKill(f func(e EventKill, err error))                             {}
func (c *client) OnMoveWindow(f func(e EventMoveWindow, err error))                 {}
func (c *client) OnMoveWindowV2(f func(e EventMoveWindowV2, err error))             {}
func (c *client) OnOpenLayer(f func(e EventOpenLayer, err error))                   {}
func (c *client) OnCloseLayer(f func(e EventCloseLayer, err error))                 {}
func (c *client) OnSubmap(f func(e EventSubmap, err error))                         {}
func (c *client) OnChangeFloatingMode(f func(e EventChangeFloatingMode, err error)) {}
func (c *client) OnUrgent(f func(e EventUrgent, err error))                         {}
func (c *client) OnScreencast(f func(e EventScreencast, err error))                 {}
func (c *client) OnScreencastV2(f func(e EventScreencastV2, err error))             {}
func (c *client) OnWindowTitle(f func(e EventWindowTitle, err error))               {}
func (c *client) OnWindowTitleV2(f func(e EventWindowTitleV2, err error))           {}
func (c *client) OnToggleGroup(f func(e EventToggleGroup, err error))               {}
func (c *client) OnMoveIntoGroup(f func(e EventMoveIntoGroup, err error))           {}
func (c *client) OnMoveOutOfGroup(f func(e EventMoveOutOfGroup, err error))         {}
func (c *client) OnIgnoreGroupLock(f func(e EventIgnoreGroupLock, err error))       {}
func (c *client) OnLockGroups(f func(e EventLockGroups, err error))                 {}
func (c *client) OnConfigReloaded(f func(e EventConfigReloaded, err error))         {}
func (c *client) OnPin(f func(e EventPin, err error))                               {}
func (c *client) OnMinimized(f func(e EventMinimized, err error))                   {}
func (c *client) OnBell(f func(e EventBell, err error))                             {}
