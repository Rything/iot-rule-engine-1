package node

type node struct {
	properties Properties
	config     []string
	execute    func()
}

func NewNode() *node {
	return &node{}
}

func (n *node) SetProperties(p Properties) {

}

func (n *node) SetConfig() {

}

func (n *node) Execute(f func()) {

}
