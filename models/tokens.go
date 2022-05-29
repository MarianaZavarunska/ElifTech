package models

type Tokens struct {
	AccessToken string   `json:"accessToken"`
	RefreshToken string  `json:"refreshToken"`
	UserEmail    string  `json:"userEmail"`
	AtExpires    int64
	RtExpires    int64
}