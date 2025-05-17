# gotk3-session-lock

A simple Go library that provides bindings for [gtk-session-lock](https://github.com/Cu3PO42/gtk-session-lock). It is intended to be used with [gotk3](https://github.com/gotk3/gotk3).
This library allows to create lockscreens for Wayland compositors using Go by utilizing the [ext-session-lock-v1](https://wayland.app/protocols/ext-session-lock-v1) protocol.

## Example

```go
package main

import (
	"log"

	"github.com/gotk3/gotk3/gdk"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"github.com/shelepuginivan/gotk3-session-lock/sessionlock"
)

var (
	app  *gtk.Application
	lock *sessionlock.Lock
)

func onFinished() {
	log.Println("session cannot be locked")
	app.Quit()
}

func onLocked() {
	log.Println("session is now locked")
}

func unlock() {
	lock.UnlockAndDestroy()

	display, err := gdk.DisplayGetDefault()
	if err != nil {
		log.Fatal(err)
	}

	display.Sync()

	log.Println("session is now unlocked")

	app.Quit()
}

func createLockWindow() (*gtk.Window, error) {
	win, err := gtk.WindowNew(gtk.WINDOW_TOPLEVEL)
	if err != nil {
		return nil, err
	}

	entry, err := gtk.EntryNew()
	if err != nil {
		return nil, err
	}

	entry.SetVAlign(gtk.ALIGN_CENTER)
	entry.SetHAlign(gtk.ALIGN_CENTER)

	entry.Connect("activate", func() {
		glib.IdleAdd(unlock)
	})

	win.Add(entry)

	return win, nil
}

func activate(app *gtk.Application) {
	lock = sessionlock.PrepareLock()

	lock.Connect("locked", onLocked)
	lock.Connect("finished", onFinished)

	display, err := gdk.DisplayGetDefault()
	if err != nil {
		log.Fatal(err)
	}

	lock.Lock()

	for i := range display.GetNMonitors() {
		mon, err := display.GetMonitor(i)
		if err != nil {
			continue
		}

		win, err := createLockWindow()
		if err != nil {
			continue
		}

		app.AddWindow(win)

		lock.NewSurface(win, mon)
		win.ShowAll()
	}
}

func main() {
	gtk.Init(nil)

	var err error

	app, err = gtk.ApplicationNew("com.github.shelepuginivan.Gotk3SessionLockExample", glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal(err)
	}

	app.Connect("activate", activate)
	app.Run(nil)
}
```
