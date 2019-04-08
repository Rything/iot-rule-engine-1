package main

import (
	"fmt"
	"time"

	"github.com/nattaponra/iot-rule-engine/network"

	"github.com/nattaponra/iot-rule-engine/node"
)

func main() {

	//Create instance node
	n := node.NewNode("MQTT-Sub", node.SourceNode)

	//Create Property form input
	formInputs := []node.FormInput{
		node.FormInput{
			InputName:  "host",
			InputType:  node.Text,
			IsRequired: true,
		},
		node.FormInput{
			InputName:  "port",
			InputType:  node.Text,
			IsRequired: true,
		},
		node.FormInput{
			InputName:  "topic",
			InputType:  node.Text,
			IsRequired: true,
		},
	}

	//Set  form input that we just creates above.
	n.SetProperties(node.Properties{
		FormInputs: formInputs,
	})

	n.SetConfig(node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	})

	n.SetExecute(func(n node.Node) {
		for {
			n.SetOutput([]string{"Hello World"})
			fmt.Println(n.Output)
			time.Sleep(time.Second)
		}

	})

	nw := network.NewNetwork()
	nw.AddNode(n)
	nw.Start()

	var e int
	fmt.Scanf("%d", &e)

}
