package logger

import (
	"github.com/zondax/golem/pkg/logger"
)

func GetSafeLogger(lg *logger.Logger) *logger.Logger {
	if lg == nil {
		productionLogger := logger.NewLogger(logger.Config{
			Level:    "info",
			Encoding: "json",
		}) // default
		return productionLogger
	}

	return lg
}
