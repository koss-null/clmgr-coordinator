package cluster

import (
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/node"
	"myproj.com/clmgr-coordinator/pkg/db"
	"strings"
	. "myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/resource"
	"github.com/coreos/etcd/mvcc/mvccpb"
)

type (
	cluster struct {
		config    Config
		nodePool  node.Pool
		agentPool resource.Pool
	}

	Cluster interface {
		Start(chan error)
		AddConfig(*Config) error
		Stop(chan error)
	}
)

var Current = New()

func New() Cluster {
	return &cluster{Config{}, node.NewPool(), resource.NewPool()}
}

func (c *cluster) Start(errChan chan error) {
	logger.Info("Starting cluster")

	// creating node only with hostname
	c.nodePool.Add(node.Node{
		Name: GetHostname(),
	})

	clnt := db.NewClient()
	result, err := clnt.Get(ClmgrKey)
	if err != nil {
		logger.Error(err.Error())
		return
	}
	logger.Infof("Got result %s", result)

	watchChan := clnt.Watch(strings.Join([]string{ClmgrKey, GetHostname()}, "/"), nil)
	// getting all existing nodes

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
	c.config = *config
	return nil
}

func (c *cluster) Stop(errChan chan error) {
	close(errChan)
}
