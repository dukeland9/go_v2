package main

import (
	"flag"
	"fmt"
	"io/ioutil"
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
	log.Printf("Pinging http://%s/\n", addr)
	resp, err := http.Get(fmt.Sprintf("http://%s/", addr))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Got pong: %s\n", body)
}
