// Code generated by protoc-gen-go.
// source: eventstore.proto
// DO NOT EDIT!

/*
Package eventstore is a generated protocol buffer package.

It is generated from these files:
	eventstore.proto

It has these top-level messages:
	GetHistoryRequest
	GetTypeHistoryRequest
	GetAggregateHistoryRequest
	Event
	StoreEventRequest
	StoreEventResponse
*/
package eventstore

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type GetHistoryRequest struct {
}

func (m *GetHistoryRequest) Reset()                    { *m = GetHistoryRequest{} }
func (m *GetHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetHistoryRequest) ProtoMessage()               {}
func (*GetHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type GetTypeHistoryRequest struct {
	Type string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
}

func (m *GetTypeHistoryRequest) Reset()                    { *m = GetTypeHistoryRequest{} }
func (m *GetTypeHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetTypeHistoryRequest) ProtoMessage()               {}
func (*GetTypeHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *GetTypeHistoryRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

type GetAggregateHistoryRequest struct {
	Type        string `protobuf:"bytes,1,opt,name=Type" json:"Type,omitempty"`
	AggregateID string `protobuf:"bytes,2,opt,name=AggregateID" json:"AggregateID,omitempty"`
}

func (m *GetAggregateHistoryRequest) Reset()                    { *m = GetAggregateHistoryRequest{} }
func (m *GetAggregateHistoryRequest) String() string            { return proto.CompactTextString(m) }
func (*GetAggregateHistoryRequest) ProtoMessage()               {}
func (*GetAggregateHistoryRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *GetAggregateHistoryRequest) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *GetAggregateHistoryRequest) GetAggregateID() string {
	if m != nil {
		return m.AggregateID
	}
	return ""
}

type Event struct {
	Data []byte `protobuf:"bytes,1,opt,name=Data,proto3" json:"Data,omitempty"`
}

func (m *Event) Reset()                    { *m = Event{} }
func (m *Event) String() string            { return proto.CompactTextString(m) }
func (*Event) ProtoMessage()               {}
func (*Event) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Event) GetData() []byte {
	if m != nil {
		return m.Data
	}
	return nil
}

type StoreEventRequest struct {
	Event []byte `protobuf:"bytes,1,opt,name=Event,proto3" json:"Event,omitempty"`
}

