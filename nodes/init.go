package nodes

import "github.com/nattaponra/iot-rule-engine/node"

type Info struct {
	Name     string
	NodeType node.NodeType
}

type INode interface {
	Info() Info
}
