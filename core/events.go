package core

import "fmt"

// Event represents a system event
type Event struct {
	Name string
	Data interface{}
}

// EventBus handles event-driven communication
type EventBus struct {}

// NewEventBus creates a new event bus
func NewEventBus() *EventBus {
	return &EventBus{}
}

// Publish publishes an event
func (e *EventBus) Publish(event Event) {
	fmt.Printf("Event '%s' published with data: %v\n", event.Name, event.Data)
}
