package node

import "sync"

type Label string

var lables map[string]interface{}
var key sync.Locker = &sync.Mutex{}

func AddLabel(l string) bool {
	key.Lock()
	defer key.Unlock()
	if _, ok := lables[l]; ok {
		return false
	}
	lables[l] = struct{}{}
	return true
}

func GetLables() []string {
	key.Lock()
	defer key.Unlock()
	s, i := make([]string, len(lables)), 0
	for k := range lables {
		s[i] = k
		i++
	}
	return s
}

func IsLable(l string) bool {
	key.Lock()
	defer key.Unlock()
	_, ok := lables[l]
	return ok
}
