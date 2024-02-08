package main

import (
    "fmt"
)

func main() {
    // Массив чисел
    numbers := []int{2, 4, 6, 8, 10}

    // Проходим по всем числам в массиве
    for _, num := range numbers {
        // Вычисляем квадрат числа и выводим его
        fmt.Println(num * num)
    }
}
