// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             v3.6.1
// source: api/parser/pb/parser.proto

package pb

import (
	context "context"

	empty "github.com/golang/protobuf/ptypes/empty"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// ParserServiceClient is the client API for ParserService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type ParserServiceClient interface {
	ParsePosts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParsePostsResponce, error)
	ParseStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParseStatusResponce, error)
}

type parserServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewParserServiceClient(cc grpc.ClientConnInterface) ParserServiceClient {
	return &parserServiceClient{cc}
}

func (c *parserServiceClient) ParsePosts(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParsePostsResponce, error) {
	out := new(ParsePostsResponce)
	err := c.cc.Invoke(ctx, "/parser.ParserService/ParsePosts", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *parserServiceClient) ParseStatus(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*ParseStatusResponce, error) {
	out := new(ParseStatusResponce)
	err := c.cc.Invoke(ctx, "/parser.ParserService/ParseStatus", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ParserServiceServer is the server API for ParserService service.
// All implementations must embed UnimplementedParserServiceServer
// for forward compatibility
type ParserServiceServer interface {
	ParsePosts(context.Context, *empty.Empty) (*ParsePostsResponce, error)
	ParseStatus(context.Context, *empty.Empty) (*ParseStatusResponce, error)
	mustEmbedUnimplementedParserServiceServer()
}

// UnimplementedParserServiceServer must be embedded to have forward compatible implementations.
type UnimplementedParserServiceServer struct{}

func (UnimplementedParserServiceServer) ParsePosts(context.Context, *empty.Empty) (*ParsePostsResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParsePosts not implemented")
}

func (UnimplementedParserServiceServer) ParseStatus(context.Context, *empty.Empty) (*ParseStatusResponce, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ParseStatus not implemented")
}
func (UnimplementedParserServiceServer) mustEmbedUnimplementedParserServiceServer() {}

// UnsafeParserServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to ParserServiceServer will
// result in compilation errors.
type UnsafeParserServiceServer interface {
	mustEmbedUnimplementedParserServiceServer()
}

func RegisterParserServiceServer(s grpc.ServiceRegistrar, srv ParserServiceServer) {
	s.RegisterService(&ParserService_ServiceDesc, srv)
}

func _ParserService_ParsePosts_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServiceServer).ParsePosts(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.ParserService/ParsePosts",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServiceServer).ParsePosts(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _ParserService_ParseStatus_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(empty.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ParserServiceServer).ParseStatus(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/parser.ParserService/ParseStatus",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ParserServiceServer).ParseStatus(ctx, req.(*empty.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// ParserService_ServiceDesc is the grpc.ServiceDesc for ParserService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var ParserService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "parser.ParserService",
	HandlerType: (*ParserServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "ParsePosts",
			Handler:    _ParserService_ParsePosts_Handler,
		},
		{
			MethodName: "ParseStatus",
			Handler:    _ParserService_ParseStatus_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/parser/pb/parser.proto",
}
