package gorm

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type gormDB struct {
	DB *gorm.DB
}

// seedDefaultAdmin cria o usuário admin padrão se não existir nenhum usuário.
// Login: admin / admin123 — altere a senha em produção.
func seedDefaultAdmin(db *gorm.DB) error {
	var count int64
	if err := db.Model(&UserModel{}).Count(&count).Error; err != nil {
		return err
	}
	if count > 0 {
		return nil
	}
	hashed, err := bcrypt.GenerateFromPassword([]byte("admin123"), bcrypt.DefaultCost)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}
	return db.Create(&UserModel{
		Username: "admin",
		Password: string(hashed),
	}).Error
}

func NewDB(dsn string) (*gormDB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		PrepareStmt: false, // evita "insufficient arguments" com o driver postgres/pgx
	})
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	// Define UTF8 na conexão (evita "simple protocol queries must be run with client_encoding=UTF8")
	if err := db.Exec("SET client_encoding = 'UTF8'").Error; err != nil {
		return nil, fmt.Errorf("set encoding: %w", err)
	}

	if err := db.AutoMigrate(&OrderModel{}, &OrderItemModel{}, &UserModel{}); err != nil {
		return nil, fmt.Errorf("migrate: %w", err)
	}

	if err := seedDefaultAdmin(db); err != nil {
		return nil, fmt.Errorf("seed admin: %w", err)
	}

	return &gormDB{DB: db}, nil
}
