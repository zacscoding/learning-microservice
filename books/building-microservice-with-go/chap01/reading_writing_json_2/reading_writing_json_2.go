package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type helloWorldResponse struct {
	// 출력 필드를 "message"로 변경
	Message string `json:"message"`
	// 이 필드는 출력하지 않음
	Author string `json:"-"`
	// 값이 비어 있으면 필드를 출력하지 않음
	Date string `json:", omitempty"`
	// 출력을 문자열로 변환하고 이름을 "id"로 바꾼다
	Id int `json:"id, string"`
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

/*
// http://localhost:8080/helloworld
{
  "Message": "HelloWorld",
  "Date": "2019",
  "id": 10
}
*/
func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := helloWorldResponse{Message: "HelloWorld", Author: "Zaccoding", Date: "2019", Id: 10}
	data, err := json.Marshal(response)
	if err != nil {
		panic("Ooops")
	}

	fmt.Fprint(w, string(data))
}
