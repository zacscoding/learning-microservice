package main

import (
	"fmt"
	"github.com/eapache/go-resiliency/breaker"
	"time"
)

func main() {
	// errorThreshold, successThreshold, timeout
	b := breaker.New(3, 1, 5*time.Second)

	for {
		result := b.Run(func() error {
			// call some service
			time.Sleep(2 * time.Second)
			return fmt.Errorf("Timeout")
		})

		switch result {
		case nil:
			// success
			fmt.Println("Success")
		case breaker.ErrBreakerOpen:
			// our function wasn`t run because the beraker was open
			fmt.Println("Breaker open")
		default:
			fmt.Println(result)
		}

		time.Sleep(500 * time.Millisecond)
	}
	// Output :
	//Timeout
	//Timeout
	//Timeout
	//Breaker open
	//Breaker open
	//Breaker open
	//Breaker open
	//Breaker open
	//Breaker open
	//Breaker open
	//Breaker open
	//Breaker open
	//Timeout
	//Breaker open
	//Breaker open
}
