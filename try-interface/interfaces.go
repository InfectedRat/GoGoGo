package main

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

func 