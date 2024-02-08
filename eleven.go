package main

import (
	"fmt"
)

// Функция для нахождения пересечения двух множеств
func intersection(set1, set2 []int) []int {
	// Создаем map для хранения элементов первого множества
	set1Map := make(map[int]bool)
	for _, num := range set1 {
		set1Map[num] = true
	}

	// Создаем map для хранения элементов второго множества
	set2Map := make(map[int]bool)
	for _, num := range set2 {
		set2Map[num] = true
	}

	// Создаем пустой срез для хранения пересечения множеств
	result := []int{}

	// Перебираем элементы первого множества и проверяем их наличие во втором множестве
	for num := range set1Map {
		if set2Map[num] {
			result = append(result, num) // Если элемент присутствует в обоих множествах, добавляем его в результат
		}
	}

	return result
}

func main() {
	// Пример неупорядоченных множеств
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}

	// Вызываем функцию для нахождения пересечения множеств
	intersect := intersection(set1, set2)

	// Выводим результат
	fmt.Println("Пересечение множеств:", intersect)
}
