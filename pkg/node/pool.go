package node

import "sync"

type (
	pool struct {
		key   sync.Once
		nodes []Node
	}

	Pool interface {
		Add()
		Remove()
		AddLabel(hostname string, labels []string)
		GetLabels(hostname string) []string
		Contains(hostname string) bool
	}
)

var NodePool Pool
