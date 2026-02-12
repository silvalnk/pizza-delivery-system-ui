package session

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Set(c *gin.Context, key string, value interface{}) error {
	s := sessions.Default(c)
	s.Set(key, value)
	return s.Save()
}

func GetString(c *gin.Context, key string) string {
	val := sessions.Default(c).Get(key)
	if val == nil {
		return ""
	}
	s, _ := val.(string)
	return s
}

func Clear(c *gin.Context) error {
	s := sessions.Default(c)
	s.Clear()
	return s.Save()
}
