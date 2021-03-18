package router

import "github.com/gorilla/mux"

func Router() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/shop", controller.Controller).Methods("GET", "OPTIONS")        // Get all the stuff
	router.HandleFunc("/shop/{name}", controller.Controller).Methods("GET", "OPTIONS") // Get stuff by cathergory
	router.HandleFunc("/shop/{id}", controller.Controller).Methods("GET", "OPTIONS")   // Get stuff by id
	//router.HandleFunc("/shop/{model}", controller.Controller).Methods("GET", "OPTIONS")   // Get stuff by model     ==> at the second step to improve my app
	router.HandleFunc("/shop/make", controller.Controller).Methods("POST", "OPTIONS")     // Create an order
	router.HandleFunc("/shop/edit", controller.Controller).Methods("PUT", "OPTIONS")      // Edit the ordeth whicht created
	router.HandleFunc("/shop/delete", controller.Controller).Methods("DELETE", "OPTIONS") // Delete the orde which created

	return router
}
