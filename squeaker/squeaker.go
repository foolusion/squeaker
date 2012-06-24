package squeaker

import (
	"sync"
	"time"
)

type Squeaker interface {
	Get(title string) []Squeak
	Squeak(title, message string)
	Topics() []string
}

type MapSqueaker struct {
	topics map[string][]Squeak
	mu     sync.RWMutex
}

type Squeak struct {
	UUID, Message string
	Time          time.Time
}

func NewMapSqueaker() *MapSqueaker {
	return &MapSqueaker{topics: make(map[string][]Squeak)}
}

func (s *MapSqueaker) Get(title string) []Squeak {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.topics[title]
}

func (s *MapSqueaker) Squeak(title, message string) {
	squeak := Squeak{UUID(), message, time.Now()}
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, present := s.topics[title]; !present {
		s.topics[title] = make([]Squeak, 0, 100)
	}
	s.topics[title] = append(s.topics[title], squeak)
}

func (s *MapSqueaker) Topics() []string {
	var topics []string
	for i := range s.topics {
		topics = append(topics, i)
	}
	return topics
}
