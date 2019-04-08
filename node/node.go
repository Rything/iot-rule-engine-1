package node

type Node struct {
	properties Properties
	config     NodeConfig
	execute    func()
}

func NewNode() *Node {
	return &Node{}
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
