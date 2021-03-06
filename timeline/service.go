package timeline

import (
	"log"

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
	owner := utf8string.NewString(rooms[0]).Slice(2, 11)
	clients := rooms[1:]

	for _, v := range clients {
		timeline := Timeline{OwnerID: owner, ClientID: utf8string.NewString(v).Slice(2, 11)}
		if err := db.Create(&timeline).Error; err != nil {
			log.Println(err)
		}
	}

	attachment = slack.Attachment{
		Text: "Create timeline success!!",
	}

	return attachment
}

func HandleMessageResponse(user, text, channel string) (attachment slack.Attachment, owner string) {
	db := connectDB()
	timelines := []Timeline{}

	db.Find(&timelines, "client_id=?", channel)
	if len(timelines) == 0 {
		return
	}
	for _, v := range timelines {
		owner = v.OwnerID
	}
	attachment = slack.Attachment{
		Text: "<@" + user + ">" + " say  " + text,
	}
	return attachment, owner
}
