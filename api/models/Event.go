package models

import (
	"time"

	"gorm.io/gorm"
)

type Event struct {
	EventID   uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`     // User name
	Locatioln string    `gorm:"type:varchar(255)" json:"location"` // Chat group name
	Details   string    `gorm:"type:varchar(255)" json:"details"`  // Message content
	Time      time.Time `json:"time"`                              // Auto-create timestamp for when the message is sent
}

func MigrateEvent(db *gorm.DB) error {
	return db.AutoMigrate(&Event{})
}
