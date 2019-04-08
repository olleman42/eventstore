package client

import (
	"io"
	"log"

	context "golang.org/x/net/context"

	"encoding/json"

	"google.golang.org/grpc"

	"eventstore/event"
	pb "eventstore/proto"
)

// ESClient a wrapper library to easily query the event store
type ESClient struct {
	Client pb.EventStoreClient
}

type receiver interface {
	Recv() (*pb.Event, error)
}

// NewESClient returns a new instance of the wrapper
func NewESClient(conn *grpc.ClientConn) *ESClient {
	eventStoreClient := pb.NewEventStoreClient(conn)
	return &ESClient{eventStoreClient}
}

func connectPipe(stream receiver) io.Reader {
	pr, pw := io.Pipe()
	go func(w *io.PipeWriter) {
		for {
			event, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			w.Write(event.Data)

		}
	}(pw)
	return pr
}

// GetAggregateHistory Gets reader with JSON data stream to deliver event data for specific aggregate
func (esc *ESClient) GetAggregateHistory(aggregateType, aggregateID string) (io.Reader, error) {
	stream, err := esc.Client.GetAggregateHistory(
		context.Background(),
		&pb.GetAggregateHistoryRequest{AggregateID: aggregateID, Type: aggregateType},
	)
	if err != nil {
		return nil, err
	}
	return connectPipe(stream), nil
}

// GetTypeHistory returns reader with JSON data stream with the history of all events of a certain type
func (esc *ESClient) GetTypeHistory(aggregateType string) (io.Reader, error) {
	stream, err := esc.Client.GetTypeHistory(
		context.Background(),
		&pb.GetTypeHistoryRequest{Type: aggregateType},
	)
	if err != nil {
		return nil, err
	}
	return connectPipe(stream), nil
}

// GetHistory returns a JSON stream of all the events in the store
func (esc *ESClient) GetHistory() (io.Reader, error) {
	stream, err := esc.Client.GetHistory(
		context.Background(),
		&pb.GetHistoryRequest{},
	)
	if err != nil {
		return nil, err
	}
	return connectPipe(stream), nil
}

// StoreEvent stores specified event in the store
func (esc *ESClient) StoreEvent(event event.Event) error {
	encoded, err := json.Marshal(event)
	if err != nil {
		return err
	}
	_, err = esc.Client.StoreEvent(context.Background(), &pb.StoreEventRequest{Event: encoded})
	return err
}
