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

func GetMajorVersion() uint {
	v := C.gtk_session_lock_get_major_version()
	return uint(v)
}

func GetMinorVersion() uint {
	v := C.gtk_session_lock_get_minor_version()
	return uint(v)
}

func GetMicroVersion() uint {
	v := C.gtk_session_lock_get_micro_version()
	return uint(v)
}

func IsSupported() bool {
	b := C.gtk_session_lock_is_supported()
	return gobool(b)
}

func GetProtocolVersion() uint {
	v := C.gtk_session_lock_get_protocol_version()
	return uint(v)
}

func IsLockWindow(window *gtk.Window) bool {
	wp := nativeWindow(window)
	b := C.gtk_session_lock_is_lock_window(wp)
	return gobool(b)
}

type Lock struct {
	*glib.Object

	ptr *C.GtkSessionLockLock
}

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

func (l *Lock) Lock() {
	C.gtk_session_lock_lock_lock(l.ptr)
}

func (l *Lock) Destroy() {
	C.gtk_session_lock_lock_destroy(l.ptr)
}

func (l *Lock) UnlockAndDestroy() {
	C.gtk_session_lock_lock_unlock_and_destroy(l.ptr)
}

func (l *Lock) NewSurface(window *gtk.Window, monitor *gdk.Monitor) {
	wp := nativeWindow(window)
	mp := nativeMonitor(monitor)

	C.gtk_session_lock_lock_new_surface(l.ptr, wp, mp)
}

func UnmapLockWindow(window *gtk.Window) {
	wp := nativeWindow(window)
	C.gtk_session_lock_unmap_lock_window(wp)
}
