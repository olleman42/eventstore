package event

import (
	"encoding/json"
	"time"
)

// Event Generic event type used to shape events to store int the store
type Event struct {
	ID            string
	AggregateID   string
	AggregateType string
	Timestamp     time.Time
	EventType     string
}

// GetGenericEvent returns a struct containing data required to categorize and store an event
func GetGenericEvent(in []byte) (event Event, err error) {
	event = Event{}
	err = json.Unmarshal(in, &event)
	return
}
