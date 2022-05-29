package desktop

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var pathToFolder string // ссылка до папки
var imageWidth int      // ширина изображения
var isRemoveOutput bool // очистить папку вывода

func Init() {
	application := app.New()

	window := application.NewWindow("Desktop image-resizer")

	selectFolderBtn := widget.NewButton("Выбрать папку", func() {
		showSelectFolder(window)
	})

	widthInput := widget.NewEntry()
	widthInput.SetPlaceHolder("Введите ширину изображения")

	rmCheck := widget.NewCheck("Очистить папку вывода", func(b bool) {
		isRemoveOutput = b
	})

	progress := widget.NewProgressBar()

	startBtn := widget.NewButton("Выполнить сжатие", func() {
		changeProgress(progress)
	})

	content := container.NewVBox(widthInput, selectFolderBtn, rmCheck, progress, startBtn)

	window.SetContent(content)
	window.Resize(fyne.NewSize(800, 500))
	window.Show()

	backgroundProcessInit()
	application.Run()
}

// открыть диалоговое окно и вернуть путь до папки
func showSelectFolder(window fyne.Window) {
	dialog.ShowFolderOpen(func(dir fyne.ListableURI, err error) {
		if err != nil {
			dialog.ShowError(err, window)
			pathToFolder = ""
		}
		pathToFolder = dir.Path()
	}, window)
}

// Запуск фонового процеса для внутренних обновлений программы
func backgroundProcessInit() {
	go func() {
		for range time.Tick(time.Second) {

		}
	}()
}

func changeProgress(progress *widget.ProgressBar) {
	for i := 0.0; i <= 1.0; i += 0.1 {
		time.Sleep(time.Millisecond * 250)
		progress.SetValue(i)
	}
}
