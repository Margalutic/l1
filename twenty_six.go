package main

import (
	"fmt"
	"strings"
)

// IsUnique проверяет, все ли символы в строке уникальны (без учета регистра).
func IsUnique(str string) bool {
	// Преобразование строки в нижний регистр
	str = strings.ToLower(str)
	// Мапа для хранения встреченных символов
	charMap := make(map[rune]bool)

	// Перебор символов строки
	for _, char := range str {
		// Если символ уже встречался, строка не уникальна
		if charMap[char] {
			return false
		}
		// Добавление символа в мапу
		charMap[char] = true
	}

	// Если все символы уникальны, возвращаем true
	return true
}

func main() {
	// Тестовые строки
	testStrings := []string{"abcd", "abCdefAaf", "aabcd"}

	// Проверка уникальности для каждой строки
	for _, str := range testStrings {
		fmt.Printf("Строка '%s' уникальна? %t\n", str, IsUnique(str))
	}
}
