package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"strconv"
)

func main() {
	// Crear la aplicación
	myApp := app.New()
	w := myApp.NewWindow("Calculadora")

	// Variables para la lógica
	var result string
	var currentOperation string
	var operand1, operand2 float64

	// Pantalla de resultados
	resultLabel := widget.NewLabel("0")

	// Función para actualizar la pantalla
	updateDisplay := func(value string) {
		resultLabel.SetText(value)
	}

	// Función para manejar los clics de botones numéricos
	handleNumberClick := func(num string) {
		result += num
		updateDisplay(result)
	}

	// Función para manejar los operadores
	handleOperatorClick := func(op string) {
		if result != "" {
			operand1, _ = strconv.ParseFloat(result, 64)
			currentOperation = op
			result = ""
		}
	}

	// Función para calcular el resultado
	handleEqualsClick := func() {
		if result != "" {
			operand2, _ = strconv.ParseFloat(result, 64)
			var res float64
			switch currentOperation {
			case "+":
				res = operand1 + operand2
			case "-":
				res = operand1 - operand2
			case "*":
				res = operand1 * operand2
			case "/":
				if operand2 != 0 {
					res = operand1 / operand2
				} else {
					updateDisplay("Error")
					return
				}
			}
			result = strconv.FormatFloat(res, 'f', -1, 64)
			updateDisplay(result)
		}
	}

	// Función para limpiar todo
	handleClearClick := func() {
		result = ""
		currentOperation = ""
		operand1, operand2 = 0, 0
		updateDisplay("0")
	}

	// Crear los botones
	buttons := []struct {
		label    string
		onTapped func()
	}{
		{"7", func() { handleNumberClick("7") }},
		{"8", func() { handleNumberClick("8") }},
		{"9", func() { handleNumberClick("9") }},
		{"/", func() { handleOperatorClick("/") }},
		{"4", func() { handleNumberClick("4") }},
		{"5", func() { handleNumberClick("5") }},
		{"6", func() { handleNumberClick("6") }},
		{"*", func() { handleOperatorClick("*") }},
		{"1", func() { handleNumberClick("1") }},
		{"2", func() { handleNumberClick("2") }},
		{"3", func() { handleNumberClick("3") }},
		{"-", func() { handleOperatorClick("-") }},
		{"0", func() { handleNumberClick("0") }},
		{".", func() { handleNumberClick(".") }},
		{"=", func() { handleEqualsClick() }},
		{"+", func() { handleOperatorClick("+") }},
		{"C", func() { handleClearClick() }},
	}

	// Crear el contenedor de botones
	grid := container.NewGridWithColumns(4)
	for _, btn := range buttons {
		b := widget.NewButton(btn.label, btn.onTapped)
		grid.Add(b)
	}

	// Crear el layout principal
	content := container.NewVBox(
		resultLabel,
		grid,
	)

	// Configurar el contenido de la ventana
	w.SetContent(content)
	w.Resize(fyne.NewSize(300, 400))
	w.ShowAndRun()
}
