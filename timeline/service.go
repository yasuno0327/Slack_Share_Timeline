package timeline

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/nlopes/slack"
	"golang.org/x/exp/utf8string"
)

func connectDB() *gorm.DB {
	DBMS := "mysql"
	USER := "root"
	//PASS := ""
	//PROTOCOL := ""
	DBNAME := "slack_timeline"

	CONNECT := USER + "@" + "/" + DBNAME + "?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open(DBMS, CONNECT)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func Create(rooms []string) (attachment slack.Attachment) {
	db := connectDB()
	owner := utf8string.NewString(rooms[0]).Slice(2, 10)
	clients := rooms[1:]
	fmt.Println(owner, clients)

	for _, v := range clients {
		timeline := Timeline{OwnerID: owner, ClientID: utf8string.NewString(v).Slice(2, 10)}
		if err := db.Create(&timeline).Error; err != nil {
			panic(err.Error())
		}
	}

	attachment = slack.Attachment{
		Text: "Create timeline success!!",
	}

	return attachment
}

func HandleMessageResponse(user, text, channel string) (attachment slack.Attachment) {
	return attachment
}
