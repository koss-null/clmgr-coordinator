package node

import (
	"encoding/json"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/db"
	"strings"
)

type Node struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
	IP     string  `json:"ip"`
	client db.Client
}

func (n *Node) Watch() {
	if n.client == nil {
		n.client = db.NewClient()
	}
	watchClusterChan := n.client.Watch(strings.Join([]string{common.ClmgrKey, "nodes", n.Name}, "/"), nil)

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
