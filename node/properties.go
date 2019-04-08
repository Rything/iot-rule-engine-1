package node

type InputType string

const (
	Text InputType = "text"
)

type FormInput struct {
	InputName  string
	InputType  InputType
	IsRequired bool
}

type Properties struct {
	FormInputs []FormInput
}
