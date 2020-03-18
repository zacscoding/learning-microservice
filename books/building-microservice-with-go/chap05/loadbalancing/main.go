package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// Strategy interface like round robin or random
type Strategy interface {
	NextEndpoint() url.URL
	SetEndpoints([] url.URL)
}

// RandomStrategy implements Strategy for random end points
type RandomStrategy struct {
	endpoints []url.URL
}

func (r *RandomStrategy) NextEndpoint() url.URL {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r.endpoints[r1.Intn(len(r.endpoints))]
}

func (r *RandomStrategy) SetEndpoints(endpoints []url.URL) {
	r.endpoints = endpoints
}

type LoadBalancer struct {
	strategy Strategy
}

func NewLoadBalancer(strategy Strategy, endpoints []url.URL) *LoadBalancer {
	strategy.SetEndpoints(endpoints)
	return &LoadBalancer{strategy: strategy}
}

func (l *LoadBalancer) GetEndpoint() url.URL {
	return l.strategy.NextEndpoint()
}

func (l *LoadBalancer) UpdateEndpoints(urls []url.URL) {
	l.strategy.SetEndpoints(urls)
}

func main() {
	endpoints := []url.URL{
		{
			Host: "www.google.com",
		},
		{
			Host: "www.google.co.kr",
		},
	}

	lb := NewLoadBalancer(&RandomStrategy{}, endpoints)

	for i := 0; i < 5; i++ {
		fmt.Println("Endpoint : ", lb.GetEndpoint())
		time.Sleep(1 * time.Millisecond)
	}
}
