package main

import (
	"fmt"
	"sort"
)

func main() {
	// Отсортированный массив
	arr := []int{2, 4, 6, 8, 10, 12, 14, 16, 18, 20}

	// Ищем индекс элемента 12 в отсортированном массиве
	index := sort.Search(len(arr), func(i int) bool {
		return arr[i] >= 12
	})

	// Проверяем, найден ли элемент и выводим его индекс
	if index < len(arr) && arr[index] == 12 {
		fmt.Println("Элемент 12 найден в позиции", index)
	} else {
		fmt.Println("Элемент 12 не найден")
	}
}
