package common

import (
	"github.com/coreos/etcd/mvcc/mvccpb"
	"github.com/satori/go.uuid"
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
		res[string(data[i].Key)] = data[i].Key
	}
	return res
}
