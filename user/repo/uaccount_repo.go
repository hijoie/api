package repo

import (
	"gorm.io/gorm"
)

type UAccount struct {
	gorm.Model
	Light int `gorm:"column:light" json:"light"`
}

// TableName sets the insert table name for this struct type
func (u *UAccount) TableName() string {
	return "u_account"
}
