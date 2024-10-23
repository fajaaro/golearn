package main

import (
	"learn/config"
	"learn/routes"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

    config.ConnectDB()

    r := routes.SetupRouter()

    r.Run(":" + os.Getenv("PORT"))
}
