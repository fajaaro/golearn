package main

import (
	"learn/config"
	"learn/routes"
)

func main() {
    config.ConnectDB()

    r := routes.SetupRouter()

    r.Run(":8080")
}
