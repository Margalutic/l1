package main

import (
	"fmt"
	"unicode/utf8"
)

// Функция для переворачивания строки
func reverseString(s string) string {
	// Преобразуем строку в массив рун
	runes := []rune(s)
	// Получаем длину строки
	length := len(runes)

	// Переворачиваем строку, меняя местами символы в начале и в конце строки
	for i, j := 0, length-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Возвращаем перевернутую строку
	return string(runes)
}

func main() {
	// Пример строки с символами Unicode
	input := "главрыба"

	// Выводим исходную строку
	fmt.Println("Исходная строка:", input)

	// Переворачиваем строку
	reversed := reverseString(input)

	// Выводим перевернутую строку
	fmt.Println("Перевернутая строка:", reversed)
}
