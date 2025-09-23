package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := NewReader(conn)
	writer := NewWriter(conn)

	for {
		val, err := reader.Read()
		if err != nil {
			fmt.Println("client disconnected:", err)
			return
		}
		resp := HandleCommand(val)
		writer.Write(resp)
	}
}

func main() {
	ln, err := net.Listen("tcp", ":6379")
	if err != nil {
		panic(err)
	}
	fmt.Println("TinyRedis listening on :6379")

	for {
		conn, _ := ln.Accept()
		go handleConnection(conn)
	}
}
