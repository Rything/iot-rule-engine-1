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

	nw := network.NewNetwork()
	nw.AddNode(mqttNode)
	nw.AddNode(filterNode)
	nw.AddNode(debugNode)

	nw.Start()

	var e int
	fmt.Scanf("%d", &e)

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
