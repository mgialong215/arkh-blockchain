// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: wasm/query.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion3 // please upgrade the proto package

func init() { proto.RegisterFile("wasm/query.proto", fileDescriptor_23bf88269acb91a9) }

var fileDescriptor_23bf88269acb91a9 = []byte{
	// 185 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0xce, 0x31, 0x0e, 0x82, 0x30,
	0x18, 0x05, 0x60, 0x18, 0xd4, 0x84, 0xc9, 0xe8, 0x46, 0x4c, 0x47, 0x07, 0x87, 0xfe, 0x41, 0x2f,
	0x60, 0xbc, 0x81, 0xab, 0xdb, 0x5f, 0x6c, 0x4a, 0xa3, 0xf4, 0xaf, 0xb4, 0xa8, 0xdc, 0xc2, 0x63,
	0x39, 0x32, 0x3a, 0x1a, 0xb8, 0x88, 0xa1, 0x24, 0x6e, 0x6f, 0xf8, 0xde, 0xcb, 0x4b, 0xe6, 0x0f,
	0x74, 0x25, 0xdc, 0x6a, 0x59, 0x35, 0xdc, 0x56, 0xe4, 0x69, 0xb1, 0xc4, 0xea, 0x52, 0xe0, 0x59,
	0xa3, 0xe1, 0x43, 0xe2, 0x03, 0x48, 0x57, 0x8a, 0x48, 0x5d, 0x25, 0xa0, 0xd5, 0x80, 0xc6, 0x90,
	0x47, 0xaf, 0xc9, 0xb8, 0xb1, 0x92, 0x6e, 0x72, 0x72, 0x25, 0x39, 0x10, 0xe8, 0xe4, 0xb8, 0x05,
	0xf7, 0x4c, 0x48, 0x8f, 0x19, 0x58, 0x54, 0xda, 0x04, 0x3c, 0xda, 0xed, 0x2c, 0x99, 0x1c, 0x07,
	0x71, 0xd8, 0xbf, 0x3b, 0x16, 0xb7, 0x1d, 0x8b, 0xbf, 0x1d, 0x8b, 0x5f, 0x3d, 0x8b, 0xda, 0x9e,
	0x45, 0x9f, 0x9e, 0x45, 0xa7, 0xb5, 0xd2, 0xbe, 0xa8, 0x05, 0xcf, 0xa9, 0x84, 0xff, 0x99, 0x90,
	0xe0, 0x09, 0xe1, 0xaf, 0x6f, 0xac, 0x74, 0x62, 0x1a, 0x16, 0x77, 0xbf, 0x00, 0x00, 0x00, 0xff,
	0xff, 0x46, 0x38, 0xb9, 0x34, 0xc4, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// QueryClient is the client API for Query service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type QueryClient interface {
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

// QueryServer is the server API for Query service.
type QueryServer interface {
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "arkhadian.arkh.wasm.Query",
	HandlerType: (*QueryServer)(nil),
	Methods:     []grpc.MethodDesc{},
	Streams:     []grpc.StreamDesc{},
	Metadata:    "wasm/query.proto",
}
