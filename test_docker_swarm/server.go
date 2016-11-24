package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", "", "ip:port of the service.")
}

func main() {
	flag.Parse()
	log.Printf("Listening on %s", addr)
	http.HandleFunc("/", pong)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal(err)
	}
}

func pong(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "<html><body>I'm %s, respond to %s.</body></html>", request.Host, request.RemoteAddr)
	log.Printf("responded to %s.", request.RemoteAddr)
}
