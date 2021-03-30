package main

import (
	"fmt"
	"log"
	"net/http"
	"shopping-app/router"
)

func main() {
	/* 	config := config.NewConfig()

	   	dbRepo := repository.StockRepository(config)
	   	shoppingController := middleware.NewController(dbRepo)

	   	router := mux.NewRouter()
	   	router.HandleFunc("/shop", shoppingController.ShowAllInStock).Methods("GET", "OPTIONS") // Get all the stuff
	   	//router.HandleFunc("/shop/{name}", middleware.NewController().Show).Methods("GET", "OPTIONS")              // Get stuff by cathergory
	   	router.HandleFunc("/shop/{id}", shoppingController.ShowTheOrderById).Methods("GET", "OPTIONS") // Get stuff by id
	   	router.HandleFunc("/shop/make", shoppingController.CreateNewOrder).Methods("POST", "OPTIONS")  // Create an order
	   	router.HandleFunc("/shop/edit", shoppingController.UpdateOrder).Methods("PUT", "OPTIONS")      // Edit the ordeth whicht created
	   	router.HandleFunc("/shop/delete", shoppingController.DeleteTheOrder).Methods("DELETE", "OPTIONS") */
	r := router.Router()
	fmt.Println("starting server on port 8080...")

	log.Fatal(http.ListenAndServe(":8080", r))
}
