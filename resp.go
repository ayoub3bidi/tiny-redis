package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

type Value struct {
	typ   string
	str   string
	num   int
	bulk  string
	array []Value
}

type Reader struct {
	r *bufio.Reader
}

func NewReader(r io.Reader) *Reader {
	return &Reader{r: bufio.NewReader(r)}
}

func (rd *Reader) Read() (Value, error) {
	b, err := rd.r.ReadByte()
	if err != nil {
		return Value{}, err
	}

	switch b {
	case '+': // simple string
		line, _ := rd.readLine()
		return Value{typ: "string", str: line}, nil
	case '$': // bulk string
		line, _ := rd.readLine()
		n, _ := strconv.Atoi(line)
		buf := make([]byte, n+2) // include \r\n
		io.ReadFull(rd.r, buf)
		return Value{typ: "bulk", bulk: string(buf[:n])}, nil
	case '*': // array
		line, _ := rd.readLine()
		n, _ := strconv.Atoi(line)
		arr := make([]Value, 0, n)
		for i := 0; i < n; i++ {
			val, _ := rd.Read()
			arr = append(arr, val)
		}
		return Value{typ: "array", array: arr}, nil
	default:
		return Value{}, fmt.Errorf("unknown RESP type %q", b)
	}
}

func (rd *Reader) readLine() (string, error) {
	line, err := rd.r.ReadString('\n')
	if err != nil {
		return "", err
	}
	return line[:len(line)-2], nil // strip \r\n
}

type Writer struct {
	w io.Writer
}

func NewWriter(w io.Writer) *Writer {
	return &Writer{w: w}
}

func (wr *Writer) Write(v Value) error {
	switch v.typ {
	case "string":
		_, err := fmt.Fprintf(wr.w, "+%s\r\n", v.str)
		return err
	case "bulk":
		_, err := fmt.Fprintf(wr.w, "$%d\r\n%s\r\n", len(v.bulk), v.bulk)
		return err
	default:
		_, err := fmt.Fprintf(wr.w, "+OK\r\n")
		return err
	}
}
