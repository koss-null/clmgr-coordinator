package node

import (
	"encoding/json"
	"github.com/google/logger"
	. "myproj.com/clmgr-coordinator/pkg/common"
	"myproj.com/clmgr-coordinator/pkg/db"
	"strings"
	"sync"
)

type (
	pool struct {
		key   sync.Locker
		etcd  db.Client
		nodes []Node
	}

	Pool interface {
		Add(Node)
		Remove(hostname string)
		AddLabel(hostname string, labels []string)
		GetLabels(hostname string) []string
		Contains(hostname string) bool
	}
)

var NodePool = NewPool()

func NewPool() Pool {
	return &pool{
		&sync.Mutex{},
		db.NewClient(),
		make([]Node, 3),
	}
}

func (p *pool) Add(n Node) {
	p.key.Lock()
	defer p.key.Unlock()
	p.nodes = append(p.nodes, n)
	curNodeKey := strings.Join([]string{ClmgrKey, "nodes", GetHostname()}, "/")
	data, err := json.Marshal(n)
	if err != nil {
		logger.Errorf("Can't marshall node info")
		return
	}
	p.etcd.Set(curNodeKey, string(data))
}

func (p *pool) Remove(hostname string) {
	for i := range p.nodes {
		if p.nodes[i].Name == hostname {
			p.nodes = append(p.nodes[0:i], p.nodes[i+1:]...)
			curNodeKey := strings.Join([]string{ClmgrKey, "nodes", p.nodes[i].Name}, "/")
			err := p.etcd.Remove(curNodeKey)
			if err != nil {
				logger.Errorf("Can't remove node, err %s", err.Error())
				return
			}
			break
		}
	}
}

func (p *pool) AddLabel(hostname string, labels []string) {
	p.key.Lock()
	defer p.key.Unlock()
	for i, nodeHN := range p.nodes {
		// found node
		if nodeHN.Name == hostname {
			for _, lbl := range labels {
				inside := false
				for _, oldLbl := range nodeHN.Labels {
					if oldLbl == Label(lbl) {
						inside = false
						break
					}
				}
				// if node doesn't contain the label
				if !inside {
					p.nodes[i].Labels = append(p.nodes[i].Labels, Label(lbl))
				}
			}
		}
	}
}

func (p *pool) GetLabels(hostname string) []string {
	for _, nodeHN := range p.nodes {
		if nodeHN.Name == hostname {
			lbls := make([]string, len(nodeHN.Labels))
			for _, lbl := range nodeHN.Labels {
				lbls = append(lbls, string(lbl))
			}
			return lbls
		}
	}
	return []string{}
}

func (p *pool) Contains(hostname string) bool {
	for i := range p.nodes {
		if p.nodes[i].Name == hostname {
			return true
		}
	}
	return false
}
