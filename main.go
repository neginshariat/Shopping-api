package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping-app/config"
	"shopping-app/controller"
	"shopping-app/middleware"
	"shopping-app/repository"
)

func main() {
	config := config.NewConfig()

	dbRepo := repository.StockRepository(config)
	shoppingService := middleware.NewServiceRepository(dbRepo)
	shoppingController := controller.NewController(shoppingService)

	router := shoppingController.Router()

	fmt.Println("starting server on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", router))
}
