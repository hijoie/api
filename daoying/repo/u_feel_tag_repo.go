package repo

import "time"

type UFeelTag struct {
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	ID         int       `gorm:"column:id;primary_key" json:"id"`
	Name       string    `gorm:"column:name" json:"name"`
	Type       int       `gorm:"column:type" json:"type"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	UserID     int       `gorm:"column:user_id" json:"user_id"`
}

// TableName sets the insert table name for this struct type
func (u *UFeelTag) TableName() string {
	return "u_feel_tag"
}
