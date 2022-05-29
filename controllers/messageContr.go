package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"

	"first-messanger/models"
	"first-messanger/repositories"
)

// ADD NEW MESSAGE

func SendMessage(context *gin.Context) {
	sendMessage := models.SentMessage{}
	receivedMessage := models.ReceiveMessage{}

	err := context.BindJSON(&sendMessage)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		fmt.Errorf("Error on save", err.Error())
	}

	receivedMessage.Author = sendMessage.Author
	receivedMessage.Receiver = sendMessage.Receiver
	receivedMessage.Body = sendMessage.Body

	repositories.SaveSendMessage(&sendMessage)
	repositories.SaveReceivedMessage(&receivedMessage)

	context.JSON(200, sendMessage)
}

// GET ALL MESSAGES

func GetAllReceivedMessages(context *gin.Context) {
	receiver := context.Param("receiver")

	messages, err := repositories.GetAllReceivedMessages(receiver)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
	}
	context.JSON(200, messages)
}

func GetAllSendMessages(context *gin.Context) {
	author := context.Param("author")

	messages, err := repositories.GetAllSendMessages(author)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err.Error())
	}
	context.JSON(200, messages)
}

// UPDATE

func UpdateSendMessage(context *gin.Context){
	id := context.Param("id")

	messageFromDB, err := repositories.GetSendMessageById(id)

	newMessage := models.SentMessage{}
	err = context.BindJSON(&newMessage)

	if err != nil {
		context.AbortWithError(http.StatusBadRequest, err)
		fmt.Println(err.Error())

	}

	res:= repositories.UpdateSendMessage(&messageFromDB, &newMessage)


   context.JSON(200, res)
}

func UpdateReceivedMessage(context *gin.Context){
	id := context.Param("id")

	messageFromDB, err := repositories. GetReceivedMessageById(id)

	newMessage := models.ReceiveMessage{}

	err = context.BindJSON(&newMessage)

	if err != nil {
		context.AbortWithError(500, err)
		fmt.Println(err.Error())
	}

	res:= repositories.UpdateReceivedMessage(&messageFromDB, &newMessage)

   context.JSON(200, res)
} 

// DELETE

func DeleteReceivedMessage(context *gin.Context) {
	id := context.Param("id")

	messages := repositories.DeleteReceivedMessage(id)

	context.JSON(200, messages)
}

func DeleteSendMessage(context *gin.Context) {
	id := context.Param("id")

	messages := repositories.DeleteSendMessage(id)

	context.JSON(200, messages)
}