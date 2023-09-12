package main

import "fmt"

type BalancingStrategy interface {
	Init([]*server)
	GetNextBackend() *server
	RegisterBackend(*server)
	PrintToplogy()
}

type RRBalancingStrategy struct {
	Index   int
	servers []*server
}

type StaticBalancingStrategy struct {
	Index   int
	servers []*server
}

type HashBalacingStrategy struct {
	OccupiedSlots []int
	servers       []*server
}

func (s *RRBalancingStrategy) Init(servers []*server) {
	s.Index = 0
	s.servers = servers
}

func (s *RRBalancingStrategy) GetNextBackend() *server {
	s.Index = (s.Index + 1) % len(s.servers)
	return s.servers[s.Index]
}

func (s *RRBalancingStrategy) RegisterBackend(server *server) {
	s.servers = append(s.servers, server)
}

func (s *RRBalancingStrategy) PrintToplogy() {
	for idx, backend := range s.servers {
		fmt.Printf("	[%d]%s", idx, backend)
	}
}

func NewRRBalancingStrategy(servers []*server) *RRBalancingStrategy {
	strategy := new(RRBalancingStrategy)
	strategy.Init(servers)
	return strategy
}
