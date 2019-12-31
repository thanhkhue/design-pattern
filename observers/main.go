package main

import (
	"log"
	"sync"
)

// Observer is a observer
type Observer struct {
	ID int
}

// ObserverInterface observer interface
type ObserverInterface interface {
	send(message string)
}

func (o *Observer) send(message string) {
	log.Printf("%d received message %s\n", o.ID, message)
}

// Subject is subject in Observer system, which is maintains a record of observers
type Subject struct {
	observers   []*Observer
	observerIDs map[int]bool
	mu          *sync.RWMutex
}

// NewSubject creates new subject
func NewSubject() *Subject {
	return &Subject{
		observers:   make([]*Observer, 0),
		observerIDs: make(map[int]bool),
		mu:          new(sync.RWMutex),
	}
}

// Add adds new observer to pool
// return false whether observer already exists in list, else return true
func (s *Subject) Add(o *Observer) bool {
	// check is o already exists
	s.mu.RLock()
	if s.observerIDs[o.ID] {
		return false
	}
	s.mu.RUnlock()

	// if not exists, add to list
	s.mu.Lock()
	s.observers = append(s.observers, o)
	s.observerIDs[o.ID] = true
	s.mu.Unlock()
	return true
}

// Remove removes an observer from list
func (s *Subject) Remove(o *Observer) bool {
	s.mu.RLock()
	if !s.observerIDs[o.ID] {
		return false
	}
	s.mu.RUnlock()

	s.mu.Lock()
	delete(s.observerIDs, o.ID)
	for i, v := range s.observers {
		if v == o {
			s.observers = append(s.observers[:i], s.observers[i+1:]...)
			break
		}
	}
	s.mu.Unlock()
	return true
}

// Notify sends message to observers
func (s *Subject) Notify(message string) {
	s.mu.RLock()
	observers := s.observers
	s.mu.RUnlock()

	for _, observer := range observers {
		if observer != nil {
			observer.send(message)
		}
	}
}

func main() {
	sub := NewSubject()
	// listen to subject
	for i := 1; i <= 5; i++ {
		sub.Add(&Observer{
			ID: i,
		})
	}

	// notifies to observers
	sub.Notify("Hi there")
}
