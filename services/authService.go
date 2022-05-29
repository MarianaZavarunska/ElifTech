package services

import (
	"errors"
	"first-messanger/models"
	"first-messanger/repositories"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser (context *gin.Context, user models.User) (models.User, models.Tokens, error) {

    //1. check if user exist with email -> call user repo 
	isExists, err := CheckUserExist(user.Email)
	
	if err != nil {
		 fmt.Println("Check error")
	}
	if isExists {
		fmt.Println("User exists")

		return models.User{}, models.Tokens{}, http.ErrBodyNotAllowed
	}

	//2. encrypt password GeneratehashPassword(user.Password) 

	user.Password, err = GeneratehashPassword(user.Password)
	if err != nil {
		log.Fatalln("error in password hash")
	}

	//3. save user

    repositories.SaveUser(&user)

	//4. creat access and refresh tokens 

	tokens, err := GenerateTokensPair(user.ID, user.Email)

	// 5. Set tokens into authorization headers
	context.Header("Authorization", tokens.AccessToken)
	context.Header("Authorization_Refresh", tokens.RefreshToken)

	return user, *tokens, err

}

func LoginUser (context *gin.Context, user models.User) (models.User, models.Tokens, error)  {
	
	//1. check if user exist with email -> call user repo 
	userFromDB, err := GetUserIfExist(user.Email)
	
	if err != nil {
		return models.User{}, models.Tokens{}, err
	}
	
	// 2. Check password
	isValid := CheckPasswordHash(user.Password, userFromDB.Password)

	if !isValid  {
		err = errors.New("user or password isn't valid")
		return models.User{}, models.Tokens{}, err
	}
	// 3. generate Tokens
	tokens, _:= GenerateTokensPair(user.ID, user.Email)

	// 4. Set tokens into authorization headers
	context.Request.Header.Add("Authorization", tokens.AccessToken)
	context.Request.Header.Add("Authorization_Refresh", tokens.RefreshToken)

	// 5. return response

	return userFromDB, *tokens, nil
}

func LogoutUser(context *gin.Context){

		// 1. Get accessToken from authorization headers
		accessToken := context.GetHeader("Authorization");
		
		if accessToken == "" {
			context.JSON(http.StatusUnauthorized, gin.H{"status": http.StatusUnauthorized, "message": "Authorization Token is required"})
			context.Abort()
			return
		}
}

//FE: check if access token exists and if it's valid. If yes, return, not - redirect refresh token 

func RefreshUserToken(context *gin.Context, tokenString string) (models.User, models.Tokens, error) {
	

	// 1. Convert tokenString to JWT 
 
	err := godotenv.Load()
	if err != nil {
	  fmt.Println("Error loading .env file")
	}
  
	VERIFICATION_KEY := os.Getenv("SECRET_REFRESH_KEY")

    claims := jwt.MapClaims{}

    token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
    return []byte(VERIFICATION_KEY), nil
    })

    if err != nil {

		return models.User{}, models.Tokens{}, err
	}

	// 2. Since token is valid, get the email:

	userEmail, userId, err := GetUserDataFromToken(token)

	if err != nil {
		fmt.Println("the error: ", err)
		context.JSON(http.StatusUnauthorized, "Email doesn't exist")
		return models.User{}, models.Tokens{}, err
	}

	// 3. Generate tokens 
	tokens, err:= GenerateTokensPair(userId, userEmail)
	 if err !=nil {
		fmt.Println("the error: ", err)
		return models.User{}, models.Tokens{}, err
	}

	// 4. Set tokens into authorization headers
	 context.Header("Authorization", tokens.AccessToken)

	var user = models.User{
		Email: userEmail,
	}
	

	return user, *tokens, nil

}


func CheckUserExist(userEmail string) (isExists bool, err error){

 user, err := repositories.GetUserByEmail(userEmail)

 if err != nil {
	 return false , err
 } else {
	return user.Email != "", nil
	 
 }
}

func GetUserIfExist(userEmail string) (models.User, error){

	user, err := repositories.GetUserByEmail(userEmail)

	if err != nil {
		return models.User{} , err
	} else {
		if user.ID > 0{
			return user, nil
		} else {
			err = errors.New("user doesn't exist")
			return models.User{} , err
		}
		
	}
}


func GeneratehashPassword(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
