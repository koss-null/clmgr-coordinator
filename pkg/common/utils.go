package common

import (
	"github.com/coreos/etcd/mvcc/mvccpb"
	"math/rand"
	"myproj.com/clmgr-coordinator/config"
	"strconv"
	"strings"
	"sync"
	"time"
)

var hostname string
var once sync.Once

var hostNames = []string{"Donkey", "Aardwolf", "Admiral", "Vanessa", "Adouri", "Limnocorax", "Snycerus", "Paraxerus",
	"Aonyx", "Anhinga", "Panthera", "Owl", "Mabuya"}
var hostPrefixes = []string{"Big", "Little", "Fancy", "Silly", "Funny", "Huge", "Dancing", "Suspicious"}

func GetHostname() string {
	if hostname == "" {
		once.Do(func() {
			r := rand.New(rand.NewSource(time.Now().UnixNano()))
			prefix := r.Intn(len(hostPrefixes))
			name := r.Intn(len(hostNames))
			id := r.Intn(1024000)
			hostname = strings.Join([]string{hostPrefixes[prefix], hostNames[name], strconv.Itoa(id)}, "-")
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
