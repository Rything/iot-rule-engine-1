package rule

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type RuleType string

const (
	Action    RuleType = "Action"
	Condition RuleType = "Condition"
	Decision  RuleType = "Decision"
)

type IRule interface {
	Name() string
	Type() RuleType
	OnMsg()
}

type Rule struct {
	ID         string
	ParentID   string
	Name       string
	Type       RuleType
	RuleEngine interface{}
}

func NewRule(name string, rtype RuleType) *Rule {
	u := uuid.NewV4()
	return &Rule{
		ID:   fmt.Sprintf("%s", u),
		Name: name,
		Type: rtype,
	}
}

func (r Rule) OnMsg() {

}

func (r Rule) AddChild(childRule *Rule) {
	childRule.ParentID = r.ID
}
