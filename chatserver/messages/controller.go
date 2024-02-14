package messages

import (
	"log"
	"net/http"

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
		err := c.AbortWithError(http.StatusInternalServerError, err)
		if err != nil {
			log.Println(err)
		}
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
	var message Message

	if err := c.BindJSON(&message); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdMessage, err := rc.rr.CreateMessage(message)
	if err != nil {
		err := c.AbortWithError(http.StatusInternalServerError, err)
		if err != nil {
			log.Println(err)
		}
		return
	}

	c.JSON(http.StatusCreated, createdMessage)
}
