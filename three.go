package main

import (
    "fmt"
)

func main() {
    // Массив чисел
    numbers := []int{2, 4, 6, 8, 10}

    // Переменная для хранения суммы квадратов чисел
    sum := 0

    // Проходим по всем числам в массиве
    for _, num := range numbers {
        // Вычисляем квадрат числа и добавляем его к сумме
        square := num * num
        sum += square
    }

    // Выводим сумму квадратов в stdout
    fmt.Println("Сумма квадратов чисел:", sum)
}
