package main

import (
	"os"
	"fmt"
	"log"
	"net/http"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	visits = prometheus.NewCounter(
		prometheus.CounterOpts{
			Name: "visits",
			Help: "Number of website visits",
		},
	)
)

func init() {
	prometheus.MustRegister(visits)
}

func GetCounter() int {
	// www.xkcd.com/221
	return 4
}

func NewCounter() <-chan int {
	counter := make(chan int, 1)
	go func() {
		for i := 1; ; i++ {
			counter <- i
		}
	}()
	return counter
}

func main() {
	owner := os.Getenv("OWNER")

	http.Handle("/metrics", promhttp.Handler())

	counter := NewCounter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s 200", r.Host, r.Method, r.RequestURI)
		w.Header().Add("Content-Type", "text/html; charset=utf-8")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>Container Demo</title></head><body><h1>Hello, %s!</h1><p>You are visitor %d</p></body></html>", owner, <-counter)
		visits.Inc()
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
