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

//func GetAnimalsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
//	fmt.Fprintf()
//}

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
