package router

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"

	"pizza-tracker-go/internal/infra/delivery/http/handler"
	"pizza-tracker-go/internal/infra/delivery/http/middleware"
)

func Setup(r *gin.Engine, h *handler.Handler, sessionStore sessions.Store) {
	r.Use(sessions.Sessions("pizza-tracker", sessionStore))

	r.GET("/", h.ServeNewOrderForm)
	r.POST("/new-order", h.HandleNewOrderPost)
	r.GET("/customer/:id", h.ServeCustomer)
	r.GET("/notifications", h.NotificationHandler)

	r.GET("/login", h.HandleLoginGet)
	r.POST("/login", h.HandleLoginPost)
	r.POST("/logout", h.HandleLogout)

	admin := r.Group("/admin")
	admin.Use(middleware.Auth(h))
	{
		admin.GET("", h.ServeAdminDashboard)
		admin.POST("/order/:id/update", h.HandleOrderPut)
		admin.POST("/order/:id/delete", h.HandleOrderDelete)
		admin.GET("/notifications", h.AdminNotificationHandler)
	}

	r.Static("/static", "./internal/infra/delivery/http/templates/static")
}
