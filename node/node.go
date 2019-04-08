package node

type Node struct {
	Name       string
	properties Properties
	config     NodeConfig
	execute    func()
}

func NewNode(name string) *Node {
	return &Node{
		Name: name,
	}
}

func (n *Node) SetProperties(p Properties) {
	n.properties = p
}

func (n *Node) SetConfig(cf NodeConfig) {
	n.config = cf
}

func (n *Node) GetConfig() NodeConfig {
	return n.config
}

func (n *Node) Execute(f func()) {
	n.execute = f
}

func (n *Node) Asset() {
	n.execute()
}
