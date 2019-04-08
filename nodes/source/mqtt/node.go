package debug

import (
	"fmt"
	"time"

	"github.com/nattaponra/iot-rule-engine/node"
	"github.com/nattaponra/iot-rule-engine/nodes"
)

var _ nodes.INode = (*MQTTNode)(nil)

type MQTTNode struct{}

func (mqtt *MQTTNode) Info() nodes.Info {
	return nodes.Info{
		Name:     "MQTT-Sub",
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
		"host": node.FormInput{
			InputType:    node.Text,
			DefaultValue: "127.0.0.1",
			IsRequired:   true,
		},
		"port": node.FormInput{
			InputType:    node.Text,
			DefaultValue: "1885",
			IsRequired:   true,
		},
		"topic": node.FormInput{
			InputType:    node.Text,
			DefaultValue: "/home/sensor",
			IsRequired:   true,
		},
	}
}

func (mqtt *MQTTNode) Execute() func(node.Node, chan interface{}) {
	return func(n node.Node, output chan interface{}) {
		pro := n.GetProperties()
		fmt.Println("Host:", pro.FormInputs["host"].GetStringValue())
		fmt.Println("Port:", pro.FormInputs["port"].GetStringValue())
		fmt.Println("Topic:", pro.FormInputs["topic"].GetStringValue())

		for {
			output <- []string{"Hello World"}
			n.SetOutput([]string{"Hello World"})
			///	fmt.Println(n.Output)
			time.Sleep(time.Second)
		}
	}
}
