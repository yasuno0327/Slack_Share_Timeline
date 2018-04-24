package timeline

import (
	"Slack_Share_Timeline/client"
	"Slack_Share_Timeline/owner"
	"github.com/jinzhu/gorm"
)

type timeline struct {
	gorm.Model
	ClientID uint
	OwnerID  uint
}
