package etcd

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdKV struct {
	client *clientv3.Client
}

func NewEtcdKV(client *clientv3.Client) *EtcdKV {
	return &EtcdKV{
		client: client,
	}
}

func (e *EtcdKV) Put(key string, val string) error {
	kv := clientv3.NewKV(e.client)
	_, err := kv.Put(context.Background(), key, val)
	return err
}

func (e *EtcdKV) Get(key string) (*clientv3.GetResponse, error) {
	kv := clientv3.NewKV(e.client)
	get, err := kv.Get(context.Background(), key)
	return get, err
}
