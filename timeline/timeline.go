package timeline

import (
	"github.com/jinzhu/gorm"
)

type timeline struct {
	gorm.Model
	ClientIDs []string
	OwnerID   string
}
