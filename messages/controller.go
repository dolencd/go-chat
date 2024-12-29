package messages

import (
	"log"
	"net/http"

	"github.com/dolencd/go-playground/chatserver/users"
	"github.com/gin-gonic/gin"
)

type MessageController struct {
	rr *MessageRepo
}

func NewMessageController(router *gin.RouterGroup, rr *MessageRepo) MessageController {

	rc := MessageController{rr: rr}

	// Basic CRUD
	router.GET("/messages", rc.HandleGetMessages)
	router.GET("/messages/:id", rc.HandleGetMessageById)
	router.POST("/messages", rc.HandleCreateMessage)

	return rc
}

func (rc *MessageController) HandleGetMessages(c *gin.Context) {
	messages, err := rc.rr.GetMessages()
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}
	c.JSON(http.StatusOK, messages)
}

func (rc *MessageController) HandleGetMessageById(c *gin.Context) {
	id := c.Param("id")

	message, isFound := rc.rr.GetMessage(id)
	if !isFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, message)
}

func (rc *MessageController) HandleCreateMessage(c *gin.Context) {
	user := c.MustGet("user").(users.User)
	var message Message

	if err := c.BindJSON(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	message.SenderUserId = user.Id

	createdMessage, err := rc.rr.CreateMessage(message)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}

	c.JSON(http.StatusCreated, createdMessage)
}
