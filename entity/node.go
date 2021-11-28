package entity

type Node struct {
	HostName string `json:"hostname"`
}

type SSH_Node struct {
	*Node
	User        string
	Pwd         string
	Private_key []byte
}

func (n *Node) GetHostName() string {
	return n.HostName
}
