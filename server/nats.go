package server

import (
	"log"

	nats "github.com/nats-io/go-nats"
	"github.com/olleman42/dinnerbot/eventstore"
)

// RegisterNATS registers a listener and a publisher on the NATS broker when provided a valid connection and event store
func RegisterNATS(nc *nats.Conn, ess *eventstore.EventStore) error {
	// TODO: Allow consumer to specify topics for receiving and dispatching events
	if err := ess.AddListener(func(in []byte) error {
		log.Printf("Dispatching event: %s", string(in))
		return nc.Publish("events", in)
	}); err != nil {
		return err
	}

	if _, err := nc.Subscribe("dispatched", func(m *nats.Msg) {
		ess.Ingest <- m.Data
	}); err != nil {
		return err
	}

	return nil
}
