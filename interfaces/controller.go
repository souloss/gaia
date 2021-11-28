package interfaces

type NodeController interface {
	AddNode(Node) error
	DelNode(Node) error
	UpdateNode(Node) error
}
