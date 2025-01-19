package models

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ChatID    uint      `gorm:"primaryKey;autoIncrement" json:"id"` // Chat ID
	UserID    uint      `gorm:"not null" json:"user_id"`
	Name      string    `gorm:"type:varchar(255)" json:"name"`       // User name
	ChatGroup string    `gorm:"type:varchar(255)" json:"chat_group"` // Chat group name
	Message   string    `gorm:"type:varchar(255)" json:"message"`    // Message content
	Sent      time.Time `gorm:"autoCreateTime" json:"sent"`          // Auto-create timestamp for when the message is sent
}

func MigrateChat(db *gorm.DB) error {
	return db.AutoMigrate(&Chat{})
}
