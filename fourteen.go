package main

import (
    "fmt"
)

func checkType(x interface{}) {
    // Проверяем тип переменной
    switch x.(type) {
    case int:
        fmt.Println("Переменная является типом int")
    case string:
        fmt.Println("Переменная является типом string")
    case bool:
        fmt.Println("Переменная является типом bool")
    case chan int:
        fmt.Println("Переменная является типом channel для int")
    default:
        fmt.Println("Неизвестный тип")
    }
}

func main() {
    var a int = 42
    var b string = "Hello"
    var c bool = true
    var d chan int = make(chan int)

    // Передаем переменные разных типов в функцию для определения типа
    checkType(a)
    checkType(b)
    checkType(c)
    checkType(d)
}
