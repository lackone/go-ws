package global

import (
	"fmt"
	"github.com/lackone/go-ws/pkg/etcd"
	clientv3 "go.etcd.io/etcd/client/v3"
)

var (
	EtcdClient *clientv3.Client
	EtcdKV     *etcd.EtcdKV
	WsEtcdDis  *etcd.EtcdDiscover
)

const (
	ETCD_WS_SERVICES = "/ws/services/"
	ETCD_WS_ACCOUNTS = "/ws/accounts/"
	ETCD_WS_MACHINES = "/ws/machines/"
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
		if err != nil {
			panic(err)
		}
	}
	if EtcdKV == nil {
		EtcdKV = etcd.NewEtcdKV(EtcdClient)
	}
	if WsEtcdDis == nil {
		var err error
		WsEtcdDis, err = etcd.NewEtcdDiscover(EtcdClient, ETCD_WS_SERVICES)
		if err != nil {
			panic(err)
		}
		fmt.Println(WsEtcdDis.ServiceList())
	}
	return nil
}
