package sessionlock

// #cgo pkg-config: gtk+-3.0 gtk-session-lock-0
// #include <gtk/gtk.h>
// #include <gtk-session-lock.h>
import "C"

import (
	"unsafe"

	"github.com/gotk3/gotk3/gdk"
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
