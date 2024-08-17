package structure

// Структура для страны
type Country struct {
	AlfaTwo   string `json:"alfaTwo"`
	AlfaThree string `json:"alfaThree"`
	Name      string `json:"name"`
	NameBrief string `json:"nameBrief"`
}

// Структура для ответа
type ResponseCountry struct {
	Countries []Country `json:"countries"`
}

type Accounts struct {
	ID          string `json:"id"`
	Type        string `json:"type"`
	Name        string `json:"name"`
	Status      string `json:"status"`
	OpenedDate  string `json:"openedDate"`
	AccessLevel string `json:"accessLevel"`
}

type ResponseAccounts struct {
	Accounts []Accounts `json:"accounts"`
}
