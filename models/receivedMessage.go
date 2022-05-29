package models

import (
	"gorm.io/gorm"
)

type ReceiveMessage struct {
	gorm.Model
	Author string `json:"author"`
	Receiver string `json:"receiver"`
	Body string `json:"body"`
	
}

func (b *ReceiveMessage) TableName() string {
	return "message_received"
}