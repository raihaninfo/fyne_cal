package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Calculator")
	w.Resize(fyne.NewSize(300, 400))

	output := ""
	input := widget.NewLabel(output)

	historyBtn := widget.NewButton("History", func() {

	})
	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}

	})
	claerBtn := widget.NewButton("clear", func() {
		output = ""
		input.SetText(output)
	})
	openBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})
	closeBtn := widget.NewButton(")", func() {
		output = output + ")"
		input.SetText(output)
	})
	divisionBtn := widget.NewButton("/", func() {
		output = output + "/"
		input.SetText(output)
	})
	sevenBtn := widget.NewButton("7", func() {
		output = output + "7"
		input.SetText(output)
	})
	eightBtn := widget.NewButton("8", func() {
		output = output + "8"
		input.SetText(output)
	})
	nineBtn := widget.NewButton("9", func() {
		output = output + "9"
		input.SetText(output)
	})
	multiplyBtn := widget.NewButton("*", func() {
		output = output + "*"
		input.SetText(output)
	})
	foreBtn := widget.NewButton("4", func() {
		output = output + "4"
		input.SetText(output)
	})
	fiveBtn := widget.NewButton("5", func() {
		output = output + "5"
		input.SetText(output)
	})
	sixBtn := widget.NewButton("6", func() {
		output = output + "6"
		input.SetText(output)
	})
	minusBtn := widget.NewButton("-", func() {
		output = output + "-"
		input.SetText(output)
	})
	oneBtn := widget.NewButton("1", func() {
		output = output + "1"
		input.SetText(output)
	})
	twoBtn := widget.NewButton("2", func() {
		output = output + "2"
		input.SetText(output)
	})
	threeBtn := widget.NewButton("3", func() {
		output = output + "3"
		input.SetText(output)
	})
	plusBtn := widget.NewButton("+", func() {
		output = output + "+"
		input.SetText(output)
	})
	zeroBtn := widget.NewButton("0", func() {
		output = output + "0"
		input.SetText(output)
	})
	dotBtn := widget.NewButton(".", func() {
		output = output + "."
		input.SetText(output)
	})
	equalBtn := widget.NewButton("=", func() {

	})

	w.SetContent(container.NewVBox(
		input,
		container.NewGridWithColumns(1,
			container.NewGridWithColumns(2,
				historyBtn,
				backBtn,
			),
			container.NewGridWithColumns(4,
				claerBtn,
				openBtn,
				closeBtn,
				divisionBtn,
			),
			container.NewGridWithColumns(4,
				sevenBtn,
				eightBtn,
				nineBtn,
				multiplyBtn,
			),
			container.NewGridWithColumns(4,
				foreBtn,
				fiveBtn,
				sixBtn,
				minusBtn,
			),
			container.NewGridWithColumns(4,
				oneBtn,
				twoBtn,
				threeBtn,
				plusBtn,
			),
			container.NewGridWithColumns(2,
				container.NewGridWithColumns(2,
					zeroBtn,
					dotBtn,
				),
				container.NewGridWithColumns(1,
					equalBtn,
				),
			),
		),
	))

	w.ShowAndRun()

}
