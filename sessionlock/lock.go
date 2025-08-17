package sessionlock

// #cgo pkg-config: gtk+-3.0 gtk-session-lock-0
// #include <gtk/gtk.h>
// #include <gtk-session-lock.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
)

func gobool(b C.gboolean) bool {
	return b != C.FALSE
}

func nativeWindow(window *gtk.Window) *C.GtkWindow {
	w := window.Native()
	wp := (*C.GtkWindow)(unsafe.Pointer(w))
	return wp
}

func nativeMonitor(monitor *gdk.Monitor) *C.GdkMonitor {
	m := monitor.Native()
	mp := (*C.GdkMonitor)(unsafe.Pointer(m))
	return mp
}

// GetMajorVersion returns the major version number of the gtk-session-lock
// library.
func GetMajorVersion() uint {
	v := C.gtk_session_lock_get_major_version()
	return uint(v)
}

// GetMinorVersion returns the minor version number of the gtk-session-lock
// library.
func GetMinorVersion() uint {
	v := C.gtk_session_lock_get_minor_version()
	return uint(v)
}

// GetMicroVersion returns the micro/patch version number of the
// gtk-session-lock library.
func GetMicroVersion() uint {
	v := C.gtk_session_lock_get_micro_version()
	return uint(v)
}

// IsSupported reports whether platform is Wayland and Wayland compositor
// supports the ext_session_lock_v1 protocol.
func IsSupported() bool {
	b := C.gtk_session_lock_is_supported()
	return gobool(b)
}

// GetProtocolVersion returns version of the ext_session_lock_v1 protocol
// supported by the compositor or 0 if the protocol is not supported.
func GetProtocolVersion() uint {
	v := C.gtk_session_lock_get_protocol_version()
	return uint(v)
}

// IsLockWindow reports whether window has been initialized as a lock surface.
func IsLockWindow(window *gtk.Window) bool {
	wp := nativeWindow(window)
	b := C.gtk_session_lock_is_lock_window(wp)
	return gobool(b)
}

// Lock performs locking operation and manages lock surfaces.
//
// Signals should be connected before [Lock.Lock] is executed. The following
// signals are supported:
//   - locked
//   - finished
//
// Lock cannot be reused after one of its methods [Lock.Destroy] or
// [Lock.UnlockAndDestroy] was called.
type Lock struct {
	*glib.Object

	ptr *C.GtkSessionLockLock
}

// PrepareLock prepares a new [Lock].
func PrepareLock() *Lock {
	gp := C.gtk_session_lock_prepare_lock()
	obj := &glib.Object{
		GObject: glib.ToGObject(unsafe.Pointer(gp)),
	}

	return &Lock{
		Object: obj,

		ptr: gp,
	}
}

// Lock performs the locking operation. Signals should be connected before
// running this method. The compositor will hide all surfaces except those
// created via [Lock.NewSurface] method.
func (l *Lock) Lock() {
	C.gtk_session_lock_lock_lock(l.ptr)
}

// Destroy destroys an inactive lock object. This method should be called only
// after finished signal is received.
func (l *Lock) Destroy() {
	C.gtk_session_lock_lock_destroy(l.ptr)
}

// UnlockAndDestroy unlocks active session lock and disposes of it.
func (l *Lock) UnlockAndDestroy() {
	C.gtk_session_lock_lock_unlock_and_destroy(l.ptr)
}

// NewSurface initializes window as a lock surface for the given monitor.
//
// This method must be called after [Lock.Lock]. If the session
// is locked successfully, the specified window will be shown on the given
// monitor. You must only ever call this method once for a given lock and
// monitor. The window will automatically be stretched to cover the entire
// screen.
func (l *Lock) NewSurface(window *gtk.Window, monitor *gdk.Monitor) {
	wp := nativeWindow(window)
	mp := nativeMonitor(monitor)

	C.gtk_session_lock_lock_new_surface(l.ptr, wp, mp)
}

// UnmapLockWindow unmaps the surface if the given window is a lock window.
//
// This method must be called before the window is unmapped (e.g. hidden).
func UnmapLockWindow(window *gtk.Window) {
	wp := nativeWindow(window)
	C.gtk_session_lock_unmap_lock_window(wp)
}