func (m *StoreEventRequest) Reset()                    { *m = StoreEventRequest{} }
func (m *StoreEventRequest) String() string            { return proto.CompactTextString(m) }
func (*StoreEventRequest) ProtoMessage()               {}
func (*StoreEventRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

func (m *StoreEventRequest) GetEvent() []byte {
	if m != nil {
		return m.Event
	}
	return nil
}

type StoreEventResponse struct {
}

func (m *StoreEventResponse) Reset()                    { *m = StoreEventResponse{} }
func (m *StoreEventResponse) String() string            { return proto.CompactTextString(m) }
func (*StoreEventResponse) ProtoMessage()               {}
func (*StoreEventResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func init() {
	proto.RegisterType((*GetHistoryRequest)(nil), "eventstore.GetHistoryRequest")
	proto.RegisterType((*GetTypeHistoryRequest)(nil), "eventstore.GetTypeHistoryRequest")
	proto.RegisterType((*GetAggregateHistoryRequest)(nil), "eventstore.GetAggregateHistoryRequest")
	proto.RegisterType((*Event)(nil), "eventstore.Event")
	proto.RegisterType((*StoreEventRequest)(nil), "eventstore.StoreEventRequest")
	proto.RegisterType((*StoreEventResponse)(nil), "eventstore.StoreEventResponse")
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// Client API for EventStore service

type EventStoreClient interface {
	GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (EventStore_GetHistoryClient, error)
	GetTypeHistory(ctx context.Context, in *GetTypeHistoryRequest, opts ...grpc.CallOption) (EventStore_GetTypeHistoryClient, error)
	GetAggregateHistory(ctx context.Context, in *GetAggregateHistoryRequest, opts ...grpc.CallOption) (EventStore_GetAggregateHistoryClient, error)
	StoreEvent(ctx context.Context, in *StoreEventRequest, opts ...grpc.CallOption) (*StoreEventResponse, error)
}

type eventStoreClient struct {
	cc *grpc.ClientConn
}

func NewEventStoreClient(cc *grpc.ClientConn) EventStoreClient {
	return &eventStoreClient{cc}
}

func (c *eventStoreClient) GetHistory(ctx context.Context, in *GetHistoryRequest, opts ...grpc.CallOption) (EventStore_GetHistoryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventStore_serviceDesc.Streams[0], c.cc, "/eventstore.EventStore/GetHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStoreGetHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStore_GetHistoryClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventStoreGetHistoryClient struct {
	grpc.ClientStream
}

func (x *eventStoreGetHistoryClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventStoreClient) GetTypeHistory(ctx context.Context, in *GetTypeHistoryRequest, opts ...grpc.CallOption) (EventStore_GetTypeHistoryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventStore_serviceDesc.Streams[1], c.cc, "/eventstore.EventStore/GetTypeHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStoreGetTypeHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStore_GetTypeHistoryClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventStoreGetTypeHistoryClient struct {
	grpc.ClientStream
}

func (x *eventStoreGetTypeHistoryClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventStoreClient) GetAggregateHistory(ctx context.Context, in *GetAggregateHistoryRequest, opts ...grpc.CallOption) (EventStore_GetAggregateHistoryClient, error) {
	stream, err := grpc.NewClientStream(ctx, &_EventStore_serviceDesc.Streams[2], c.cc, "/eventstore.EventStore/GetAggregateHistory", opts...)
	if err != nil {
		return nil, err
	}
	x := &eventStoreGetAggregateHistoryClient{stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

type EventStore_GetAggregateHistoryClient interface {
	Recv() (*Event, error)
	grpc.ClientStream
}

type eventStoreGetAggregateHistoryClient struct {
	grpc.ClientStream
}

func (x *eventStoreGetAggregateHistoryClient) Recv() (*Event, error) {
	m := new(Event)
	if err := x.ClientStream.RecvMsg(m); err != nil {
		return nil, err
	}
	return m, nil
}

func (c *eventStoreClient) StoreEvent(ctx context.Context, in *StoreEventRequest, opts ...grpc.CallOption) (*StoreEventResponse, error) {
	out := new(StoreEventResponse)
	err := grpc.Invoke(ctx, "/eventstore.EventStore/StoreEvent", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for EventStore service

type EventStoreServer interface {
	GetHistory(*GetHistoryRequest, EventStore_GetHistoryServer) error
	GetTypeHistory(*GetTypeHistoryRequest, EventStore_GetTypeHistoryServer) error
	GetAggregateHistory(*GetAggregateHistoryRequest, EventStore_GetAggregateHistoryServer) error
	StoreEvent(context.Context, *StoreEventRequest) (*StoreEventResponse, error)
}

func RegisterEventStoreServer(s *grpc.Server, srv EventStoreServer) {
	s.RegisterService(&_EventStore_serviceDesc, srv)
}

func _EventStore_GetHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStoreServer).GetHistory(m, &eventStoreGetHistoryServer{stream})
}

type EventStore_GetHistoryServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventStoreGetHistoryServer struct {
	grpc.ServerStream
}

func (x *eventStoreGetHistoryServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventStore_GetTypeHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetTypeHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStoreServer).GetTypeHistory(m, &eventStoreGetTypeHistoryServer{stream})
}

type EventStore_GetTypeHistoryServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventStoreGetTypeHistoryServer struct {
	grpc.ServerStream
}

func (x *eventStoreGetTypeHistoryServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventStore_GetAggregateHistory_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(GetAggregateHistoryRequest)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(EventStoreServer).GetAggregateHistory(m, &eventStoreGetAggregateHistoryServer{stream})
}

type EventStore_GetAggregateHistoryServer interface {
	Send(*Event) error
	grpc.ServerStream
}

type eventStoreGetAggregateHistoryServer struct {
	grpc.ServerStream
}

func (x *eventStoreGetAggregateHistoryServer) Send(m *Event) error {
	return x.ServerStream.SendMsg(m)
}

func _EventStore_StoreEvent_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(StoreEventRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(EventStoreServer).StoreEvent(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/eventstore.EventStore/StoreEvent",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(EventStoreServer).StoreEvent(ctx, req.(*StoreEventRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _EventStore_serviceDesc = grpc.ServiceDesc{
	ServiceName: "eventstore.EventStore",
	HandlerType: (*EventStoreServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "StoreEvent",
			Handler:    _EventStore_StoreEvent_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "GetHistory",
			Handler:       _EventStore_GetHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetTypeHistory",
			Handler:       _EventStore_GetTypeHistory_Handler,
			ServerStreams: true,
		},
		{
			StreamName:    "GetAggregateHistory",
			Handler:       _EventStore_GetAggregateHistory_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "eventstore.proto",
}

func init() { proto.RegisterFile("eventstore.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 259 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x92, 0x5f, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x9b, 0xa2, 0x82, 0x9f, 0x22, 0x66, 0x5a, 0xa1, 0x44, 0x94, 0xba, 0x0f, 0xa2, 0x08,
	0x45, 0xf4, 0x04, 0x4a, 0x25, 0x2a, 0xf8, 0x12, 0x7b, 0x81, 0x08, 0x43, 0xf0, 0xa5, 0x89, 0xc9,
	0x28, 0xf4, 0x78, 0xde, 0x4c, 0x76, 0x52, 0xcd, 0x26, 0x71, 0xa1, 0x6f, 0xbb, 0xf3, 0xfd, 0xd9,
	0xc9, 0x8f, 0xe0, 0x90, 0xbf, 0x78, 0x29, 0x95, 0xe4, 0x25, 0xcf, 0x8a, 0x32, 0x97, 0x9c, 0xd0,
	0x4c, 0xcc, 0x08, 0x61, 0xcc, 0xf2, 0xf8, 0x6e, 0x6f, 0xab, 0x84, 0x3f, 0x3e, 0xb9, 0x12, 0x73,
	0x85, 0xa3, 0x98, 0x65, 0xb1, 0x2a, 0xb8, 0x2d, 0x10, 0x61, 0xcb, 0x4e, 0x27, 0xc1, 0x34, 0xb8,
	0xd8, 0x4d, 0xf4, 0x6c, 0x12, 0x44, 0x31, 0xcb, 0x5d, 0x96, 0x95, 0x9c, 0xa5, 0xb2, 0x41, 0x82,
	0xa6, 0xd8, 0xfb, 0xb3, 0x3f, 0xcd, 0x27, 0x43, 0x95, 0xdc, 0x91, 0x39, 0xc6, 0xf6, 0x83, 0xdd,
	0xd1, 0xc6, 0xe7, 0xa9, 0xa4, 0x1a, 0xdf, 0x4f, 0xf4, 0x6c, 0x2e, 0x11, 0xbe, 0xda, 0xdd, 0xd5,
	0xf1, 0xfb, 0xce, 0x78, 0x9d, 0x58, 0x3b, 0xeb, 0x8b, 0x19, 0x83, 0x5c, 0x6b, 0x55, 0xe4, 0xcb,
	0x8a, 0x6f, 0xbe, 0x87, 0x80, 0x4e, 0x54, 0xa3, 0x7b, 0xa0, 0x41, 0x40, 0x27, 0x33, 0x87, 0x57,
	0x0f, 0x4d, 0x14, 0xba, 0x72, 0xfd, 0xc8, 0xe0, 0x3a, 0xa0, 0x67, 0x1c, 0xb4, 0x89, 0xd1, 0x59,
	0xa7, 0xa7, 0x4f, 0xd3, 0xd7, 0xb5, 0xc0, 0xe8, 0x1f, 0xa0, 0x74, 0xde, 0x29, 0xf4, 0x10, 0xf7,
	0xb5, 0xbe, 0x00, 0x0d, 0x8a, 0xf6, 0x57, 0xf6, 0x68, 0x46, 0xa7, 0x3e, 0xb9, 0x26, 0x68, 0x06,
	0x6f, 0x3b, 0xfa, 0x2b, 0xdd, 0xfe, 0x04, 0x00, 0x00, 0xff, 0xff, 0x92, 0x1b, 0x76, 0x3e, 0x5e,
	0x02, 0x00, 0x00,
}
