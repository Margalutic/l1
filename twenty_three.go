package main

import "fmt"

func main() {
	// Исходный срез
	slice := []int{1, 2, 3, 4, 5}

	// Индекс элемента, который нужно удалить
	index := 2

	// Удаление элемента из среза
	if index >= 0 && index < len(slice) {
		// Удаление элемента путем сдвига всех элементов влево
		copy(slice[index:], slice[index+1:])
		// Уменьшение длины среза на 1
		slice = slice[:len(slice)-1]
		fmt.Println("Элемент удален успешно:", slice)
	} else {
		fmt.Println("Индекс выходит за пределы длины среза")
	}
}
