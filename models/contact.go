package models

import "time"

type Contact struct {
	ID          uint64 `gorm:"primaryKey;autoIncrement"`
	Name        string `gorm:"type:varchar(255);not null"`
	Email       string `gorm:"type:varchar(255);email;not null;"`
	Description string `gorm:"type:text;not null"`
	Status      string `gorm:"type:ENUM('active','deleted');default:'active'"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   time.Time `gorm:"type:TIMESTAMP;default:null"`
}
