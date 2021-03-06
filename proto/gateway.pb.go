// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/gateway.proto

package proto

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	descriptor "github.com/golang/protobuf/protoc-gen-go/descriptor"
	any "github.com/golang/protobuf/ptypes/any"
	empty "github.com/golang/protobuf/ptypes/empty"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Error struct {
	Code                 int64    `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Reason               string   `protobuf:"bytes,2,opt,name=reason,proto3" json:"reason,omitempty"`
	Detail               *any.Any `protobuf:"bytes,3,opt,name=detail,proto3" json:"detail,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_85acbde2a6adc437, []int{0}
}

func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (m *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(m, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetCode() int64 {
	if m != nil {
		return m.Code
	}
	return 0
}

func (m *Error) GetReason() string {
	if m != nil {
		return m.Reason
	}
	return ""
}

func (m *Error) GetDetail() *any.Any {
	if m != nil {
		return m.Detail
	}
	return nil
}

type ProdRequest struct {
	ProdId               int32    `protobuf:"varint,1,opt,name=prod_id,json=prodId,proto3" json:"prod_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdRequest) Reset()         { *m = ProdRequest{} }
func (m *ProdRequest) String() string { return proto.CompactTextString(m) }
func (*ProdRequest) ProtoMessage()    {}
func (*ProdRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85acbde2a6adc437, []int{1}
}

func (m *ProdRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdRequest.Unmarshal(m, b)
}
func (m *ProdRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdRequest.Marshal(b, m, deterministic)
}
func (m *ProdRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdRequest.Merge(m, src)
}
func (m *ProdRequest) XXX_Size() int {
	return xxx_messageInfo_ProdRequest.Size(m)
}
func (m *ProdRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdRequest.DiscardUnknown(m)
}

var xxx_messageInfo_ProdRequest proto.InternalMessageInfo

func (m *ProdRequest) GetProdId() int32 {
	if m != nil {
		return m.ProdId
	}
	return 0
}

type ProdResponse struct {
	ProdStock            int32    `protobuf:"varint,1,opt,name=prod_stock,json=prodStock,proto3" json:"prod_stock,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ProdResponse) Reset()         { *m = ProdResponse{} }
func (m *ProdResponse) String() string { return proto.CompactTextString(m) }
func (*ProdResponse) ProtoMessage()    {}
func (*ProdResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85acbde2a6adc437, []int{2}
}

func (m *ProdResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ProdResponse.Unmarshal(m, b)
}
func (m *ProdResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ProdResponse.Marshal(b, m, deterministic)
}
func (m *ProdResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ProdResponse.Merge(m, src)
}
func (m *ProdResponse) XXX_Size() int {
	return xxx_messageInfo_ProdResponse.Size(m)
}
func (m *ProdResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_ProdResponse.DiscardUnknown(m)
}

var xxx_messageInfo_ProdResponse proto.InternalMessageInfo

func (m *ProdResponse) GetProdStock() int32 {
	if m != nil {
		return m.ProdStock
	}
	return 0
}

type OrderRequest struct {
	OrderId              int64    `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *OrderRequest) Reset()         { *m = OrderRequest{} }
func (m *OrderRequest) String() string { return proto.CompactTextString(m) }
func (*OrderRequest) ProtoMessage()    {}
func (*OrderRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_85acbde2a6adc437, []int{3}
}

func (m *OrderRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderRequest.Unmarshal(m, b)
}
func (m *OrderRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderRequest.Marshal(b, m, deterministic)
}
func (m *OrderRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderRequest.Merge(m, src)
}
func (m *OrderRequest) XXX_Size() int {
	return xxx_messageInfo_OrderRequest.Size(m)
}
func (m *OrderRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderRequest.DiscardUnknown(m)
}

var xxx_messageInfo_OrderRequest proto.InternalMessageInfo

func (m *OrderRequest) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

type OrderResponse struct {
	OrderId              int64                `protobuf:"varint,1,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	Name                 string               `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	M                    map[string]string    `protobuf:"bytes,4,rep,name=m,proto3" json:"m,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Time                 *timestamp.Timestamp `protobuf:"bytes,3,opt,name=time,proto3" json:"time,omitempty"`
	Any                  *any.Any             `protobuf:"bytes,5,opt,name=any,proto3" json:"any,omitempty"`
	Empty                *empty.Empty         `protobuf:"bytes,6,opt,name=empty,proto3" json:"empty,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *OrderResponse) Reset()         { *m = OrderResponse{} }
func (m *OrderResponse) String() string { return proto.CompactTextString(m) }
func (*OrderResponse) ProtoMessage()    {}
func (*OrderResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_85acbde2a6adc437, []int{4}
}

func (m *OrderResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OrderResponse.Unmarshal(m, b)
}
func (m *OrderResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OrderResponse.Marshal(b, m, deterministic)
}
func (m *OrderResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OrderResponse.Merge(m, src)
}
func (m *OrderResponse) XXX_Size() int {
	return xxx_messageInfo_OrderResponse.Size(m)
}
func (m *OrderResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_OrderResponse.DiscardUnknown(m)
}

var xxx_messageInfo_OrderResponse proto.InternalMessageInfo

func (m *OrderResponse) GetOrderId() int64 {
	if m != nil {
		return m.OrderId
	}
	return 0
}

func (m *OrderResponse) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *OrderResponse) GetM() map[string]string {
	if m != nil {
		return m.M
	}
	return nil
}

func (m *OrderResponse) GetTime() *timestamp.Timestamp {
	if m != nil {
		return m.Time
	}
	return nil
}

func (m *OrderResponse) GetAny() *any.Any {
	if m != nil {
		return m.Any
	}
	return nil
}

func (m *OrderResponse) GetEmpty() *empty.Empty {
	if m != nil {
		return m.Empty
	}
	return nil
}

var E_DefaultString = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*string)(nil),
	Field:         50000,
	Name:          "proto.default_string",
	Tag:           "bytes,50000,opt,name=default_string",
	Filename:      "proto/gateway.proto",
}

var E_DefaultInt = &proto.ExtensionDesc{
	ExtendedType:  (*descriptor.FieldOptions)(nil),
	ExtensionType: (*int64)(nil),
	Field:         50001,
	Name:          "proto.default_int",
	Tag:           "varint,50001,opt,name=default_int",
	Filename:      "proto/gateway.proto",
}

func init() {
	proto.RegisterType((*Error)(nil), "proto.Error")
	proto.RegisterType((*ProdRequest)(nil), "proto.ProdRequest")
	proto.RegisterType((*ProdResponse)(nil), "proto.ProdResponse")
	proto.RegisterType((*OrderRequest)(nil), "proto.OrderRequest")
	proto.RegisterType((*OrderResponse)(nil), "proto.OrderResponse")
	proto.RegisterMapType((map[string]string)(nil), "proto.OrderResponse.MEntry")
	proto.RegisterExtension(E_DefaultString)
	proto.RegisterExtension(E_DefaultInt)
}

func init() { proto.RegisterFile("proto/gateway.proto", fileDescriptor_85acbde2a6adc437) }

var fileDescriptor_85acbde2a6adc437 = []byte{
	// 571 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x93, 0x4f, 0x6f, 0xd3, 0x30,
	0x18, 0xc6, 0x95, 0xa6, 0xc9, 0xb6, 0xb7, 0x1b, 0x02, 0xb7, 0x82, 0x2c, 0x63, 0xa2, 0xca, 0x61,
	0xea, 0xa4, 0x91, 0x88, 0xc2, 0x01, 0xed, 0x04, 0x93, 0xba, 0xa9, 0x07, 0xb4, 0xca, 0xe5, 0xc2,
	0x69, 0x72, 0x6b, 0xaf, 0x44, 0x6b, 0xec, 0xe0, 0xb8, 0x45, 0xd5, 0xb4, 0xcb, 0x0e, 0x88, 0x3b,
	0x9f, 0x81, 0x8f, 0xb0, 0xef, 0x01, 0xe2, 0x2b, 0xf0, 0x41, 0x50, 0x6c, 0x47, 0xa2, 0x2d, 0x7f,
	0x4e, 0xf5, 0xeb, 0xf7, 0xe7, 0xe7, 0xed, 0xfb, 0x3c, 0x2d, 0x34, 0x73, 0x29, 0x94, 0x48, 0x26,
	0x44, 0xb1, 0x8f, 0x64, 0x11, 0xeb, 0x0a, 0x79, 0xfa, 0x23, 0x7c, 0x3c, 0x11, 0x62, 0x32, 0x65,
	0x09, 0xc9, 0xd3, 0x84, 0x70, 0x2e, 0x14, 0x51, 0xa9, 0xe0, 0x85, 0x81, 0xc2, 0x27, 0xb6, 0xab,
	0xab, 0xd1, 0xec, 0x32, 0x51, 0x69, 0xc6, 0x0a, 0x45, 0xb2, 0xdc, 0x02, 0xbb, 0xab, 0x00, 0xe1,
	0x76, 0x40, 0xb8, 0xb7, 0xda, 0x62, 0x59, 0xae, 0xaa, 0x66, 0x7b, 0xb5, 0x49, 0x59, 0x31, 0x96,
	0x69, 0xae, 0x84, 0x34, 0x44, 0x54, 0x80, 0xd7, 0x93, 0x52, 0x48, 0x14, 0x40, 0x7d, 0x2c, 0x28,
	0x0b, 0x9c, 0xb6, 0xd3, 0x71, 0x4f, 0xea, 0x9f, 0xef, 0x02, 0xc0, 0xfa, 0x06, 0x45, 0xe0, 0x4b,
	0x46, 0x0a, 0xc1, 0x83, 0x5a, 0xdb, 0xe9, 0x6c, 0x9d, 0xc0, 0xed, 0x5d, 0xe0, 0x4f, 0x44, 0xfe,
	0x9e, 0x49, 0x6c, 0x3b, 0xe8, 0x08, 0x7c, 0xca, 0x14, 0x49, 0xa7, 0x81, 0xdb, 0x76, 0x3a, 0x8d,
	0x6e, 0x2b, 0x36, 0x93, 0xe3, 0x6a, 0x72, 0xfc, 0x9a, 0x2f, 0xb0, 0x65, 0xa2, 0x03, 0x68, 0x0c,
	0xa4, 0xa0, 0x98, 0x7d, 0x98, 0xb1, 0x42, 0xa1, 0x47, 0xb0, 0x91, 0x4b, 0x41, 0x2f, 0x52, 0xaa,
	0xa7, 0x7b, 0xd8, 0x2f, 0xcb, 0x3e, 0x8d, 0x9e, 0xc2, 0xb6, 0xe1, 0x8a, 0x5c, 0xf0, 0x82, 0xa1,
	0x7d, 0x00, 0x0d, 0x16, 0x4a, 0x8c, 0xaf, 0x2c, 0xbb, 0x55, 0xde, 0x0c, 0xcb, 0x8b, 0xe8, 0x10,
	0xb6, 0xcf, 0x25, 0x65, 0xb2, 0xd2, 0xdd, 0x85, 0x4d, 0x51, 0xd6, 0x95, 0xb0, 0x8b, 0x37, 0x74,
	0xdd, 0xa7, 0xd1, 0xd7, 0x1a, 0xec, 0x58, 0xd6, 0x6a, 0xff, 0x1d, 0x46, 0x08, 0xea, 0x9c, 0x64,
	0xcc, 0xac, 0x8f, 0xf5, 0x19, 0x1d, 0x82, 0x93, 0x05, 0xf5, 0xb6, 0xdb, 0x69, 0x74, 0xf7, 0xcc,
	0x92, 0xf1, 0x92, 0x5e, 0xfc, 0xa6, 0xc7, 0x95, 0x5c, 0x60, 0x27, 0x43, 0x31, 0xd4, 0xcb, 0x3c,
	0xad, 0x33, 0xe1, 0x9a, 0x33, 0x6f, 0xab, 0xb0, 0xb1, 0xe6, 0xd0, 0x01, 0xb8, 0x84, 0x2f, 0x02,
	0xef, 0x1f, 0x46, 0x96, 0x00, 0x3a, 0x02, 0x4f, 0x67, 0x1d, 0xf8, 0x9a, 0x7c, 0xb8, 0x46, 0xf6,
	0xca, 0x2e, 0x36, 0x50, 0xf8, 0x02, 0x7c, 0xf3, 0x95, 0xd0, 0x7d, 0x70, 0xaf, 0xd8, 0x42, 0x2f,
	0xb9, 0x85, 0xcb, 0x23, 0x6a, 0x81, 0x37, 0x27, 0xd3, 0x59, 0xb5, 0xa1, 0x29, 0x8e, 0x6b, 0x2f,
	0x9d, 0xee, 0xc8, 0x24, 0x35, 0x64, 0x72, 0x9e, 0x8e, 0x19, 0x1a, 0xc2, 0xf6, 0x19, 0x53, 0x83,
	0xca, 0x71, 0x84, 0xec, 0xea, 0xbf, 0xa5, 0x19, 0x36, 0x97, 0xee, 0x8c, 0x1b, 0x51, 0x78, 0xfb,
	0xe3, 0xe7, 0x97, 0x5a, 0x0b, 0xa1, 0x64, 0xfe, 0xac, 0xfc, 0x35, 0xd2, 0xe4, 0xda, 0x46, 0x7e,
	0xd3, 0x7d, 0x07, 0x9e, 0xb6, 0x0e, 0x0d, 0x60, 0xf3, 0x8c, 0x29, 0x73, 0x6e, 0x2e, 0x9b, 0x6a,
	0xa4, 0x5b, 0x7f, 0x72, 0x3a, 0xda, 0xd5, 0xda, 0x4d, 0xf4, 0x20, 0xd1, 0x81, 0x25, 0xd7, 0x55,
	0x8e, 0x37, 0xc7, 0xa7, 0x70, 0x8f, 0xb2, 0x4b, 0x32, 0x9b, 0xaa, 0x8b, 0x42, 0xc9, 0x94, 0x4f,
	0xd0, 0xfe, 0x9a, 0x4b, 0xa7, 0x29, 0x9b, 0xd2, 0xf3, 0x5c, 0xff, 0x1f, 0x83, 0x6f, 0x9f, 0x5c,
	0x6d, 0xc1, 0x8e, 0x7d, 0x36, 0xd4, 0xaf, 0x8e, 0x5f, 0x41, 0xa3, 0xd2, 0x49, 0xb9, 0xfa, 0x9f,
	0xc8, 0x77, 0x2d, 0xe2, 0x62, 0xb0, 0x6f, 0xfa, 0x5c, 0x8d, 0x7c, 0x8d, 0x3e, 0xff, 0x15, 0x00,
	0x00, 0xff, 0xff, 0xd8, 0x95, 0xda, 0x67, 0x25, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ProdServiceClient is the client API for ProdService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ProdServiceClient interface {
	GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error)
}

type prodServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewProdServiceClient(cc grpc.ClientConnInterface) ProdServiceClient {
	return &prodServiceClient{cc}
}

func (c *prodServiceClient) GetProdStock(ctx context.Context, in *ProdRequest, opts ...grpc.CallOption) (*ProdResponse, error) {
	out := new(ProdResponse)
	err := c.cc.Invoke(ctx, "/proto.ProdService/GetProdStock", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ProdServiceServer is the server API for ProdService service.
type ProdServiceServer interface {
	GetProdStock(context.Context, *ProdRequest) (*ProdResponse, error)
}

// UnimplementedProdServiceServer can be embedded to have forward compatible implementations.
type UnimplementedProdServiceServer struct {
}

func (*UnimplementedProdServiceServer) GetProdStock(ctx context.Context, req *ProdRequest) (*ProdResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetProdStock not implemented")
}

func RegisterProdServiceServer(s *grpc.Server, srv ProdServiceServer) {
	s.RegisterService(&_ProdService_serviceDesc, srv)
}

func _ProdService_GetProdStock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ProdRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ProdServiceServer).GetProdStock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.ProdService/GetProdStock",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ProdServiceServer).GetProdStock(ctx, req.(*ProdRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ProdService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.ProdService",
	HandlerType: (*ProdServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetProdStock",
			Handler:    _ProdService_GetProdStock_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gateway.proto",
}

// OrderClient is the client API for Order service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type OrderClient interface {
	GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error)
}

type orderClient struct {
	cc grpc.ClientConnInterface
}

func NewOrderClient(cc grpc.ClientConnInterface) OrderClient {
	return &orderClient{cc}
}

func (c *orderClient) GetOrder(ctx context.Context, in *OrderRequest, opts ...grpc.CallOption) (*OrderResponse, error) {
	out := new(OrderResponse)
	err := c.cc.Invoke(ctx, "/proto.Order/GetOrder", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// OrderServer is the server API for Order service.
type OrderServer interface {
	GetOrder(context.Context, *OrderRequest) (*OrderResponse, error)
}

// UnimplementedOrderServer can be embedded to have forward compatible implementations.
type UnimplementedOrderServer struct {
}

func (*UnimplementedOrderServer) GetOrder(ctx context.Context, req *OrderRequest) (*OrderResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetOrder not implemented")
}

func RegisterOrderServer(s *grpc.Server, srv OrderServer) {
	s.RegisterService(&_Order_serviceDesc, srv)
}

func _Order_GetOrder_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(OrderRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(OrderServer).GetOrder(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/proto.Order/GetOrder",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(OrderServer).GetOrder(ctx, req.(*OrderRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Order_serviceDesc = grpc.ServiceDesc{
	ServiceName: "proto.Order",
	HandlerType: (*OrderServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetOrder",
			Handler:    _Order_GetOrder_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/gateway.proto",
}
