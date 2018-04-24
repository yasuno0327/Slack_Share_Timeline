package timeline

import (
  "fmt"
  "regexp"

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
  regex := regexp.MustCompile(`[A-Z0-9]*{9}`)
  owner := regex.FindAllStringSubmatch(rooms[0], -1)
  clients := rooms[1:]
  fmt.Println(owner, clients)

  for _, v := range clients {
    timeline := Timeline{OwnerID: owner, ClientID: v}
    if err := db.Create(&timeline).Error; err != nil {
      panic(err.Error())
    }
  }

  attachment = slack.Attachment{
    Text: "Create timeline success!!",
  }

  return attachment
}
