package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	lb "github.com/ashokdey/loadbalancer"
	"github.com/go-co-op/gocron"
)

var (
	serverList = []*lb.Server{
		lb.NewServer("server-1", "http://127.0.0.1:5001"),
		lb.NewServer("server-2", "http://127.0.0.1:5002"),
		lb.NewServer("server-3", "http://127.0.0.1:5003"),
		lb.NewServer("server-4", "http://127.0.0.1:5004"),
		lb.NewServer("server-5", "http://127.0.0.1:5005"),
	}

	lastIndex = 0
)

func main() {
	http.HandleFunc("/", forwardRequests)
	go startHealthCheck()
	log.Fatal(http.ListenAndServe(":8050", nil))
}

func forwardRequests(res http.ResponseWriter, req *http.Request) {
	server, err := getHealthyServer()
	if err != nil {
		http.Error(res, "Couldn't process request: "+err.Error(), http.StatusServiceUnavailable)
		return
	}
	server.ReverseProxy.ServeHTTP(res, req)
}

func getHealthyServer() (*lb.Server, error) {
	for i := 0; i < len(serverList); i++ {
		server := getServer()
		if server.Health {
			return server, nil
		}
	}
	return nil, fmt.Errorf("no healthy server")
}

func getServer() *lb.Server {
	nextIndex := (lastIndex + 1) % len(serverList)
	server := serverList[nextIndex]
	lastIndex = nextIndex
	return server
}

func startHealthCheck() {
	cron := gocron.NewScheduler(time.Local)

	for _, host := range serverList {
		_, err := cron.Every(2).Second().Do(func(s *lb.Server) {
			healthy := s.CheckHealth()
			if healthy {
				log.Printf("'%s' is healthy!", s.Name)
			} else {
				log.Printf("'%s' is not healthy", s.Name)
			}
		}, host)
		if err != nil {
			log.Fatalln(err)
		}
	}
	cron.StartAsync()
}
