package node

type IONodeDataType int

const (
	Int IONodeDataType = iota
	String
	Bool
)

type ConnectionType int

const (
	Single ConnectionType = iota
	Multiple
)

type NodeConfig struct {
	InputNodeType      ConnectionType
	OutputNodeType     ConnectionType
	InputNodeDataType  IONodeDataType
	OutputNodeDataType IONodeDataType
}

func (nc NodeConfig) IsOutputMultipleConnect() bool {
	return nc.OutputNodeType == Multiple
}

func (nc NodeConfig) IsInputSingleConnect() bool {
	return nc.InputNodeType == Single
}
