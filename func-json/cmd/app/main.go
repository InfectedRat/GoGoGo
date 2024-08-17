package application

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/mattn/go-sqlite3"

	structure "func-json/internal/models"
)

var Token string

func GetAccounts() (*structure.ResponseAccounts, error) {
	url := "https://sandbox-invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts"

	err := godotenv.Load("/Users/maximbabichev/ProjectGolang/GoGoGo/func-json/internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token := os.Getenv("API_TOKEN")

	requestBody, err := json.Marshal(map[string]interface{}{})
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании тела запроса %v", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, fmt.Errorf("ошибка при создании запроса: %v", err)
	}

	req.Header.Set("Authorization", "Bearer "+Token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("ошибка при получении ответа: %s", err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ошибка при чтении ответа: %v", err)
	}

	var response structure.ResponseAccounts
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе ответа: %v", err)
	}

	return &response, nil
}

// GetCountries выполняет запрос к API для получения списка стран
func GetCountries() (*structure.ResponseCountry, error) {
	url := "https://sandbox-invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.InstrumentsService/GetCountries"

	err := godotenv.Load("/Users/maximbabichev/ProjectGolang/GoGoGo/func-json/internal/config/.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	Token := os.Getenv("API_TOKEN")

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
	req.Header.Set("Authorization", "Bearer "+Token)
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

	// Парсим JSON-ответ
	var response structure.ResponseCountry
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, fmt.Errorf("ошибка при разборе ответа: %v", err)
	}

	return &response, nil
}

// connectDB подключается к базе данных SQLite
func ConnectDB() *sql.DB {
	db, err := sql.Open("sqlite3", "/Users/maximbabichev/Library/DBeaverData/workspace6/.metadata/sample-database-sqlite-1/Chinook.db")
	if err != nil {
		log.Fatalf("Ошибка при подключении к базе данных: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("Ошибка при проверке соединения с БД: %v", err)
	}

	return db
}

// createTable создает таблицу countries, если она еще не создана
func CreateTable(db *sql.DB) {
	querys := map[string]string{"countries": `CREATE TABLE IF NOT EXISTS countries (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		alfaTwo TEXT NOT NULL,
		alfaThree TEXT NOT NULL,
		name TEXT NOT NULL,
		nameBrief TEXT NOT NULL);`,
		"accounts": `CREATE TABLE IF NOT EXISTS accounts (
		id TEXT PRIMARY KEY,
		type TEXT NOT NULL,
		name TEXT NOT NULL,
		status TEXT NOT NULL,
		openedDate TIMESTAMP NOT NULL,
		accessLevel TEXT NOT NULL
	);`}

	for name, query := range querys {
		_, err := db.Exec(query)
		if err != nil {
			log.Fatalf("Ошибка создания запроса: %v", err)
		}
		log.Printf("Таблица %s создана", name)
	}
}

// saveCountries сохраняет данные стран в таблицу countries
func SaveCountries(db *sql.DB, countries []structure.Country) error {
	for _, country := range countries {
		query := `INSERT INTO countries (alfaTwo, alfaThree, name, nameBrief) VALUES (?, ?, ?, ?)`

		_, err := db.Exec(query, country.AlfaTwo, country.AlfaThree, country.Name, country.NameBrief)
		if err != nil {
			return fmt.Errorf("ошибка при сохранении страны %s: %v", country.Name, err)
		}
	}
	return nil
}

func SaveAccounts(db *sql.DB, accounts []structure.Accounts) error {
	for _, account := range accounts {
		query := `INSERT INTO accounts (
		type ,
		name ,
		status ,
		openedDate ,
		accessLevel ) VALUES (?,?,?,?,?)`

		_, err := db.Exec(query, account.Type, account.Name, account.Status, account.OpenedDate, account.AccessLevel)
		if err != nil {
			return fmt.Errorf("ошибка при сохранении аккаунта %s: %v", account.Name, err)
		}
	}
	return nil
}

func RunApp() {

	// Подключаемся к базе данных
	db := ConnectDB()
	defer db.Close()

	// Создаем таблицу countries, если она еще не создана
	CreateTable(db)

	// Получаем список стран от API
	response, err := GetCountries()
	if err != nil {
		log.Fatalf("Ошибка при получении данных стран: %v", err)
	}

	// Сохраняем страны в базу данных
	err = SaveCountries(db, response.Countries)
	if err != nil {
		log.Fatalf("Ошибка при сохранении данных стран: %v", err)
	}

	fmt.Println("Данные стран успешно сохранены в базу данных")

	responseAcc, err := GetAccounts()
	if err != nil {
		log.Fatalf("Ошибка при получении аккаунтов %s", err)
	}

	err = SaveAccounts(db, responseAcc.Accounts)
	if err != nil {
		log.Fatalf("Ошибка при сохранении данных аккаунтов %s", err)
	}

	fmt.Println("Данные аккаунта успешно сохранены")

}
