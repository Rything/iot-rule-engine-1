package node

type node struct {
	properties Properties
	config     NodeConfig
	execute    func()
}

func NewNode() *node {
	return &node{}
}

func (n *node) SetProperties(p Properties) {
	n.properties = p
}

func (n *node) SetConfig(cf NodeConfig) {
	n.config = cf
}

func (n *node) Execute(f func()) {
	f()
}
