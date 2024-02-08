package main

import (
	"fmt"
)

// Функция для установки i-го бита в числе num в значение bit
func setBit(num int64, i int, bit int) int64 {
	// Маска для установки i-го бита в 1
	mask := int64(1 << i)

	if bit == 1 {
		// Устанавливаем i-й бит в 1
		num |= mask
	} else {
		// Устанавливаем i-й бит в 0
		num &^= mask
	}

	return num
}

func main() {
	// Пример числа типа int64
	num := int64(10) // 1010 в двоичной системе

	// Устанавливаем 3-й бит в 1
	num = setBit(num, 3, 1)

	// Выводим результат
	fmt.Printf("Число после установки бита: %d\n", num) // Ожидаемый результат: 14 (1110 в двоичной системе)
}
