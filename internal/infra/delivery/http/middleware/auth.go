package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pizza-tracker-go/internal/infra/delivery/http/handler"
	"pizza-tracker-go/internal/infra/delivery/http/session"
)

func Auth(h *handler.Handler) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := session.GetString(c, "userID")
		if userID == "" {
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		_, err := h.AuthService.GetUserByID(userID)
		if err != nil {
			_ = session.Clear(c)
			c.Redirect(http.StatusSeeOther, "/login")
			c.Abort()
			return
		}

		c.Next()
	}
}
