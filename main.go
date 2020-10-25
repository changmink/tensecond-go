package main

import (
	"./router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/auth", router.Auth)
	r.GET("/friends/:id", router.Friend)
	r.GET("/users/:id", router.GetUserInfoFunc)
	r.PUT("/users/:id", router.UpdateUserFunc)
	r.GET("/users/:id/chats", router.GetUserChats)
	r.GET("/users/:id/chats/:roomId/messages", router.GetChatMessages)
	r.GET("/users/:id/chats/:roomId/messages/date", router.GetChatMessageList)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
