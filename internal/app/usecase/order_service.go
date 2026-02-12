package usecase

import (
	"pizza-tracker-go/internal/domain/entity"
	"pizza-tracker-go/internal/domain/repository"
)

type OrderService struct {
	orderRepo repository.OrderRepository
}

func NewOrderService(orderRepo repository.OrderRepository) *OrderService {
	return &OrderService{orderRepo: orderRepo}
}

func (s *OrderService) CreateOrder(order *entity.Order) (*entity.Order, error) {
	return s.orderRepo.Create(order)
}

func (s *OrderService) GetOrder(id string) (*entity.Order, error) {
	return s.orderRepo.GetByID(id)
}

func (s *OrderService) GetAllOrders() ([]entity.Order, error) {
	return s.orderRepo.GetAll()
}

func (s *OrderService) UpdateOrderStatus(id, status string) error {
	return s.orderRepo.UpdateStatus(id, status)
}

func (s *OrderService) DeleteOrder(id string) error {
	return s.orderRepo.Delete(id)
}
