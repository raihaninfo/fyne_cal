package main

import (
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/Knetic/govaluate"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne Calculator")
	w.Resize(fyne.NewSize(300, 400))

	output := ""
	input := widget.NewLabel(output)
	historyOutput := ""
	historyLabel := widget.NewLabel(historyOutput)
	errorOutput := "Calculator"
	ErrorLabel := widget.NewLabel(errorOutput)

	copyright := "Made With ♥ by Raihan"
	copyrightLabel := widget.NewLabel(copyright)

	var historyArr []string
	isHistory := false

	// history button
	historyBtn := widget.NewButton("History", func() {
		if isHistory {
			historyOutput = ""
		} else {
			for i := len(historyArr) - 1; i >= 0; i-- {
				historyOutput = historyOutput + historyArr[i]
				historyOutput += "\n"
			}
		}
		isHistory = !isHistory
		historyLabel.SetText(historyOutput)
	})

	// back button
	backBtn := widget.NewButton("Back", func() {
		if len(output) > 0 {
			output = output[:len(output)-1]
			input.SetText(output)
		}

	})

	// clear button
	claerBtn := widget.NewButton("clear", func() {
		output = ""
		historyOutput = ""
		errorOutput = "Calculator"
		input.SetText(output)
		historyLabel.SetText(historyOutput)
		ErrorLabel.SetText(errorOutput)
	})

	// open button
	openBtn := widget.NewButton("(", func() {
		output = output + "("
		input.SetText(output)
	})

	//close button
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
		expression, err := govaluate.NewEvaluableExpression(output)
		if err == nil {
			result, err := expression.Evaluate(nil)
			if err == nil {
				ans := strconv.FormatFloat(result.(float64), 'f', -1, 64)
				strToAppend := output + " + " + ans
				historyArr = append(historyArr, strToAppend)
				output = ans
			} else {
				errorOutput = "error"
			}
		} else {
			errorOutput = "error"

		}
		ErrorLabel.SetText(errorOutput)
		input.SetText(output)
	})

	w.SetContent(container.NewVBox(
		ErrorLabel,
		input,
		historyLabel,
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
		copyrightLabel,
	))

	w.ShowAndRun()

}
