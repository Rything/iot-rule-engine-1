package rule

type ReadNumberRole struct{}

func (r ReadNumberRole) Name() string {
	return "ReadNumber"
}

func (r ReadNumberRole) Type() RuleType {
	return Action
}

func (r ReadNumberRole) Execute() string {
	return "Executed"
}
