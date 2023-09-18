package main

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	IsHealthy    bool
}

func newServer(name string, urlstr string) *server {
	u, _ := url.Parse(urlstr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		Name:         name,
		URL:          urlstr,
		ReverseProxy: rp,
		IsHealthy:    true,
	}
}

func (s *server) checkHealth() bool {
	resp, err := http.Head(s.URL)

	if err != nil {
		s.IsHealthy = false
		return s.IsHealthy
	}

	if resp.StatusCode != http.StatusOK {
		s.IsHealthy = false
		return s.IsHealthy
	}
	s.IsHealthy = true
	return s.IsHealthy
}

func (s *server) String() string {
	return s.Name + " : " + s.URL
}
