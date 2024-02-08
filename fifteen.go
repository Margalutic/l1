package main

import (
	"fmt"
)

var justString string

// createHugeString - пример функции, которая возвращает строку.
func createHugeString(size int) string {
	// В данном примере используем просто строку "huge string"
	return "huge string"
}

func someFunc() {
	// Получаем строку от функции createHugeString
	v := createHugeString(1 << 10)

	// Проверяем, что длина строки не меньше 100 символов
	if len(v) >= 100 {
		// Присваиваем justString первые 100 символов строки v
		justString = v[:100]
	} else {
		// В случае, если длина строки меньше 100 символов, присваиваем всю строку
		justString = v
	}
}

func main() {
	someFunc()
	fmt.Println(justString)
}
