package repo

import "gorm.io/gorm"

type Introduction struct {
	gorm.Model
	Content string `gorm:"column:content" json:"content,omitempty"`
}
