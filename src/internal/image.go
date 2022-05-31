package internal

import (
	"errors"
	"image"
	"image/jpeg"
	"image/png"
	"os"
	"strconv"
	"strings"

	"github.com/nfnt/resize"
)

/*
	Принимаю путь до файла, проверяю что это картинка, выполняю сжатие
*/
func ImageResize(filePath string, imageWidth int, isPrefix bool, slicePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	imageExtension := strings.Split(filePath, ".")[len(strings.Split(filePath, "."))-1]

	var imageName string
	if isPrefix {
		imageName = strings.Split(strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1], ".")[0] + "__w" + strconv.Itoa(imageWidth)
	} else {
		imageName = strings.Split(strings.Split(filePath, "/")[len(strings.Split(filePath, "/"))-1], ".")[0]
	}

	imagePathSlice := strings.Split(strings.Replace(filePath, slicePath, "", 1), "/")
	imagePathList := []string{"./output"}
	for index, element := range imagePathSlice {
		if index < len(imagePathSlice)-1 && index != 0 {
			imagePathList = append(imagePathList, element)
		}
	}
	imagePath := strings.Join(imagePathList, "/")

	if !IsJpg(imageExtension) && !IsPng(imageExtension) {
		return "", errors.New("не найдено валидное расширение файла")
	}

	var imageContent image.Image
	var compressImage image.Image

	if IsJpg(imageExtension) {
		imgC, err := jpeg.Decode(file)
		if err != nil {
			return "", err
		} else {
			imageContent = imgC
			compressImage = resize.Resize(uint(imageWidth), 0, imageContent, resize.MitchellNetravali)
		}
	}
	if IsPng(imageExtension) {
		imgC, err := png.Decode(file)
		if err != nil {
			return "", err
		} else {
			imageContent = imgC
			compressImage = resize.Resize(uint(imageWidth), 0, imageContent, resize.Bicubic)
		}
	}

	file.Close()

	if err := CheckFolder(imagePath); err != nil {
		CreateFolder(imagePath)
	}

	compressFile, err := os.Create(imagePath + "/" + imageName + "." + imageExtension)

	if err != nil {
		return "", err
	}
	defer compressFile.Close()

	if IsJpg(imageExtension) {
		jpeg.Encode(compressFile, compressImage, nil)
	}
	if IsPng(imageExtension) {
		png.Encode(compressFile, compressImage)
	}

	return "./output/" + imageName + "." + imageExtension, nil
}

// проверка, на то что изображение формата jpg
func IsJpg(argExtension string) bool {
	for _, element := range []string{"jpeg", "jpg"} {
		if argExtension == element {
			return true
		}
	}
	return false
}

// проверка, на то что изображение формата png
func IsPng(argExtension string) bool {
	for _, element := range []string{"png"} {
		if argExtension == element {
			return true
		}
	}
	return false
}
