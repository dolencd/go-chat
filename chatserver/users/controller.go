package users

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	ur *UserRepo
}

func NewUserController(router *gin.RouterGroup, ur *UserRepo) UserController {

	uc := UserController{ur: ur}

	router.GET("/users", uc.HandleGetUsers)
	router.GET("/users/:id", uc.HandleGetUserById)
	router.POST("/users", uc.HandleCreateUser)
	router.PUT("/users/:id", uc.HandleUpdateUser)
	router.DELETE("/users/:id", uc.HandleDeleteUser)

	return uc
}

func (uc *UserController) HandleGetUsers(c *gin.Context) {
	users, err := uc.ur.GetUsers()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
	}
	c.JSON(http.StatusOK, users)
}

func (uc *UserController) HandleGetUserById(c *gin.Context) {
	id := c.Param("id")

	user, isFound := uc.ur.GetUser(id)
	if !isFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *UserController) HandleCreateUser(c *gin.Context) {
	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUser, err := uc.ur.CreateUser(user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusCreated, createdUser)
}

func (uc *UserController) HandleUpdateUser(c *gin.Context) {
	id := c.Param("id")

	var user User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUser, err := uc.ur.UpdateUser(id, user)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, updatedUser)
}

func (uc *UserController) HandleDeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := uc.ur.DeleteUser(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.Status(http.StatusNoContent)
}
