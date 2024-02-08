package main

import (
	"fmt"
)

func main() {
	// Создаем map для хранения групп температурных значений
	temperatureGroups := make(map[int][]float64)

	// Данная последовательность температурных колебаний
	temperatures := []float64{-25.4, -27.0, 13.0, 19.0, 15.5, 24.5, -21.0, 32.5}

	// Проходим по каждой температуре из последовательности
	for _, temp := range temperatures {
		// Определяем номер группы температур по формуле: температура / шаг
		group := int(temp / 10)

		// Добавляем температуру в соответствующую группу в map
		temperatureGroups[group] = append(temperatureGroups[group], temp)
	}

	// Выводим содержимое каждой группы
	for group, temps := range temperatureGroups {
		fmt.Printf("%d градусов:\n", group*10)
		fmt.Println(temps)
	}
}
