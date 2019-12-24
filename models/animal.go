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
	Error       bool    `json:"error,omitempty"`
	Description string  `json:"description,omitempty"`
	Animal      *Animal `json:"animal,omitempty"`
}

type GetAnimalsResponse struct {
	Error       bool      `json:"error"`
	Description string    `json:"description,omitempty"`
	Animals     []*Animal `json:"animals,omitempty"`
}
