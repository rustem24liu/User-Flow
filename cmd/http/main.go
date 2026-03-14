package main

import (
	"log"
	"user-flow/api"
)

func main() {
	server := api.NewServer()

	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
