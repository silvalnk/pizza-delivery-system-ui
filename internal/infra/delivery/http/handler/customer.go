package handler

import (
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"pizza-tracker-go/internal/domain/entity"
	"pizza-tracker-go/internal/infra/delivery/http/viewmodel"
)

func (h *Handler) ServeNewOrderForm(c *gin.Context) {
	c.HTML(http.StatusOK, "order.tmpl", viewmodel.OrderFormViewModel{
		PizzaTypes: entity.PizzaTypes,
		PizzaSizes: entity.PizzaSizes,
	})
}

func (h *Handler) HandleNewOrderPost(c *gin.Context) {
	var form viewmodel.OrderRequest

	if err := c.ShouldBind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	items := make([]entity.OrderItem, len(form.Sizes))
	for i := range items {
		instructions := ""
		if i < len(form.Instructions) {
			instructions = form.Instructions[i]
		}
		items[i] = entity.OrderItem{
			Size:         form.Sizes[i],
			Pizza:        form.PizzaTypes[i],
			Instructions: instructions,
		}
	}

	order := &entity.Order{
		CustomerName: form.Name,
		Phone:        form.Phone,
		Address:      form.Address,
		Status:       entity.OrderStatuses[0],
		Items:        items,
	}

	created, err := h.OrderService.CreateOrder(order)
	if err != nil {
		slog.Error("Failed to create order", "error", err)
		c.String(http.StatusInternalServerError, "Something went wrong")
		return
	}

	slog.Info("Order created", "orderId", created.ID, "customer", created.CustomerName)
	h.Notifier.Notify("admin:new_orders", "new_order")

	c.Redirect(http.StatusSeeOther, "/customer/"+created.ID)
}

func (h *Handler) ServeCustomer(c *gin.Context) {
	orderID := c.Param("id")
	if orderID == "" {
		c.String(http.StatusBadRequest, "Order ID is required")
		return
	}

	order, err := h.OrderService.GetOrder(orderID)
	if err != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}

	c.HTML(http.StatusOK, "customer.tmpl", viewmodel.CustomerViewModel{
		Title:    "Pizza Order Status " + orderID,
		Order:    *order,
		Statuses: entity.OrderStatuses,
	})
}
