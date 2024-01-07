package external

import (
	"log"

	"go.uber.org/zap"
)

func GeZapLogger() *zap.SugaredLogger {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	return logger.Sugar()
}
