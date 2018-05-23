package node

type Node struct {
	Name   string  `json:"name"`
	Labels []Label `json:"labels"`
	IP     string  `json:"ip"`
}
