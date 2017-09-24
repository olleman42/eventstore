eventstore [![GoDoc](https://godoc.org/github.com/olleman42/eventstore?status.svg)](https://godoc.org/github.com/olleman42/eventstore)
==========

eventstore is a boltdb-backed JSON event store with a small API to facilitate making simple CQRS applications

The package is meant to be used as a library while letting the user dictate how the event store is to be exposed. There are helper server and client bindings to enable simpe communication with the store from other processes

The package is in an early stage of development, so tests are missing and optimization is lacking.

Events are currently indexed by three criteria (aggregate type, aggregate id and date) and the data store is not optimized, so one stored event takes up three times the size of the event itself.

The required fields for an event store properly are the following:

```json
{
    "EventType": "Event type",
    "AggregateID": "Aggregate ID",
    "AggregateType": "Aggregate Type",
    "Timestamp": "2016-11-05T14:53:24.487Z",
}
```

Events being handled in other processes can leverage embedding the ```Event``` type and related helpers from the ```event``` package to simplify serialization of events for the store