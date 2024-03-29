// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.3.0
// - protoc             v4.25.2
// source: service_verify_card.proto

package pb

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

const (
	CardValidator_VerifyCard_FullMethodName = "/pb.CardValidator/VerifyCard"
)

// CardValidatorClient is the client API for CardValidator service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type CardValidatorClient interface {
	VerifyCard(ctx context.Context, in *VerifyCardRequest, opts ...grpc.CallOption) (*VerifyCardResponse, error)
}

type cardValidatorClient struct {
	cc grpc.ClientConnInterface
}

func NewCardValidatorClient(cc grpc.ClientConnInterface) CardValidatorClient {
	return &cardValidatorClient{cc}
}

func (c *cardValidatorClient) VerifyCard(ctx context.Context, in *VerifyCardRequest, opts ...grpc.CallOption) (*VerifyCardResponse, error) {
	out := new(VerifyCardResponse)
	err := c.cc.Invoke(ctx, CardValidator_VerifyCard_FullMethodName, in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CardValidatorServer is the server API for CardValidator service.
// All implementations must embed UnimplementedCardValidatorServer
// for forward compatibility
type CardValidatorServer interface {
	VerifyCard(context.Context, *VerifyCardRequest) (*VerifyCardResponse, error)
	mustEmbedUnimplementedCardValidatorServer()
}

// UnimplementedCardValidatorServer must be embedded to have forward compatible implementations.
type UnimplementedCardValidatorServer struct {
}

func (UnimplementedCardValidatorServer) VerifyCard(context.Context, *VerifyCardRequest) (*VerifyCardResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method VerifyCard not implemented")
}
func (UnimplementedCardValidatorServer) mustEmbedUnimplementedCardValidatorServer() {}

// UnsafeCardValidatorServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to CardValidatorServer will
// result in compilation errors.
type UnsafeCardValidatorServer interface {
	mustEmbedUnimplementedCardValidatorServer()
}

func RegisterCardValidatorServer(s grpc.ServiceRegistrar, srv CardValidatorServer) {
	s.RegisterService(&CardValidator_ServiceDesc, srv)
}

func _CardValidator_VerifyCard_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VerifyCardRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CardValidatorServer).VerifyCard(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: CardValidator_VerifyCard_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CardValidatorServer).VerifyCard(ctx, req.(*VerifyCardRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// CardValidator_ServiceDesc is the grpc.ServiceDesc for CardValidator service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var CardValidator_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "pb.CardValidator",
	HandlerType: (*CardValidatorServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "VerifyCard",
			Handler:    _CardValidator_VerifyCard_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "service_verify_card.proto",
}
