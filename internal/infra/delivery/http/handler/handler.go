package handler

import (
	"pizza-tracker-go/internal/app/notification"
	"pizza-tracker-go/internal/app/usecase"
)

type Handler struct {
	OrderService *usecase.OrderService
	AuthService  *usecase.AuthService
	Notifier     notification.Manager
}

func NewHandler(
	orderService *usecase.OrderService,
	authService *usecase.AuthService,
	notifier notification.Manager,
) *Handler {
	return &Handler{
		OrderService: orderService,
		AuthService:  authService,
		Notifier:     notifier,
	}
}
