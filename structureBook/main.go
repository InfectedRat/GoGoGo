package main

import (
	"fmt"
)

type Book struct {
	Name    string
	Author  string
	YearPub int
}

func ShowBooks(b []Book) {
	for _, book := range b {
		fmt.Printf("Название: %s Автор: %s Год: %v \n", book.Name, book.Author, book.YearPub)
	}
}

func main() {

	book := []Book{{Name: "Test", Author: "Test", YearPub: 1987},
		{Name: "sdsd", Author: "sdsdsd", YearPub: 1987}}

	ShowBooks(book)
}
