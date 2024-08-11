package main

import (
	"log"
)

func main() {
	if err := initAPI(); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
