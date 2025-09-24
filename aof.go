package main

import (
	"os"
	"sync"
)

var (
	aofFile *os.File
	aofOnce sync.Once
)

func appendToAOF(entry string) {
	aofOnce.Do(func() {
		var err error
		aofFile, err = os.OpenFile("appendonly.aof", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
	})

	if aofFile != nil {
		aofFile.WriteString(entry)
	}
}
