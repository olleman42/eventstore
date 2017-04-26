package event

import (
	"encoding/json"
	"time"
)

type Event struct {
	ID            string
	AggregateID   string
	AggregateType string
	Timestamp     time.Time
	EventType     string
}

func GetGenericEvent(in []byte) (event Event, err error) {
	event = Event{}
	err = json.Unmarshal(in, &event)
	return
}
