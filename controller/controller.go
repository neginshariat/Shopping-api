package controller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shopping-app/middleware"
	"shopping-app/models"
	"strconv"

	"github.com/gorilla/mux"
)

func (c *Controller) Router() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/shop", c.ShowAllInStock).Methods("GET", "OPTIONS") // Get all the stuff
	router.HandleFunc("/shop/order", c.ShowAllOrders).Methods("GET", "OPTIONS")
	router.HandleFunc("/shop/{id}", c.ShowTheOrderById).Methods("GET", "OPTIONS")    // Get stuff by id
	router.HandleFunc("/shop/make", c.CreateNewOrder).Methods("POST", "OPTIONS")     // Create an order
	router.HandleFunc("/shop/edit", c.UpdateOrder).Methods("PUT", "OPTIONS")         // Edit the ordeth whicht created
	router.HandleFunc("/shop/delete", c.DeleteTheOrder).Methods("DELETE", "OPTIONS") // Delete the orde which created

	return router
}

type Controller struct {
	cs middleware.ServiceRepository
}

func NewController(cs middleware.ServiceRepository) *Controller {
	return &Controller{cs: cs}
}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func (c *Controller) CreateNewOrder(w http.ResponseWriter, req *http.Request) {
	fmt.Println(req.Body)
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var order models.Order

	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		fmt.Printf("Unable to decode the request body. %v\n", err)
	}
	insertId := c.cs.CreateOrderService(order)

	res := response{
		ID:      insertId,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)
}
func (c *Controller) ShowAllOrders(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	weather, err := c.cs.ShowAllOrdersService()
	if err != nil {
		log.Fatalf("Unable to get all wethear temp.%v", err)
	}

	json.NewEncoder(w).Encode(weather)
}

func (c *Controller) ShowAllInStock(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	weather, err := c.cs.ShowAllStockService()
	if err != nil {
		log.Fatalf("Unable to get all wethear temp.%v", err)
	}

	json.NewEncoder(w).Encode(weather)
}

func (c *Controller) ShowTheOrderById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string to int. %v", err)
	}
	weather, err := c.cs.ShowOrdeByIdService(int64(id))
	if err != nil {
		log.Fatalf("Unable to get weather. %v", err)
	}

	json.NewEncoder(w).Encode(weather)
}
func (c *Controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "PUT")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		fmt.Sprintf("Unable to convert string to int. %v", err)
	}

	var order models.Order

	err = json.NewDecoder(r.Body).Decode(&order)
	if err != nil {
		log.Fatal("unale to decode the request body", err)
	}

	updatedRows := c.cs.EditOrderService(order)

	msg := fmt.Sprintf(" User updated successfully . Total rows/record affected %v.", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func (c *Controller) DeleteTheOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert string to int", err)
	}

	deletedRows := c.cs.DeleteOrderService(int64(id))

	msg := fmt.Sprintf("User updated successfully %d\n", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}
