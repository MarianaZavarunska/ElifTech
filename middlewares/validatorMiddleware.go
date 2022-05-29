package middlewares

import (
	"first-messanger/models"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)


type IError struct {
    Field string
    Tag   string
    Value string
}

var Validator = validator.New()

func  ValidatorMiddleware ( )  gin.HandlerFunc {

	return func (context *gin.Context)  {
	
		var errors []*IError

		user := models.User{}

		context.BindJSON(&user)

		err := Validator.Struct(user)

		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				var el IError
				el.Field = err.Field()
				el.Tag = err.Tag()
				el.Value = err.Param()
				errors = append(errors, &el)
			}
			context.JSON(500, errors)
			context.Abort()
			return
		}
		 context.Next()

	}
}