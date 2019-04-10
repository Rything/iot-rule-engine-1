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

type Info struct {
	Name     string
	NodeType NodeType
}

type PluginNode interface {
	Info() Info
	Config() NodeConfig
	Properties() Properties
	Execute() func(Node, chan interface{})
}

type Plugin struct {
	Plugin PluginNode
}

func NewNodeWithPlugin(p Plugin) *Node {

	node := Node{
		Name: p.Plugin.Info().Name,
		Type: p.Plugin.Info().NodeType,
	}
	node.SetProperties(p.Plugin.Properties())
	node.SetConfig(p.Plugin.Config())
	node.SetExecute(p.Plugin.Execute())

	return &node
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

func (n *Node) SetFormInput(key string, value interface{}) {
	if _, ok := n.properties.FormInputs[key]; ok {
		n.properties.FormInputs[key].Value = value
	}
}

func (n *Node) GetFormInput(key string) *FormInput {
	if val, ok := n.properties.FormInputs[key]; ok {
		return val
	}

	return nil
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
