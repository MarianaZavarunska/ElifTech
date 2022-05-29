package models

import (
	"gorm.io/gorm"
)

type Tokens struct {
	AccessToken string   `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	UserEmail    string  `json:"userEmail"`
	AtExpires    int64
	RtExpires    int64
}

type BlackList struct {
	gorm.Model
	Token string   `json:"token"`
	TokenType string  `json:"tokenType"`
}

func (b *BlackList) TableName() string {
	return "blacklist"
}
