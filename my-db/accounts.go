package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func GetAccounts() (*Response, error) {
	// URL для запроса
	url := "https://sandbox-invest-public-api.tinkoff.ru/rest/tinkoff.public.invest.api.contract.v1.UsersService/GetAccounts"

	// Токен для авторизации
	token := ""

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
