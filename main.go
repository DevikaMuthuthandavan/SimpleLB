package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Simple Load Balancer")
	http.HandleFunc("/", forwardRequest)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func forwardRequest(res http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(res, "Hello from load balancer")
}
