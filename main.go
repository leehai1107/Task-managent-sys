package main

import (
	"log"

	"scheduleSystem/infra"
	"scheduleSystem/router"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can not read .env file.")
	}

	infra.ConnectDB()
	defer infra.CloseDB()

	routes := router.NewRouter()
	routes.Register().Run("localhost:8080")
}
