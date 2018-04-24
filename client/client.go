package client

import (
	"Slack_Share_Timeline/timeline"
	"github.com/jinzhu/gorm"
)

type Client struct {
	gorm.Model
	RoomId    string
	Timelines []timeline.Timeline
}
