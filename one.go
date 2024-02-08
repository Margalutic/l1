package main
import (
    "fmt"
)

// Определение структуры Human
type Human struct {
    Name   string
    Age    int
    Gender string
}

// Метод для структуры Human
func (h *Human) Introduce() {
    fmt.Printf("Привет, меня зовут %s, мне %d лет, я %s\n", h.Name, h.Age, h.Gender)
}

// Определение структуры Action с встраиванием типа Human
type Action struct {
    Human // Встраивание типа Human
    ActionName string
}

// Метод для структуры Action
func (a *Action) DoAction() {
    fmt.Printf("%s %s\n", a.Name, a.ActionName)
}

func main() {
    // Создание объекта типа Action
    action := Action{
        Human: Human{
            Name:   "Иван",
            Age:    30,
            Gender: "мужчина",
        },
        ActionName: "готовит завтрак",
    }

    // Вызов метода Introduce внутри метода DoAction
    action.DoAction()
    action.Introduce()
}