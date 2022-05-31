package internal

import "strings"

// Валидатор расширения файла
func ExtensionValidator(fileName string, accessExtension []string) bool {
	for _, extension := range accessExtension {
		if strings.Contains(fileName, extension) {
			return true
		}
	}
	return false
}
