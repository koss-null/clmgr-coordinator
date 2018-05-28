package connection

import (
	"encoding/json"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/db"
	"net"
	"strings"
)

type (
	node struct {
		HN     string `json:"hostname"`
		Ip     net.IP `json:"ip,omitempty"`
		client db.Client
	}

	Node interface {
		Hostname() string
		IP() net.IP
	}
)

func NewNode(hostName string, ip net.IP) Node {
	return &node{hostName, ip, db.NewClient()}
}

func (n *node) Hostname() string {
	return n.HN
}

func (n *node) IP() net.IP {
	return n.Ip
}

func (n *node) Watch() {
	watchClusterChan := n.client.Watch(strings.Join([]string{common.ClmgrKey, n.HN}, "/"), nil)

	// watching cluster config changes
	go func() {
		for r := range watchClusterChan {
			logger.Infof("Got node changing %+v", r)
			for _, e := range r.Events {
				if e.IsModify() || e.IsCreate() {
					data := e.Kv.Value
					json.Unmarshal(data, n)
				}
			}
		}
	}()
}
