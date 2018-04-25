package timeline

import (
	"github.com/jinzhu/gorm"
)

type Timeline struct {
	gorm.Model
	ClientID string
	OwnerID  string
}
