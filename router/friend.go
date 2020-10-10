package router

import (
	"database/sql"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func Friend(c *gin.Context) {
	id := c.Param("id")
	fmt.Println(id)
	idValue, err := strconv.ParseInt(id, 10, 64)
	checkErr(err)
	friends := GetFriendsById(idValue)
	c.JSON(200, friends)
}

func GetFriendsById(id int64) []ResponseUser {
	db, err := sql.Open("sqlite3", "./10s.db")
	checkErr(err)
	rows, err := db.Query(fmt.Sprintf("SELECT friend_id FROM friend WHERE user_id=%d", id))
	checkErr(err)

	var friendSlice []ResponseUser
	var friendId int64
	for rows.Next() {
		err = rows.Scan(&friendId)
		checkErr(err)
		friend := GetResponseUserById(friendId)
		friendSlice = append(friendSlice, friend)
	}

	return friendSlice
}
