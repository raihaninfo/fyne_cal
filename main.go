package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func main() {
	a := app.New()
	w := a.NewWindow("Fyne_cal")
	w.Resize(fyne.NewSize(300, 400))

	w.ShowAndRun()

}
