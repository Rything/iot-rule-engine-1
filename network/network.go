package network

import (
	"errors"
	"fmt"

	"github.com/nattaponra/iot-rule-engine/node"
)

const (
	NodeNotFoundInNetwork     string = "Not not found in network"
	InputParamOutOfLength     string = "Input param out of length"
	InputParamIsInvalidFormat string = "Input param is invalid format"
)

type NetworkNode struct {
	started bool
	Nodes   []node.Node
}

func NewNetwork() *NetworkNode {
	return &NetworkNode{}
}

func (nw *NetworkNode) AddNode(n *node.Node) {
	nw.Nodes = append(nw.Nodes, *n)
}

func (nw *NetworkNode) Input(inputParams ...interface{}) error {

	if len(nw.Nodes) == 0 {
		return errors.New(NodeNotFoundInNetwork)
	}

	//SourceNode
	config := nw.Nodes[0].GetConfig()

	if config.IsInputSingleConnect() && len(inputParams) > 1 {
		return errors.New(InputParamOutOfLength)
	}

	if IsInvalidInputFormat(config, inputParams) {
		return errors.New(InputParamIsInvalidFormat)
	}

	return nil
}

func (nw *NetworkNode) Start() {
	nw.started = true

	if nw.Nodes[0].Type == node.SourceNode {
		go func() {
			nw.Nodes[0].Execute()
		}()
	}

}

//IsInvalidInputFormat เป็น func ที่ตรวจสอบว่า input parameter ตรง formatที่ node config ไว้หรือไม่
func IsInvalidInputFormat(config node.NodeConfig, inputParams []interface{}) bool {
	for _, value := range inputParams {

		var inputNodeDataType node.IONodeDataType

		switch v := value.(type) {
		case int:
			inputNodeDataType = node.Int
		default:
			fmt.Println("inputNodeDataType not cover", v)
			return true
		}

		if inputNodeDataType != config.InputNodeDataType {
			return true
		}
	}
	return false
}
