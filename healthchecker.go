package main

import (
	"log"
	"time"

	"github.com/go-co-op/gocron"
)

func startHealthCheck(lb *LoadBalancer) {
	s := gocron.NewScheduler(time.Local)
	for _, host := range lb.backends {
		_, err := s.Every(2).Seconds().Do(func(s *server) {
			healthy := s.checkHealth()
			if healthy {
				log.Printf("'%s : %s' is healthy!", s.Name, s.URL)
			} else {
				log.Printf("'%s : %s' is not healthy", s.Name, s.URL)
			}
		}, host)
		if err != nil {
			log.Fatalln(err)
		}
	}
	s.StartAsync()
}
