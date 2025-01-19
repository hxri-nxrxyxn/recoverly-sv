package models

import (
	"gorm.io/gorm"
)

type Template struct {
	Name string   `gorm:"type:varchar(255)" json:"Name"`
	Pages int 	`gorm:"type:integer" json:"pages"`
	Url string 	`gorm:"type:varchar(255)" json:"url"`
}

func MigrateTemplate(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
