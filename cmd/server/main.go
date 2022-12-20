package main

import (
	"flag"
	"fmt"

	"github.com/danigmx/chess-api/cmd/server/api"
)

func main() {
	listeningPort := flag.String("port", ":8080", "Server port listening")
	flag.Parse()

	server := api.NewServer(*listeningPort)

	fmt.Println("server running on port", *listeningPort)
	server.Start()
}
