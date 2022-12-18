package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/danigmx/chess-api/cmd/server/api"
)

func main() {
	listenAddress := flag.String("port", ":8080", "Server address")
	flag.Parse()

	server := api.NewServer(*listenAddress)
	fmt.Println("server running on port:", *listenAddress)
	log.Fatal(server.Start())
}
