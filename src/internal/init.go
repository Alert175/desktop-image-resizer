package internal

import (
	"fmt"
)

// список разрешенных расширений файлов
var AccessExtensions = []string{"jpg", "jpeg", "png"}

// Получение данных об содержании каталога input
func InitApp() {
	fmt.Println("Сканирование...")
	exitStr := ""
	fileList, errorFileList := ScanFolder("./input", true)
	if errorFileList != nil {
		fmt.Println("Ошибка чтения каталога input ", errorFileList)
	}
	if len(fileList) == 0 {
		fmt.Println("Файлы не обнаружены")
		fmt.Println("Введите что-либо для продолжения")
		fmt.Scan(&exitStr)
		return
	}

	var imageWidth int
	var resultFileCounter int
	var isAddSuffix string
	var isDeleteOutput string

	fmt.Println("Найдено файлов:", len(fileList))

	fmt.Println("Укажите ширину изображений: ")
	fmt.Scan(&imageWidth)
	fmt.Println("Добавлять суффикс к названию файла? y/n: ")
	fmt.Scan(&isAddSuffix)
	fmt.Println("Очистить папку output? y/n: ")
	fmt.Scan(&isDeleteOutput)

	// очистка папки output
	if isDeleteOutput == "y" {
		err := CheckFolder("./output")
		if err != nil {
			CreateFolder("./output")
		} else {
			RemoveFolder("./output")
			CreateFolder("./output")
		}
	}

	for _, path := range fileList {
		_, err := ImageResize(path, imageWidth, isAddSuffix == "y", "")
		if err != nil {
			fmt.Println(err)
		} else {
			resultFileCounter = resultFileCounter + 1
			exitStr = exitStr + "-"
			fmt.Print("-")
		}
	}

	_, resultErrorFileList := ScanFolder("./output", true)
	if resultErrorFileList != nil {
		fmt.Println("Ошибка чтения каталога input ", resultErrorFileList)
	}
	fmt.Println("")
	fmt.Println("Создано файлов: ", resultFileCounter)
	fmt.Println("Введите что-либо для продолжения")
	fmt.Scan(&exitStr)
}
