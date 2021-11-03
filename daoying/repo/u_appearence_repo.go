package repo

import (
	"gorm.io/gorm"
)

type UAppearence struct {
	gorm.Model
	UserID            string `gorm:"column:user_id;type:varchar(32);index" json:"user_id"`
	AppearenceCode    string `gorm:"column:appearence_code" json:"appearence_code"`
	AppearenceVersionName string `gorm:"column:appearence_version_name;type:varchar(16)" json:"appearence_version_name"`
	Status            int `gorm:"column:status;type:smallint" json:"status"`
}

// TableName sets the insert table name for this struct type
func (u *UAppearence) TableName() string {
	return "u_appearence"
}
