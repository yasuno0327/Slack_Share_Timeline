package timeline

import (
  "strings"

  "github.com/nlopes/slack"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectDB() *gorm.DB {
  DBMS := "mysql"
  USER := "root"
  PASS := ""
  //PROTOCOL := ""
  DBNAME := "slack_timeline"

  CONNECT := USER+":"+PASS+"@"+"/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local"
  db, err := gorm.Open(DBMS, CONNECT)
  if err != nil {
    panic(err.Error())
  }
  return db
}

func Create(rooms []string) (attachment slack.Attachment) {
  db := connectDB()
  owner := rooms[0]
  clients := rooms[1:]
  timeline := Timeline{OwnerID: owner, ClientID: strings.Join(clients, ",")}

  db.Create(&timeline)

  attachment = slack.Attachment{
    Pretext: "Create timeline success!!",
  }

  return attachment
}
