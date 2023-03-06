package global

import (
	"github.com/lackone/go-ws/pkg/logger"
	"go.uber.org/zap"
)

var (
	Logger *logger.Logger
)

func InitLogger() error {
	if Logger == nil {
		Logger = logger.NewLogger(LogSetting.Folder, LogSetting.File, zap.InfoLevel, LogSetting.MaxSize, LogSetting.MaxBackups, LogSetting.MaxAge, LogSetting.Compress)
	}
	return nil
}
