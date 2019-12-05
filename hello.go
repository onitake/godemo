package main

import (
	"fmt"
	"log"
	"net/http"
)

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
	counter := NewCounter()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s %s 200", r.Host, r.Method, r.RequestURI)
		w.Header().Add("Content-Type", "text/html")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "<!DOCTYPE html><html><head><title>Container Demo</title></head><body><h1>Hello, %d!</h1></body></html>", <-counter)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
