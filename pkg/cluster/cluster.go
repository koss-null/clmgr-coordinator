package cluster

import (
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/node"
	"myproj.com/clmgr-coordinator/pkg/db"
	"strings"
	. "myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/resource"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"encoding/json"
	"errors"
)

type (
	cluster struct {
		config    Config
		nodePool  node.Pool
		agentPool resource.Pool
		clnt      db.Client
	}

	Cluster interface {
		Start(chan error)
		AddConfig(*Config) error
		Stop(chan error)
		GetConfig() Config
	}
)

var Current = New()

func New() Cluster {
	return &cluster{
		DefaultConfig(),
		node.NewPool(),
		resource.NewPool(),
		nil,
	}
}

func (c *cluster) Start(errChan chan error) {
	logger.Info("Starting cluster")

	// creating node only with hostname
	c.nodePool.Add(node.Node{
		Name: GetHostname(),
	})

	c.clnt = db.NewClient()
	result, err := c.clnt.Get(ClmgrKey)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("Got result %s", result)

	data, err := json.Marshal(c.config)
	if err != nil {
		errChan <- err
		return
	}
	err = c.clnt.Set(strings.Join([]string{ClmgrKey, "config"}, "/"), string(data))
	if err != nil {
		errChan <- err
		return
	}

	watchClusterChan := c.clnt.Watch(strings.Join([]string{ClmgrKey, "config"}, "/"), nil)

	// watching cluster config changes
	go func() {
		for r := range watchClusterChan {
			logger.Infof("Got key changing %+v", r)
			for _, e := range r.Events {
				if e.IsModify() || e.IsCreate() {
					data = e.Kv.Value
					err := json.Unmarshal(data, &(c.config))
					if err != nil {
						errChan <- err
					}
				}
			}
		}
	}()

	watchChan := c.clnt.Watch(strings.Join([]string{ClmgrKey, GetHostname()}, "/"), nil)
	// watching nodes changing
	go func() {
		for r := range watchChan {
			logger.Infof("Got key changing %+v", r)
			for _, e := range r.Events {
				// todo: check if it's work, looks bad
				if e.Type == mvccpb.DELETE {
					logger.Info("This node was deleted from cluster")
					close(errChan)
					return
				}
			}
		}
	}()
}

func (c *cluster) AddConfig(config *Config) error {
	logger.Info("Adding config to cluster")
	if !config.Check() {
		return errors.New("config contains bad value")
	}

	data, err := json.Marshal(config)
	if err != nil {
		return err
	}
	err = c.clnt.Set(strings.Join([]string{ClmgrKey, "config"}, "/"), string(data))
	if err != nil {
		return err
	}
	c.config = *config
	return nil
}

func (c *cluster) Stop(errChan chan error) {
	close(errChan)
}

func (c *cluster) GetConfig() Config {
	return c.config
}
