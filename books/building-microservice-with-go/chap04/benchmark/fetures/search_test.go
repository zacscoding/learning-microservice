package features

import (
	"bytes"
	"fmt"
	"github.com/DATA-DOG/godog"
	"github.com/zacscoding/learning-microservice-with-go/chap04/step5/data"
	"io/ioutil"
	"net/http"
	"os/exec"
	"syscall"
	"time"
)

// godog ./
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

	if criteria != nil {
		request = []byte(criteria.(string))
	}

	response, err = http.Post("http://localhost:8323", "application/json", bytes.NewReader(request))
	return err
}

func iShouldReceiveABadRequestMessage() error {
	if response.StatusCode != http.StatusBadRequest {
		return fmt.Errorf("Should have recieved a bad response")
	}

	return nil
}

func iHaveAValidSearchCriteria() error {
	criteria = `{ "query": "Fat Freddy's Cat" }`

	return nil
}

func iShouldReceiveAListOfKittens() error {
	var body []byte
	body, err := ioutil.ReadAll(response.Body)

	if len(body) < 1 || err != nil {
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
		clearDB()
		setupData()
	})

	s.AfterScenario(func(interface{}, error) {
		server.Process.Signal(syscall.SIGINT)
	})

	waitForDB()
}

var server *exec.Cmd
var store *data.MongoStore

func startServer() {
	server = exec.Command("go", "build", "../main.go")
	server.Run()

	server = exec.Command("./main")
	go server.Run()

	// 실제로는 health 엔드 포인트를 점검
	time.Sleep(3 * time.Second)
	fmt.Printf("Server running with pid: %v", server.Process.Pid)
}

func waitForDB() {
	var err error

	for i := 0; i < 10; i++ {
		store, err = data.NewMongoStore("192.168.79.130")
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
}

func clearDB() {
	store.DeleteAllKittens()
}

func setupData() {
	store.InsertKittens(
		[]data.Kitten{
			{
				Id:     "1",
				Name:   "Felix",
				Weight: 12.3,
			},
			{
				Id:     "2",
				Name:   "Fat Freddy's Cat",
				Weight: 20.0,
			},
			{
				Id:     "3",
				Name:   "Garfield",
				Weight: 35.0,
			},
		})
}
