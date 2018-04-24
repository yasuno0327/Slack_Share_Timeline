package timeline

import (
  "github.com/nlopes/slack"
  "github.com/jinzhu/gorm"
  _ "github.com/jinzhu/gorm/dialects/mysql"
)

func connectDB() *gorm.DB {
  DBMS := "mysql"
  USER := "root"
  //PASS := ""
  //PROTOCOL := ""
  DBNAME := "slack_timeline"

  CONNECT := USER+"@"+"/"+DBNAME+"?charset=utf8&parseTime=True&loc=Local"
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

  for _, v := range clients {
    timeline := Timeline{OwnerID: owner, ClientID: v}
    if err := db.Create(&timeline).Error; err != nil {
      panic(err.Error())
    }
  }

  attachment = slack.Attachment{
    Title: "Create timeline success!!",
  }

  return attachment
}
