package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Prueba canvas")
	myWidget := widget.NewEntry()

	window.SetContent(myWidget)
	window.Resize(fyne.NewSize(400, 400))

	window.ShowAndRun()

}
