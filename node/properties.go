package node

type InputType string

const (
	Text InputType = "text"
)

type FormInput struct {
	InputType    InputType
	Value        interface{}
	DefaultValue interface{}
	IsRequired   bool
}

// Properties
type Properties struct {
	FormInputs map[string]FormInput
}

func (fi FormInput) GetStringValue() string {
	if fi.Value != nil {
		return fi.Value.(string)
	}
	return fi.DefaultValue.(string)
}
