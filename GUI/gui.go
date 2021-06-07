package main

import (
	"github.com/mattn/go-gtk/glib"
	"github.com/mattn/go-gtk/gtk"
)

func main() {

	gtk.Init(nil)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetPosition(gtk.WIN_POS_CENTER)
	window.SetTitle("DHA application")
	window.Connect("destroy", func(ctx *glib.CallbackContext) {
		println("got destroy!", ctx.Data().(string))
		gtk.MainQuit()
	}, "foo")

	// Box and window control
	vbox := gtk.NewVBox(false, 1)

	vpaned := gtk.NewVPaned()
	vbox.Add(vpaned)

	// Frames for encrypt decrypt
	frame1 := gtk.NewFrame("Encryption")
	framebox1 := gtk.NewVBox(false, 1)
	frame1.Add(framebox1)

	frame2 := gtk.NewFrame("Decryption")
	framebox2 := gtk.NewVBox(false, 1)
	frame2.Add(framebox2)

	vpaned.Pack1(frame1, false, false)
	vpaned.Pack2(frame2, false, false)

	// Encryption pane
	label := gtk.NewLabel("public key")

	framebox1.PackStart(label, false, true, 0)

	entry := gtk.NewEntry()

	framebox1.Add(entry)

	window.Add(vbox)
	window.SetSizeRequest(900, 400)
	window.ShowAll()
	gtk.Main()

}
