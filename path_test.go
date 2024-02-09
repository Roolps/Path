package path

import (
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var pathTest = &APIClient{}

func newTest() {
	wd, _ := os.Getwd()

	env, err := godotenv.Read(fmt.Sprintf("%v/.env", wd))
	if err != nil {
		log.Fatalf("couldn't open .env: %v", err)
	}
	pathTest.Username = env["PATH_USERNAME"]
	pathTest.Password = env["PATH_PASSWORD"]

	pathTest.New()
}

func TestNewTokenRequestWithInvalidLogin(t *testing.T) {
	err := pathTest.New()
	if err == nil {
		t.Errorf(err.Error())
	}
}

func TestNewTokenRequestWithValidLogin(t *testing.T) {
	newTest()
	err := pathTest.New()
	if err != nil {
		t.Error(err)
	}
}
