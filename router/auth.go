package router

import (
	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func Auth(c *gin.Context) {
	var user RequestUser
	c.BindJSON(&user)
	resUser := AddUser(user)
	c.JSON(200, resUser)
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
