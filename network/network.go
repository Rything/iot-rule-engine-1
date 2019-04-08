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
	Nodes []node.Node
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

	output := make(chan interface{}, 1)
	var bufferOutput interface{}

	for i := 0; i < len(nw.Nodes); i++ {
		if nw.Nodes[i].Type == node.SourceNode {
			go nw.Nodes[i].Execute(output)

		} else {
			nw.Nodes[i].SetInput(bufferOutput)
			go nw.Nodes[i].Execute(output)

		}

		//Send current output node to next input node.
		bufferOutput = <-output
		fmt.Println(bufferOutput)
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
