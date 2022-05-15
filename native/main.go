package main

import (
	"github.com/bhojpur/gui/pkg/engine/app"
	"github.com/bhojpur/gui/pkg/engine/theme"
)

func main() {
	// initilize the Bhojpur GUI application
	erp := app.NewWithID("net.bhojpur.erp")
	erp.SetIcon(theme.BhojpurLogo())
	wm := erp.NewWindow("Bhojpur ERP")
	wm.ShowAndRun()
}
