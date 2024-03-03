package exec

import (
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func Run() error {
  logger,_ := zap.NewProduction()
  defer logger.Sync()
	err := godotenv.Load()
  if err != nil {
    return err
  }

  return nil

}
