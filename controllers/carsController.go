package controllers

import (
	"adam/models"
	"encoding/json"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

//"CreateCarsHandler создается, для подачи логики"
func CreateCarsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Content-Type", "application/json;charset=utf-8")
	var responseBody models.CarResponse

	var requestBody models.CarRequest
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

	responseBody.Car = &models.Car{
		ID:      "1",
		Mark:    requestBody.Mark,
		Model:   "123",
		Country: "Germany",
	}

	json.NewEncoder(w).Encode(&responseBody)
}
