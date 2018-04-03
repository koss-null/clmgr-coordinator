package connection

import (
	"net"
)

type (
	node struct {
		hostname string
		ip       net.IP
	}

	Node interface {
		Hostname() string
		IP() net.IP
		Ping() (bool, error)
	}
)

func NewNode(hostName string, ip net.IP) Node {
	return &node{hostName, ip}
}

func (n node) Hostname() string {
	return n.hostname
}

func (n node) IP() net.IP {
	return n.ip
}

func (n node) Ping() (bool, error) {

	return true, nil
}
