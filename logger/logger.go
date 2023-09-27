package logger

import "go.uber.org/zap"

func GetSafeLogger(zapLogger *zap.Logger) *zap.Logger {
	if zapLogger == nil {
		productionLogger, _ := zap.NewProduction() // default
		return productionLogger
	}

	return zapLogger
}
