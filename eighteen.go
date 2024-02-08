package main

import (
	"fmt"
	"sync"
)

// Counter - структура счетчика
type Counter struct {
	mu    sync.Mutex // Мьютекс для защиты доступа к счетчику
	value int        // Значение счетчика
}

// Increment - метод для инкрементирования счетчика на 1
func (c *Counter) Increment() {
	c.mu.Lock()         // Блокируем доступ к счетчику для других горутин
	defer c.mu.Unlock() // После выполнения функции снимаем блокировку

	c.value++ // Инкрементируем значение счетчика
}

func main() {
	var wg sync.WaitGroup // Создаем WaitGroup для ожидания завершения всех горутин
	counter := Counter{}   // Создаем экземпляр счетчика

	// Запускаем 1000 горутин, каждая из которых инкрементирует счетчик 1000 раз
	for i := 0; i < 1000; i++ {
		wg.Add(1) // Добавляем в WaitGroup каждую горутину

		go func() {
			defer wg.Done() // Убираем один элемент из WaitGroup при завершении горутины
			counter.Increment()
		}()
	}

	wg.Wait() // Ждем завершения всех горутин

	// Выводим итоговое значение счетчика
	fmt.Println("Итоговое значение счетчика:", counter.value)
}
