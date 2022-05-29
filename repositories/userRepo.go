package repositories

import (
	"first-messanger/config"
	"first-messanger/models"
)

func GetUserByEmail(email string) (user models.User, err error) {
	user = models.User{}

	result := config.DB.Where("Email = ?", email).Find(&user);

	if result.Error != nil {
		err := result.Error 
		return  models.User{}, err;
   }

	return  user, nil
}

func SaveUser(user *models.User) (err error) {

	result := config.DB.Create(user)

	return result.Error

}

