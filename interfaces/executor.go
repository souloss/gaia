package interfaces

type Node interface {
	GetHostName() string
}

type VMExecutor interface {
	Run(node Node, command string) (string, error)
	Upload(node Node, localPath string, remotePath string) error
	Download(node Node, remotePath string, localPath string) error
}
