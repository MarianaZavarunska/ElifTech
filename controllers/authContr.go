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
	token := context.GetHeader("Authorization")

	if token == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
		context.Abort()
		return
	}

	blacklistItem := models.BlackList{TokenType: "access", Token: token}

	 err := services.SetTokenToBlacklist(&blacklistItem)

	if err != nil {
		context.JSON(500, "Token wasn't set in blacklist")
		context.Abort()
		return
	}

   context.JSON(200, "Token was set in blacklist")

}
 
func Refresh(context *gin.Context) {

	// 1. Get tokens
	accessToken := context.GetHeader("Authorization")
	refreshToken := context.GetHeader("Authorization_Refresh")


	if refreshToken == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
		context.Abort()
		return
	}
	
	// 2. Generate new token's pair
	user, tokens, err := services.RefreshUserToken(context, refreshToken)

	// 3. Set old tokens to blacklist

	blacklist1 := models.BlackList{TokenType: "access", Token: accessToken}
	blacklist2 :=  models.BlackList{TokenType: "refresh", Token: refreshToken}

	err1 := services.SetTokenToBlacklist(&blacklist1)

	if err1 != nil {

		context.JSON(500, "Token wasn't set in blacklist")
		context.Abort()
		return

  }


	err2 := services.SetTokenToBlacklist(&blacklist2)

	if err2 != nil {

		  context.JSON(500, "Token wasn't set in blacklist")
		  context.Abort()
		  return

	}

	// TODO : Fix slice of blacklist's items
	

	// blacklist := []models.BlackList{}

	// blacklist[0] = models.BlackList{TokenType: "access", Token: accessToken}
	// blacklist[1] = models.BlackList{TokenType: "refresh", Token: refreshToken}

	// for key := range blacklist{
	// 	err := services.SetTokenToBlacklist(&blacklist[key])

	// 	if err != nil {

	// 	  context.JSON(500, "Token wasn't set in blacklist")
	// 	  context.Abort()
	// 	  return
	// 	}
	// }

    // 4. Build response 

	var response helper.ResponseAuth 

	if err != nil {

		response = helper.BuildErrorResponseAuth(user, tokens, err.Error())
		context.JSON(http.StatusUnauthorized,response.Errors)
	} else {
		response = helper.BuildResponseAuth(user, tokens, true)
		context.JSON(200, response) 
	}
	
}