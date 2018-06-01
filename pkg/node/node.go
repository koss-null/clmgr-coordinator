package node

import (
	"encoding/json"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/db"
	"strings"
)

type NodeStatus string

const (
	ns_green  NodeStatus = "green"
	ns_yellow            = "yellow"
	ns_red               = "red"
)

type Node struct {
	Name   string     `json:"name"`
	Labels []Label    `json:"labels"`
	IP     string     `json:"ip"`
	Status NodeStatus `json:"status"`
	client db.Client
}

/*
	parseClusterHealth() is an inner function, which gets cluster health output
	from ETCD, and returns an array of healthy nodes
*/
// func parseClusterHealth(resp string) (goodNodes []string) {
// 	respLines := strings.Split(resp, "\n")
// 	r := regexp.MustCompile("member ([a-f0-9])+ is healthy(.)+http://([0-9])+\\.([0-9])+\\.([0-9])+\\.([0-9])+")
// 	var subm [][]string
// 	for _, line := range respLines {
// 		subm = r.FindAllStringSubmatch(line, -1)
// 		goodNodes = append(goodNodes, strings.Split(subm[0][0], "http://")[1])
// 	}
// 	return
// }

func (n *Node) Watch() {
	if n.client == nil {
		n.client = db.NewClient()
	}
	watchClusterChan := n.client.Watch(strings.Join([]string{common.ClmgrKey, "nodes", n.Name}, "/"))

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

	// checking nodes state
	watchNodeKeyChan := n.client.Watch(strings.Join([]string{common.ClmgrKey, "nodes", n.Name, common.IsAliveKey}, "/"))
	go func() {
		for r := range watchNodeKeyChan {
			for _, e := range r.Events {
				if e.IsModify() || e.IsCreate() {
					continue
				}
				// if we are here, that means that the key was expired
				logger.Infof("Found live key event, %+v", e)
				n.Status = ns_red
				data, err := json.Marshal(n)
				if err != nil {
					logger.Errorf("failed to marshall node into json, err: %s", err.Error())
					continue
				}
				err = n.client.Set(strings.Join([]string{common.ClmgrKey, "nodes", n.Name}, "/"), string(data))
				if err != nil {
					logger.Errorf("failed to change node state ¡in etcd, err: %s", err.Error())
					continue
				}
			}
		}
	}()
}
