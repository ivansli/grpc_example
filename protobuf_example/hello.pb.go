// Code generated by protoc-gen-go. DO NOT EDIT.
// source: protobuf_example/hello.proto

package hello

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

// 定义 message 消息结构，对用go语言struct
type HelloRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Age                  int32    `protobuf:"varint,2,opt,name=age,proto3" json:"age,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloRequest) Reset()         { *m = HelloRequest{} }
func (m *HelloRequest) String() string { return proto.CompactTextString(m) }
func (*HelloRequest) ProtoMessage()    {}
func (*HelloRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_45a558161c551c34, []int{0}
}

func (m *HelloRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloRequest.Unmarshal(m, b)
}
func (m *HelloRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloRequest.Marshal(b, m, deterministic)
}
func (m *HelloRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloRequest.Merge(m, src)
}
func (m *HelloRequest) XXX_Size() int {
	return xxx_messageInfo_HelloRequest.Size(m)
}
func (m *HelloRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloRequest.DiscardUnknown(m)
}

var xxx_messageInfo_HelloRequest proto.InternalMessageInfo

func (m *HelloRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *HelloRequest) GetAge() int32 {
	if m != nil {
		return m.Age
	}
	return 0
}

// 定义 message 消息结构，对用go语言struct
type HelloResponse struct {
	Address              string   `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
	Money                int32    `protobuf:"varint,2,opt,name=money,proto3" json:"money,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *HelloResponse) Reset()         { *m = HelloResponse{} }
func (m *HelloResponse) String() string { return proto.CompactTextString(m) }
func (*HelloResponse) ProtoMessage()    {}
func (*HelloResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_45a558161c551c34, []int{1}
}

func (m *HelloResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_HelloResponse.Unmarshal(m, b)
}
func (m *HelloResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_HelloResponse.Marshal(b, m, deterministic)
}
func (m *HelloResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_HelloResponse.Merge(m, src)
}
func (m *HelloResponse) XXX_Size() int {
	return xxx_messageInfo_HelloResponse.Size(m)
}
func (m *HelloResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_HelloResponse.DiscardUnknown(m)
}

var xxx_messageInfo_HelloResponse proto.InternalMessageInfo

func (m *HelloResponse) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func (m *HelloResponse) GetMoney() int32 {
	if m != nil {
		return m.Money
	}
	return 0
}

var E_Author = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FileOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1000,
	Name:          "hello.author",
	Tag:           "bytes,1000,opt,name=author",
	Filename:      "protobuf_example/hello.proto",
}

var E_WhichMsg = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MessageOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1000,
	Name:          "hello.which_msg",
	Tag:           "bytes,1000,opt,name=which_msg",
	Filename:      "protobuf_example/hello.proto",
}

var E_DefaultString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         51000,
	Name:          "hello.default_string",
	Tag:           "bytes,51000,opt,name=default_string",
	Filename:      "protobuf_example/hello.proto",
}

var E_DefaultInt = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int32)(nil),
	Field:         51001,
	Name:          "hello.default_int",
	Tag:           "varint,51001,opt,name=default_int",
	Filename:      "protobuf_example/hello.proto",
}

var E_ServiceVersion = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.ServiceOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1000,
	Name:          "hello.service_version",
	Tag:           "bytes,1000,opt,name=service_version",
	Filename:      "protobuf_example/hello.proto",
}

var E_MethodVersion = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.MethodOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         1000,
	Name:          "hello.method_version",
	Tag:           "bytes,1000,opt,name=method_version",
	Filename:      "protobuf_example/hello.proto",
}


func init() {
	proto.RegisterType((*HelloRequest)(nil), "hello.HelloRequest")
	proto.RegisterType((*HelloResponse)(nil), "hello.HelloResponse")
	proto.RegisterExtension(E_Author)
	proto.RegisterExtension(E_WhichMsg)
	proto.RegisterExtension(E_DefaultString)
	proto.RegisterExtension(E_DefaultInt)
	proto.RegisterExtension(E_ServiceVersion)
	proto.RegisterExtension(E_MethodVersion)
}

func init() { proto.RegisterFile("protobuf_example/hello.proto", fileDescriptor_45a558161c551c34) }

var fileDescriptor_45a558161c551c34 = []byte{
	// 510 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x8b, 0xd3, 0x40,
	0x18, 0xc6, 0xd9, 0x3f, 0x6d, 0xb7, 0xa3, 0x5d, 0x71, 0xac, 0x6b, 0x28, 0x5d, 0x0d, 0x45, 0x61,
	0x41, 0x48, 0x76, 0x15, 0x2f, 0x91, 0x2d, 0xda, 0xc3, 0xe2, 0x1e, 0x16, 0x21, 0x05, 0x0f, 0x22,
	0x84, 0xd9, 0xe6, 0xdd, 0x64, 0x20, 0x9d, 0x89, 0x79, 0xa7, 0xd1, 0x22, 0x5e, 0xc4, 0xb3, 0x17,
	0xbf, 0x8c, 0xde, 0x24, 0xa7, 0xbd, 0xfb, 0x15, 0x3c, 0x78, 0xef, 0x17, 0x90, 0xcc, 0x24, 0x4b,
	0xad, 0x05, 0x6f, 0xd3, 0xb7, 0xcf, 0xf3, 0x7b, 0xf2, 0x3e, 0x93, 0x90, 0x7e, 0x9a, 0x49, 0x25,
	0xcf, 0x67, 0x17, 0x01, 0xbc, 0x67, 0xd3, 0x34, 0x01, 0x37, 0x86, 0x24, 0x91, 0x8e, 0x1e, 0xd3,
	0x86, 0xfe, 0xd1, 0xeb, 0x47, 0x52, 0x46, 0x09, 0xb8, 0x2c, 0xe5, 0x2e, 0x13, 0x42, 0x2a, 0xa6,
	0xb8, 0x14, 0x68, 0x44, 0x3d, 0xbb, 0xfa, 0xb7, 0x26, 0xb9, 0x21, 0xe0, 0x24, 0xe3, 0xa9, 0x92,
	0x99, 0x51, 0x0c, 0x02, 0x72, 0xfd, 0x45, 0x09, 0xf2, 0xe1, 0xed, 0x0c, 0x50, 0xd1, 0x7d, 0xb2,
	0x2d, 0xd8, 0x14, 0xac, 0x0d, 0x7b, 0xe3, 0xa0, 0x3d, 0x6a, 0x17, 0x0b, 0xab, 0xc1, 0x73, 0x26,
	0xd0, 0xd7, 0x63, 0xba, 0x47, 0xb6, 0x58, 0x04, 0xd6, 0xa6, 0xbd, 0x71, 0xd0, 0x18, 0x6d, 0x5f,
	0x2e, 0x2c, 0xea, 0x97, 0x03, 0xcf, 0x2a, 0x86, 0xb7, 0x55, 0xcc, 0xd1, 0xe6, 0x68, 0xeb, 0x07,
	0xb3, 0x33, 0x03, 0x1c, 0x20, 0xe9, 0x54, 0x01, 0x98, 0x4a, 0x81, 0x40, 0x1f, 0x92, 0x16, 0x0b,
	0xc3, 0x0c, 0x10, 0xab, 0x90, 0x9b, 0xc5, 0xc2, 0xea, 0x3c, 0x47, 0xce, 0xdc, 0x71, 0xcc, 0x44,
	0x14, 0x33, 0xee, 0xd7, 0x0a, 0xda, 0x27, 0x8d, 0xa9, 0x14, 0x30, 0xaf, 0x12, 0x9b, 0x97, 0x0b,
	0xeb, 0xc7, 0xe7, 0x63, 0xdf, 0x0c, 0xbd, 0x3b, 0xc5, 0xb0, 0xcb, 0xd5, 0x72, 0xa6, 0xc9, 0x78,
	0x84, 0x64, 0x67, 0xcc, 0xe6, 0x3a, 0x97, 0xbe, 0x59, 0x3a, 0xdf, 0x72, 0x4c, 0x85, 0xcb, 0x2b,
	0xf7, 0xba, 0x7f, 0x0f, 0x0d, 0x62, 0xf0, 0xe0, 0xd3, 0xcf, 0x5f, 0x5f, 0x37, 0xf7, 0x68, 0xd7,
	0xcd, 0x8f, 0x5c, 0x64, 0xf3, 0x40, 0x8b, 0xdc, 0x0f, 0x65, 0x0d, 0x1f, 0x8b, 0x61, 0x33, 0x3f,
	0x74, 0x0e, 0x9d, 0xa3, 0x5e, 0xfb, 0xea, 0xe8, 0x3d, 0x21, 0x4d, 0x36, 0x53, 0xb1, 0xcc, 0x68,
	0xdf, 0x31, 0xbd, 0x3b, 0x75, 0xef, 0xce, 0x09, 0x4f, 0xe0, 0x65, 0xaa, 0xaf, 0xc6, 0xfa, 0xdd,
	0x2a, 0xd7, 0xf6, 0x2b, 0xb1, 0x77, 0x4c, 0xda, 0xef, 0x62, 0x3e, 0x89, 0x83, 0x29, 0x46, 0xf4,
	0xde, 0x3f, 0xce, 0x33, 0x40, 0x64, 0xd1, 0xaa, 0x79, 0x47, 0x5b, 0xce, 0x30, 0xf2, 0x4e, 0xc8,
	0x6e, 0x08, 0x17, 0x6c, 0x96, 0xa8, 0x00, 0x55, 0xc6, 0x45, 0x44, 0xf7, 0xd7, 0xa4, 0x43, 0x12,
	0xd6, 0x84, 0x6f, 0x5f, 0xb6, 0x34, 0xa2, 0x53, 0xd9, 0xc6, 0xda, 0xe5, 0x3d, 0x23, 0xd7, 0x6a,
	0x0e, 0x17, 0xea, 0x7f, 0x90, 0xef, 0x1a, 0xd2, 0xf0, 0x49, 0xe5, 0x39, 0x15, 0xca, 0x3b, 0x25,
	0x37, 0x10, 0xb2, 0x9c, 0x4f, 0x20, 0xc8, 0x21, 0x43, 0x2e, 0xc5, 0x9a, 0x75, 0xc6, 0x46, 0xb1,
	0xb2, 0xce, 0x6e, 0x65, 0x7c, 0x65, 0x7c, 0xe5, 0x52, 0x53, 0x50, 0xb1, 0x0c, 0xaf, 0x48, 0x77,
	0xd7, 0x14, 0x53, 0x0a, 0x56, 0x40, 0x1d, 0x63, 0xab, 0x38, 0xa3, 0xfb, 0xaf, 0x2d, 0xc7, 0x5d,
	0xfd, 0x8c, 0x9e, 0xea, 0x9b, 0x2c, 0x86, 0x2d, 0xfd, 0x62, 0x27, 0xfc, 0xbc, 0xa9, 0x15, 0x8f,
	0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xd8, 0x39, 0xe5, 0x42, 0x71, 0x03, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// SayHelloClient is the client API for SayHello service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type SayHelloClient interface {
	// rpc对应服务方法
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error)
}

type sayHelloClient struct {
	cc grpc.ClientConnInterface
}

func NewSayHelloClient(cc grpc.ClientConnInterface) SayHelloClient {
	return &sayHelloClient{cc}
}

func (c *sayHelloClient) SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloResponse, error) {
	out := new(HelloResponse)
	err := c.cc.Invoke(ctx, "/hello.SayHello/SayHello", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// SayHelloServer is the server API for SayHello service.
type SayHelloServer interface {
	// rpc对应服务方法
	SayHello(context.Context, *HelloRequest) (*HelloResponse, error)
}

// UnimplementedSayHelloServer can be embedded to have forward compatible implementations.
type UnimplementedSayHelloServer struct {
}

func (*UnimplementedSayHelloServer) SayHello(ctx context.Context, req *HelloRequest) (*HelloResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SayHello not implemented")
}

func RegisterSayHelloServer(s *grpc.Server, srv SayHelloServer) {
	s.RegisterService(&_SayHello_serviceDesc, srv)
}

func _SayHello_SayHello_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HelloRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(SayHelloServer).SayHello(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/hello.SayHello/SayHello",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(SayHelloServer).SayHello(ctx, req.(*HelloRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _SayHello_serviceDesc = grpc.ServiceDesc{
	ServiceName: "hello.SayHello",
	HandlerType: (*SayHelloServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SayHello",
			Handler:    _SayHello_SayHello_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "protobuf_example/hello.proto",
}
