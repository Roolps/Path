package path

import (
	"log"
	"testing"
)

func TestGetFiltersFunction(t *testing.T) {
	newTest()
	filters, err := pathTest.GetRules()
	if err != nil {
		t.Error(err)
	}
	log.Println(filters)
}
