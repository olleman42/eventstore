package server

import (
	context "golang.org/x/net/context"

	"google.golang.org/grpc"

	"github.com/olleman42/dinnerbot/eventstore"
	pb "github.com/olleman42/dinnerbot/eventstore/proto"
)

type streamSender interface {
	Send(*pb.Event) error
}

type grpcEventWriter struct {
	sender streamSender
}

func (gw *grpcEventWriter) Write(p []byte) (int, error) {
	return len(p), gw.sender.Send(&pb.Event{Data: p})
}

type gserver struct {
	store *eventstore.EventStore
}

func (g *gserver) GetHistory(req *pb.GetHistoryRequest, stream pb.EventStore_GetHistoryServer) error {
	grpcWriter := &grpcEventWriter{sender: stream}
	return g.store.GetHistory(grpcWriter)
}

func (g *gserver) GetTypeHistory(req *pb.GetTypeHistoryRequest, stream pb.EventStore_GetTypeHistoryServer) error {
	grpcWriter := &grpcEventWriter{sender: stream}
	return g.store.GetTypeHistory(req.Type, grpcWriter)
}

func (g *gserver) GetAggregateHistory(req *pb.GetAggregateHistoryRequest, stream pb.EventStore_GetAggregateHistoryServer) error {
	grpcWriter := &grpcEventWriter{sender: stream}
	return g.store.GetAggregateHistory(req.Type, req.AggregateID, grpcWriter)
}

func (g *gserver) StoreEvent(ctx context.Context, req *pb.StoreEventRequest) (*pb.StoreEventResponse, error) {
	err := g.store.StoreEvent(req.Event)
	if err != nil {
		return &pb.StoreEventResponse{}, err
	}
	return &pb.StoreEventResponse{}, nil
}

// RegisterGRPCServer Get a grpc server connection (wrapping a listener) and an event store to provide client-compatible endpoints for querying event store
func RegisterGRPCServer(grpcServer *grpc.Server, ess *eventstore.EventStore) {
	gserver := &gserver{store: ess}
	pb.RegisterEventStoreServer(grpcServer, gserver)
}
