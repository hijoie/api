package repo

import "time"

type URecord struct {
	ScenesID   int       `gorm:"column:scenes_id" json:"scenes_id"`
	Score      int       `gorm:"column:score" json:"score"`
	UpdateTime time.Time `gorm:"column:update_time" json:"update_time"`
	Content    string    `gorm:"column:content" json:"content"`
	CreateTime time.Time `gorm:"column:create_time" json:"create_time"`
	FeelID     int       `gorm:"column:feel_id" json:"feel_id"`
	ID         int       `gorm:"column:id;primary_key" json:"id"`
}

// TableName sets the insert table name for this struct type
func (u *URecord) TableName() string {
	return "u_record"
}
