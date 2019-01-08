package main

import (
	"flag"
	"gitlab.okta-solutions.com/mashroom/backend/rightmove/impl"
)

func main() {
	addr := ":10000"
	flag.StringVar(&addr, "addr", addr, "Listen address")
	flag.Parse()

	server := impl.NewServer()
	server.Serve(addr)
}