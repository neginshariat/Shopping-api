package controller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"shopping-app/middleware"
	"shopping-app/models"
	"testing"

	"github.com/stretchr/testify/require"
)

type MockDB struct {
}

func (m MockDB) ShowAllStock() ([]models.Store, error) {
	return []models.Store{{

		Pants:   " blue",
		Shoes:   "white",
		TShirts: "black",
	}}, nil
}
func (m MockDB) ShowAllOrders() ([]models.Order, error) {
	return []models.Order{{
		OrID:    2,
		Pants:   " blue",
		Shoes:   "white",
		TShirts: "black",
	}}, nil
}
func (m MockDB) ShowOrderById(id int64) (models.Order, error) {
	return models.Order{
		OrID:    2,
		Pants:   " blue",
		Shoes:   "white",
		TShirts: "black",
	}, nil
}
func (m MockDB) CreateOrder(models.Order) int64 {
	return 13
}
func (m MockDB) EditOrder(md models.Order) int64 {
	return 13

}
func (m MockDB) DeleteOrder(id int64) int64 {
	return 13
}
func TestCreateNewOrder(t *testing.T) {

	shoppingService := middleware.NewServiceRepository(MockDB{})
	shoppingController := NewController(shoppingService)

	order := models.Order{OrID: 2,
		Pants:   " blue",
		Shoes:   "white",
		TShirts: "black"}

	data, _ := json.Marshal(&order)
	res := httptest.NewRecorder()
	req, err := http.NewRequest("POST", "/shop/make", bytes.NewBuffer(data))
	require.NoError(t, err)

	r := shoppingController.Router(order)
	r.ServeHTTP(res, req)

	fmt.Println(req)

	require.Equal(t, res.Code, http.StatusOK)

}
func TestShowAllInStock(t *testing.T) {

	shoppingService := middleware.NewServiceRepository(MockDB{})
	shoppingController := NewController(shoppingService)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/shop", nil)
	r := shoppingController.Router()
	r.ServeHTTP(res, req)

	require.Equal(t, res.Code, req, http.StatusOK)

}
func TestShowAllOrders(t *testing.T) {
	shoppingService := middleware.NewServiceRepository(MockDB{})
	shoppingController := NewController(shoppingService)
	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/shop/order", nil)

	r := Router(shoppingController)
	r.ServeHTTP(res, req)

	require.Equal(t, res.Code, http.StatusOK)

}
func TestShowTheOrderById(t *testing.T) {
	shoppingService := middleware.NewServiceRepository(MockDB{})
	shoppingController := NewController(shoppingService)

	orderID := 2

	data, _ := json.Marshal(&orderID)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/shop/{id}", bytes.NewBuffer(data))

	r := Router(shoppingController)
	r.ServeHTTP(res, req)

	require.Equal(t, res.Code, http.StatusOK)
}
func TestUpdateOrder(t *testing.T) {
	shoppingService := middleware.NewServiceRepository(MockDB{})
	shoppingController := NewController(shoppingService)

	order := models.Order{OrID: 2,
		Pants:   " blue",
		Shoes:   "white",
		TShirts: "black"}

	data, _ := json.Marshal(&order)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("PUT", "/shop/edit", bytes.NewBuffer(data))

	r := Router(shoppingController)
	r.ServeHTTP(res, req)

	require.Equal(t, res.Code, http.StatusOK)
}
func TestDeleteTheOrder(t *testing.T) {
	shoppingService := middleware.NewServiceRepository(MockDB{})
	shoppingController := NewController(shoppingService)

	orderID := 2
	data, _ := json.Marshal(&orderID)

	res := httptest.NewRecorder()
	req, _ := http.NewRequest("DELETE", "/shop/delete", bytes.NewBuffer(data))

	r := Router(shoppingController)
	r.ServeHTTP(res, req)

	require.Equal(t, res.Code, http.StatusOK)

}
