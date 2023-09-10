package main

import (
	"net/http/httputil"
	"net/url"
)

type server struct {
	Url          string
	ReverseProxy *httputil.ReverseProxy
	IsHealthy    bool
}

func newServer(urlstr string) *server {
	u, _ := url.Parse(urlstr)
	rp := httputil.NewSingleHostReverseProxy(u)
	return &server{
		Url:          urlstr,
		ReverseProxy: rp,
		IsHealthy:    true,
	}
}
