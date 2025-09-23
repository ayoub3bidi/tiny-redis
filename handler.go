// handler.go
package main

import (
	"strings"
)

var store = map[string]string{}

func HandleCommand(v Value) Value {
	if v.typ != "array" || len(v.array) == 0 {
		return Value{typ: "string", str: "ERR invalid command"}
	}

	cmd := strings.ToUpper(v.array[0].bulk)

	switch cmd {
	case "PING":
		return Value{typ: "string", str: "PONG"}
	case "SET":
		if len(v.array) < 3 {
			return Value{typ: "string", str: "ERR wrong number of arguments"}
		}
		key := v.array[1].bulk
		val := v.array[2].bulk
		store[key] = val
		return Value{typ: "string", str: "OK"}
	case "GET":
		if len(v.array) < 2 {
			return Value{typ: "string", str: "ERR wrong number of arguments"}
		}
		key := v.array[1].bulk
		if val, ok := store[key]; ok {
			return Value{typ: "bulk", bulk: val}
		}
		return Value{typ: "bulk", bulk: ""} // nil in Redis would be "$-1"
	default:
		return Value{typ: "string", str: "ERR unknown command"}
	}
}
