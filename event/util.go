package event

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

var (
	// ErrorUnknownEvent signals that there is no valid handler for incoming event type
	ErrorUnknownEvent = errors.New("Unknown event type added")
)

// FeedArrayToJSONStream Takes a json array and writes its contents to a writer
func FeedArrayToJSONStream(in []byte, w io.Writer) error {
	events := &[]json.RawMessage{}
	err := json.Unmarshal(in, events)
	if err != nil {
		return err
	}
	for _, v := range *events {
		w.Write(v)
	}
	return nil
}

// DecodeJSONArrayToEvents decodes JSON encoded event stream
func DecodeJSONArrayToEvents(in []byte, deserializeEvent func([]byte) (interface{}, error)) ([]interface{}, error) {

	genEvents := &[]json.RawMessage{}
	outEvents := []interface{}{}

	json.Unmarshal(in, &genEvents)

	for _, v := range *genEvents {
		typedEvent, err := deserializeEvent(v)
		if err != nil {
			return nil, err
		}
		outEvents = append(outEvents, typedEvent)
	}

	return outEvents, nil
}

// DecodeStreamToEvents decodes a stream of json events (not array) to typed events
func DecodeStreamToEvents(r io.Reader, deserializeEvent func([]byte) (interface{}, error)) ([]interface{}, error) {
	outEvents := []interface{}{}
	dec := json.NewDecoder(r)

	for dec.More() {
		var rawEvent json.RawMessage
		if err := dec.Decode(&rawEvent); err != nil {
			return nil, err
		}

		typedEvent, err := deserializeEvent(rawEvent)
		if err != nil {
			if err == ErrorUnknownEvent {
				fmt.Println(err)
			} else {
				return nil, err
			}
		}
		outEvents = append(outEvents, typedEvent)
	}

	return outEvents, nil
}

// DecodeStreamToApplyableEvents decodes a stream of json events (inline JSON) to typed applyable events
func DecodeStreamToApplyableEvents(r io.Reader, deserializeEvent func([]byte) (Applyable, error)) ([]Applyable, error) {
	outEvents := []Applyable{}
	dec := json.NewDecoder(r)

	for dec.More() {
		var rawEvent json.RawMessage
		if err := dec.Decode(&rawEvent); err != nil {
			return nil, err
		}

		typedEvent, err := deserializeEvent(rawEvent)
		if err != nil {
			if err == ErrorUnknownEvent {
				fmt.Println(err)
			} else {
				return nil, err
			}
		}
		outEvents = append(outEvents, typedEvent.(Applyable))
	}

	return outEvents, nil
}

// DecodeStreamToChan decodes and event stream to typed events that are then published on a specified channel
func DecodeStreamToChan(r io.Reader, sink chan Applyable, deserializeEvent func([]byte) (Applyable, error)) {
	dec := json.NewDecoder(r)

	for dec.More() {
		var rawEvent json.RawMessage
		if err := dec.Decode(&rawEvent); err != nil {
			fmt.Println("Failed decoing event JSON, resetting encoder")
			dec = json.NewDecoder(r)
		}
		typedEvent, err := deserializeEvent(rawEvent)
		if err != nil {
			if err == ErrorUnknownEvent {
				fmt.Println(err)
				continue
			}
			fmt.Println(err)
			continue
		}
		sink <- typedEvent.(Applyable)
	}
	close(sink)
}
