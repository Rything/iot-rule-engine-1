package main

import (
	"fmt"

	"github.com/nattaponra/iot-rule-engine/rule"
)

func main() {
	var chainRule []*rule.Rule
	r1 := rule.NewRule("GetTempSensor", rule.Action)
	r2 := rule.NewRule("DebugValue", rule.Action)
	r3 := rule.NewRule("CheckIsHigh", rule.Condition)
	r4 := rule.NewRule("SendEmail", rule.Action)

	r1.AddChild(r2)
	r2.AddChild(r3)
	r3.AddChild(r4)

	chainRule = append(chainRule, r1)
	chainRule = append(chainRule, r2)
	chainRule = append(chainRule, r3)
	chainRule = append(chainRule, r4)

	//Execute  chainRule
	for _, v := range chainRule {
		fmt.Println(v.ID, v.Name, v.ParentID)
	}

}
