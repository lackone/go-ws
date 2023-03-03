package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lackone/go-ws/global"
	"github.com/spf13/viper"
	"strings"
	"time"
)

var sections = make(map[string]interface{})

type WsSetting struct {
	HttpPort         int
	HttpReadTimeout  time.Duration
	HttpWriteTimeout time.Duration
	WsPort           int
	GrpcPort         int
	IsCluster        bool
}

type EtcdSetting struct {
	DialTimeout time.Duration
	Endpoints   []string
	Username    string
	Password    string
}

type LogSetting struct {
	Folder     string
	File       string
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type Setting struct {
	vp *viper.Viper
}

func NewSetting(paths ...string) (*Setting, error) {
	v := viper.New()
	v.SetConfigName("app")
	v.SetConfigType("yaml")
	for _, path := range paths {
		if path != "" {
			v.AddConfigPath(path)
		}
	}
	err := v.ReadInConfig()
	if err != nil {
		return nil, err
	}
	s := &Setting{
		vp: v,
	}
	s.WatchConfig()
	return s, nil
}

func (s *Setting) WatchConfig() {
	go func() {
		s.vp.WatchConfig()
		s.vp.OnConfigChange(func(e fsnotify.Event) {
			s.ReloadSection()
		})
	}()
}

func (s *Setting) ReadSection(k string, v interface{}) error {
	err := s.vp.UnmarshalKey(k, v)
	if err != nil {
		return err
	}

	if _, ok := sections[k]; !ok {
		sections[k] = v
	}

	return nil
}

func (s *Setting) ReloadSection() error {
	for k, v := range sections {
		err := s.ReadSection(k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func InitSetting(conf string) error {
	s, err := NewSetting(strings.Split(conf, ",")...)
	err = s.ReadSection("ws", &global.WsSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("etcd", &global.EtcdSetting)
	if err != nil {
		return err
	}
	err = s.ReadSection("log", &global.LogSetting)
	if err != nil {
		return err
	}
	global.WsSetting.HttpReadTimeout *= time.Second
	global.WsSetting.HttpWriteTimeout *= time.Second
	global.EtcdSetting.DialTimeout *= time.Second
	return nil
}
