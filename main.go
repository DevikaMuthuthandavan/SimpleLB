package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Simple Load Balancer")
	lb := InitLB()
	lb.proxy()
	//http.HandleFunc("/", forwardRequest)
	go startHealthCheck(lb)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
