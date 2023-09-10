package main

import (
	"fmt"
	"time"

	"github.com/go-co-op/gocron"
)

func startHealthCheck() {
	s := gocron.NewScheduler(time.Local)

	for _, server := range serverlist {
		s.Every(2).Seconds().Do(func() {})
		fmt.Println(server.Url)
	}
}

func checkHealth() {

}
