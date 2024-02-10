package path

import (
	"log"
	"testing"
)

func TestGetHistoryFunction(t *testing.T) {
	newTest()
	history, err := pathTest.GetHistory()
	if err != nil {
		t.Error(err)
	}
	log.Println(history)
}
