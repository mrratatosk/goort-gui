package main

import "fyne.io/fyne/v2"

func mainMenu() *fyne.MainMenu {
	file := fyne.NewMenu("File", fyne.NewMenuItem("Open ROM...", func() {}))
	tools := fyne.NewMenu("Tools", fyne.NewMenuItem("Memory", func() {
		w := fyne.CurrentApp().NewWindow("Memory")
		w.Resize(fyne.NewSize(320, 240))
		w.Show()
	}))
	help := fyne.NewMenu("Help", fyne.NewMenuItem("About", func() {}))
	return fyne.NewMainMenu(file, tools, help)
}
