package main

import (
	"fmt"
	"time"
)

// Sleep - функция, которая приостанавливает выполнение программы на заданное количество миллисекунд.
func Sleep(milliseconds int) {
	time.Sleep(time.Duration(milliseconds) * time.Millisecond)
}

func main() {
	fmt.Println("Начало выполнения программы.")
	Sleep(2000) // Приостановка выполнения на 2 секунды
	fmt.Println("Программа продолжает выполнение после паузы.")
}
