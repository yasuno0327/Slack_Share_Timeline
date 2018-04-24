package owner

import (
	"Slack_Share_Timeline/timeline"
	"github.com/jinzhu/gorm"
)

// Slack Owner channel
type Owner struct {
	gorm.Model
	RoomID    string
	Timelines []timeline.Timeline
}
