package services

import (
	"first-messanger/models"
	"first-messanger/repositories"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) ([]models.User, error){
	
	users, err := repositories.GetAllUsers()

  if err != nil {
	  return []models.User{}, err
  }

  return users , nil
}