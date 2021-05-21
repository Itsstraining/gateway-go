package main

import (
	"log"
	"main/core"
	"main/core/routers"
)

func main() {
	server, err := core.NewServer()
	if err != nil {
		log.Fatal(err)
	}
	// Assigning routers
	routers.NewUser(server, "/v1/user")
	routers.NewChallenge(server, "/v1/challenge")

	server.Start()
}
