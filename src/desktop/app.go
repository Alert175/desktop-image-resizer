package desktop

import (
	"desktop-image-resizer/src/internal"
	"errors"
	"strconv"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
)

var pathToFolder string   // ссылка до папки
var imageWidth int        // ширина изображения
var isRemoveOutput bool   // очистить папку вывода
var isRecursiveCheck bool // Проверять вложенные папки

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

	recCheck := widget.NewCheck("Сканирование вложенных папок", func(b bool) {
		isRecursiveCheck = b
	})

	progress := widget.NewProgressBar()

	startBtn := widget.NewButton("Выполнить сжатие", func() {
		if len(pathToFolder) == 0 {
			err := errors.New("Не указан путь для сжатия")
			dialog.ShowError(err, window)
			return
		}
		imageWidth, errImageWidth := strconv.Atoi(widthInput.Text)
		if errImageWidth != nil || imageWidth == 0 {
			err := errors.New("Не указана ширина изображения")
			dialog.ShowError(err, window)
			return
		}
		fileList, _ := internal.ScanFolder(pathToFolder, isRecursiveCheck)
		if len(fileList) == 0 {
			err := errors.New("Не найдены изображения для сжатия")
			dialog.ShowError(err, window)
			return
		}
		errOutputFolder := internal.CheckFolder("./output")
		if errOutputFolder != nil {
			internal.CreateFolder("./output")
		}
		if isRemoveOutput == true {
			internal.RemoveFolder("./output")
			internal.CreateFolder("./output")
		}
		var errCounter int
		progress.Min = 0
		progress.Max = float64(len(fileList))
		for index, path := range fileList {
			_, err := internal.ImageResize(path, imageWidth, false, pathToFolder)
			if err != nil {
				errCounter += 1
			}
			progress.SetValue(float64(index + 1))
		}
		if errCounter > 0 {
			err := errors.New("Не удалось сжать " + strconv.Itoa(errCounter) + "изображений")
			dialog.ShowError(err, window)
		}
		internal.OpenWidthExplorer("./output")
	})

	content := container.NewVBox(widthInput, selectFolderBtn, rmCheck, recCheck, progress, startBtn)

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
		// time.Sleep(time.Millisecond * 250)
		progress.SetValue(i)
	}
}
