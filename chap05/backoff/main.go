package main

import (
	"fmt"
	"github.com/eapache/go-resiliency/retrier"
	"time"
)

func main() {
	retry := 3
	n := 0
	r := retrier.New(retrier.ConstantBackoff(retry, 1*time.Second), nil)

	err := r.Run(func() error {
		fmt.Println("Attempt: ", n)
		n++
		return fmt.Errorf("Failed")
	})

	if err != nil {
		fmt.Println(err)
	}
}
