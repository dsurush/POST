package controllers

import (
	"adam/models"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// Научились использовать параметры
func GetAnimalHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {

	fmt.Fprintf(w, p.ByName("id"))
}

//TO DO
func GetAnimalsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	res := models.GetAnimalsResponse{}

	res.Animals, _ = getAnimals()

	json.NewEncoder(w).Encode(res)

}

// HelloHandler - dfgdfg
func CreateAnimalHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	// TODO:
	//
	var responseBody models.AddAnimalResponse

	var requestBody models.AddAnimalRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		responseBody.Error = true
		responseBody.Description = err.Error()

		log.Println(err)

		// Установил статус кода 400
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(&responseBody)
		return
	}

	// TODO:
	// 1. Сохранение в бд
	// 2. Возвращение сохраненных данных
	responseBody.Animal = &models.Animal{
		ID:     "21312344",
		Name:   requestBody.Name,
		Gender: requestBody.Gender,
	}

	json.NewEncoder(w).Encode(&responseBody)
}

//
func getAnimals() ([]*models.Animal, error) {
	animals := []*models.Animal{
		&models.Animal{
			ID:     "2323",
			Name:   "Vasya",
			Gender: "Khan",
		},
		&models.Animal{
			ID:     "34r5w34rw34",
			Name:   "Anatoliy",
			Gender: "Vasserman",
		},
	}

	return animals, nil
}

// Повторяем вчерашнюю тему
func GettestaAnimalHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	fmt.Fprintf(w, p.ByName("id"))
}

//Creat new Animal
