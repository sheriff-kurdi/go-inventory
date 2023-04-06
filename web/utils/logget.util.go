package utils

import (
	"fmt"

	"go.uber.org/zap"
)

func Logger() *zap.Logger {
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Print("logger failed")
	}
	defer logger.Sync() // flushes buffer, if any
	return logger
}
