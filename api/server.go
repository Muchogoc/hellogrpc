package api

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// PingServer is the server API for Ping service.
// All implementations must embed ImplementedPingServer
// for forward compatibility
type PingServer interface {
	SayHello(ctx context.Context, message *PingMessage) (*PingMessage, error)
	mustEmbedImplementedPingServer()
}

// ImplementedPingServer must be embedded to have forward compatible implementations.
type ImplementedPingServer struct {
}

func (ImplementedPingServer) SayHello(ctx context.Context, message *PingMessage) (*PingMessage, error) {
	log.Printf("Received body from client: %s", message)
	returnMessage := PingMessage{
		Greeting: "Hello from server",
	}
	return &returnMessage, nil
}
func (ImplementedPingServer) mustEmbedImplementedPingServer() {}

// UnsafePingServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to PingServer will
// result in compilation errors.
type UnsafePingServer interface {
	mustEmbedImplementedPingServer()
}

func RegisterPingServer(s *grpc.Server, srv PingServer) {
	s.RegisterService(&_Ping_serviceDesc, srv)
}

func _Ping_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PingMessage)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(PingServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/api.Ping/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(PingServer).SayHello(ctx, req.(*PingMessage))
	}
	return interceptor(ctx, in, info, handler)
}

var _Ping_serviceDesc = grpc.ServiceDesc{
	ServiceName: "api.Ping",
	HandlerType: (*PingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _Ping_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api.proto",
}
