package interfaces

type NodeStorage interface {
	Put(e Node) error
	Get(hostname string, nodetype interface{}) (Node, error)
	List(nodetype interface{}) ([]Node, error)
}
