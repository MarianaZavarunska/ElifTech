package services

import (
	"errors"
	"first-messanger/models"
	"first-messanger/repositories"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/joho/godotenv"
)

//pass email/id and set in token
func GenerateTokensPair(userId uint, userEmail string) (*models.Tokens, error) {
	td := &models.Tokens{}
	td.AtExpires = time.Now().Add(time.Minute * 15).Unix()
	td.RtExpires = time.Now().Add(time.Hour * 24 * 7).Unix()
	

	var err error

	godotenv.Load()

	//Creating Access Token
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["user_email"] = userEmail
	atClaims["exp"] = td.AtExpires
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	td.AccessToken, err = at.SignedString([]byte(os.Getenv("SECRET_ACCESS_KEY")))
	if err != nil {
		return nil, err
	}

	//Creating Refresh Token
	rtClaims := jwt.MapClaims{}
	rtClaims["user_email"] = userEmail
	rtClaims["user_id"] = userId
	rtClaims["exp"] = td.RtExpires
	rt := jwt.NewWithClaims(jwt.SigningMethodHS256, rtClaims)
	td.RefreshToken, err = rt.SignedString([]byte(os.Getenv("SECRET_REFRESH_KEY")))
	if err != nil {
		return nil, err
	}
	return td, nil
}

func VerifyToken(tokenFromCookie string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenFromCookie, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET_REFRESH_KEY")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func GetUserDataFromToken(refreshToken *jwt.Token ) (string, uint, error) {
	claims, ok := refreshToken.Claims.(jwt.MapClaims)
	if ok  {
		userEmail, _ := claims["user_email"].(string) //convert the interface to string
	
		id, _ := strconv.ParseUint(fmt.Sprintf("%.f", claims["user_id"]), 10, 32)
		userId:= uint(id)
		return userEmail, userId, nil 

	} else {
	   err := errors.New("refresh expired")
       return "" , 0,  err

		// c.JSON(http.StatusUnauthorized, "refresh expired")
	}
}

func GetTokenFromBlacklist(tokenType string, token string) bool {
   
	isTokenExist, _ := repositories.FindTokenInBlacklist(tokenType,token)

	return isTokenExist
   
}

func SetTokenToBlacklist(blacklistItem *models.BlackList) ( error){
    err := repositories.SetTokenToBlacklist(blacklistItem)

	if err != nil {
		return  err
	}
	
	return  nil
}