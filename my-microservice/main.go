package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// Определение структуры
type Note struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func main() {
	// Подключение к базе данных
	databasePath := "/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db"
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Выполнение запроса
	rows, err := db.Query("SELECT id, name FROM note_test")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Слайс для хранения результатов
	var notes []Note

	// Обработка результата
	for rows.Next() {
		var note Note
		err := rows.Scan(&note.ID, &note.Name)
		if err != nil {
			log.Fatal(err)
		}
		notes = append(notes, note)
	}

	// Проверка на ошибки после итерации
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	// Конвертация в JSON
	notesJSON, err := json.Marshal(notes)
	if err != nil {
		log.Fatal(err)
	}

	// Вывод результата
	fmt.Println(string(notesJSON))
}
