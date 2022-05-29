package repositories

import (
	"first-messanger/config"
	"first-messanger/models"
	"fmt"
)

func GetAllSendMessages(author string) (messages []models.SentMessage, err error) {
	messages = []models.SentMessage{}

	result := config.DB.Where("Author = ?", author).Find(&messages);

	// if result := config.DB.Where("Author = ?", author).Find(&messages); result.Error != nil {
	// 	err := result.Error 
	// 	fmt.Println(err)
	// 	return  nil, err;
	// }
	// fmt.Println("check result")
	// jsonBytes, err := json.Marshal(messages)
	// fmt.Println(err)
	// fmt.Println(string(jsonBytes))
	// fmt.Println(len(string(jsonBytes)))

	// if errors.Is(result.Error, gorm.ErrRecordNotFound) {
	// 	err := "not found"
	// 	fmt.Println(err)
	// 	return  nil, result.Error
	// }

	if result.Error != nil {
		err := result.Error 
		fmt.Println(err)
		return  nil, err;
	}


	return messages, nil;
}

func GetAllReceivedMessages(receiver string) (messages []models.ReceiveMessage, err error) {
	messages = []models.ReceiveMessage{}

	result := config.DB.Where("Receiver = ?", receiver).Find(&messages);

	if result.Error != nil {
		err := result.Error 
		return  nil, err;
   }

   return messages, nil;
}

func GetSendMessageById(id string) (message models.SentMessage, err error) {
	message = models.SentMessage{}

	result := config.DB.Where("Id = ?", id).Find(&message);

	if result.Error != nil {
		err := result.Error 
		return  models.SentMessage{}, err;
   }

	return  message, nil
}

func GetReceivedMessageById(id string) (message models.ReceiveMessage, err error) {
	message = models.ReceiveMessage{}

	result := config.DB.Where("Id = ?", id).Find(&message);

	if result.Error != nil {
		err := result.Error 
		return  models.ReceiveMessage{}, err;
   }

	return  message, nil
}

func SaveSendMessage(message *models.SentMessage) (err error) {
	
	result := config.DB.Create(message)
	return result.Error
}

func SaveReceivedMessage(message *models.ReceiveMessage) (err error) {
	
	result := config.DB.Create(message)
	return result.Error
}


func UpdateSendMessage(messageFromDB, newMessage *models.SentMessage) (string) {

	config.DB.Model(&messageFromDB).Updates(models.SentMessage{Author: newMessage.Author, Receiver: newMessage.Receiver, Body: newMessage.Body})

	return "ok"
}

func UpdateReceivedMessage(messageFromDB, newMessage *models.ReceiveMessage) (string) {

	config.DB.Model(&messageFromDB).Updates(models.ReceiveMessage{Author: newMessage.Author, Receiver: newMessage.Receiver, Body: newMessage.Body})

	return "ok"
}

func DeleteSendMessage(id string) (err error) {
	
	config.DB.Delete(&models.SentMessage{}, id)
	return nil
}

func DeleteReceivedMessage(id string) (err error) {
	
	config.DB.Delete(&models.ReceiveMessage{}, id)
	return nil
}

