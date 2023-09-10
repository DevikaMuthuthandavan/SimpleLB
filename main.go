package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	serverlist = []string{
		"http://127.0.0.1:5000",
		"http://127.0.0.1:5001",
		"http://127.0.0.1:5002",
		"http://127.0.0.1:5003",
		"http://127.0.0.1:5004",
	}

	lastServedIdx = 0
)

func main() {
	fmt.Println("Simple Load Balancer")
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	url := getServer()
	rProxy := httputil.NewSingleHostReverseProxy(url)
	rProxy.ServeHTTP(res, req)
	//fmt.Fprintln(res, "Hello from load balancer")
}

func getServer() *url.URL {
	serverUrl := serverlist[(lastServedIdx+1)%5]
	url, _ := url.Parse(serverUrl)
	lastServedIdx++
	return url
}
