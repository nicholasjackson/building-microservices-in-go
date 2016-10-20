package features

import (
	"bytes"
	"fmt"
	"net/http"
	"os/exec"
	"time"

	"github.com/DATA-DOG/godog"
)

var criteria interface{}
var response *http.Response
var err error

func iHaveNoSearchCriteria() error {
	if criteria != nil {
		return fmt.Errorf("Criteria should be nil")
	}

	return nil
}

func iCallTheSearchEndpoint() error {
	var request []byte

	response, err = http.Post("http://localhost:2323", "application/json", bytes.NewReader(request))
	return err
}

func iShouldReceiveABadRequestMessage() error {
	if response.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("Should have recieved a bad response")
	}

	return nil
}

func iHaveAValidSearchCriteria() error {
	if criteria == nil {
		return fmt.Errorf("Do not have a valid criteria")
	}

	return nil
}

func iShouldReceiveAListOfKittens() error {
	var body []byte
	response.Body.Read(body)

	if len(body) < 1 {
		return fmt.Errorf("Should have received a list of kittens")
	}

	return nil
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^I have no search criteria$`, iHaveNoSearchCriteria)
	s.Step(`^I call the search endpoint$`, iCallTheSearchEndpoint)
	s.Step(`^I should receive a bad request message$`, iShouldReceiveABadRequestMessage)
	s.Step(`^I have a valid search criteria$`, iHaveAValidSearchCriteria)
	s.Step(`^I should receive a list of kittens$`, iShouldReceiveAListOfKittens)

	s.BeforeScenario(func(interface{}) {
		startServer()
		fmt.Printf("Server running with pid: %v", server.Process.Pid)
	})

	s.AfterScenario(func(interface{}, error) {
		server.Process.Kill()
	})
}

var server *exec.Cmd

func startServer() {
	server = exec.Command("go", "run", "../main.go")
	go server.Run()
	time.Sleep(3 * time.Second)
}
