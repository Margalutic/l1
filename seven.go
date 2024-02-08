package main

import (
	"fmt"
	"sync"
)

func main() {
	// Создаем мапу для хранения данных
	data := make(map[string]int)

	// Создаем мьютекс для синхронизации доступа к map
	var mutex sync.Mutex

	// Создаем WaitGroup для ожидания завершения всех горутин
	var wg sync.WaitGroup

	// Количество горутин для записи в map
	numGoroutines := 5

	// Запускаем горутины для записи данных в map
	for i := 0; i < numGoroutines; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// Блокируем мьютекс перед доступом к map
			mutex.Lock()
			defer mutex.Unlock()

			// Записываем данные в map
			key := fmt.Sprintf("Key%d", id)
			data[key] = id
			fmt.Printf("Горутина %d записала данные в map\n", id)
		}(i)
	}

	// Ждем завершения всех горутин
	wg.Wait()

	// Выводим содержимое map после записи
	fmt.Println("Содержимое map после записи:")
	for key, value := range data {
		fmt.Printf("%s: %d\n", key, value)
	}
}
