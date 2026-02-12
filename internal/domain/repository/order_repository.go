package repository

import "pizza-tracker-go/internal/domain/entity"

type OrderRepository interface {
	Create(order *entity.Order) (*entity.Order, error)
	GetByID(id string) (*entity.Order, error)
	GetAll() ([]entity.Order, error)
	UpdateStatus(id, status string) error
	Delete(id string) error
}
