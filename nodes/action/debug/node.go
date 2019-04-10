package debug

import (
	"fmt"

	"github.com/nattaponra/iot-rule-engine/node"
)

var _ node.PluginNode = (*DebugNode)(nil)

type DebugNode struct{}

func (d DebugNode) Info() node.Info {
	return node.Info{
		Name:     "Debug",
		NodeType: node.ActionNode,
	}
}

func (d DebugNode) Config() node.NodeConfig {
	return node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	}
}

func (d DebugNode) Properties() node.Properties {
	return node.Properties{
		FormInputs: map[string]*node.FormInput{
			"format": &node.FormInput{
				InputType:    node.Text,
				DefaultValue: "json",
				IsRequired:   true,
			},
		},
	}

}

func (d DebugNode) Execute() func(node.Node, chan interface{}) {
	return func(n node.Node, output chan interface{}) {
		pro := n.GetProperties()
		fmt.Println("Prevouise Node Output:", n.Input)
		fmt.Println("Format:", pro.FormInputs["format"].GetStringValue())
		output <- n.Input
	}
}
