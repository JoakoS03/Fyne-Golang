package main

import (
	"fmt"
	"pruebas/utils"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type numericalEntry struct {
	widget.Entry
}

type Gasto struct {
	servicio string
	monto    float64
}

func newNumericalEntry() *numericalEntry {
	entry := &numericalEntry{}
	entry.ExtendBaseWidget(entry)

	return entry
}

func (e *numericalEntry) TypedRune(r rune) {
	if (r >= '0' && r <= '9') || r == '.' || r == ',' {
		e.Entry.TypedRune(r)
	}
}

func (e *numericalEntry) TypedShortcut(shortcut fyne.Shortcut) {
	paste, ok := shortcut.(*fyne.ShortcutPaste)
	if !ok {
		e.Entry.TypedShortcut(shortcut)
		return
	}
	content := paste.Clipboard.Content()
	if _, err := strconv.ParseFloat(content, 64); err == nil {
		e.Entry.TypedShortcut(shortcut)
	}
}

func main() {
	var service utils.Serve
	var servicios []string

	err := service.ConfigCredential()
	if err != nil {
		fmt.Println(err)
	}

	// Selector de servicios
	widgetSelect := widget.NewSelect(servicios, func(selected string) {})
	widgetSelect.PlaceHolder = ""

	myApp := app.New()
	window := myApp.NewWindow("Gastos mensuales")

	// Entrada de servicio
	inputService := widget.NewEntry()
	inputService.SetPlaceHolder("Ingrese un servicio")

	svrs, err := service.GetService()
	if err != nil {
		fmt.Println(err)
	}

	for _, value := range svrs {
		servicios = append(servicios, value.(string))
		widgetSelect.Options = servicios
	}

	// Botón para agregar servicio
	agregarServicio := widget.NewButton("Agregar servicio", func() {
		if inputService.Text != "" {
			// Agregar servicio a la hoja
			err := service.AddService(inputService.Text)
			if err == nil {
				// Actualizar las opciones del selector
				servicios = append(servicios, inputService.Text)
				widgetSelect.Options = servicios
				widgetSelect.Refresh()
				inputService.SetText("")
			} else {
				fmt.Println("Error al agregar servicio:", err)
			}
		}
	})

	// Entrada numérica para gastos
	inputNumber := newNumericalEntry()
	inputNumber.SetPlaceHolder("Ingrese monto")

	// Contenedor para mostrar gastos
	innerContainer := container.NewVBox()

	gastos, _ := service.GetGastos()
	str := " " + gastos[0][0].(string) + "  " + gastos[0][1].(string)
	lbl := widget.NewLabel(str)
	innerContainer.Add(lbl)
	innerContainer.Refresh()

	if len(gastos) > 1 {
		for _, value := range gastos {
			str := " " + value[0].(string) + "  " + value[1].(string)
			lbl := widget.NewLabel(str)
			innerContainer.Add(lbl)
			innerContainer.Refresh()
		}
	}

	// Botón para agregar gasto
	agregarGasto := widget.NewButton("Agregar gasto", func() {
		if widgetSelect.Selected != "" && inputNumber.Text != "" {
			err := service.AddGasto(widgetSelect.Selected, inputNumber.Text)
			if err == nil {
				// Agregar gasto al contenedor dinámicamente
				str := "" + widgetSelect.Selected + " $" + inputNumber.Text

				lbl := widget.NewLabel(str)
				innerContainer.Add(lbl)
				innerContainer.Refresh()
				inputNumber.SetText("")
			} else {
				fmt.Println(err)
			}
		}
	})

	// Layout principal
	content := container.NewGridWithColumns(2,
		container.NewVBox(
			inputService,
			agregarServicio,
			widgetSelect,
			inputNumber,
			agregarGasto,
		),
		innerContainer,
	)

	window.SetContent(content)
	window.Resize(fyne.NewSize(500, 300))
	window.ShowAndRun()
}
