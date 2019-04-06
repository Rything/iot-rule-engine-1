package rule

import (
	"fmt"

	uuid "github.com/satori/go.uuid"
)

type ActionResult struct {
	FloatVaule float64
}

type Rule struct {
	ID       string
	ParentID string
	Name     string
	Type     RuleType
	Execute  func(interface{}) interface{}
}

func NewRule(name string, rtype RuleType) *Rule {
	u := uuid.NewV4()
	return &Rule{
		ID:   fmt.Sprintf("%s", u),
		Name: name,
		Type: rtype,
	}
}

func ExecuteRule(chainRules []*Rule) {
	//Execute  chainRule
	var index int
	var result interface{}
	var currentRuleID string
	for {
		if index == len(chainRules) {
			break
		}

		r := chainRules[index]

		if r.ParentID == "" { // IsFirstRule
			result = r.Execute(nil)
			currentRuleID = r.ID
			index++

		} else {
			//Find child
			for j := 0; j < len(chainRules); j++ {
				rj := chainRules[j]
				if currentRuleID == rj.ParentID {
					result = rj.Execute(result)
					currentRuleID = rj.ID
					break
				}
			}
			index++
		}

	}
}

func (r Rule) AddChild(childRule *Rule) {
	childRule.ParentID = r.ID
}
