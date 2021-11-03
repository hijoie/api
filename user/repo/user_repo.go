package repo

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	UserId   string `gorm:"column:user_id;type:varchar(32);uniqueindex" json:"user_id"`
	Password string `gorm:"column:password;type:varchar(32)" json:"password"`
	Phone    string `gorm:"column:phone;index;type:varchar(16)" json:"phone"`
	UserName string `gorm:"column:user_name;type:varchar(32)" json:"user_name"`
}

// TableName sets the insert table name for this struct type
func (u *User) TableName() string {
	return "user"
}

