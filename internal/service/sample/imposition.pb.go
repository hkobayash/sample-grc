// Code generated by protoc-gen-go. DO NOT EDIT.
// source: sample/imposition.proto

package sample

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ImpositionItem struct {
	Id                   int32    `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Bucket               string   `protobuf:"bytes,2,opt,name=bucket,proto3" json:"bucket,omitempty"`
	ObjectKey            string   `protobuf:"bytes,3,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ImpositionItem) Reset()         { *m = ImpositionItem{} }
func (m *ImpositionItem) String() string { return proto.CompactTextString(m) }
func (*ImpositionItem) ProtoMessage()    {}
func (*ImpositionItem) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d82f240b6d5a47e, []int{0}
}

func (m *ImpositionItem) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ImpositionItem.Unmarshal(m, b)
}
func (m *ImpositionItem) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ImpositionItem.Marshal(b, m, deterministic)
}
func (m *ImpositionItem) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ImpositionItem.Merge(m, src)
}
func (m *ImpositionItem) XXX_Size() int {
	return xxx_messageInfo_ImpositionItem.Size(m)
}
func (m *ImpositionItem) XXX_DiscardUnknown() {
	xxx_messageInfo_ImpositionItem.DiscardUnknown(m)
}

var xxx_messageInfo_ImpositionItem proto.InternalMessageInfo

func (m *ImpositionItem) GetId() int32 {
	if m != nil {
		return m.Id
	}
	return 0
}

func (m *ImpositionItem) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *ImpositionItem) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

