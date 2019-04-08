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
	formInputs := map[string]node.FormInput{
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
		pro := n.GetProperties()
		fmt.Println("Host:", pro.FormInputs["host"].GetStringValue())
		fmt.Println("Port:", pro.FormInputs["port"].GetStringValue())
		fmt.Println("Topic:", pro.FormInputs["topic"].GetStringValue())

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
