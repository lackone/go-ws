package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdRegister struct {
	client         *clientv3.Client
	lease          clientv3.Lease
	leaseId        clientv3.LeaseID
	leaseKeepAlive <-chan *clientv3.LeaseKeepAliveResponse
}

func NewEtcdRegister(client *clientv3.Client, ttl int64) (*EtcdRegister, error) {
	reg := &EtcdRegister{
		client: client,
	}

	if err := reg.setLease(ttl); err != nil {
		return nil, err
	}

	go reg.listenLease()

	return reg, nil
}

// 设置租约
func (e *EtcdRegister) setLease(ttl int64) error {
	newLease := clientv3.NewLease(e.client)

	lease, err := newLease.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}

	alive, err := newLease.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		return err
	}

	e.lease = newLease
	e.leaseId = lease.ID
	e.leaseKeepAlive = alive

	return nil
}

// 监听续租情况
func (e *EtcdRegister) listenLease() {
	for {
		v, ok := <-e.leaseKeepAlive
		if !ok {
			fmt.Println("关闭续租")
			return
		}
		fmt.Println("续租成功", v)
	}
}

// 注册服务
func (e *EtcdRegister) RegService(key, val string) error {
	kv := clientv3.NewKV(e.client)
	_, err := kv.Put(context.Background(), key, val, clientv3.WithLease(e.leaseId))
	return err
}

// 删除服务
func (e *EtcdRegister) UnRegService(key string) error {
	kv := clientv3.NewKV(e.client)
	_, err := kv.Delete(context.Background(), key)
	return err
}

// 撤销租约
func (e *EtcdRegister) RevokeLease() error {
	_, err := e.lease.Revoke(context.Background(), e.leaseId)
	return err
}
