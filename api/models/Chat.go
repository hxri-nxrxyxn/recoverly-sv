package models

import (
	"time"

	"gorm.io/gorm"
)

type Chat struct {
	ChatID      uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID      uint      `gorm:"not null" json:"user_id"`               
	User        User      `gorm:"constraint:OnDelete:CASCADE;" json:"user"` // Reference to User, with cascading delete
	ChatGroupID uint      `gorm:"not null" json:"chat_group_id"`           // Foreign Key to ChatGroup
	ChatGroup   ChatGroup `gorm:"constraint:OnDelete:CASCADE;" json:"chat_group"` // Reference to ChatGroup, with cascading delete
	Message     string    `gorm:"type:varchar(255)" json:"message"`        // Message content
	Sent        time.Time `gorm:"autoCreateTime" json:"sent"`             // Auto-create timestamp for when the message is sent
}

func MigrateChat(db *gorm.DB) error {
	return db.AutoMigrate(&Chat{})
}
