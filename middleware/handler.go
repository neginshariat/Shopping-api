package middleware

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"shopping-app/models"
	"shopping-app/repository"
	"strconv"

	"github.com/gorilla/mux"
)

type controller struct {
	sr repository.Repository
}

func NewController() *controller {
	return &controller{}
}

type response struct {
	ID      int64  `json:"id,omitempty"`
	Message string `json:"message,omitempty"`
}

func (c *controller) CreateNewOrder(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	var order models.Order

	err := json.NewDecoder(req.Body).Decode(&order)
	if err != nil {
		fmt.Printf("Unable to decode the request body. %v\n", err)
	}
	insertId := c.sr.CreateOrder(order)

	res := response{
		ID:      insertId,
		Message: "User created successfully",
	}
	json.NewEncoder(w).Encode(res)
}

func (c *controller) ShowAllInStock(w http.ResponseWriter, req *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	weather, err := c.sr.ShowAllStock()
	if err != nil {
		log.Fatalf("Unable to get all wethear temp.%v", err)
	}

	json.NewEncoder(w).Encode(weather)
}

func (c *controller) ShowTheOrderById(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	params := mux.Vars(req)

	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatalf("Unable to convert the string to int. %v", err)
	}
	weather, err := c.sr.ShowOrdeById(int64(id))
	if err != nil {
		log.Fatalf("Unable to get weather. %v", err)
	}

	json.NewEncoder(w).Encode(weather)
}
func (c *controller) UpdateOrder(w http.ResponseWriter, r *http.Request) {
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

	updatedRows := c.sr.EditOrder(int64(id), order)

	msg := fmt.Sprintf(" User updated successfully . Total rows/record affected %v.", updatedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)
}

func (c *controller) DeleteTheOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Context-Type", "application/x-www-form-urlencoded")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "DELETE")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		log.Fatal("Unable to convert string to int", err)
	}

	deletedRows := c.sr.DeleteOrder(int64(id))

	msg := fmt.Sprintf("User updated successfully %d\n", deletedRows)

	res := response{
		ID:      int64(id),
		Message: msg,
	}

	json.NewEncoder(w).Encode(res)

}
