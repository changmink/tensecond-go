package router

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type ChatRoom struct {
	Id       int64     `json:"id"`
	Name     string    `json:"name"`
	Created  time.Time `json:"created"`
	Modified time.Time `json:"modified"`
	Count    int       `json:"count"`
}

func GetUserChats(c *gin.Context) {
	id := c.Param("id")
	chats := GetChatsByUserId(id)
	c.JSON(200, chats)
}

func GetChatsByUserId(id string) []ChatRoom {
	db, err := sql.Open("sqlite3", "./10s.db")
	defer db.Close()
	checkErr(err)

	sql := fmt.Sprintf("SELECT id, room_name, chatroom.create_date, chatroom.modified_date, count FROM chat_user, chatroom, (SELECT room_id, COUNT(*) as count FROM message GROUP BY room_id HAVING read='FALSE' AND user_id !=%s) AS c WHERE user_id=%s AND chat_user.room_id = chatroom.id AND chat_user.room_id = c.room_id", id, id)
	rows, err := db.Query(sql)
	checkErr(err)

	chatRooms := make([]ChatRoom, 0)
	for rows.Next() {
		var chatRoom ChatRoom
		err := rows.Scan(&chatRoom.Id, &chatRoom.Name, &chatRoom.Created, &chatRoom.Modified, &chatRoom.Count)
		checkErr(err)
		chatRooms = append(chatRooms, chatRoom)
	}

	return chatRooms
}
