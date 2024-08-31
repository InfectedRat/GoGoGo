package main

import "fmt"

// Определяем интерфейс Animal
type Animal interface {
	Speak() string
}

// Структура Dog реализует интерфейс Animal
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

// Структура Cat реализует интерфейс Animal
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow!"
}

// Функция, которая принимает интерфейс Animal
func MakeSound(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{Name: "Rex"}
	cat := Cat{Name: "Whiskers"}

	// Вызываем MakeSound для Dog и Cat
	MakeSound(dog)
	MakeSound(cat)
}
