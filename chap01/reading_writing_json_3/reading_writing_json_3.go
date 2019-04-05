package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HelloWorldResponse struct {
	Message string
}

func main() {
	port := 8080

	http.HandleFunc("/helloworld", helloWorldHandler)

	log.Printf("Server starting on port %v\n", 8080)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}

func helloWorldHandler(w http.ResponseWriter, r *http.Request) {
	response := HelloWorldResponse{Message: "HelloWorld"}
	encoder := json.NewEncoder(w)

	encoder.Encode(&response)
}
