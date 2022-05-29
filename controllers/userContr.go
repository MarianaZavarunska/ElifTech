package controllers

import (
	"first-messanger/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllUsers(context *gin.Context) {
	
   users, err := services.GetAllUsers(context)

   if err != nil {
	context.AbortWithError(http.StatusBadRequest, err)
	fmt.Println(err.Error())
  }
  context.JSON(200, users)
}