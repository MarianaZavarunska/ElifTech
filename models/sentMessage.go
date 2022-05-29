package models

import (
	"gorm.io/gorm"
)

type SentMessage struct {
	gorm.Model
	Author string `json:"author"`
	Receiver string `json:"receiver"`
	Body string `json:"body"`
}


func (b *SentMessage) TableName() string {
	return "message_sent"
}
