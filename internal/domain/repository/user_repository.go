package repository

import "pizza-tracker-go/internal/domain/entity"

type UserRepository interface {
	Authenticate(username, password string) (*entity.User, error)
	GetByID(id string) (*entity.User, error)
}
