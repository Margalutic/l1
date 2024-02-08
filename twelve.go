package main

import "fmt"

func main() {
    // Исходная последовательность строк
    sequence := []string{"cat", "cat", "dog", "cat", "tree"}

    // Создаем пустой map для хранения уникальных строк
    set := make(map[string]bool)

    // Добавляем каждую строку в множество
    for _, str := range sequence {
        set[str] = true
    }

    // Выводим элементы множества (уникальные строки)
    fmt.Println("Собственное множество:")
    for str := range set {
        fmt.Println(str)
    }
}
