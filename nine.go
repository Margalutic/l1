package main

import (
	"fmt"
)

// Функция, которая принимает числа из первого канала,
// умножает их на 2 и отправляет во второй канал
func multiplier(input <-chan int, output chan<- int) {
	for num := range input {
		output <- num * 2
	}
	close(output)
}

func main() {
	// Создаем два канала: для ввода чисел и для вывода результатов
	inputChan := make(chan int)
	outputChan := make(chan int)

	// Запускаем горутину, которая умножает числа на 2
	go multiplier(inputChan, outputChan)

	// Записываем числа в первый канал
	numbers := []int{1, 2, 3, 4, 5}
	go func() {
		for _, num := range numbers {
			inputChan <- num
		}
		close(inputChan)
	}()

	// Выводим результаты из второго канала
	for result := range outputChan {
		fmt.Println("Результат:", result)
	}
}
