package main

import "fmt"

func main() {
    // Исходные числа
    a, b := 5, 10

    fmt.Println("Исходные числа:", a, b)

    // Меняем местами числа без создания временной переменной
    a = a ^ b
    b = a ^ b
    a = a ^ b

    fmt.Println("Числа после обмена:", a, b)
}
