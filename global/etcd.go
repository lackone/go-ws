package global

import (
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	EtcdClient *clientv3.Client
)

const (
	ETCD_WS_SERVERS = "/ws/servers/"
)

func InitEtcdClient() error {
	if EtcdClient == nil {
		var err error
		EtcdClient, err = clientv3.New(clientv3.Config{
			Endpoints:   EtcdSetting.Endpoints,
			DialTimeout: EtcdSetting.DialTimeout,
			Username:    EtcdSetting.Username,
			Password:    EtcdSetting.Password,
		})
		return err
	}
	return nil
}
