package main

import (
	"net/http"

	"github.com/dolencd/go-playground/chatserver/common"
	"github.com/dolencd/go-playground/chatserver/messages"
	"github.com/dolencd/go-playground/chatserver/rooms"
	"github.com/dolencd/go-playground/chatserver/users"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	err := godotenv.Load("../.env")
	if err != nil {
		panic(err)
	}
	conn, err := common.InitializeConnection()
	if err != nil {
		panic(err)
	}
	r := gin.Default()

	public := r.Group("/api/")
	public.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	messageRepo := messages.NewMessageRepo(conn)
	userRepo := users.NewUserRepo(conn)
	roomRepo := rooms.NewRoomRepo(conn)
	private := r.Group("/api/")
	private.Use(common.PopulateUserMiddleware(&userRepo))
	private.Use(common.RequireUserMiddleware())
	messages.NewMessageController(private, &messageRepo)
	users.NewUserController(private, &userRepo)
	rooms.NewRoomController(private, &roomRepo)

	return r
}

func main() {
	r := setupRouter()
	// Listen and Server in 0.0.0.0:8080
	err := r.Run(":8080")
	if err != nil {
		panic(err)
	}
}
