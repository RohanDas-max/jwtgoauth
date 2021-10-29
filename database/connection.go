package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"github.com/rohandas-max/shoping-site/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connection() {

	godotenv.Load(".env")
	envMap, envErr := godotenv.Read(".env")

	if envErr != nil {
		panic(envErr)
	}

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", envMap["HOST"], envMap["USER"], envMap["PASSWORD"], envMap["NAME"], envMap["DBPORT"])

	//* Starting/Opening Database connection via gorm
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("DB Connected Successfully!")
	}

	DB = connection

	connection.AutoMigrate(&models.User{})

}
