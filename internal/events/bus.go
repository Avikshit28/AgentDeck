package events

import (
	"errors"
	"sync"
)

// ErrClosed is returned by Publish after the bus is closed .
var ErrClosed = errors.New("events: bus closed")

// Bus carries events from many agents to one reader.
type Bus struct {
	ch   chan Event
	done chan struct{}
	once sync.Once
}

// NewBus creates a bus that can hold 'buffer' unread events.
func NewBus(buffer int) *Bus {
	return &Bus{ch: make(chan Event, buffer), done: make(chan struct{})}
}

// Publish sends e, waiting if the bus is full. Fails once the bus is closed.
func (b *Bus) Publish(e Event) error {
	select {
	case <-b.done:
		return ErrClosed
	default:
	}
	select {
	case b.ch <- e:
		return nil
	case <-b.done:
		return ErrClosed
	}
}

// Events is where the reader receives events.
func (b *Bus) Events() <-chan Event { return b.ch }

// Done is closed when the bus is closed
func (b *Bus) Done() <-chan struct{} { return b.done }

// Close shuts the bus. Safe to call more than once.
func (b *Bus) Close() {
	b.once.Do(func() { close(b.done) })
}
