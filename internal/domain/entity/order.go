package entity

import "time"

var (
	OrderStatuses = []string{"Pedido realizado", "Preparando", "Assando", "Controle de qualidade", "Pronto"}

	PizzaTypes = []string{
		"Margherita",
		"Pepperoni",
		"Vegetariana",
		"Havaiana",
		"Frango com BBQ",
		"Meat Lovers",
		"Frango Buffalo",
		"Suprema",
		"Cogumelo com Trufa",
		"Quatro Queijos",
	}

	PizzaSizes = []string{
		"Pequena", "Média", "Grande", "Extra Grande",
	}
)

type Order struct {
	ID           string      `json:"id"`
	Status       string      `json:"status"`
	CustomerName string      `json:"customerName"`
	Phone        string      `json:"phone"`
	Address      string      `json:"address"`
	Items        []OrderItem `json:"pizzas"`
	CreatedAt    time.Time   `json:"createdAt"`
}

type OrderItem struct {
	ID           string `json:"id"`
	OrderID      string `json:"orderId"`
	Size         string `json:"size"`
	Pizza        string `json:"pizza"`
	Instructions string `json:"instructions"`
}
