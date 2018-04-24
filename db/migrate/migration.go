package main

import (
  "Slack_Share_Timeline/timeline"

  "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
  db, err := gorm.Open("mysql", "root@/slack_timeline?charset=utf8&parseTime=True&loc=Local")
  if err != nil {
    panic(err.Error())
  }
  defer db.Close()

  db.AutoMigrate(&timeline.Timeline{})

}
