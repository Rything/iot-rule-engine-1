# IoT Rule Engine

## Example
- Flow
### MQTTSub -> Script -> Debug Result
```
package main

import (
	"fmt"
	"time"

	"github.com/nattaponra/iot-rule-engine/network"
	"github.com/nattaponra/iot-rule-engine/node"
	"github.com/robertkrimen/otto"
)

func main() {

	//######################## MQTT NODE ###################################
	mqttNode := node.NewNode("MQTT-Sub", node.SourceNode)

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
	mqttNode.SetProperties(node.Properties{
		FormInputs: formInputs,
	})

	mqttNode.SetConfig(node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	})

	mqttNode.SetExecute(func(n node.Node, output chan interface{}) {
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

	})

	//######################## Script NODE ###################################
	scriptNode := node.NewNode("scriptNode", node.ConditionNode)
	//Create Property form input
	formInputs = map[string]node.FormInput{
		"script": node.FormInput{
			InputType:    node.Text,
			DefaultValue: "return input==='Hello World';",
			IsRequired:   true,
		},
	}

	//Set  form input that we just creates above.
	scriptNode.SetProperties(node.Properties{
		FormInputs: formInputs,
	})

	scriptNode.SetConfig(node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	})

	scriptNode.SetExecute(func(n node.Node, output chan interface{}) {

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
	})

	//######################## Debug NODE ###################################
	debugNode := node.NewNode("DebugNode", node.ActionNode)
	//Create Property form input
	formInputs = map[string]node.FormInput{
		"format": node.FormInput{
			InputType:    node.Text,
			DefaultValue: "json",
			IsRequired:   true,
		},
	}

	//Set  form input that we just creates above.
	debugNode.SetProperties(node.Properties{
		FormInputs: formInputs,
	})

	debugNode.SetConfig(node.NodeConfig{
		InputNodeType:      node.Single,
		InputNodeDataType:  node.String,
		OutputNodeType:     node.Single,
		OutputNodeDataType: node.String,
	})

	debugNode.SetExecute(func(n node.Node, output chan interface{}) {
		pro := n.GetProperties()
		fmt.Println("Prevouise Node Output:", n.Input)
		fmt.Println("Format:", pro.FormInputs["format"].GetStringValue())
		output <- n.Input
	})

	nw := network.NewNetwork()
	nw.AddNode(mqttNode)
	nw.AddNode(scriptNode)
	nw.AddNode(debugNode)
	nw.Start()

	var e int
	fmt.Scanf("%d", &e)

}
```
### Result
```
-------Execute MQTT-Sub node----------
Host: 127.0.0.1
Port: 1885
Topic: /home/sensor
-------Execute scriptNode node----------
RunScript: 
                (function(input){
                 return input==='Hello World';
                })(input);
 
-------Execute DebugNode node----------
Prevouise Node Output: true
Format: json
```

## ToDo
- Send an email when device attribute changes.[on progress]
- Create an alarm when telemetry value exceeds a certain threshold.[on progress]
- Forward telemetry data to Kafka, RabbitMQ or external RESTful server.[on progress]
 
