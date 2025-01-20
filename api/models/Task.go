package models

import (
	"gorm.io/gorm"
)

type Task struct {
	Task string   `gorm:"type:varchar(255)" json:"task"`
	Todo []string `gorm:"type:varchar(255)" json:"todo"`
}

func MigrateTask(db *gorm.DB) error {
	return db.AutoMigrate(&Task{})
}
