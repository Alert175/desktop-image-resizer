package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

func main() {
	application := app.New()
	window := application.NewWindow("Desktop image-resizer")

	window.SetContent(widget.NewLabel("Hello World!"))
	window.Resize(fyne.NewSize(800, 500))
	window.Show()
	application.Run()
}
