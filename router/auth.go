package router

import (
	"database/sql"
	"fmt"
	"strings"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

type RequestUser struct {
	Id           int64  `json:"id"`
	Name         string `json:"name"`
	Auth         string `json:"auth"`
	ProfileImage string `json:"profileImage"`
}
type ResponseUser struct {
	Id             int64  `json:"id"`
	Name           string `json:"name"`
	ProfileImage   string `json:"profileImage"`
	ProfileMessage string `json:"profileMessage"`
}

func Auth(c *gin.Context) {
	var user RequestUser
	c.BindJSON(&user)
	resUser := AddUser(user)
	c.JSON(200, resUser)
}

func AddUser(user RequestUser) ResponseUser {
	db, err := sql.Open("sqlite3", "./10s.db")
	checkErr(err)

	stmt, err := db.Prepare("INSERT INTO user(id, name, auth, profile_image, profile_message, register_date, modified_date) VALUES(?, ?, ?, ?, ?, DATETIME('now'), DATETIME('now'))")
	checkErr(err)

	res, err := stmt.Exec(user.Id, user.Name, user.Auth, user.ProfileImage, " ")
	db.Close()
	// 중복일 경우 그냥 그대로 리턴함
	if err != nil && strings.HasPrefix(err.Error(), "UNIQUE constraint failed") {
		result := GetResponseUserById(user.Id)
		return result
	}
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)
	result := GetResponseUserById(id)
	return result
}

func GetResponseUserById(id int64) ResponseUser {
	db, err := sql.Open("sqlite3", "./10s.db")
	checkErr(err)
	rows, err := db.Query(fmt.Sprintf("SELECT id, name, profile_image, profile_message FROM user WHERE id=%d", id))
	checkErr(err)
	var result ResponseUser
	for rows.Next() {
		err = rows.Scan(&result.Id, &result.Name, &result.ProfileImage, &result.ProfileMessage)
		checkErr(err)
	}
	return result
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
