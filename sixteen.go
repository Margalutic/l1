package main

import (
	"fmt"
	"sort"
)

// Функция для быстрой сортировки массива
func quickSort(arr []int) {
	if len(arr) < 2 {
		return
	}

	left, right := 0, len(arr)-1

	// Выбираем опорный элемент
	pivot := arr[right]

	// Разделение массива на подмассивы, в которых элементы слева меньше опорного,
	// а элементы справа больше опорного
	for i := range arr {
		if arr[i] < pivot {
			arr[i], arr[left] = arr[left], arr[i]
			left++
		}
	}

	// Помещаем опорный элемент на правильную позицию
	arr[left], arr[right] = arr[right], arr[left]

	// Рекурсивно сортируем подмассивы слева и справа от опорного элемента
	quickSort(arr[:left])
	quickSort(arr[left+1:])
}

func main() {
	// Исходный массив
	arr := []int{3, 1, 4, 1, 5, 9, 2, 6, 5}

	// Встроенная функция для быстрой сортировки с использованием пользовательской функции сравнения
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	// Вывод отсортированного массива
	fmt.Println("Сортировка с использованием встроенных методов:")
	fmt.Println(arr)

	// Вызываем функцию быстрой сортировки
	quickSort(arr)

	// Вывод отсортированного массива
	fmt.Println("Быстрая сортировка:")
	fmt.Println(arr)
}
