package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverlist = []*httputil.ReverseProxy{
		createHost("http://127.0.0.1:5000"),
		createHost("http://127.0.0.1:5001"),
		createHost("http://127.0.0.1:5002"),
		createHost("http://127.0.0.1:5003"),
		createHost("http://127.0.0.1:5004"),
	}

	lastServedIdx = 0
)

func main() {
	fmt.Println("Simple Load Balancer")
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	server := getServer()
	server.ServeHTTP(res, req)
	//fmt.Fprintln(res, "Hello from load balancer")
}

func getServer() *httputil.ReverseProxy {
	nextIndex := (lastServedIdx + 1) % len(serverlist)
	server := serverlist[nextIndex]
	lastServedIdx = nextIndex
	return server
}

func createHost(urlstr string) *httputil.ReverseProxy {
	u, _ := url.Parse(urlstr)
	return httputil.NewSingleHostReverseProxy(u)
}
