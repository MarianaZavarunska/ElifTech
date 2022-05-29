package controllers

import (
	"first-messanger/helper"
	"first-messanger/models"
	"first-messanger/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Register (context *gin.Context) {
	user := models.User{}

	context.BindJSON(&user)

	user, tokens,  err := services.CreateUser(context,user)
 
	var response helper.ResponseAuth

	if err != nil {

		response = helper.BuildErrorResponseAuth(user, tokens, err.Error())
		context.JSON(http.StatusBadRequest,response.Errors)
	} else {
		response = helper.BuildResponseAuth(user, tokens, true)
		context.JSON(200, response) 
	}

}

func Login (context *gin.Context) {

	user := models.User{}

	context.BindJSON(&user)

	user, tokens, err := services.LoginUser(context,user)

	var response helper.ResponseAuth

	if err != nil {

		response = helper.BuildErrorResponseAuth(user, tokens, err.Error())
		context.JSON(http.StatusUnauthorized,response.Errors)
	} else {
		response = helper.BuildResponseAuth(user, tokens, true)
		context.JSON(200, response) 
	}
}

func Logout(context *gin.Context) {
	services.LogoutUser(context);
	context.JSON(200, "logout") 
}
 
func Refresh(context *gin.Context) {

	user, tokens, err := services.RefreshUserToken(context)

	var response helper.ResponseAuth

	if err != nil {

		response = helper.BuildErrorResponseAuth(user, tokens, err.Error())
		context.JSON(http.StatusUnauthorized,response.Errors)
	} else {
		response = helper.BuildResponseAuth(user, tokens, true)
		context.JSON(200, response) 
	}
	
}