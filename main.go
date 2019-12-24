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
	router.GET("/animals/:id", controllers.GetAnimalHandler)
	router.POST("/animals", controllers.CreateAnimalHandler)
	router.GET("/animals", controllers.GetAnimalsHandler)
	//////TEST////////
	router.GET("/testanimals/:id", controllers.GettestaAnimalHandler)
	router.POST("/cars", controllers.CreateCarsHandler)

	http.ListenAndServe(":8080", router)
}
