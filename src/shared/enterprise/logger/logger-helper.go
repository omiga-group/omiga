package logger

import (
	"log"

	"go.uber.org/zap"
)

func CreateLogger() *zap.SugaredLogger {
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	return logger.Sugar()
}
