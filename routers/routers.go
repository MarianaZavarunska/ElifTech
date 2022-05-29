package routers

import (
	"github.com/gin-gonic/gin"

	"first-messanger/controllers"
)


func SetupRouter() *gin.Engine {
  router := gin.Default()

  // router.Use(cors.New(cors.Config{
  //   AllowOrigins:     []string{"https://gorm.io"},
  //   AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "DELETE"},
  //   AllowHeaders:     []string{"Origin"},
  //   ExposeHeaders:    []string{"Content-Length"},
  //   AllowCredentials: true,
  //   MaxAge: 12 * time.Hour,
  // }))

  v1 := router.Group("/messages")
  {
    v1.GET("send/:author", controllers.GetAllSendMessages)
	  v1.GET("received/:receiver", controllers.GetAllReceivedMessages)

    v1.POST("", controllers.SendMessage)

    v1.PATCH("send/:id", controllers.UpdateSendMessage)
    v1.PATCH("received/:id", controllers.UpdateReceivedMessage)

    v1.DELETE("send/:id", controllers.DeleteSendMessage)
    v1.DELETE("received/:id", controllers.DeleteReceivedMessage)
  }

  v2 := router.Group("api/auth")
  {
	  v2.POST("/register", controllers.Register)
	  v2.POST("/login", controllers.Login)
	  v2.POST("/logout", controllers.Logout)
	  v2.POST("/refresh", controllers.Refresh)
  }

  return router
}