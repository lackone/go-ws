package global

import (
	"github.com/lackone/go-ws/pkg/logger"
	"github.com/lackone/go-ws/pkg/setting"
)

var (
	WsSetting   *setting.WsSetting
	EtcdSetting *setting.EtcdSetting
	LogSetting  *setting.LogSetting
	Logger      *logger.Logger
)
