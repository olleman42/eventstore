package eventstore

import (
	"errors"
	"io"

	"github.com/boltdb/bolt"
)

// GetAggregateHistory returns a byte stream of timestime ordered aggregate events
// AggregateID and AggregateType are required
func (e *EventStore) GetAggregateHistory(aggType, aggID string, w io.Writer) error {
	return e.Connection.View(func(tx *bolt.Tx) error {
		tb := tx.Bucket([]byte(aggType))

		if tb == nil {
			return errors.New("Type does not exist in store: " + aggType)
		}

		ab := tb.Bucket([]byte(aggID))

		if ab == nil {
			return errors.New("Aggregate does not exist: " + aggID)
		}

		return dumpBucket(ab, w)
	})
}

// GetTypeHistory returns a byte stream of timestamp ordered aggregate event of a certain type
func (e *EventStore) GetTypeHistory(aggType string, w io.Writer) error {
	return e.Connection.View(func(tx *bolt.Tx) error {
		tb := tx.Bucket([]byte(aggType))

		if tb == nil {
			return errors.New("Type does not exist in store: " + aggType)
		}

		return dumpBucket(tb, w)
	})
}

// GetHistory return a byte stream of all the events in the system ordered by timestamp
func (e *EventStore) GetHistory(w io.Writer) error {
	// TODO Add functionality to also trickle out events that might have arrived since a read started, so no stray events disappear
	return e.Connection.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("events"))

		return dumpBucket(b, w)

	})
}
