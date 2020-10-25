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

type Message struct {
	UserId   int64     `json:"userId"`
	Length   int       `json:"length"`
	Audio    string    `json:"audio"`
	Read     bool      `json:"read"`
	Outdated bool      `json:"outdated"`
	Modified time.Time `json:"modified"`
}

type MessageList struct {
	Date     string `json:"date"`
	Messages []Message
}

func GetChatMessageList(c *gin.Context) {
	roomId := c.Param("roomId")

	db, err := sql.Open("sqlite3", "./10s.db")
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT date(modified_date) as date FROM message WHERE room_id=" + roomId + " GROUP BY date(modified_date)")
	checkErr(err)

	messageList := make([]MessageList, 0)
	for rows.Next() {
		var msgList MessageList
		err := rows.Scan(&msgList.Date)
		checkErr(err)
		messages := GetChatMessagesFromIdAndDate(roomId, msgList.Date)
		msgList.Messages = messages
		messageList = append(messageList, msgList)
	}
	c.JSON(200, messageList)
}

func GetChatMessagesFromIdAndDate(roomId string, date string) []Message {
	db, err := sql.Open("sqlite3", "./10s.db")
	defer db.Close()
	checkErr(err)
	sql := "SELECT user_id, length, audio, read, outdate, modified_date FROM message WHERE room_id=" + roomId + " AND date(modified_date)='" + date + "' ORDER BY modified_date"
	fmt.Println(sql)
	rows, err := db.Query(sql)
	checkErr(err)

	messages := make([]Message, 0)
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.UserId, &msg.Length, &msg.Audio, &msg.Read, &msg.Outdated, &msg.Modified)
		checkErr(err)
		messages = append(messages, msg)
	}

	return messages
}

func GetChatMessages(c *gin.Context) {
	roomId := c.Param("roomId")

	db, err := sql.Open("sqlite3", "./10s.db")
	defer db.Close()
	checkErr(err)

	rows, err := db.Query("SELECT user_id, length, audio, read, outdate, modified_date FROM message WHERE room_id=" + roomId + " ORDER BY modified_date")
	checkErr(err)

	messages := make([]Message, 0)
	for rows.Next() {
		var msg Message
		err := rows.Scan(&msg.UserId, &msg.Length, &msg.Audio, &msg.Read, &msg.Outdated, &msg.Modified)
		checkErr(err)
		messages = append(messages, msg)
	}

	c.JSON(200, messages)
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
