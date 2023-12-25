package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func AssignRoutes(r *gin.RouterGroup) {
	r.GET("/users", handleGetUsers)
	r.GET("/users/:id", handleGetUserById)
	r.POST("/users", handleCreateUser)
	r.PUT("/users/:id", handleUpdateUser)
	r.DELETE("/users/:id", handleDeleteUser)
}

func handleGetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, GetUsers())
}

func handleGetUserById(c *gin.Context) {
	id := c.Param("id")

	user, isFound := GetUser(id)
	if !isFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

func handleCreateUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newId, _ := uuid.NewV7()
	user.Id = newId.String()
	users[user.Id] = user

	c.JSON(http.StatusCreated, user)
}

func handleUpdateUser(c *gin.Context) {
	id := c.Param("id")

	_, ok := users[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	var updatedUser User

	if err := c.BindJSON(&updatedUser); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser.Id = id

	users[id] = updatedUser

	c.JSON(http.StatusOK, updatedUser)
}

func handleDeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, ok := users[id]
	if !ok {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	delete(users, id)

	c.Status(http.StatusNoContent)
}
