package gohypr

type Event interface{}

type (
	EventWorkspace          struct{}
	EventWorkspaceV2        struct{}
	EventFocusedMon         struct{}
	EventFocusedMonV2       struct{}
	EventActiveWindow       struct{}
	EventActiveWindowV2     struct{}
	EventFullscreen         struct{}
	EventMonitorRemoved     struct{}
	EventMonitorRemovedV2   struct{}
	EventMonitorAdded       struct{}
	EventMonitorAddedV2     struct{}
	EventCreateWorkspace    struct{}
	EventCreateWorkspaceV2  struct{}
	EventDestroyWorkspace   struct{}
	EventDestroyWorkspaceV2 struct{}
	EventMoveWorkspace      struct{}
	EventMoveWorkspaceV2    struct{}
	EventRenameWorkspace    struct{}
	EventActiveSpecial      struct{}
	EventActiveSpecialV2    struct{}
	EventActiveLayout       struct{}
	EventOpenWindow         struct{}
	EventCloseWindow        struct{}
	EventKill               struct{}
	EventMoveWindow         struct{}
	EventMoveWindowV2       struct{}
	EventOpenLayer          struct{}
	EventCloseLayer         struct{}
	EventSubmap             struct{}
	EventChangeFloatingMode struct{}
	EventUrgent             struct{}
	EventScreencast         struct{}
	EventScreencastV2       struct{}
	EventWindowTitle        struct{}
	EventWindowTitleV2      struct{}
	EventToggleGroup        struct{}
	EventMoveIntoGroup      struct{}
	EventMoveOutOfGroup     struct{}
	EventIgnoreGroupLock    struct{}
	EventLockGroups         struct{}
	EventConfigReloaded     struct{}
	EventPin                struct{}
	EventMinimized          struct{}
	EventBell               struct{}
)

func AsType[T Event](event Event) (T, bool) {
	e, ok := event.(T)
	return e, ok
}
