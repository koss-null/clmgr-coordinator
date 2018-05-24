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

	Client interface {
		Set(key string, value string) error
		Get(key string) (map[string][]byte, error)
		Remove(key string) error
		Watch(key string, wo []clientv3.OpOption) clientv3.WatchChan
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
	resp, err := kvc.Put(context.Background(), key, value, nil)
	if err != nil {
		return err
	}
	logger.Infof("Set is done. Metadata is %q", resp)
	return nil
}

func (c *dbclient) Get(key string) (map[string][]byte, error) {
	logger.Infof("Getting %s key", key)
	resp, err := kvc.Get(context.Background(), key, nil)
	if err != nil {
		return map[string][]byte{}, err
	}
	return KV2Map(resp.Kvs), nil
}

func (c *dbclient) Remove(key string) error {
	logger.Infof("Removing etcd key %s", key)
	_, err := kvc.Delete(context.Background(), key, nil)
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

func (c *dbclient) Watch(key string, wo []clientv3.OpOption) clientv3.WatchChan {
	return etcdClient.Watcher.Watch(context.Background(), key, wo...)
}
