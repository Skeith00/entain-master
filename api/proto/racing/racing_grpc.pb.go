// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v4.23.0
// source: racing/racing.proto

package racing

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

// RacingClient is the client API for Racing service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type RacingClient interface {
	// ListRaces returns a list of all races.
	ListRaces(ctx context.Context, in *ListRacesRequest, opts ...grpc.CallOption) (*ListRacesResponse, error)
	// GetRace will return a single race by its ID.
	GetRace(ctx context.Context, in *GetRaceRequest, opts ...grpc.CallOption) (*Race, error)
}

type racingClient struct {
	cc grpc.ClientConnInterface
}

func NewRacingClient(cc grpc.ClientConnInterface) RacingClient {
	return &racingClient{cc}
}

func (c *racingClient) ListRaces(ctx context.Context, in *ListRacesRequest, opts ...grpc.CallOption) (*ListRacesResponse, error) {
	out := new(ListRacesResponse)
	err := c.cc.Invoke(ctx, "/racing.Racing/ListRaces", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *racingClient) GetRace(ctx context.Context, in *GetRaceRequest, opts ...grpc.CallOption) (*Race, error) {
	out := new(Race)
	err := c.cc.Invoke(ctx, "/racing.Racing/GetRace", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RacingServer is the server API for Racing service.
// All implementations must embed UnimplementedRacingServer
// for forward compatibility
type RacingServer interface {
	// ListRaces returns a list of all races.
	ListRaces(context.Context, *ListRacesRequest) (*ListRacesResponse, error)
	// GetRace will return a single race by its ID.
	GetRace(context.Context, *GetRaceRequest) (*Race, error)
	mustEmbedUnimplementedRacingServer()
}

// UnimplementedRacingServer must be embedded to have forward compatible implementations.
type UnimplementedRacingServer struct {
}

func (UnimplementedRacingServer) ListRaces(context.Context, *ListRacesRequest) (*ListRacesResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListRaces not implemented")
}
func (UnimplementedRacingServer) GetRace(context.Context, *GetRaceRequest) (*Race, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetRace not implemented")
}
func (UnimplementedRacingServer) mustEmbedUnimplementedRacingServer() {}

// UnsafeRacingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RacingServer will
// result in compilation errors.
type UnsafeRacingServer interface {
	mustEmbedUnimplementedRacingServer()
}

func RegisterRacingServer(s grpc.ServiceRegistrar, srv RacingServer) {
	s.RegisterService(&Racing_ServiceDesc, srv)
}

func _Racing_ListRaces_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListRacesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RacingServer).ListRaces(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/racing.Racing/ListRaces",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RacingServer).ListRaces(ctx, req.(*ListRacesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Racing_GetRace_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetRaceRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RacingServer).GetRace(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/racing.Racing/GetRace",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RacingServer).GetRace(ctx, req.(*GetRaceRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// Racing_ServiceDesc is the grpc.ServiceDesc for Racing service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Racing_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "racing.Racing",
	HandlerType: (*RacingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ListRaces",
			Handler:    _Racing_ListRaces_Handler,
		},
		{
			MethodName: "GetRace",
			Handler:    _Racing_GetRace_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "racing/racing.proto",
}
