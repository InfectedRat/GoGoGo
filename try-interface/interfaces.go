package main

import "fmt"

type Animal interface {
	Move()
	Speak()
}

type Cat struct {
	Name string
}

type Dog struct {
	Name string
}

func (c Cat) Move() {
	fmt.Println("Кошка по имени %s прыгает", c.Name)
}

func (c Cat) Speak() {
	fmt.Println("Кошка по имени %s говорит Mewooo", c.Name)
}

func (d Dog) Move() {
	fmt.Println("")
}

func (d Dog) Speak() {
	fmt.Println("")
}
