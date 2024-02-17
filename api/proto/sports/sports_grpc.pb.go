// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: sports/sports.proto

package sports

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// SportsClient is the client API for Sports service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SportsClient interface {
	// ListEvents returns a list of all events.
	ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error)
}

type sportsClient struct {
	cc grpc.ClientConnInterface
}

func NewSportsClient(cc grpc.ClientConnInterface) SportsClient {
	return &sportsClient{cc}
}

func (c *sportsClient) ListEvents(ctx context.Context, in *ListEventsRequest, opts ...grpc.CallOption) (*ListEventsResponse, error) {
	out := new(ListEventsResponse)
	err := c.cc.Invoke(ctx, "/sports.Sports/ListEvents", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SportsServer is the server API for Sports service.
// All implementations must embed UnimplementedSportsServer
// for forward compatibility
type SportsServer interface {
	// ListEvents returns a list of all events.
	ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error)
	mustEmbedUnimplementedSportsServer()
}

// UnimplementedSportsServer must be embedded to have forward compatible implementations.
type UnimplementedSportsServer struct {
}

func (UnimplementedSportsServer) ListEvents(context.Context, *ListEventsRequest) (*ListEventsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListEvents not implemented")
}
func (UnimplementedSportsServer) mustEmbedUnimplementedSportsServer() {}

// UnsafeSportsServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SportsServer will
// result in compilation errors.
type UnsafeSportsServer interface {
	mustEmbedUnimplementedSportsServer()
}

func RegisterSportsServer(s grpc.ServiceRegistrar, srv SportsServer) {
	s.RegisterService(&Sports_ServiceDesc, srv)
}

func _Sports_ListEvents_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListEventsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SportsServer).ListEvents(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sports.Sports/ListEvents",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SportsServer).ListEvents(ctx, req.(*ListEventsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Sports_ServiceDesc is the grpc.ServiceDesc for Sports service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Sports_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "sports.Sports",
	HandlerType: (*SportsServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListEvents",
			Handler:    _Sports_ListEvents_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sports/sports.proto",
}
