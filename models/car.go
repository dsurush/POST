package models

// Создаем саму структуру машины
type Car struct {
	ID      string `json:"id"`
	Mark    string `json:"mark"`
	Model   string `json:"model"`
	Country string `json:"country"`
}

// "CarResponse создается"
type CarResponse struct {
	Error       bool   `json:"error"`
	Description string `json:"description"`
	Car         *Car   `json:"car"`
}

// "CarRequest создается"
type CarRequest struct {
	Mark string `json:"mark"`
}
