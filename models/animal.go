package models

type Animal struct {
	ID     string `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Gender string `json:"gender,omitempty"`
	// TODO: Добавить поля
}

type AddAnimalRequest struct {
	Name   string `json:"name"`
	Gender string `json:"gender"`
}

type AddAnimalResponse struct {
	Error       bool   `json:"error"`
	Description string `json:"description"`
	Animal      Animal `json:"animal"`
}
