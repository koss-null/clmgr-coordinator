package db

import (
	"context"
	"github.com/coreos/etcd/clientv3"
	"github.com/google/logger"
	. "myproj.com/clmgr-coordinator/pkg/common"
	"sync"
	"time"
)

type (
	dbclient struct{}

	ttlclient struct {
		client *clientv3.Client
		ID     clientv3.LeaseID
	}

	Client interface {
		Set(key string, value string) error
		Get(key string) (map[string][]byte, error)
		Remove(key string) error
		Watch(key string, wo ...clientv3.OpOption) clientv3.WatchChan
		Close()
	}
)

const (
	etcdEndpoint = "http://127.0.0.1:2379"
)

var once sync.Once
var etcdClient *clientv3.Client
var kvc clientv3.KV

func NewClient() Client {
	initClient()
	return &dbclient{}
}

/*
	initClient() performs singletone pattern to create client and kapi instances
*/
func initClient() error {
	var err error
	once.Do(func() {
		cfg := clientv3.Config{
			Endpoints:   []string{etcdEndpoint},
			DialTimeout: 2 * time.Second,
		}
		etcdClient, err = clientv3.New(cfg)
		if err != nil {
			logger.Errorf("Can't create new client, err: %s", err.Error())
			return
		}
		kvc = clientv3.NewKV(etcdClient)
	})
	return err
}

func (c *dbclient) Set(key string, value string) error {
	logger.Infof("Setting %s key with %s value", key, value)
	resp, err := kvc.Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	logger.Infof("Set is done. Metadata is %q", resp)
	return nil
}

func (c *dbclient) Get(key string) (map[string][]byte, error) {
	logger.Infof("Getting %s key", key)
	resp, err := kvc.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return map[string][]byte{}, err
	}
	return KV2Map(resp.Kvs), nil
}

func (c *dbclient) Remove(key string) error {
	logger.Infof("Removing etcd key %s", key)
	_, err := kvc.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	logger.Infof("Remove is done")
	return nil
}

const (
	Action_get    = "get"
	Action_set    = "set"
	Action_delete = "delete"
	Action_upd    = "update"
	Action_create = "create"
	Action_cas    = "compareAndSwap"
	Action_cad    = "compareAndDelete"
	Action_expire = "expire"
)

type Response struct {
	Action string
	Value  string
}

func (c *dbclient) Watch(key string, wo ...clientv3.OpOption) clientv3.WatchChan {
	return etcdClient.Watcher.Watch(context.Background(), key, wo...)
}

func (c *dbclient) Close() {
	logger.Error("You are trying to close shared client. Please don't do it")
}

/*
	GetTTLClient() returns custom client for careful usage =)
	it's often used to deal with ttl
	ATTENTION
	This client need to be closed by yourself
*/
func GetTTLClient(duration int) Client {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcdEndpoint},
		DialTimeout: time.Second * time.Duration(duration),
	})
	if err != nil {
		logger.Errorf("can't create custom client, err: %s", err.Error())
		return nil
	}

	resp, err := client.Grant(context.Background(), int64(duration))
	if err != nil {
		logger.Errorf("can't create custom client, err: %s", err.Error())
		return nil
	}

	return &ttlclient{
		client,
		resp.ID,
	}
}

func (c *ttlclient) Set(key string, value string) error {
	_, err := c.client.Put(context.TODO(), key, value, clientv3.WithLease(c.ID))
	return err
}

func (c *ttlclient) Get(key string) (map[string][]byte, error) {
	logger.Infof("Getting %s key", key)
	resp, err := c.client.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return map[string][]byte{}, err
	}
	return KV2Map(resp.Kvs), nil
}

func (c *ttlclient) Remove(key string) error {
	logger.Infof("Removing etcd key %s", key)
	_, err := c.client.Delete(context.Background(), key)
	if err != nil {
		return err
	}
	logger.Infof("Remove is done")
	return nil
}

func (c *ttlclient) Watch(key string, wo ...clientv3.OpOption) clientv3.WatchChan {
	return c.client.Watcher.Watch(context.Background(), key, wo...)
}

func (c *ttlclient) Close() {
	c.client.Close()
}
