package main

import (
	"sync"
	"time"
)

type Squeaker struct {
	topics map[string][]Squeak
	mu     sync.RWMutex
}

type Squeak struct {
	UUID, Message string
	Time          time.Time
}

func NewSqueaker() *Squeaker {
	return &Squeaker{topics: make(map[string][]Squeak)}
}

func (s *Squeaker) Get(title string) []Squeak {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.topics[title]
}

func (s *Squeaker) Add(title string, message string) {
	squeak := Squeak{genUUID(), message, time.Now()}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.topics[title]; !present {
		s.topics[title] = make([]Squeak, 0, 100)
	}
	s.topics[title] = append(s.topics[title], squeak)
}
