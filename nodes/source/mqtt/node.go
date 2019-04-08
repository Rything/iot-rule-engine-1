package debug

import (
	"fmt"

	"github.com/nattaponra/iot-rule-engine/node"
	"github.com/nattaponra/iot-rule-engine/nodes"
)

var _ nodes.INode = (*MQTTNode)(nil)

type MQTTNode struct{}

func (mqtt *MQTTNode) Info() nodes.Info {
	return nodes.Info{
		Name:     "Debug",
		NodeType: node.ActionNode,
	}
}

func (mqtt *MQTTNode) Config() node.NodeConfig {
	return node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	}
}

func (mqtt *MQTTNode) FormInput() map[string]node.FormInput {
	return map[string]node.FormInput{
		"format": node.FormInput{
			InputType:    node.Text,
			DefaultValue: "json",
			IsRequired:   true,
		},
	}
}

func (mqtt *MQTTNode) Execute() func(node.Node, chan interface{}) {
	return func(n node.Node, output chan interface{}) {
		pro := n.GetProperties()
		fmt.Println("Prevouise Node Output:", n.Input)
		fmt.Println("Format:", pro.FormInputs["format"].GetStringValue())
		output <- n.Input
	}
}
