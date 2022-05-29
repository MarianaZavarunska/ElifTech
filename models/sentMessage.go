package models

import (
	"gorm.io/gorm"
)

type SentMessage struct {
	gorm.Model
	// ID   uint   `gorm:"primaryKey"`
	Author string `json:"author"`
	Receiver string `json:"receiver"`
	Body string `json:"body"`
	// CreatedAt time.Time
    // UpdatedAt time.Time
    // DeletedAt gorm.DeletedAt `gorm:"index"`
}


func (b *SentMessage) TableName() string {
	return "message_sent"
}
