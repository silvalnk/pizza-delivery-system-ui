package gorm

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"pizza-tracker-go/internal/domain/entity"
	"pizza-tracker-go/internal/domain/repository"
)

type UserRepository struct {
	db *gormDB
}

func NewUserRepository(db *gormDB) repository.UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) Authenticate(username, password string) (*entity.User, error) {
	var m UserModel
	if err := r.db.DB.Where("username = ?", username).First(&m).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid credentials")
		}
		return nil, err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(m.Password), []byte(password)); err != nil {
		return nil, errors.New("invalid credentials")
	}
	return &entity.User{
		ID:       m.ID,
		Username: m.Username,
		Password: "",
	}, nil
}

func (r *UserRepository) GetByID(id string) (*entity.User, error) {
	var m UserModel
	if err := r.db.DB.First(&m, "id = ?", id).Error; err != nil {
		return nil, err
	}
	return &entity.User{
		ID:       m.ID,
		Username: m.Username,
		Password: "",
	}, nil
}
