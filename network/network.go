package network

import (
	"fmt"

	"github.com/nattaponra/iot-rule-engine/node"
)

const (
	NodeNotFoundInNetwork     string = "Not not found in network"
	InputParamOutOfLength     string = "Input param out of length"
	InputParamIsInvalidFormat string = "Input param is invalid format"
)

type NetworkNode struct {
	Nodes []node.Node
}

func NewNetwork() *NetworkNode {
	return &NetworkNode{}
}

func (nw *NetworkNode) AddNode(n *node.Node) {
	nw.Nodes = append(nw.Nodes, *n)
}

func (nw *NetworkNode) Start() {

	output := make(chan interface{}, 1)
	var bufferOutput interface{}

	for i := 0; i < len(nw.Nodes); i++ {
		fmt.Printf("-------Execute %s node----------", nw.Nodes[i].Name)
		fmt.Println()

		if nw.Nodes[i].Type == node.SourceNode {
			go nw.Nodes[i].Execute(output)

		} else {
			nw.Nodes[i].SetInput(bufferOutput)
			go nw.Nodes[i].Execute(output)

		}

		//Send current output node to next input node.
		bufferOutput = <-output

	}

}
