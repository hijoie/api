package repo

import (
	"gorm.io/gorm"
)

type LightChangeLog struct {
	gorm.Model
	LogDesc    string `gorm:"column:log_desc;type:varchar(128)" json:"log_desc"`
	OpenID     string `gorm:"column:open_id;type:varchar(128);uniqueindex" json:"open_id"`
	UserID     int    `gorm:"column:user_id;index" json:"user_id"`
	Account    int    `gorm:"column:account" json:"account"`
	ChangeType int    `gorm:"column:change_type" json:"change_type"`
}

// TableName sets the insert table name for this struct type
func (l *LightChangeLog) TableName() string {
	return "light_change_log"
}
