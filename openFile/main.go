package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
)

type Student struct {
	LastName   string `json:"LastName"`
	FirstName  string `json:"FirstName"`
	MiddleName string `json:"MiddleName"`
	Birthday   string `json:"Birthday"`
	Address    string `json:"Address"`
	Phone      string `json:"Phone"`
	Rating     []int  `json:"Rating"`
}

type BaseInformation struct {
	ID       int       `json:"ID"`
	Number   string    `json:"Number"`
	Year     int       `json:"Year"`
	Students []Student `json:"Students"`
}

func main() {
	data, err := os.Open("text.json")
	if err != nil {
		log.Fatalf("Невозможно открыть файл: %v", err)
	}

	dataFile, err := io.ReadAll(data)
	if err != nil {
		log.Fatalf("Ошибка при чтении файла: %v", err)
	}

	var student BaseInformation
	err = json.Unmarshal(dataFile, &student)
	if err != nil {
		log.Fatalf("Ошибка при разборе JSON: %v", err)
	}
	fmt.Print(student)
}