type CreateItemRequest struct {
	Bucket               string   `protobuf:"bytes,1,opt,name=bucket,proto3" json:"bucket,omitempty"`
	ObjectKey            string   `protobuf:"bytes,2,opt,name=object_key,json=objectKey,proto3" json:"object_key,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *CreateItemRequest) Reset()         { *m = CreateItemRequest{} }
func (m *CreateItemRequest) String() string { return proto.CompactTextString(m) }
func (*CreateItemRequest) ProtoMessage()    {}
func (*CreateItemRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d82f240b6d5a47e, []int{1}
}

func (m *CreateItemRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateItemRequest.Unmarshal(m, b)
}
func (m *CreateItemRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateItemRequest.Marshal(b, m, deterministic)
}
func (m *CreateItemRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateItemRequest.Merge(m, src)
}
func (m *CreateItemRequest) XXX_Size() int {
	return xxx_messageInfo_CreateItemRequest.Size(m)
}
func (m *CreateItemRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateItemRequest.DiscardUnknown(m)
}

var xxx_messageInfo_CreateItemRequest proto.InternalMessageInfo

func (m *CreateItemRequest) GetBucket() string {
	if m != nil {
		return m.Bucket
	}
	return ""
}

func (m *CreateItemRequest) GetObjectKey() string {
	if m != nil {
		return m.ObjectKey
	}
	return ""
}

type CreateItemResponse struct {
	Item                 *ImpositionItem `protobuf:"bytes,1,opt,name=item,proto3" json:"item,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *CreateItemResponse) Reset()         { *m = CreateItemResponse{} }
func (m *CreateItemResponse) String() string { return proto.CompactTextString(m) }
func (*CreateItemResponse) ProtoMessage()    {}
func (*CreateItemResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0d82f240b6d5a47e, []int{2}
}

func (m *CreateItemResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CreateItemResponse.Unmarshal(m, b)
}
func (m *CreateItemResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CreateItemResponse.Marshal(b, m, deterministic)
}
func (m *CreateItemResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CreateItemResponse.Merge(m, src)
}
func (m *CreateItemResponse) XXX_Size() int {
	return xxx_messageInfo_CreateItemResponse.Size(m)
}
func (m *CreateItemResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_CreateItemResponse.DiscardUnknown(m)
}

var xxx_messageInfo_CreateItemResponse proto.InternalMessageInfo

func (m *CreateItemResponse) GetItem() *ImpositionItem {
	if m != nil {
		return m.Item
	}
	return nil
}

func init() {
	proto.RegisterType((*ImpositionItem)(nil), "sample.ImpositionItem")
	proto.RegisterType((*CreateItemRequest)(nil), "sample.CreateItemRequest")
	proto.RegisterType((*CreateItemResponse)(nil), "sample.CreateItemResponse")
}

func init() { proto.RegisterFile("sample/imposition.proto", fileDescriptor_0d82f240b6d5a47e) }

var fileDescriptor_0d82f240b6d5a47e = []byte{
	// 222 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2f, 0x4e, 0xcc, 0x2d,
	0xc8, 0x49, 0xd5, 0xcf, 0xcc, 0x2d, 0xc8, 0x2f, 0xce, 0x2c, 0xc9, 0xcc, 0xcf, 0xd3, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x62, 0x83, 0x48, 0x28, 0x85, 0x73, 0xf1, 0x79, 0xc2, 0xe5, 0x3c, 0x4b,
	0x52, 0x73, 0x85, 0xf8, 0xb8, 0x98, 0x32, 0x53, 0x24, 0x18, 0x15, 0x18, 0x35, 0x58, 0x83, 0x98,
	0x32, 0x53, 0x84, 0xc4, 0xb8, 0xd8, 0x92, 0x4a, 0x93, 0xb3, 0x53, 0x4b, 0x24, 0x98, 0x14, 0x18,
	0x35, 0x38, 0x83, 0xa0, 0x3c, 0x21, 0x59, 0x2e, 0xae, 0xfc, 0xa4, 0xac, 0xd4, 0xe4, 0x92, 0xf8,
	0xec, 0xd4, 0x4a, 0x09, 0x66, 0xb0, 0x1c, 0x27, 0x44, 0xc4, 0x3b, 0xb5, 0x52, 0xc9, 0x8b, 0x4b,
	0xd0, 0xb9, 0x28, 0x35, 0xb1, 0x24, 0x15, 0x64, 0x68, 0x50, 0x6a, 0x61, 0x69, 0x6a, 0x71, 0x09,
	0x92, 0x59, 0x8c, 0x78, 0xcc, 0x62, 0x42, 0x37, 0xcb, 0x81, 0x4b, 0x08, 0xd9, 0xac, 0xe2, 0x82,
	0xfc, 0xbc, 0xe2, 0x54, 0x21, 0x2d, 0x2e, 0x96, 0xcc, 0x92, 0xd4, 0x5c, 0xb0, 0x51, 0xdc, 0x46,
	0x62, 0x7a, 0x10, 0x1f, 0xe9, 0xa1, 0x7a, 0x27, 0x08, 0xac, 0xc6, 0x28, 0x82, 0x4b, 0x10, 0x21,
	0x1e, 0x9c, 0x5a, 0x54, 0x96, 0x99, 0x9c, 0x2a, 0xe4, 0xcc, 0xc5, 0x85, 0x30, 0x56, 0x48, 0x12,
	0x66, 0x00, 0x86, 0xb3, 0xa5, 0xa4, 0xb0, 0x49, 0x41, 0x5c, 0x91, 0xc4, 0x06, 0x0e, 0x4f, 0x63,
	0x40, 0x00, 0x00, 0x00, 0xff, 0xff, 0xc0, 0x19, 0xf3, 0x93, 0x6a, 0x01, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// ImpositionServiceClient is the client API for ImpositionService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ImpositionServiceClient interface {
	CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error)
}

type impositionServiceClient struct {
	cc *grpc.ClientConn
}

func NewImpositionServiceClient(cc *grpc.ClientConn) ImpositionServiceClient {
	return &impositionServiceClient{cc}
}

func (c *impositionServiceClient) CreateItem(ctx context.Context, in *CreateItemRequest, opts ...grpc.CallOption) (*CreateItemResponse, error) {
	out := new(CreateItemResponse)
	err := c.cc.Invoke(ctx, "/sample.ImpositionService/CreateItem", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ImpositionServiceServer is the server API for ImpositionService service.
type ImpositionServiceServer interface {
	CreateItem(context.Context, *CreateItemRequest) (*CreateItemResponse, error)
}

func RegisterImpositionServiceServer(s *grpc.Server, srv ImpositionServiceServer) {
	s.RegisterService(&_ImpositionService_serviceDesc, srv)
}

func _ImpositionService_CreateItem_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateItemRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ImpositionServiceServer).CreateItem(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/sample.ImpositionService/CreateItem",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ImpositionServiceServer).CreateItem(ctx, req.(*CreateItemRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _ImpositionService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "sample.ImpositionService",
	HandlerType: (*ImpositionServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateItem",
			Handler:    _ImpositionService_CreateItem_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "sample/imposition.proto",
}
