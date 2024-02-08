package main

import (
	"fmt"
)

// Shape интерфейс определяет метод для вычисления площади фигуры.
type Shape interface {
	Area() float64
}

// Rectangle структура представляет прямоугольник.
type Rectangle struct {
	Width  float64
	Height float64
}

// Area вычисляет площадь прямоугольника.
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle структура представляет круг.
type Circle struct {
	Radius float64
}

// Area вычисляет площадь круга.
func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

// Triangle структура представляет треугольник.
type Triangle struct {
	Base   float64
	Height float64
}

// GetArea вычисляет площадь треугольника (не реализует интерфейс Shape).
func (t Triangle) GetArea() float64 {
	return 0.5 * t.Base * t.Height
}

// TriangleAdapter - адаптер для преобразования Triangle к интерфейсу Shape.
type TriangleAdapter struct {
	Triangle
}

// Area вызывает метод GetArea структуры Triangle и преобразует его к интерфейсу Shape.
func (ta TriangleAdapter) Area() float64 {
	return ta.GetArea()
}

func main() {
	rect := Rectangle{Width: 5, Height: 10}
	circ := Circle{Radius: 7}
	tri := Triangle{Base: 3, Height: 6}

	shapes := []Shape{rect, circ, TriangleAdapter{tri}}

	for _, shape := range shapes {
		fmt.Printf("Площадь фигуры: %.2f\n", shape.Area())
	}
}
