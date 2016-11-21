package main

import (
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

// Strategy is an interface to be implemented by loadbalancing
// strategies like round robin or random.
type Strategy interface {
	NextEndpoint() url.URL
	SetEndpoints([]url.URL)
}

// RandomStrategy implements Strategy for random endopoint selection
type RandomStrategy struct {
	endpoints []url.URL
}

// NextEndpoint returns an endpoint using a random strategy
func (r *RandomStrategy) NextEndpoint() url.URL {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	return r.endpoints[r1.Intn(len(r.endpoints))]
}

// SetEndpoints sets the available endpoints for use by the strategy
func (r *RandomStrategy) SetEndpoints(endpoints []url.URL) {
	r.endpoints = endpoints
}

// LoadBalancer is returns endpoints for downstream calls
type LoadBalancer struct {
	strategy Strategy
}

// NewLoadBalancer creates a new loadbalancer and setting the given strategy
func NewLoadBalancer(strategy Strategy, endpoints []url.URL) *LoadBalancer {
	strategy.SetEndpoints(endpoints)
	return &LoadBalancer{strategy: strategy}
}

// GetEndpoint gets an endpoint based on the given strategy
func (l *LoadBalancer) GetEndpoint() url.URL {
	return l.strategy.NextEndpoint()
}

// UpdateEndpoints updates the endpoints available to the strategy
func (l *LoadBalancer) UpdateEndpoints(urls []url.URL) {
	l.strategy.SetEndpoints(urls)
}

func main() {
	endpoints := []url.URL{
		url.URL{Host: "www.google.com"},
		url.URL{Host: "www.google.co.uk"},
	}

	lb := NewLoadBalancer(&RandomStrategy{}, endpoints)

	fmt.Println(lb.GetEndpoint())
}
