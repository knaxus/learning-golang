package loadbalancer

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	Name         string
	URL          string
	ReverseProxy *httputil.ReverseProxy
	Health       bool
}

func NewServer(name, urlString string) *Server {
	u, _ := url.Parse(urlString)
	r := httputil.NewSingleHostReverseProxy(u)
	return &Server{
		Name:         name,
		URL:          urlString,
		ReverseProxy: r,
		Health:       true,
	}
}

func (s *Server) CheckHealth() bool {
	resp, err := http.Head(s.URL)

	if err != nil {
		s.Health = false
		return s.Health
	}

	if resp.StatusCode != http.StatusOK {
		s.Health = false
		return s.Health
	}
	s.Health = true
	return s.Health
}
