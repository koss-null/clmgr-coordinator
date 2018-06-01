package common

import (
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/satori/go.uuid"
	"myproj.com/clmgr-coordinator/config"
	"strings"
	"sync"
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
	command := config.Config.HNPath
	exec := NewExecutor()
	exec.SetOp([]string{"/bin/bash", "-c", command})
	iter := 0
	res := ""
	for {
		var err error
		res, err = exec.Exec()
		iter++
		if iter > 5 || err == nil {
			break
		}
	}
	return strings.Trim(res, "\n")
}
