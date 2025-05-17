package sessionlock

// #cgo pkg-config: gtk+-3.0 gtk-session-lock-0
// #include <gtk/gtk.h>
// #include <gtk-session-lock.h>
import "C"

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
