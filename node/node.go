package node

type NodeType int

const (
	SourceNode = iota
	ActionNode
	ConditionNode
)

type Node struct {
	Type       NodeType
	Name       string
	properties Properties
	config     NodeConfig
	execute    func(Node, chan interface{})
	Input      interface{}
	Output     interface{}
}

func NewNode(name string, nodeType NodeType) *Node {
	return &Node{
		Name: name,
		Type: nodeType,
	}
}

func (n *Node) SetInput(input interface{}) {
	n.Input = input
}

func (n *Node) SetOutput(output interface{}) {
	n.Output = output
}

func (n *Node) SetProperties(p Properties) {
	n.properties = p
}

func (n *Node) GetProperties() Properties {
	return n.properties
}

func (n *Node) SetConfig(cf NodeConfig) {
	n.config = cf
}

func (n *Node) GetConfig() NodeConfig {
	return n.config
}

func (n *Node) SetExecute(f func(n Node, output chan interface{})) {
	n.execute = f
}

func (n *Node) Execute(output chan interface{}) {
	n.execute(*n, output)
}
