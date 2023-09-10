package main

import (
	"fmt"
	"log"
	"net/http"
)

var (
	serverlist = []*server{
		newServer("http://127.0.0.1:5000"),
		newServer("http://127.0.0.1:5001"),
		newServer("http://127.0.0.1:5002"),
		newServer("http://127.0.0.1:5003"),
		newServer("http://127.0.0.1:5004"),
	}

	lastServedIdx = 0
)

func main() {
	fmt.Println("Simple Load Balancer")
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	s := getServer()
	s.ReverseProxy.ServeHTTP(res, req)
	//fmt.Fprintln(res, "Hello from load balancer")
}

func getServer() *server {
	nextIndex := (lastServedIdx + 1) % len(serverlist)
	s := serverlist[nextIndex]
	lastServedIdx = nextIndex
	return s
}

// func createHost(urlstr string) *httputil.ReverseProxy {
// 	u, _ := url.Parse(urlstr)
// 	return httputil.NewSingleHostReverseProxy(u)
// }
