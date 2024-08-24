package main

import (
	"flag"
	"log"
)

func main() {
	port := flag.String("port", "3000", "Port to start the server on")

	flag.Parse()

	if err := initAPI(port); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
