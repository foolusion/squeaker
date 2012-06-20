package main

import (
	"bufio"
	"encoding/json"
    "fmt"
    "os"
	"time"
)

type Topic struct {
    Title string
    Squeaks []Squeak
}

type Squeak struct {
    UUID, Message string
    Timestamp time.Time
}

func Get(title string) ([]Squeak, error) {
    filename := fmt.Sprintf("%s.json", title)
    f, err := os.Open(filename)
    if err != nil {
        if os.IsNotExist(err) {
            return nil, nil
        }
        return nil, err
    }
	var t Topic
	err = json.NewDecoder(bufio.NewReader(f)).Decode(&t)
	if err != nil {
		return nil, err
	}
	return t.Squeaks, nil
}

func main() {
	s, err := Get("test")
	if err != nil {
		panic(err)
	}
	for _, v := range s {
		fmt.Println(v.Message)
	}
}
