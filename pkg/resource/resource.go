package resource

import "myproj.com/clmgr-coordinator/pkg/node"

type resourceType string

const (
	primitive   resourceType = "primitive"
	clone                    = "clone"
	clusterwide              = "clusterwide"
	masterSleve              = "masterslave"
)

type Resource struct {
	Name       string       `json:"name"`
	ResTypes   resourceType `json:"resource-type"`
	NodeLabels []node.Label `json:"node-labels"`
	Deps       []string     `json:"resource-deps"`
	ETCDKey    string       `json:"etcd-key,omitempty"`
}

/*
	check() validates new resource
	name shouldn't already be added
	resType should be correct
	nodeLabels should be ok
	deps should be already added as resources
*/
func (r *Resource) Check() bool {
	// todo: implement
	return true
}

/*
	addKey() generates and adds ETCDKey to the resource
*/
func (r *Resource) AddKey() {

}
