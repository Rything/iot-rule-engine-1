package main

import (
	"fmt"

	"github.com/nattaponra/iot-rule-engine/nodes/source/mqtt"

	"github.com/nattaponra/iot-rule-engine/network"

	"github.com/nattaponra/iot-rule-engine/node"
	"github.com/nattaponra/iot-rule-engine/nodes/action/debug"
	"github.com/nattaponra/iot-rule-engine/nodes/condition/filter"
)

func main() {

	//Register Plugin
	mqttNode := node.NewNodeWithPlugin(node.Plugin{mqtt.MQTTNode{}})
	filterNode := node.NewNodeWithPlugin(node.Plugin{filter.FilterNode{}})
	debugNode := node.NewNodeWithPlugin(node.Plugin{debug.DebugNode{}})

	//Create Network
	nw := network.NewNetwork()

	//MQTT Node
	mqttNode.SetFormInput("host", "192.168.1.30")
	mqttNode.SetFormInput("port", "2000")
	mqttNode.SetFormInput("topic", "/app/home/door")
	nw.AddNode(mqttNode)

	//Condition Node
	nw.AddNode(filterNode)

	//Debug Node
	nw.AddNode(debugNode)

	//Run Network
	nw.Start()

	// var e int
	// fmt.Scanf("%d", &e)

}

type MQTT struct{}

func (MQTT) Info() node.Info {
	return node.Info{
		Name:     "ScriptNode",
		NodeType: node.ActionNode,
	}
}

type Debug struct{}

func (Debug) Info() {
	fmt.Println("Debug")
}

type Plugin interface {
	Info()
}

type Operation struct {
	Plugin Plugin
}
