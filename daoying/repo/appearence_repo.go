package repo

import (
	"gorm.io/gorm"
)

type Appearence struct {
	gorm.Model
	Url         string `gorm:"column:url;type:varchar(255)" json:"url"'`
	Version     int    `gorm:"column:version;type:integer" json:"version"`
	VersionName string `gorm:"column:version_name;type:varchar(16)" json:"version_name"`
	Code        string `gorm:"column:code;type:varchar(32)" json:"code"`
	Name        string `gorm:"column:name;type:varchar(32)" json:"name"`
	Desc        string `gorm:"column:desc;type:varchar(64)" json:"desc"`
	Value       int    `gorm:"column:value;type:integer" json:"value"`
}

// TableName sets the insert table name for this struct type
func (a *Appearence) TableName() string {
	return "appearence"
}
