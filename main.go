package main

import (
	"learn/config"
	"learn/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	file, err := os.OpenFile("storage/logs/go.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalf("Failed to open log file: %v", err)
    }
    log.SetOutput(file)
    // log Gin's logs to the file
    gin.DefaultWriter = file

    config.ConnectDB()

    r := routes.SetupRouter()

    r.Run(":" + os.Getenv("PORT"))
}
