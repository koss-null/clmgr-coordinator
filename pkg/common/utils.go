package common

import (
	"github.com/satori/go.uuid"
	"sync"
)

var hostname string
var once sync.Once

func CreateHostname() string {
	if hostname == "" {
		once.Do(func() {
			id, _ := uuid.NewV1()
			hostname = id.String()
		})
	}
	return hostname
}
