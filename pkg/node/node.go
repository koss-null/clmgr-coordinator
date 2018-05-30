package node

import (
	"encoding/json"
	"github.com/google/logger"
	"myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/db"
	"strings"
	"time"
	"github.com/coreos/etcd/client"
	"regexp"
)

type NodeStatus string

const (
	ns_green NodeStatus = "green"
	ns_yellow = "yellow"
	ns_red = "red"
)

type Node struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
	IP     string  `json:"ip"`
	Status NodeStatus `json:"status,omitempty"`
	client db.Client
}

/*
	parseClusterHealth() is an inner function, which gets cluster health output
	from ETCD, and returns an array of healthy nodes
 */
func parseClusterHealth(resp string) (goodNodes []string) {
	respLines := strings.Split(resp, "\n")
	r := regexp.MustCompile("member ([a-f0-9])+ is healthy(.)+http://([0-9])+\\.([0-9])+\\.([0-9])+\\.([0-9])+")
	var subm [][]string
	for _, line := range respLines {
		subm = r.FindAllStringSubmatch(line, -1)
		goodNodes = append(goodNodes, strings.Split(subm[0][0], "http://")[1])
	}
	return
}

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

	// checking etcd state
	exec := common.NewExecutor()
	exec.SetOp([]string{"etcdctl", "cluster-health"})
	go func() {
		for {
			time.Sleep(10 * time.Second)
			resp, err := exec.Exec()
			if err != nil {
				logger.Errorf("Error during cluster-health: %s", err.Error())
				continue
			}
			healthyNodes := parseClusterHealth(resp)
			n.client.Get()
		}
	}()
}
