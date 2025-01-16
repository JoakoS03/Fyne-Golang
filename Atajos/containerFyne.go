package main


/**

A fyne.Layout implementa un método para organizar artículos dentro de un contenedor.
 Al descomentar el container.New() 
 línea en este ejemplo usted altere el contenedor para usar un diseño de cuadrícula con 2 columnas.
 Ejecute este código e intente cambiar el tamaño de la ventana para ver cómo se 
 configura automáticamente el diseño el contenido de la ventana. Observe también que la posición 
 manual de text2 es ignorado por el código de diseño.

*/

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func main() {
	myApp := app.New()
	window := myApp.NewWindow("Prueba canvas")

	green := color.NRGBA{R: 0, G: 180, B: 0, A: 255}

	text1 := canvas.NewText("Hello", green)
	text2 := canvas.NewText("Mundo", green)

	//content := container.NewWithoutLayout(text1, text2)
	content := container.New(layout.NewGridLayout(2), text1, text2)

	window.Resize(fyne.NewSize(400, 400))
	window.SetContent(content)

	window.ShowAndRun()

}
