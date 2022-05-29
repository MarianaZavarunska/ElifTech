package repositories

import (
	"first-messanger/config"
	"first-messanger/models"
)


func FindTokenInBlacklist(tokenType, token string) (bool, error) {

	tokenFromBl := []models.BlackList{}

	result := config.DB.Where("token_type = ?", tokenType).Where("Token = ?", token).Find(&tokenFromBl);

	if result.Error != nil {
		err := result.Error 
		return false, err;
   }

	return len(tokenFromBl) > 0, nil
}

func SetTokenToBlacklist(blacklistItem *models.BlackList ) (error) {

	err := config.DB.Create(&blacklistItem)

	if err != nil {
		return err.Error
	}

    return nil	
}
