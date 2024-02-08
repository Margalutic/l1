package main

import (
	"fmt"
	"math"
)

// Point - структура, представляющая точку в двумерном пространстве.
type Point struct {
	x float64 // Координата X
	y float64 // Координата Y
}

// NewPoint - конструктор для создания новой точки.
func NewPoint(x, y float64) *Point {
	return &Point{x: x, y: y}
}

// Distance - функция для вычисления расстояния между двумя точками.
func Distance(p1, p2 *Point) float64 {
	dx := p2.x - p1.x
	dy := p2.y - p1.y
	return math.Sqrt(dx*dx + dy*dy)
}

func main() {
	// Создаем две точки
	point1 := NewPoint(1, 2)
	point2 := NewPoint(4, 6)

	// Вычисляем расстояние между ними
	distance := Distance(point1, point2)

	// Выводим результат
	fmt.Printf("Расстояние между точками: %.2f\n", distance)
}
