package main

import (
	"net/http"
)

type LoadBalancer struct {
	backends []*server
	strategy BalancingStrategy
}

func InitLB() *LoadBalancer {
	servers := []*server{
		newServer("server-1", "http://127.0.0.1:5000"),
		newServer("server-2", "http://127.0.0.1:5001"),
		newServer("server-3", "http://127.0.0.1:5002"),
		newServer("server-4", "http://127.0.0.1:5003"),
		newServer("server-5", "http://127.0.0.1:5004"),
	}

	lb := new(LoadBalancer)
	lb.backends = servers
	lb.strategy = NewRRBalancingStrategy(servers)
	return lb
}

func (lb *LoadBalancer) proxy() {
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		server := lb.strategy.GetNextBackend()

		if server != nil {
			server.ReverseProxy.ServeHTTP(res, req)
		} else {
			http.Error(res, "Couldn't process request: "+"No hosts available", http.StatusServiceUnavailable)
			return
		}

	})
}
