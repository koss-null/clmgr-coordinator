package connection

import "net"

type (
	node struct {
		hostname string
		ip net.IP
	}

	Node interface {
		Hostname() string
		IP() net.IP
		Ping() (bool, error)
	}
)

func (n node) Hostname() string {
	return n.hostname
}

func (n node) IP() net.IP {
	return n.ip
}

func (n node) Ping() (bool, error) {
	// TODO implement
	return true, nil
}

