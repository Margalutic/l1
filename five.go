package main

import (
    "fmt"
    "time"
)

func main() {
    // Создаем канал для передачи данных
    dataChan := make(chan int)

    // Запускаем горутину для отправки данных в канал
    go sendData(dataChan)

    // Читаем данные из канала
    for {
        select {
        case data := <-dataChan:
            // Выводим полученное значение
            fmt.Println("Прочитано из канала:", data)
        case <-time.After(5 * time.Second): // Ждем 5 секунд
            // После истечения времени завершаем программу
            fmt.Println("Время истекло. Программа завершена.")
            return
        }
    }
}

// Функция для отправки данных в канал
func sendData(ch chan<- int) {
    // Отправляем значения от 1 до 10 в канал
    for i := 1; i <= 10; i++ {
        ch <- i
        // Задержка перед отправкой следующего значения
        time.Sleep(1 * time.Second)
    }
    // Закрываем канал после отправки всех значений
    close(ch)
}
