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

// HelloHandler - dfgdfg
func CreateAnimalHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO:
	//
	var requestBody models.AddAnimalRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)

	if err != nil {
		log.Println(err)

		// Установил статус кода 400
		w.WriteHeader(http.StatusBadRequest)

		fmt.Fprintf(w, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	json.NewEncoder(w).Encode(&requestBody)
}
