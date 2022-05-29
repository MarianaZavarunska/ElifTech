package middlewares

import (
	"first-messanger/services"
	"fmt"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)



func AuthMiddleware(tokenType string) gin.HandlerFunc {
	 
  return func(context *gin.Context) {
	var clientToken string
	// 1. Get Token 
	if tokenType == "access"{
	clientToken = context.GetHeader("Authorization")
	} else {
	clientToken = context.GetHeader("Authorization_refresh")
	}

	// clientToken = strings.Split(clientToken, "Bearer ")[1]
	if clientToken == "" {
		context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
		context.Abort()
		return
	}

	// TODO: Ask where write a func, which will be looking for token in blacklist

	// 2. Check Blacklist 
	isTokenInBlacklist := services.GetTokenFromBlacklist(tokenType,clientToken)

	if isTokenInBlacklist {
		fmt.Println("blacklist")
		fmt.Println("found in Blacklist")
		context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
		context.Abort()
		return
	}
	claims := jwt.MapClaims{}

	// 3. Validate Token

	err := godotenv.Load()
	if err != nil {
	  fmt.Println("Error loading .env file")
	}

	var VERIFICATION_KEY string
	
	switch tokenType {
	case "access" : 
	   VERIFICATION_KEY = os.Getenv("SECRET_ACCESS_KEY")
	case "refresh": 
	   VERIFICATION_KEY = os.Getenv("SECRET_REFRESH_KEY")
	}

	parsedToken, err := jwt.ParseWithClaims(clientToken, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(VERIFICATION_KEY), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token Signature"})
			context.Abort()
			return
		}

	context.JSON(http.StatusBadRequest, gin.H{"status": http.StatusBadRequest, "message": err.Error()})
	fmt.Println(err)
	context.Abort()
	 return

	}

	if !parsedToken.Valid {
	    fmt.Println("parsed")
		context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Invalid Token"})
		context.Abort()
		return
	}

    context.Next()
  }
}