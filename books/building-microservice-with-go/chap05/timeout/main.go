package main

import (
	"fmt"
	"github.com/eapache/go-resiliency/deadline"
	"os"
	"time"
)

// go run main.go slow
func main() {
	switch os.Args[1] {
	case "slow":
		fmt.Println("running slow..")
		makeNormalRequest()
	case "timeout":
		fmt.Println("running timeout..")
		makeTimeoutRequest()
	}
}

func makeNormalRequest() {
	slowFunction()
}

// 2초 후 취소
func makeTimeoutRequest() {
	dl := deadline.New(1 * time.Second)
	err := dl.Run(func(stopper <-chan struct{}) error {
		slowFunction()
		return nil
	})

	switch err {
	case deadline.ErrTimedOut:
		fmt.Println("Timeout")
	default:
		fmt.Println(err)
	}
}

func slowFunction() {
	for i := 0; i < 100; i++ {
		fmt.Println("Loop: ", i)
		time.Sleep(1 * time.Second)
	}
}
