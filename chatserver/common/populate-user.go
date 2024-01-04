package common

import (
	"github.com/dolencd/go-playground/chatserver/users"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func PopulateUserMiddleware(ur *users.UserRepo) gin.HandlerFunc {

	return func(c *gin.Context) {
		// Obviously far from secure... It'd normally be a JWT or session token
		id := c.Request.Header.Get("USER_ID")
		if len(id) != 0 {
			_, err := uuid.Parse(id)
			if err == nil {
				user, ok := ur.GetUser(id)
				if ok {
					c.Set("user", user)
				}
			}
		}

		c.Next()
	}
}
