package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// RESP Value types
const (
	TypeString = '+'
	TypeError  = '-'
	TypeInt    = ':'
	TypeBulk   = '$'
	TypeArray  = '*'
)

type Value struct {
	typ   byte
	str   string
	bulk  []byte
	array []Value
}

// Parse RESP input into Value
func ReadValue(r *bufio.Reader) (Value, error) {
	prefix, err := r.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch prefix {
	case TypeString:
		line, _ := r.ReadString('\n')
		return Value{typ: TypeString, str: strings.TrimSpace(line)}, nil
	case TypeBulk:
		lenLine, _ := r.ReadString('\n')
		l, _ := strconv.Atoi(strings.TrimSpace(lenLine))
		buf := make([]byte, l+2) // +2 for CRLF
		_, err := io.ReadFull(r, buf)
		if err != nil {
			return Value{}, err
		}
		return Value{typ: TypeBulk, bulk: buf[:l]}, nil
	case TypeArray:
		lenLine, _ := r.ReadString('\n')
		count, _ := strconv.Atoi(strings.TrimSpace(lenLine))
		arr := make([]Value, count)
		for i := 0; i < count; i++ {
			v, err := ReadValue(r)
			if err != nil {
				return Value{}, err
			}
			arr[i] = v
		}
		return Value{typ: TypeArray, array: arr}, nil
	default:
		return Value{}, errors.New("unknown RESP type")
	}
}

// Marshal converts Value back to RESP format
func (v Value) Marshal() string {
	switch v.typ {
	case TypeString:
		return fmt.Sprintf("+%s\r\n", v.str)
	case TypeError:
		return fmt.Sprintf("-%s\r\n", v.str)
	case TypeBulk:
		return fmt.Sprintf("$%d\r\n%s\r\n", len(v.bulk), v.bulk)
	case TypeArray:
		var sb strings.Builder
		sb.WriteString(fmt.Sprintf("*%d\r\n", len(v.array)))
		for _, item := range v.array {
			sb.WriteString(item.Marshal())
		}
		return sb.String()
	default:
		return "-ERR unknown type\r\n"
	}
}
