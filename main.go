package main

import (
	"log"

	"github.com/joho/godotenv"
	_ "github.com/leehai1107/Task-managent-sys/docs"
	"github.com/leehai1107/Task-managent-sys/infra"
	"github.com/leehai1107/Task-managent-sys/router"
)

// @title           Swagger Task-Management-System API
// @version         1.0
// @description     This is Task-Management-System server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
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
