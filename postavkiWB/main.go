package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/olekukonko/tablewriter"
)

// CoefficientResponse представляет структуру одного элемента ответа
type CoefficientResponse struct {
	Date          time.Time `json:"date"`
	Coefficient   int       `json:"coefficient"`
	WarehouseID   int       `json:"warehouseID"`
	WarehouseName string    `json:"warehouseName"`
	BoxTypeName   string    `json:"boxTypeName"`
	BoxTypeID     int       `json:"boxTypeID"`
}

func main() {
	// IDs складов, которые нас интересуют
	warehouseIDs := []int{130744, 117986, 507, 120762, 208277, 206348}

	// Преобразуем slice в строку через запятую
	warehouseIDStr := strings.Trim(strings.Replace(fmt.Sprint(warehouseIDs), " ", ",", -1), "[]")

	// URL для запроса с параметром warehouseIDs
	url := fmt.Sprintf("https://supplies-api.wildberries.ru/api/v1/acceptance/coefficients?warehouseIDs=%s", warehouseIDStr)

	// Создаем новый запрос
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatalf("Ошибка создания запроса: %v", err)
	}

	// Добавляем заголовок Authorization
	req.Header.Set("Authorization", "")

	// Выполняем запрос
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Ошибка выполнения запроса: %v", err)
	}
	defer resp.Body.Close()

	// Проверяем успешность запроса
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("Неудачный запрос, код статуса: %d", resp.StatusCode)
	}

	// Декодируем JSON-ответ
	var coefficients []CoefficientResponse
	if err := json.NewDecoder(resp.Body).Decode(&coefficients); err != nil {
		log.Fatalf("Ошибка декодирования JSON: %v", err)
	}

	// Подготавливаем таблицу для вывода
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Date", "Coefficient", "WarehouseID", "WarehouseName", "BoxTypeName", "BoxTypeID"})

	// Заполняем таблицу данными, фильтруя по BoxTypeName = "Короба"
	for _, c := range coefficients {
		if c.BoxTypeName == "Короба" {
			table.Append([]string{
				c.Date.Format("2006-01-02"),
				fmt.Sprintf("%d", c.Coefficient),
				fmt.Sprintf("%d", c.WarehouseID),
				c.WarehouseName,
				c.BoxTypeName,
				fmt.Sprintf("%d", c.BoxTypeID),
			})
		}
	}

	// Выводим таблицу
	table.Render()
}
