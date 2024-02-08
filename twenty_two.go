package main

import (
	"fmt"
)

func main() {
	// Значение переменных a и b больше 2^20
	a := 2 << 20
	b := 3 << 20

	// Умножение
	multiplication := a * b
	fmt.Printf("Умножение: %d\n", multiplication)

	// Деление
	division := a / b
	fmt.Printf("Деление: %d\n", division)

	// Сложение
	addition := a + b
	fmt.Printf("Сложение: %d\n", addition)

	// Вычитание
	subtraction := a - b
	fmt.Printf("Вычитание: %d\n", subtraction)
}
