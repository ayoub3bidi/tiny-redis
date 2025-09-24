package main

import (
	"bufio"
	"fmt"
	"net"
	"strings"
	"sync"
)

var (
	dataStore = make(map[string]string)
	mu        sync.RWMutex
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	reader := bufio.NewReader(conn)

	for {
		val, err := ReadValue(reader)
		if err != nil {
			return
		}

		resp := handleCommand(val)
		conn.Write([]byte(resp.Marshal()))
	}
}

func handleCommand(v Value) Value {
	if v.typ != TypeArray || len(v.array) == 0 {
		return Value{typ: TypeError, str: "invalid command"}
	}

	cmd := strings.ToUpper(string(v.array[0].bulk))

	switch cmd {
	case "PING":
		if len(v.array) > 1 {
			return Value{typ: TypeString, str: string(v.array[1].bulk)}
		}
		return Value{typ: TypeString, str: "PONG"}

	case "SET":
		if len(v.array) != 3 {
			return Value{typ: TypeError, str: "wrong number of arguments for SET"}
		}
		key := string(v.array[1].bulk)
		val := string(v.array[2].bulk)

		mu.Lock()
		dataStore[key] = val
		appendToAOF(fmt.Sprintf("SET %s %s\n", key, val))
		mu.Unlock()

		return Value{typ: TypeString, str: "OK"}

	case "GET":
		if len(v.array) != 2 {
			return Value{typ: TypeError, str: "wrong number of arguments for GET"}
		}
		key := string(v.array[1].bulk)

		mu.RLock()
		val, ok := dataStore[key]
		mu.RUnlock()

		if !ok {
			return Value{typ: TypeBulk, bulk: nil}
		}
		return Value{typ: TypeBulk, bulk: []byte(val)}

	default:
		return Value{typ: TypeError, str: "unknown command"}
	}
}
