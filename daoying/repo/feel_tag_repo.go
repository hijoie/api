package repo

import (
	"gorm.io/gorm"
)

type FeelTag struct {
	gorm.Model
	Name string `gorm:"column:name" json:"name"`
	Sort int    `gorm:"column:sort" json:"sort"`
	Type int    `gorm:"column:type" json:"type"`
}

// TableName sets the insert table name for this struct type
func (f *FeelTag) TableName() string {
	return "feel_tag"
}
