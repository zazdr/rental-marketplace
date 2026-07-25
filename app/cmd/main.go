package main

import (
	"app/bootstrap"
	"log"
)

func main() {
	server, err := bootstrap.New()
	if err != nil {
		log.Fatal(err)
	}

	log.Fatal(server.Start())
}
