package main

import (
	"encoding/json"
	"fmt"
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
	data := []byte(`{
		"ID": 134,
		"Number": "ИЛМ-1274",
		"Year": 2,
		"Students": [
			{
				"LastName": "Вещий",
				"FirstName": "Лифон",
				"MiddleName": "Вениаминович",
				"Birthday": "4 апреля 1970 года",
				"Address": "632432, г.Тобольск, ул.Киевская, дом 6, квартира 23",
				"Phone": "+7(948)709-47-24",
				"Rating": [1, 2, 3]
			},
			{
				"LastName": "Ien",
				"FirstName": "ccc",
				"MiddleName": "Вениаминович",
				"Birthday": "4 апреля 1970 года",
				"Address": "632432, г.Тобольск, ул.Киевская, дом 6, квартира 23",
				"Phone": "+7(948)709-47-24",
				"Rating": [5, 2]
			}
		]
	}`)

	var info BaseInformation
	err := json.Unmarshal(data, &info)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	fmt.Printf("Результат: %+v\n", info)
	fmt.Printf("Содержимое: %s", data)
}
