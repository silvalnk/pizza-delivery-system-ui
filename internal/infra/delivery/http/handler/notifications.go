package handler

import (
	"io"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *Handler) NotificationHandler(c *gin.Context) {
	orderID := c.Query("orderId")
	if orderID == "" {
		c.String(http.StatusBadRequest, "Invalid orderId")
		return
	}

	_, err := h.OrderService.GetOrder(orderID)
	if err != nil {
		c.String(http.StatusNotFound, "Order not found")
		return
	}

	key := "order:" + orderID
	client := make(chan string, 10)
	h.Notifier.AddClient(key, client)

	defer func() {
		h.Notifier.RemoveClient(key, client)
		slog.Info("Customer client disconnected", "orderId", orderID)
	}()

	h.streamSSE(c, client)
}

func (h *Handler) AdminNotificationHandler(c *gin.Context) {
	key := "admin:new_orders"
	client := make(chan string, 10)
	h.Notifier.AddClient(key, client)

	defer func() {
		h.Notifier.RemoveClient(key, client)
		slog.Info("Admin client disconnected")
	}()

	h.streamSSE(c, client)
}

func (h *Handler) streamSSE(c *gin.Context, client chan string) {
	c.Header("Content-Type", "text/event-stream")
	c.Header("Cache-Control", "no-cache")
	c.Header("Connection", "keep-alive")

	c.Stream(func(w io.Writer) bool {
		if msg, ok := <-client; ok {
			c.SSEvent("message", msg)
			return true
		}
		return false
	})
}
