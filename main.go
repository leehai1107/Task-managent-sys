package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/leehai1107/Task-managent-sys/docs"
	"github.com/leehai1107/Task-managent-sys/infra"
	"github.com/leehai1107/Task-managent-sys/router"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Can not read .env file.")
	}

	infra.ConnectDB()
	defer infra.CloseDB()

	routes := router.NewRouter()
	routes.Register().Run(":8080")
}
