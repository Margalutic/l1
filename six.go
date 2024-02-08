package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Println("Горутина: Начало работы")
    time.Sleep(2 * time.Second)
    fmt.Println("Горутина: Завершение работы")
    done <- true
}

func main() {
    done := make(chan bool, 1)
    go worker(done)

    <-done
    fmt.Println("Главная горутина: Программа завершена")
}
