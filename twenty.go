package main

import (
	"fmt"
	"strings"
)

// Функция для переворачивания слов в строке
func reverseWords(sentence string) string {
	// Разбиваем строку на слова
	words := strings.Fields(sentence)
	// Создаем пустую строку для хранения результата
	reversed := ""

	// Переворачиваем порядок слов в строке
	for i := len(words) - 1; i >= 0; i-- {
		reversed += words[i] + " "
	}

	// Удаляем лишний пробел в конце строки
	reversed = strings.TrimSpace(reversed)

	return reversed
}

func main() {
	// Пример строки
	input := "snow dog sun"

	// Выводим исходную строку
	fmt.Println("Исходная строка:", input)

	// Переворачиваем слова в строке
	reversed := reverseWords(input)

	// Выводим перевернутую строку
	fmt.Println("Перевернутая строка:", reversed)
}
