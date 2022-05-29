package helper

import (
	"first-messanger/models"
	"strings"
)

//Response is used for static shape json return
type ResponseAuth struct{
	Status  bool         `json:"status"`
	User models.User     `json:"user"`
	Tokens models.Tokens `json:"tokens"`
	Errors  interface{}  `json:"errors"`
}

//BuildResponse method is to inject data value to dynamic success response
func BuildResponseAuth(user models.User, tokens models.Tokens, status bool) ResponseAuth{
	res := ResponseAuth {
		Status: status,
		User: user,
	    Tokens: tokens,
		Errors: nil,
	}

	return res
}

//BuildErrorResponse method is to inject data value to dynamic failed response
func BuildErrorResponseAuth(user models.User, tokens models.Tokens, err string ) ResponseAuth{
	splittedError := strings.Split(err, "\n")
	res := ResponseAuth{
		Status:  false,
		User: user,
	    Tokens: tokens,
		Errors:  splittedError,
	}
	return res
}
 
