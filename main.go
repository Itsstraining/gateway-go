package main

import (
	"log"
	"main/core"
)

func main() {
	server, err := core.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	server.Start()
}
