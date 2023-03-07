package global

import (
	"github.com/lackone/go-ws/pkg/setting"
	"strings"
	"time"
)

var (
	HttpSetting      *setting.HttpSetting
	WsSetting        *setting.WsSetting
	GrpcSetting      *setting.GrpcSetting
	SnowflakeSetting *setting.SnowflakeSetting
	EtcdSetting      *setting.EtcdSetting
	LogSetting       *setting.LogSetting
)

func InitSetting(conf string) error {
	s, err := setting.NewSetting(strings.Split(conf, ",")...)
	err = s.ReadSection("http", &HttpSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("ws", &WsSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("grpc", &GrpcSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("snowflake", &SnowflakeSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("etcd", &EtcdSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("log", &LogSetting)
	if err != nil {
		return err
	}
	HttpSetting.HttpReadTimeout *= time.Second
	HttpSetting.HttpWriteTimeout *= time.Second
	WsSetting.HttpReadTimeout *= time.Second
	WsSetting.HttpWriteTimeout *= time.Second
	WsSetting.HeartbeatInterval *= time.Second
	WsSetting.ReadDeadline *= time.Second
	WsSetting.WriteDeadline *= time.Second
	EtcdSetting.DialTimeout *= time.Second
	return nil
}

func IsCluster() bool {
	return WsSetting.IsCluster
}
