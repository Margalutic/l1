package main

import (
    "fmt"
    "os"
    "os/signal"
    "sync"
    "syscall"
)

func main() {
    // Создаем канал для передачи данных
    dataChan := make(chan int)

    // Количество воркеров
    numWorkers := 5

    // Создаем WaitGroup для ожидания завершения всех воркеров
    var wg sync.WaitGroup
    wg.Add(numWorkers)

    // Функция для чтения данных из канала и вывода их
    worker := func() {
        defer wg.Done() // Уменьшаем счетчик WaitGroup при завершении воркера
        for data := range dataChan {
            fmt.Println("Worker:", data)
        }
    }

    // Запускаем воркеров
    for i := 0; i < numWorkers; i++ {
        go worker()
    }

    // Канал для обработки сигналов завершения программы
    quitChan := make(chan os.Signal, 1)
    signal.Notify(quitChan, os.Interrupt, syscall.SIGTERM)

    // Ожидаем сигнал завершения программы
    <-quitChan

    // Закрываем канал с данными
    close(dataChan)

    // Ожидаем завершения всех воркеров
    wg.Wait()

    fmt.Println("Программа завершена")
}
