package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Oort GUI")

	w.SetMainMenu(mainMenu())
	w.SetMaster()

	w.Resize(fyne.NewSize(640, 460))
	w.ShowAndRun()
}
