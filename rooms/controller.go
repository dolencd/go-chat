package rooms

import (
	"log"
	"net/http"

	"github.com/dolencd/go-playground/chatserver/users"
	"github.com/gin-gonic/gin"
)

type RoomController struct {
	rr *RoomRepo
}

func NewRoomController(router *gin.RouterGroup, rr *RoomRepo) RoomController {

	rc := RoomController{rr: rr}

	// Basic CRUD
	router.GET("/rooms", rc.HandleGetRooms)
	router.GET("/rooms/:id", rc.HandleGetRoomById)
	router.POST("/rooms", rc.HandleCreateRoom)
	router.PUT("/rooms/:id", rc.HandleUpdateRoom)
	router.DELETE("/rooms/:id", rc.HandleDeleteRoom)

	// Join/leave
	router.POST("/rooms/:id/join", rc.HandleJoinRoom)
	router.POST("/rooms/:id/leave", rc.HandleLeaveRoom)

	router.GET("/rooms/:id/messages", rc.HandleGetRoomMessages)

	return rc
}

func (rc *RoomController) HandleGetRooms(c *gin.Context) {
	rooms, err := rc.rr.GetRooms()
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}
	c.JSON(http.StatusOK, rooms)
}

func (rc *RoomController) HandleGetRoomById(c *gin.Context) {
	id := c.Param("id")

	room, isFound := rc.rr.GetRoom(id)
	if !isFound {
		c.AbortWithStatus(http.StatusNotFound)
		return
	}

	c.JSON(http.StatusOK, room)
}

func (rc *RoomController) HandleCreateRoom(c *gin.Context) {
	var room Room

	if err := c.BindJSON(&room); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdRoom, err := rc.rr.CreateRoom(room)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}

	c.JSON(http.StatusCreated, createdRoom)
}

func (rc *RoomController) HandleUpdateRoom(c *gin.Context) {
	id := c.Param("id")

	var room Room

	if err := c.BindJSON(&room); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedRoom, err := rc.rr.UpdateRoom(id, room)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}

	c.JSON(http.StatusOK, updatedRoom)
}

func (rc *RoomController) HandleDeleteRoom(c *gin.Context) {
	id := c.Param("id")

	err := rc.rr.DeleteRoom(id)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}

	c.Status(http.StatusNoContent)
}

func (rc *RoomController) HandleJoinRoom(c *gin.Context) {
	roomId := c.Param("id")
	user := c.MustGet("user").(users.User)
	err := rc.rr.AddUserToRoom(user.Id, roomId)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}
	c.Status(http.StatusNoContent)
}

func (rc *RoomController) HandleLeaveRoom(c *gin.Context) {
	roomId := c.Param("id")
	user := c.MustGet("user").(users.User)
	err := rc.rr.RemoveUserFromRoom(user.Id, roomId)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}
	c.Status(http.StatusNoContent)
}

func (rc *RoomController) HandleGetRoomMessages(c *gin.Context) {
	roomId := c.Param("id")
	messages, err := rc.rr.GetRoomMessages(roomId)
	if err != nil {
		if err := c.AbortWithError(http.StatusInternalServerError, err); err != nil {
			log.Printf("Failed to abort with error: %v", err)
		}
		return
	}
	c.JSON(http.StatusOK, messages)
}
