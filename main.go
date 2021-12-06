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

	historyBtn := widget.NewButton("History", func() {

	})
	backBtn := widget.NewButton("Back", func() {

	})
	claerBtn := widget.NewButton("clear", func() {

	})
	openBtn := widget.NewButton("(", func() {})
	closeBtn := widget.NewButton(")", func() {})
	divisionBtn := widget.NewButton("/", func() {

	})
	sevenBtn := widget.NewButton("7", func() {

	})
	eightBtn := widget.NewButton("8", func() {

	})
	nineBtn := widget.NewButton("9", func() {

	})
	multiplyBtn := widget.NewButton("*", func() {

	})
	foreBtn := widget.NewButton("4", func() {

	})
	fiveBtn := widget.NewButton("5", func() {

	})
	sixBtn := widget.NewButton("6", func() {

	})
	minusBtn := widget.NewButton("-", func() {

	})
	oneBtn := widget.NewButton("1", func() {

	})
	twoBtn := widget.NewButton("2", func() {

	})
	threeBtn := widget.NewButton("3", func() {

	})
	plusBtn := widget.NewButton("+", func() {

	})
	zeroBtn := widget.NewButton("0", func() {

	})
	dotBtn := widget.NewButton(".", func() {

	})
	equalBtn := widget.NewButton("=", func() {

	})

	hello := widget.NewLabel("Hello Fyne!")
	w.SetContent(container.NewVBox(
		hello,
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
