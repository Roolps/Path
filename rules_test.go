package path

import (
	"log"
	"testing"
)

func TestGetRulesFunction(t *testing.T) {
	newTest()
	_, err := pathTest.GetRules()
	if err != nil {
		t.Error(err)
	}
}

func TestCreateRuleFunction(t *testing.T) {
	newTest()
	rule, err := pathTest.CreateRule(&CreateRule{
		Source: "0.0.0.0/0",

		Destination:     "",
		DestinationPort: 1,

		Protocol:  "tcp",
		Whitelist: true,
		Comment:   "testing new library",
	})
	if err != nil {
		t.Error(err)
	} else {
		log.Println(rule)
	}
}

func TestDeleteRuleFunction(t *testing.T) {
	newTest()
	err := pathTest.DeleteRule("")
	if err != nil {
		t.Error(err)
	}
}
