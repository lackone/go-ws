package setting

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"time"
)

var sections = make(map[string]interface{})

type HttpSetting struct {
	Mode             string
	HttpPort         int
	HttpReadTimeout  time.Duration
	HttpWriteTimeout time.Duration
	IsTLS            bool
	TLSCertFile      string
	TLSKeyFile       string
}

type WsSetting struct {
	Mode              string
	WsPort            int
	HttpReadTimeout   time.Duration
	HttpWriteTimeout  time.Duration
	ReadLimit         int64
	HeartbeatInterval time.Duration
	ReadDeadline      time.Duration
	WriteDeadline     time.Duration
	PingMessage       string
	IsCluster         bool
	IsTLS             bool
	TLSCertFile       string
	TLSKeyFile        string
	ReadBufferSize    int
	WriteBufferSize   int
}

type GrpcSetting struct {
	GrpcPort int
}

type SnowflakeSetting struct {
	NodeId int64
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
