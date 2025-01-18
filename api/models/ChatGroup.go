package models

import (
	"gorm.io/gorm"
)

type ChatGroup struct {
	ChatGroupID uint    `gorm:"primaryKey;autoIncrement" json:"id"`
	Name        string  `gorm:"type:varchar(255)" json:"name"`  // Name of the group
	Online      uint    `gorm:"type:int" json:"online"`         // Number of users online
	Chats       []Chat  `gorm:"foreignKey:ChatGroupID;constraint:OnDelete:CASCADE;" json:"chats"` // One-to-many relation with Chat
}

func MigrateChatGroup(db *gorm.DB) error {
	return db.AutoMigrate(&ChatGroup{})
}
