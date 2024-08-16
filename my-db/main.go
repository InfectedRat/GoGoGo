package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

// Account представляет структуру данных аккаунта
type Account struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	OpenedDate  string `json:"openedDate"`
	AccessLevel string `json:"accessLevel"`
}

// Response представляет структуру ответа от API
type Response struct {
	Accounts []Account `json:"accounts"`
}

// GetAccounts выполняет запрос к API и возвращает распарсенный JSON-ответ
func GetAccounts() (*Response, error) {
	// URL для запроса
	url := "https://sandbox-invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts"

	// Токен для авторизации
	token := "t.KZN0RKZTqvlYBJxL3fM0EqZbZ6zsIJSSD8H0TGeWnTvcWgsBk0M0fzmX8dW4p_i5GIT7uZclp1TqjXafXtmBOQ"

	// Создаем пустое тело запроса
	requestBody, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании тела запроса: %v", err)
	}

	// Создаем новый POST-запрос
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	// Устанавливаем заголовки
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	// Отправляем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при выполнении запроса: %v", err)
	}
	defer resp.Body.Close()

	// Чтение ответа
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	// Проверка статуса ответа
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("запрос завершился с ошибкой: %s", resp.Status)
	}

	// Парсим JSON-ответ
	var response Response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе ответа: %v", err)
	}

	return &response, nil
}

// SaveAccounts сохраняет данные аккаунтов в базу данных
func SaveAccounts(db *sql.DB, accounts []Account) error {
	for _, account := range accounts {
		query := `INSERT INTO accounts (id, type, name, status, openedDate, accessLevel) 
		          VALUES (?, ?, ?, ?, ?, ?)`

		_, err := db.Exec(query, account.ID, account.Type, account.Name, account.Status, account.OpenedDate, account.AccessLevel)
		if err != nil {
			return fmt.Errorf("ошибка при сохранении аккаунта %s: %v", account.ID, err)
		}
	}
	return nil
}

func main() {
	// Подключаемся к базе данных
	db, err := sql.Open("sqlite3", "/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db")
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}
	defer db.Close()

	// Создаем таблицу accounts (если она еще не создана)
	createAccountsTable(db)

	// Получаем данные аккаунтов от API
	accountsResponse, err := GetAccounts()
	if err != nil {
		log.Fatalf("Ошибка при получении данных аккаунтов: %v", err)
	}

	// Сохраняем данные аккаунтов в базу данных
	err = SaveAccounts(db, accountsResponse.Accounts)
	if err != nil {
		log.Fatalf("Ошибка при сохранении данных аккаунтов: %v", err)
	}

	fmt.Println("Данные аккаунтов успешно сохранены в базу данных")
}

// createAccountsTable создает таблицу для хранения аккаунтов
func createAccountsTable(db *sql.DB) {
	query := `CREATE TABLE IF NOT EXISTS accounts (
		id TEXT PRIMARY KEY,
		type TEXT NOT NULL,
		name TEXT NOT NULL,
		status TEXT NOT NULL,
		openedDate TIMESTAMP NOT NULL,
		accessLevel TEXT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatalf("Ошибка при создании таблицы: %v", err)
	}
}
