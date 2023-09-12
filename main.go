package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	serverlist = []*server{
		newServer("server-1", "http://127.0.0.1:5000"),
		newServer("server-2", "http://127.0.0.1:5001"),
		newServer("server-3", "http://127.0.0.1:5002"),
		newServer("server-4", "http://127.0.0.1:5003"),
		newServer("server-5", "http://127.0.0.1:5004"),
	}

	lastServedIdx = 0
)

func main() {
	fmt.Println("Simple Load Balancer")
	http.HandleFunc("/", forwardRequest)
	go startHealthCheck()
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server, err := getHealthyServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(res, req)
}

func getHealthyServer() (*server, error) {
	for i := 0; i < len(serverlist); i++ {
		server := getServer()
		if server.IsHealthy {
			return server, nil
		}
	}
	return nil, fmt.Errorf("no healthy hosts")
}

func getServer() *server {
	nextIndex := (lastServedIdx + 1) % len(serverlist)
	server := serverlist[nextIndex]
	lastServedIdx = nextIndex
	return server
}
