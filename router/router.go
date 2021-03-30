package router

import (
	"shopping-app/middleware"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shop", middleware.NewController().ShowAllInStock).Methods("GET", "OPTIONS") // Get all the stuff
	//router.HandleFunc("/shop/{name}", middleware.NewController().Show).Methods("GET", "OPTIONS")              // Get stuff by cathergory
	router.HandleFunc("/shop/{id}", middleware.NewController().ShowTheOrderById).Methods("GET", "OPTIONS")    // Get stuff by id
	router.HandleFunc("/shop/make", middleware.NewController().CreateNewOrder).Methods("POST", "OPTIONS")     // Create an order
	router.HandleFunc("/shop/edit", middleware.NewController().UpdateOrder).Methods("PUT", "OPTIONS")         // Edit the ordeth whicht created
	router.HandleFunc("/shop/delete", middleware.NewController().DeleteTheOrder).Methods("DELETE", "OPTIONS") // Delete the orde which created

	return router
}
