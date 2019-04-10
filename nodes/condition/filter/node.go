package filter

import (
	"fmt"

	"github.com/robertkrimen/otto"

	"github.com/nattaponra/iot-rule-engine/node"
)

var _ node.PluginNode = (*FilterNode)(nil)

type FilterNode struct{}

func (f FilterNode) Info() node.Info {
	return node.Info{
		Name:     "ScriptNode",
		NodeType: node.ActionNode,
	}
}

func (f FilterNode) Config() node.NodeConfig {
	return node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	}
}

func (f FilterNode) Properties() node.Properties {
	return node.Properties{
		FormInputs: map[string]node.FormInput{
			"script": node.FormInput{
				InputType:    node.Text,
				DefaultValue: "return input==='Hello World';",
				IsRequired:   true,
			},
		},
	}

}

func (f FilterNode) Execute() func(node.Node, chan interface{}) {
	return func(n node.Node, output chan interface{}) {

		pro := n.GetProperties()
		injectSctipt := pro.FormInputs["script"].GetStringValue()
		vm := otto.New()
		vm.Set("input", n.Input.([]string)[0])
		var script = `
		(function(input){
		 ` + injectSctipt + `
 		})(input);
 `
		fmt.Println("RunScript:", script)

		value, err := vm.Run(script)

		if err != nil {
			fmt.Println(err.Error())
		}

		var result bool
		if value.IsBoolean() {
			if result, err = value.ToBoolean(); err != nil {
				fmt.Println(err.Error())
			}

		}
		output <- result
	}
}
