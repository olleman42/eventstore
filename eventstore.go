package eventstore

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"

	"bytes"

	"os"

	"time"

	"sync"

	"github.com/boltdb/bolt"
	"github.com/olleman42/dinnerbot/eventstore/event"
)

type safeListeners struct {
	*sync.RWMutex
	Listeners []func([]byte) error
}

func (s *safeListeners) DoAction(action func(listeners *safeListeners) error) error {
	s.Lock()
	defer s.Unlock()
	return action(s)
}

// EventStore holds helper methods to store events
type EventStore struct {
	Connection *bolt.DB
	Ingest     chan []byte
	errors     chan error
	broadcast  chan []byte
	listeners  *safeListeners
}

// NewEventStore creates and returns new event store
func NewEventStore() (*EventStore, error) {
	store := &EventStore{}
	if err := store.connect(); err != nil {
		return store, err
	}

	store.Ingest = make(chan []byte)
	go store.consumeEvents()

	store.listeners = &safeListeners{&sync.RWMutex{}, make([]func([]byte) error, 0)}
	store.broadcast = make(chan []byte)
	go store.feedListeners()

	store.errors = make(chan error)
	go store.handleErrors()

	return store, nil
}

func (e *EventStore) consumeEvents() {
	buffer := bytes.NewBuffer([]byte{})
	decoder := json.NewDecoder(buffer)

	for v := range e.Ingest {
		buffer.Write(v)

		for decoder.More() {
			var rawEvent json.RawMessage
			if err := decoder.Decode(&rawEvent); err != nil {
				e.errors <- err
				decoder = json.NewDecoder(buffer)
				continue
			}

			if err := e.StoreEvent(rawEvent); err != nil {
				e.errors <- err
			}
		}

	}
}

func (e *EventStore) handleErrors() {
	for err := range e.errors {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

// Creates the database connection
func (e *EventStore) connect() error {
	db, err := bolt.Open("events.db", 0600, nil)
	e.Connection = db

	return err
}

// ImportEvents - Imports events from json stream
func (e *EventStore) ImportEvents(r io.Reader) error {
	decoder := json.NewDecoder(r)

	for decoder.More() {
		var rawEvent json.RawMessage
		if err := decoder.Decode(&rawEvent); err != nil {
			return err
		}

		if err := e.StoreEvent(rawEvent); err != nil {
			return err
		}
	}

	return nil
}

// Close - allow closing manually
func (e *EventStore) Close() {
	e.Connection.Close()
}

func (e *EventStore) Write(in []byte) (int, error) {
	e.Ingest <- in
	return len(in), nil
}

func (e *EventStore) cleanEvent(in []byte) ([]byte, error) {
	var generic interface{}
	if err := json.Unmarshal(in, generic); err != nil {
		return nil, err
	}

	return json.MarshalIndent(generic, "", "  ")

}

// StoreEvent - store an event in its appropriate bucket
// Relevant metadata is required
func (e *EventStore) StoreEvent(in []byte) error {
	// get generic event
	// TODO: Check event version of latest event version to ensure that no conflict occurs
	event, err := event.GetGenericEvent(in) // TODO: Break out event-specific behaviour to interface
	if err != nil {
		return err
	}

	// validate that all required values are set
	if event.AggregateID == "" {
		return errors.New("Invalid Aggregate ID")
	}

	if event.AggregateType == "" {
		return errors.New("Invalid Aggregate Type")
	}

	if event.Timestamp.IsZero() {
		return errors.New("Invalid event timestamp")
	}

	// clean up the event before storage
	in, err = e.cleanEvent(in)
	if err != nil {
		return err
	}

	err = e.Connection.Update(func(tx *bolt.Tx) error {
		// create key based on unix timestamp of event (for easier querying)
		key := []byte(fmt.Sprintf("%v", event.Timestamp.Format(time.RFC3339)))
		// key := []byte(fmt.Sprintf("%v", event.Timestamp.Unix()))

		// we need to index by aggtype->aggid->time, global->time, aggtype->time

		// get global bucket
		ebucket, err := tx.CreateBucketIfNotExists([]byte("events"))
		if err != nil {
			return err
		}

		// get aggtype bucket
		b, err := tx.CreateBucketIfNotExists([]byte(event.AggregateType))
		if err != nil {
			return fmt.Errorf("Failed opening bucket: %s", err)
		}

		// get type->id bucket
		sb, err := b.CreateBucketIfNotExists([]byte(event.AggregateID))
		if err != nil {
			return fmt.Errorf("Failed opening sub-bucket: %s", err)
		}

		// append on global bucket
		if err := ebucket.Put(key, in); err != nil {
			return err
		}

		// append event by id on type bucket
		if err := b.Put(key, in); err != nil {
			return err
		}

		// append on type->id buccket
		return sb.Put(key, in)
	})
	if err != nil {
		return err
	}

	return e.broadcastEvent(in)
}

// DumpEvents - dump events to writer
func (e *EventStore) DumpEvents(w io.Writer) error {
	return e.Connection.View(func(tx *bolt.Tx) error {
		eb := tx.Bucket([]byte("events"))

		return eb.ForEach(func(k, v []byte) error {
			_, err := fmt.Fprint(w, string(v))
			return err
		})
	})
}

func dumpBucket(b *bolt.Bucket, w io.Writer) error {
	return b.ForEach(func(k, v []byte) error {
		_, err := w.Write(v)
		return err
	})
}

func (e *EventStore) broadcastEvent(event []byte) error {
	e.broadcast <- event
	return nil
}

func (e *EventStore) feedListeners() {
	for v := range e.broadcast {
		err := e.listeners.DoAction(func(listeners *safeListeners) error {
			for _, listener := range e.listeners.Listeners {
				err := listener(v)
				if err != nil {
					return err
				}
			}
			return nil
		})
		if err != nil {
			e.errors <- err
		}
	}
}

// AddListener adds a callback function to be run when a new event gets emitted
func (e *EventStore) AddListener(listener func([]byte) error) error {
	return e.listeners.DoAction(func(listeners *safeListeners) error {
		listeners.Listeners = append(listeners.Listeners, listener)
		return nil
	})
}
