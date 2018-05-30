package common

import (
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/satori/go.uuid"
	"sync"
	"myproj.com/clmgr-coordinator/config"
	"strings"
)

var hostname string
var once sync.Once

func GetHostname() string {
	if hostname == "" {
		once.Do(func() {
			id, _ := uuid.NewV1()
			hostname = id.String()
		})
	}
	return hostname
}

func KV2Map(data []*mvccpb.KeyValue) map[string][]byte {
	res := make(map[string][]byte)
	for i := range data {
		res[string(data[i].Key)] = data[i].Value
	}
	return res
}

func GetIpFromETCD() string {
	// todo: think how to fix it
	command := config.Config.HNPath
	exec := NewExecutor()
	exec.SetOp([]string{"/bin/bash", "-c", command})
	res, _ := exec.Exec()
	return strings.Trim(res, "\n")
}
