package logger

import (
	"github.com/zondax/golem/pkg/logger"
)

func GetSafeLogger(zapLogger *logger.Logger) *logger.Logger {
	if zapLogger == nil {
		productionLogger := logger.NewLogger(logger.Config{
			Level:    "info",
			Encoding: "json",
		}) // default
		return productionLogger
	}

	return zapLogger
}
