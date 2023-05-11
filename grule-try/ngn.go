package main

import (
	"fmt"

	"github.com/hyperjumptech/grule-rule-engine/ast"
	"github.com/hyperjumptech/grule-rule-engine/builder"
	"github.com/hyperjumptech/grule-rule-engine/engine"
	"github.com/hyperjumptech/grule-rule-engine/pkg"
)

var rules string = `
rule CheckSeniorAge "Senior Citzen Check Rule" salience 2 {
	when
	  person.Age > 60
	then 
	  person.Senior = true;
	  Retract("CheckSeniorAge");
}
rule CheckAddress "Address Road should be set" salience 2 {
	when
	  person.IsAddressCityNotSet()
	then 
	  person.Address.City = "Vatican City";
	  Retract("CheckAddress");
}
`

var newrules string = `
rule CheckMinorAge "Minor Person Addres is to hidden" salience 3 {
	when 
	  person.Age < 14
	then
	  person.HideAddress();
	  Retract("CheckMinorAge");
}
rule CheckAddress "Address Road should be set" salience 2 {
	when
	  person.IsAddressCityNotSet()
	then 
	  person.Address.City = "Mumbai";
	  Retract("CheckAddress");
}
`

var Ngn *engine.GruleEngine
var RuleLib *ast.KnowledgeLibrary
var MyRules *ast.KnowledgeBase
var ruleBuilder *builder.RuleBuilder

func init() {
	fmt.Println("Initializing ngn..")
	RuleLib = ast.NewKnowledgeLibrary()
	ruleBuilder = builder.NewRuleBuilder(RuleLib)
	ruleDefs := pkg.NewBytesResource([]byte(rules))
	err := ruleBuilder.BuildRuleFromResource("MyRules", "v1.0", ruleDefs)

	if err != nil {
		fmt.Println("Failed to parse rules")
		panic(err)
	}
	Ngn = engine.NewGruleEngine()
	MyRules = RuleLib.NewKnowledgeBaseInstance("MyRules", "v1.0")
	fmt.Println("Initialized ngn..")
}

func AddNewRules() {
	ruleDefs := pkg.NewBytesResource([]byte(newrules))
	err := ruleBuilder.BuildRuleFromResource("MyRules", "v1.0", ruleDefs)
	if err != nil {
		fmt.Println("Failed to parse rules")
		panic(err)
	}
}

func RemoveRule(rulename string) {
	kbbp := RuleLib.GetKnowledgeBase("MyRules", "v1.0")
	kbbp.RemoveRuleEntry(rulename)
}
