package main

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func setContentToText(c fyne.Canvas) {
	color := color.NRGBA{R: 0, G: 180, B: 0, A: 255}
	text := canvas.NewText("Hola soy un circulo", color)
	text.TextStyle.Bold = true
	c.SetContent(text)
}

func setContentToCircle(c fyne.Canvas) {
	red := color.NRGBA{R: 0xff, G: 0x33, B: 0x33, A: 0xff}
	circle := canvas.NewCircle(color.White)
	circle.StrokeWidth = 4
	circle.StrokeColor = red

	c.SetContent(circle)

}

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Prueba canvas")

	myCanvas := window.Canvas()

	color1 := color.NRGBA{R: 0, G: 100, B: 50, A: 255}
	dibujo := canvas.NewRectangle(color1)
	myCanvas.SetContent(dibujo)

	go func() {
		time.Sleep(time.Second)
		setContentToText(myCanvas)
		canvas.Refresh(myCanvas.Content())
	}()

	go func() {
		time.Sleep(time.Second)
		setContentToCircle(myCanvas)
		canvas.Refresh(myCanvas.Content())
	}()

	window.Resize(fyne.NewSize(400, 400))
	window.ShowAndRun()

}
