package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const webPort int = 80

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home(time.Now()))
	mux.HandleFunc("/demo", demo(time.Now()))

	fmt.Println("Starting backend server")

	log.Println(http.ListenAndServe(fmt.Sprintf(":%d", webPort), mux))
}

func home(t time.Time) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(w, "Hey ! Its home response at the time %s", t)
	}
}

func demo(t time.Time) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, "{}")
	}
}
