package main

import (
	"fmt"

	"github.com/nattaponra/iot-rule-engine/node"
)

func main() {

	//Create instance node
	n := node.NewNode()

	//Create Property form input
	formInputs := []node.FormInput{
		node.FormInput{
			InputName:  "script",
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
		InputNodeDataType:  node.Int,
		OutputNodeType:     node.Multiple,
		OutputNodeDataType: node.Bool,
	})

	n.Execute(func() {
		fmt.Println("Execute")
	})
}
