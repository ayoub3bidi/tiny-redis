package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		log.Fatalf("Failed to bind to port 6379: %v", err)
	}
	defer ln.Close()

	fmt.Println("Mini Redis server running on port 6379...")

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Printf("Connection error: %v", err)
			continue
		}
		go handleConnection(conn)
	}
}
