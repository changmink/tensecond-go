package main

import (
	"./router"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/auth", router.Auth)
	r.GET("/friends/:id", router.Friend)
	r.PUT("/users/:id", router.UpdateUserFunc)
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
