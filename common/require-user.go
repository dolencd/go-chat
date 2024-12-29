package common

import (
	"fmt"
	"net/http"

	"github.com/dolencd/go-playground/chatserver/users"
	"github.com/gin-gonic/gin"
)

func RequireUserMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, ok := c.Value("user").(users.User)
		if !ok {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		fmt.Printf("user: %v\n", user)
		c.Next()
	}
}
