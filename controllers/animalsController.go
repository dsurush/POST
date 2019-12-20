package controllers

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// HelloHandler - dfgdfg
func CreateAnimalHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// TODO:
	//

	fmt.Fprintf(w, "Hello!")
}
