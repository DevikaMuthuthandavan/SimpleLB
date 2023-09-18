package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"

	"github.com/google/uuid"
)

type BalancingStrategy interface {
	Init([]*server)
	GetNextBackend() *server
	RegisterBackend(*server)
	//PrintToplogy()
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
	for i := 0; i < len(s.servers); i++ {
		s.Index = (s.Index + 1) % len(s.servers)
		backend := s.servers[s.Index]
		if backend.IsHealthy {
			return backend
		}
	}

	return nil
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

func (s *StaticBalancingStrategy) Init(servers []*server) {
	s.Index = 0
	s.servers = servers
}

func (s *StaticBalancingStrategy) GetNextBackend() *server {
	for i := 0; i < len(s.servers); i++ {
		if s.servers[s.Index].IsHealthy {
			return s.servers[s.Index]
		} else {
			s.Index = (s.Index + 1) % len(s.servers)
		}
	}
	return nil
}

func (s *StaticBalancingStrategy) RegisterBackend(server *server) {
	s.servers = append(s.servers, server)
}

func (s *StaticBalancingStrategy) PrintToplogy() {
	for index, backend := range s.servers {
		if index == s.Index {
			fmt.Printf("  [%s] %s", "x", backend)
		} else {
			fmt.Printf("  [%s] %s", " ", backend)
		}
	}
}

func NewStaticBalancingStrategy(servers []*server) *StaticBalancingStrategy {
	strategy := new(StaticBalancingStrategy)
	strategy.Init(servers)
	return strategy
}

func NewHashBalancingStrategy(servers []*server) *HashBalacingStrategy {
	strategy := new(HashBalacingStrategy)
	strategy.Init(servers)
	return strategy
}

func hash(s string) int {
	h := md5.New()
	var sum int = 0
	io.WriteString(h, s)
	for _, b := range h.Sum(nil) {
		sum += int(b)
	}

	return sum % 19
}

func (s *HashBalacingStrategy) Init(servers []*server) {
	s.OccupiedSlots = []int{}
	s.servers = servers
	for _, backend := range servers {
		key := hash(backend.String())

		if len(s.OccupiedSlots) == 0 {
			s.OccupiedSlots = append(s.OccupiedSlots, key)
			s.servers = servers
			continue
		}

		index := sort.Search(len(s.OccupiedSlots), func(i int) bool {
			return s.OccupiedSlots[i] >= key
		})

		if index == len(s.OccupiedSlots) {
			s.OccupiedSlots = append(s.OccupiedSlots, key)
		} else {
			s.OccupiedSlots = append(s.OccupiedSlots[:index+1], s.OccupiedSlots[index:]...)
			s.OccupiedSlots[index] = key
		}

		if index == len(servers) {
			s.servers = append(s.servers, backend)
		} else {
			s.servers = append(s.servers[:index+1], s.servers[index:]...)
			s.servers[index] = backend
		}
	}
}

func (s *HashBalacingStrategy) GetNextBackend() *server {
	for i := 0; i < 19; i++ {
		reqId := uuid.NewString()
		slot := hash(reqId)
		index := sort.Search(len(s.OccupiedSlots), func(i int) bool {
			return s.OccupiedSlots[i] > slot
		})
		server := s.servers[index%len(s.servers)]
		if server.IsHealthy {
			return server
		}
	}

	return nil
}

func (s *HashBalacingStrategy) RegisterBackend(backend *server) {
	key := hash(backend.String())
	index := sort.Search(len(s.OccupiedSlots), func(i int) bool { return s.OccupiedSlots[i] >= key })
	if index == len(s.OccupiedSlots) {
		s.OccupiedSlots = append(s.OccupiedSlots, key)
	} else {
		s.OccupiedSlots = append(s.OccupiedSlots[:index+1], s.OccupiedSlots[:index]...)
		s.OccupiedSlots[index] = key
	}

	if index == len(s.OccupiedSlots) {
		s.servers = append(s.servers, backend)
	} else {
		s.servers = append(s.servers[:index+1], s.servers[index:]...)
		s.servers[index] = backend
	}
}
