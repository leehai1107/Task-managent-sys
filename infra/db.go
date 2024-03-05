package infra

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbSingleton *gorm.DB

func ConnectDB() {
	logger, _ := zap.NewProduction()
	defer logger.Sync()

	var (
		host     = os.Getenv("POSTGRES_HOST")
		port, _  = strconv.Atoi(os.Getenv("POSTGRES_PORT"))
		user     = os.Getenv("POSTGRES_USER")
		password = os.Getenv("POSTGRES_PASSWORD")
		dbname   = os.Getenv("POSTGRES_DBNAME")
	)

	pgsqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable TimeZone=Asia/Shanghai",
		host, port, user, password, dbname)

	db, err := gorm.Open(postgres.Open(pgsqlInfo), &gorm.Config{})
	if err != nil {
		logger.Error("Unable to ConnectDB postgres", zap.Error(err))
	}

	dbSingleton = db
}

func GetDB() *gorm.DB {
	if dbSingleton == nil {
		log.Printf("Connect to DB not setup")
	}
	return dbSingleton
}

func CloseDB() error {
	sqlDB, err := dbSingleton.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}
