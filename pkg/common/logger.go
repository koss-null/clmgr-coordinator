package common

import (
	"github.com/google/logger"
	"os"
)

func InitLogger() error {
	lf, err := os.OpenFile(Config.LogCoordPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0660)
	if err != nil {
		logger.Fatalf("Failed to open log file: %v", err)
	}

	logger.Init("Logger", false, true, lf)

	return nil
}
