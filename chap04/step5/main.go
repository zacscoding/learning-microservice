package main

import (
	"github.com/zacscoding/learning-microservice-with-go/chap04/step5/data"
	"github.com/zacscoding/learning-microservice-with-go/chap04/step5/handlers"
	"log"
	"net/http"
)

func main() {
	store, err := data.NewMongoStore("192.168.79.130")
	if err != nil {
		log.Fatal(err)
	}

	handler := handlers.Search{DataStore: store}
	err = http.ListenAndServe(":8323", &handler)
	if err != nil {
		log.Fatal(err)
	}
}
