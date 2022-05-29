package services

import (
	"errors"
	"first-messanger/models"
	"first-messanger/repositories"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
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

	 // 5. Set Cookies
	 SetCookie(context, "accessToken", tokens.AccessToken, 900)
	 SetCookie(context, "refreshToken", tokens.RefreshToken, 64000)

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

    // 4. set Cookies

	// t := time.Duration(15 * time.Minute)
	// t  :=time.Now().Add(time.Minute * 15).Unix()
	// int(12 * time.Hour)
	SetCookie(context, "accessToken", tokens.AccessToken, 900)
	SetCookie(context, "refreshToken", tokens.RefreshToken, 64000)

	// 5. return response

	return userFromDB, *tokens, nil
}

func LogoutUser(context *gin.Context){
	SetCookie(context, "accessToken", "", -1);
}

//FE: check if access token exists and if it's valid. If yes, return, not - redirect refresh token 

func RefreshUserToken(context *gin.Context) (models.User, models.Tokens, error) {

	 //1. Check if refresh token valid.
	val, err :=context.Cookie("refreshToken")

	if err != nil {
		fmt.Println("the error: ", err , http.StatusUnauthorized, "Refresh token expired")
		return models.User{}, models.Tokens{}, err
	} 

	refreshToken, err := VerifyToken(val)

	if err != nil {
		fmt.Println("the error: ", err)
		context.JSON(http.StatusUnauthorized, "Refresh isn't valid")
		return models.User{}, models.Tokens{}, err
	}

	// 2. Since token is valid, get the email:

	userEmail, userId, err := GetUserDataFromToken(refreshToken)

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


    // 4.  set Cookies

	SetCookie(context, "accessToken", tokens.AccessToken, 900)
	SetCookie(context, "refreshToken", tokens.RefreshToken, 64000)

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

	// err := godotenv.Load()
    // if err != nil {
    // fmt.Println("Error loading .env file")
    // }
    // salt := os.Getenv("HASH_SALT")

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
