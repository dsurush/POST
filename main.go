package main

import (
	"adam/controllers"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var (
	router = httprouter.New()
)

func main() {

	//database.Connect()
	router.GET("/test", controllers.HelloHandler)
	router.POST("/animals", controllers.CreateAnimalHandler)
	http.ListenAndServe(":8080", router)
}
