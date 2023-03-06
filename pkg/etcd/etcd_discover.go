package etcd

import (
	"context"
	"go.etcd.io/etcd/api/v3/mvccpb"
	clientv3 "go.etcd.io/etcd/client/v3"
	"sync"
)

type EtcdDiscover struct {
	client *clientv3.Client
	list   map[string]string
	lock   sync.Mutex
}

func NewEtcdDiscover(client *clientv3.Client, prefix string) (*EtcdDiscover, error) {
	dis := &EtcdDiscover{
		client: client,
		list:   make(map[string]string),
		lock:   sync.Mutex{},
	}

	if err := dis.watchService(prefix); err != nil {
		return nil, err
	}

	go dis.watch(prefix)

	return dis, nil
}

// 监控服务
func (e *EtcdDiscover) watchService(prefix string) error {
	kv := clientv3.NewKV(e.client)

	get, err := kv.Get(context.Background(), prefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, ev := range get.Kvs {
		if ev.Value != nil {
			e.AddService(string(ev.Key), string(ev.Value))
		}
	}

	return nil
}

// 监控key
func (e *EtcdDiscover) watch(prefix string) {
	watcher := clientv3.NewWatcher(e.client)

	watch := watcher.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for v := range watch {
		for _, ev := range v.Events {
			switch ev.Type {
			case mvccpb.PUT:
				e.AddService(string(ev.Kv.Key), string(ev.Kv.Value))
			case mvccpb.DELETE:
				e.DelService(string(ev.Kv.Key))
			}
		}
	}
}

// 添加服务
func (e *EtcdDiscover) AddService(key, value string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	e.list[key] = value
}

// 删除服务
func (e *EtcdDiscover) DelService(key string) {
	e.lock.Lock()
	defer e.lock.Unlock()
	delete(e.list, key)
}

// 服务列表
func (e *EtcdDiscover) ServiceList() []string {
	e.lock.Lock()
	defer e.lock.Unlock()
	list := make([]string, 0)
	for _, v := range e.list {
		list = append(list, v)
	}
	return list
}
