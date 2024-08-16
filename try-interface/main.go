package main

import "fmt"

// Определение интерфейса Animal
type Animal interface {
	Speak() string
	Move() string
}

// Тип Dog реализует интерфейс Animal
type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof!"
}

func (d Dog) Move() string {
	return "Run"
}

// Тип Cat реализует интерфейс Animal
type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow"
}

func (c Cat) Move() string {
	return "Jump"
}

// Функция, принимающая интерфейс Animal
func PrintAnimalActions(a Animal) {
	fmt.Println("Animal says:", a.Speak())
	fmt.Println("Animal moves by:", a.Move())
}

func main() {
	var myDog Animal = Dog{Name: "Buddy"}
	var myCat Animal = Cat{Name: "Whiskers"}

	// Работа с типами через интерфейс Animal
	PrintAnimalActions(myDog)
	PrintAnimalActions(myCat)
}
