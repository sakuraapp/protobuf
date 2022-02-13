// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package supervisorpb

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

// SupervisorServiceClient is the client API for SupervisorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type SupervisorServiceClient interface {
	Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployResponse, error)
}

type supervisorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewSupervisorServiceClient(cc grpc.ClientConnInterface) SupervisorServiceClient {
	return &supervisorServiceClient{cc}
}

func (c *supervisorServiceClient) Deploy(ctx context.Context, in *DeployRequest, opts ...grpc.CallOption) (*DeployResponse, error) {
	out := new(DeployResponse)
	err := c.cc.Invoke(ctx, "/service.SupervisorService/Deploy", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SupervisorServiceServer is the server API for SupervisorService service.
// All implementations must embed UnimplementedSupervisorServiceServer
// for forward compatibility
type SupervisorServiceServer interface {
	Deploy(context.Context, *DeployRequest) (*DeployResponse, error)
	mustEmbedUnimplementedSupervisorServiceServer()
}

// UnimplementedSupervisorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedSupervisorServiceServer struct {
}

func (UnimplementedSupervisorServiceServer) Deploy(context.Context, *DeployRequest) (*DeployResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Deploy not implemented")
}
func (UnimplementedSupervisorServiceServer) mustEmbedUnimplementedSupervisorServiceServer() {}

// UnsafeSupervisorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to SupervisorServiceServer will
// result in compilation errors.
type UnsafeSupervisorServiceServer interface {
	mustEmbedUnimplementedSupervisorServiceServer()
}

func RegisterSupervisorServiceServer(s grpc.ServiceRegistrar, srv SupervisorServiceServer) {
	s.RegisterService(&SupervisorService_ServiceDesc, srv)
}

func _SupervisorService_Deploy_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DeployRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SupervisorServiceServer).Deploy(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/service.SupervisorService/Deploy",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SupervisorServiceServer).Deploy(ctx, req.(*DeployRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// SupervisorService_ServiceDesc is the grpc.ServiceDesc for SupervisorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var SupervisorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "service.SupervisorService",
	HandlerType: (*SupervisorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Deploy",
			Handler:    _SupervisorService_Deploy_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "supervisor/service.proto",
}
