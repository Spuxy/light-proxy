package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home(time.Now()))
	mux.HandleFunc("/demo", demo(time.Now()))

	fmt.Println("Starting backend server")

	log.Println(http.ListenAndServe(":80", mux))
}

func home(t time.Time) func(w http.ResponseWriter, r *http.Request) {
	log.Println("Log: home func")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Log: home func is executing")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Hey ! Its home response at the time %s", t)
	}
}

func demo(t time.Time) func(w http.ResponseWriter, r *http.Request) {
	log.Println("Log: demo func")
	return func(w http.ResponseWriter, r *http.Request) {
		log.Println("Log: demo func is executing")
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{}")
	}
}
