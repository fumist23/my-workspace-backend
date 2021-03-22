package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	addr string
)

func init() {
	flag.StringVar(&addr, "addr", ":8080", "tcp host:port to connect")
	flag.Parse()
}

func main() {
	log.Println("Server is running...")
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatalf("listen and serve failed  %+v", err)
	}
}
