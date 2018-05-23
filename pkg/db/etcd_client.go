package db

import (
	"context"
	"github.com/coreos/etcd/client"
	"github.com/google/logger"
	"sync"
	"time"
)

type (
	dbclient struct{}

	Client interface {
		Set(key string, value string) error
		Get(key string) (string, error)
		Watch(key string, wo client.WatcherOptions) (chan string, chan error)
	}
)

const (
	etcdEndpoint = "http://127.0.0.1:2379"
)

var etcdClient client.Client
var kapi client.KeysAPI
var once sync.Once

func NewClient() Client {
	if etcdClient == nil {
		initClient()
	}
	return &dbclient{}
}

/*
	initClient() performs singletone pattern to create client and kapi instances
*/
func initClient() error {
	var err error
	once.Do(func() {
		cfg := client.Config{
			Endpoints: []string{etcdEndpoint},
			Transport: client.DefaultTransport,
			// set timeout per request to fail fast when the target endpoint is unavailable
			HeaderTimeoutPerRequest: time.Second,
		}
		etcdClient, err = client.New(cfg)
		if err != nil {
			logger.Errorf("Can't create new client, err: %s", err.Error())
			return
		}
		kapi = client.NewKeysAPI(etcdClient)
	})
	return err
}

func (c *dbclient) Set(key string, value string) error {
	logger.Infof("Setting %s key with %s value", key, value)
	resp, err := kapi.Set(context.Background(), key, value, nil)
	if err != nil {
		return err
	}
	logger.Infof("Set is done. Metadata is %q", resp)
	return nil
}

func (c *dbclient) Get(key string) (string, error) {
	logger.Infof("Getting %s key", key)
	resp, err := kapi.Get(context.Background(), "/foo", nil)
	if err != nil {
		return "", err
	}
	return resp.Node.Value, nil
}

func (c *dbclient) Watch(key string, wo client.WatcherOptions) (chan string, chan error) {
	watcher := kapi.Watcher(key, &wo)
	response, errCh := make(chan string), make(chan error)
	go func() {
		for {
			resp, err := watcher.Next(context.Background())
			if err != nil {
				errCh <- err
				break
			}
			response <- resp.Node.Value
		}
	}()
	return response, errCh
}
