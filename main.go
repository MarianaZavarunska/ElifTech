package main

import (
	"database/sql"
	"first-messanger/config"
	"first-messanger/models"
	"first-messanger/routers"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
  configDB();
  setupRouter();
}

func configDB () {
	err := godotenv.Load()
	if err != nil {
	  fmt.Println("Error loading .env file")
	}
  
	host := os.Getenv("HOST")
	user := os.Getenv("USER")
	dbname := os.Getenv("DBNAME")
	password := os.Getenv("PASSWORD")
	sslmode := os.Getenv("SSLMODE")
	port := os.Getenv("PORT")

	// sqlDB, err := sql.Open("pgx", "host=host user=user dbname=dbname password=password sslmode=sslmode port=port"
	sqlDB, err := sql.Open("pgx", "host="+host+" user="+user+" dbname="+dbname+" password="+password+ " sslmode="+sslmode+" port="+port)
  
	if err != nil {
		  fmt.Println("statuse 1: ", err)
	  }
  
	config.DB, err = gorm.Open(postgres.New(postgres.Config{
	  Conn: sqlDB,
	}), &gorm.Config{})
	
	if err != nil {
		  fmt.Println("statuse 2: ", err)
	  }
  
	config.DB.AutoMigrate(&models.SentMessage{})
	config.DB.AutoMigrate(&models.ReceiveMessage{})
	config.DB.AutoMigrate(&models.User{})
}

func setupRouter () {
	router := routers.SetupRouter()
  
	router.Run(":8080")
}
