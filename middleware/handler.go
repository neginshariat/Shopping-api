package middleware

import (
	"log"
	"shopping-app/models"
	"shopping-app/repository"
)

type service struct {
	repo repository.Repository
}
type ServiceRepository interface {
	ShowAllStockService() ([]models.Store, error)
	ShowAllOrdersService() ([]models.Order, error)
	ShowOrdeByIdService(id int64) (models.Order, error)
	CreateOrderService(order models.Order) int64
	EditOrderService(order models.Order) int64
	DeleteOrderService(id int64) int64
}

func NewServiceRepository(repo repository.Repository) *service {
	return &service{repo: repo}
}
func (sr *service) ShowAllStockService() ([]models.Store, error) {
	services, err := sr.repo.ShowAllStock()
	if err != nil {
		log.Fatalf("Unable to get the methode in service layer. %v", err)
	}
	return services, nil
}
func (sr *service) ShowAllOrdersService() ([]models.Order, error) {
	services, err := sr.repo.ShowAllOrders()
	if err != nil {
		log.Fatalf("Unable to get the methode in service layer. %v", err)
	}
	return services, nil
}
func (sr *service) ShowOrdeByIdService(id int64) (models.Order, error) {
	services, err := sr.repo.ShowOrderById(id)
	if err != nil {
		log.Fatalf("Unable to get the methode in service layer. %v", err)
	}
	return services, nil
}
func (sr *service) CreateOrderService(order models.Order) int64 {
	services := sr.repo.CreateOrder(order)
	return services
}

func (sr *service) EditOrderService(order models.Order) int64 {
	services := sr.repo.EditOrder(order)
	return services
}
func (sr *service) DeleteOrderService(id int64) int64 {
	services := sr.repo.DeleteOrder(id)
	return services
}
