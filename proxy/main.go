package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
)

const lbport int = 8080

func main() {
	url, err := url.Parse("http://backend-service")
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	// proxy := httputil.NewSingleHostReverseProxy(url)
	proxy := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		r.Host = url.Host
		r.URL.Host = url.Host
		r.URL.Scheme = url.Scheme
		r.RequestURI = ""

		host, _, _ := net.SplitHostPort(r.RemoteAddr)

		fmt.Println("HOSTIK", host)
		r.Header.Set("X-Forwarded-For", host)

		client := &http.Client{}
		resp, err := client.Do(r)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, err.Error())
		}

		for k, v := range resp.Header {
			for _, item := range v {
				w.Header().Set(k, item)
			}
		}

		fmt.Println(resp.Header.Get("Content-Type"))
		w.Header().Set("Content-Type", resp.Header.Get("Content-Type"))
		body, _ := ioutil.ReadAll(resp.Body)
		w.Write(body)
	})

	fmt.Println("Starting proxy server")

	http.ListenAndServe(fmt.Sprintf(":%d", lbport), proxy)
}
