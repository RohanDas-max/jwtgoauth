package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/rohandas-max/shoping-site/database"
	"github.com/rohandas-max/shoping-site/routes"
)

func main() {

	app := fiber.New()
	routes.Routes(app)

	database.Connection()

	//* securing port and listening
	godotenv.Load(".env")
	envMap, envErr := godotenv.Read(".env")
	if envErr != nil {
		fmt.Println("error loading environment variable", envErr)
	} else {
		app.Listen(envMap["PORT"])
	}

}
