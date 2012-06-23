package main

import (
	"fmt"
	"os"
)

func genUUID() string {
	f, err := os.Open("/dev/urandom")
	if err != nil {
		panic("You need to have access to /dev/urandom")
	}
	defer f.Close()
	b := make([]byte, 16)
	f.Read(b)
	return fmt.Sprintf("%x-%x-%x-%x-%x", b[0:4], b[4:6], b[6:8], b[8:10], b[10:])
}
