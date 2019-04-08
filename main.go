package main

import (
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
}
