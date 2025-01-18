package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	UserID    uint       `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      *string    `gorm:"type:varchar(255)" json:"name"`
	Email     string     `gorm:"type:varchar(255);unique;not null" json:"email"`
	Password  string     `gorm:"type:varchar(255);not null" json:"password"`
	Birthdate *time.Time `json:"birthdate"`
	Gender    *string    `gorm:"varchar(255)" json:"gender"`
	Height    *int       `json:"height"`
	Weight    *int       `json:"weight"`
	Severity  *int       `json:"severity"`
	Condition *string    `gorm:"varchar(255)" json:"condition"`
	Addiction *string    `gorm:"varchar(255)" json:"addiction"`
	CreatedAt time.Time  `gorm:"autoCreateTime" json:"created_at"`
}

func MigrateUser(db *gorm.DB) error{
	err := db.AutoMigrate(&User{})
	return err
}
