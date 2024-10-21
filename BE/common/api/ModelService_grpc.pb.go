// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: api/ModelService.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	ModelService_TextEmbedding_FullMethodName  = "/ModelService/TextEmbedding"
	ModelService_ImageEmbedding_FullMethodName = "/ModelService/ImageEmbedding"
)

// ModelServiceClient is the client API for ModelService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ModelServiceClient interface {
	TextEmbedding(ctx context.Context, in *TextEmbeddingRequest, opts ...grpc.CallOption) (*TextEmbeddingResponse, error)
	ImageEmbedding(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ImageEmbeddingRequest, ImageEmbeddingResponse], error)
}

type modelServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewModelServiceClient(cc grpc.ClientConnInterface) ModelServiceClient {
	return &modelServiceClient{cc}
}

func (c *modelServiceClient) TextEmbedding(ctx context.Context, in *TextEmbeddingRequest, opts ...grpc.CallOption) (*TextEmbeddingResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(TextEmbeddingResponse)
	err := c.cc.Invoke(ctx, ModelService_TextEmbedding_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *modelServiceClient) ImageEmbedding(ctx context.Context, opts ...grpc.CallOption) (grpc.ClientStreamingClient[ImageEmbeddingRequest, ImageEmbeddingResponse], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &ModelService_ServiceDesc.Streams[0], ModelService_ImageEmbedding_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[ImageEmbeddingRequest, ImageEmbeddingResponse]{ClientStream: stream}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ModelService_ImageEmbeddingClient = grpc.ClientStreamingClient[ImageEmbeddingRequest, ImageEmbeddingResponse]

// ModelServiceServer is the server API for ModelService service.
// All implementations must embed UnimplementedModelServiceServer
// for forward compatibility.
type ModelServiceServer interface {
	TextEmbedding(context.Context, *TextEmbeddingRequest) (*TextEmbeddingResponse, error)
	ImageEmbedding(grpc.ClientStreamingServer[ImageEmbeddingRequest, ImageEmbeddingResponse]) error
	mustEmbedUnimplementedModelServiceServer()
}

// UnimplementedModelServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedModelServiceServer struct{}

func (UnimplementedModelServiceServer) TextEmbedding(context.Context, *TextEmbeddingRequest) (*TextEmbeddingResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TextEmbedding not implemented")
}
func (UnimplementedModelServiceServer) ImageEmbedding(grpc.ClientStreamingServer[ImageEmbeddingRequest, ImageEmbeddingResponse]) error {
	return status.Errorf(codes.Unimplemented, "method ImageEmbedding not implemented")
}
func (UnimplementedModelServiceServer) mustEmbedUnimplementedModelServiceServer() {}
func (UnimplementedModelServiceServer) testEmbeddedByValue()                      {}

// UnsafeModelServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ModelServiceServer will
// result in compilation errors.
type UnsafeModelServiceServer interface {
	mustEmbedUnimplementedModelServiceServer()
}

func RegisterModelServiceServer(s grpc.ServiceRegistrar, srv ModelServiceServer) {
	// If the following call pancis, it indicates UnimplementedModelServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&ModelService_ServiceDesc, srv)
}

func _ModelService_TextEmbedding_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TextEmbeddingRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ModelServiceServer).TextEmbedding(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: ModelService_TextEmbedding_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ModelServiceServer).TextEmbedding(ctx, req.(*TextEmbeddingRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _ModelService_ImageEmbedding_Handler(srv interface{}, stream grpc.ServerStream) error {
	return srv.(ModelServiceServer).ImageEmbedding(&grpc.GenericServerStream[ImageEmbeddingRequest, ImageEmbeddingResponse]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type ModelService_ImageEmbeddingServer = grpc.ClientStreamingServer[ImageEmbeddingRequest, ImageEmbeddingResponse]

// ModelService_ServiceDesc is the grpc.ServiceDesc for ModelService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ModelService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "ModelService",
	HandlerType: (*ModelServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TextEmbedding",
			Handler:    _ModelService_TextEmbedding_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "ImageEmbedding",
			Handler:       _ModelService_ImageEmbedding_Handler,
			ClientStreams: true,
		},
	},
	Metadata: "api/ModelService.proto",
}
