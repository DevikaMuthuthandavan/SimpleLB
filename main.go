package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Simple Load Balancer")
	fmt.Println("Enter 1 - Round Robin \n 2 - Static Strategy \n 3 - Hash Based Strategy")
	var strategy int
	fmt.Scanln(&strategy)
	lb := InitLB()
	lb.proxy()
	lb.setStrategy(strategy)
	go startHealthCheck(lb)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
