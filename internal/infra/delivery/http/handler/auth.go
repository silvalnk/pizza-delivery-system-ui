package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"pizza-tracker-go/internal/domain/entity"
	"pizza-tracker-go/internal/infra/delivery/http/session"
	"pizza-tracker-go/internal/infra/delivery/http/viewmodel"
)

func (h *Handler) HandleLoginGet(c *gin.Context) {
	c.HTML(http.StatusOK, "login.tmpl", viewmodel.LoginViewModel{})
}

func (h *Handler) HandleLoginPost(c *gin.Context) {
	var form struct {
		Username string `form:"username" binding:"required,min=3,max=50"`
		Password string `form:"password" binding:"required,min=6"`
	}

	if err := c.ShouldBind(&form); err != nil {
		c.HTML(http.StatusOK, "login.tmpl", viewmodel.LoginViewModel{Error: "Invalid input: " + err.Error()})
		return
	}

	user, err := h.AuthService.Authenticate(form.Username, form.Password)
	if err != nil {
		c.HTML(http.StatusOK, "login.tmpl", viewmodel.LoginViewModel{Error: "Invalid credentials"})
		return
	}

	_ = session.Set(c, "userID", fmt.Sprintf("%v", user.ID))
	_ = session.Set(c, "username", user.Username)

	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) HandleLogout(c *gin.Context) {
	if err := session.Clear(c); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/login")
}

func (h *Handler) ServeAdminDashboard(c *gin.Context) {
	orders, err := h.OrderService.GetAllOrders()
	if err != nil {
		c.String(http.StatusInternalServerError, "Error fetching orders")
		return
	}
	username := session.GetString(c, "username")

	c.HTML(http.StatusOK, "admin.tmpl", viewmodel.AdminDashboardViewModel{
		Orders:   orders,
		Statuses: entity.OrderStatuses,
		Username: username,
	})
}

func (h *Handler) HandleOrderPut(c *gin.Context) {
	orderID := c.Param("id")
	newStatus := c.PostForm("status")

	if err := h.OrderService.UpdateOrderStatus(orderID, newStatus); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}

	h.Notifier.Notify("order:"+orderID, "order_updated")
	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *Handler) HandleOrderDelete(c *gin.Context) {
	orderID := c.Param("id")

	if err := h.OrderService.DeleteOrder(orderID); err != nil {
		c.String(http.StatusInternalServerError, err.Error())
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin")
}
