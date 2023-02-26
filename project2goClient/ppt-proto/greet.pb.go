// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ppt-proto/greet.proto

package templates

import (
	context "context"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

type TemplateRequest struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateRequest) Reset()         { *m = TemplateRequest{} }
func (m *TemplateRequest) String() string { return proto.CompactTextString(m) }
func (*TemplateRequest) ProtoMessage()    {}
func (*TemplateRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_f000de12cbbd68ad, []int{0}
}

func (m *TemplateRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateRequest.Unmarshal(m, b)
}
func (m *TemplateRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateRequest.Marshal(b, m, deterministic)
}
func (m *TemplateRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateRequest.Merge(m, src)
}
func (m *TemplateRequest) XXX_Size() int {
	return xxx_messageInfo_TemplateRequest.Size(m)
}
func (m *TemplateRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateRequest.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateRequest proto.InternalMessageInfo

func (m *TemplateRequest) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type TemplateResponse struct {
	Message              string   `protobuf:"bytes,1,opt,name=message,proto3" json:"message,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TemplateResponse) Reset()         { *m = TemplateResponse{} }
func (m *TemplateResponse) String() string { return proto.CompactTextString(m) }
func (*TemplateResponse) ProtoMessage()    {}
func (*TemplateResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_f000de12cbbd68ad, []int{1}
}

func (m *TemplateResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TemplateResponse.Unmarshal(m, b)
}
func (m *TemplateResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TemplateResponse.Marshal(b, m, deterministic)
}
func (m *TemplateResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TemplateResponse.Merge(m, src)
}
func (m *TemplateResponse) XXX_Size() int {
	return xxx_messageInfo_TemplateResponse.Size(m)
}
func (m *TemplateResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_TemplateResponse.DiscardUnknown(m)
}

var xxx_messageInfo_TemplateResponse proto.InternalMessageInfo

func (m *TemplateResponse) GetMessage() string {
	if m != nil {
		return m.Message
	}
	return ""
}

func init() {
	proto.RegisterType((*TemplateRequest)(nil), "templates.TemplateRequest")
	proto.RegisterType((*TemplateResponse)(nil), "templates.TemplateResponse")
}

func init() {
	proto.RegisterFile("ppt-proto/greet.proto", fileDescriptor_f000de12cbbd68ad)
}

var fileDescriptor_f000de12cbbd68ad = []byte{
	// 169 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x2d, 0x28, 0x28, 0xd1,
	0x2d, 0x28, 0xca, 0x2f, 0xc9, 0xd7, 0x4f, 0x2f, 0x4a, 0x4d, 0x2d, 0xd1, 0x03, 0xb3, 0x85, 0x38,
	0x4b, 0x52, 0x73, 0x0b, 0x72, 0x12, 0x4b, 0x52, 0x8b, 0x95, 0x54, 0xb9, 0xf8, 0x43, 0xa0, 0x9c,
	0xa0, 0xd4, 0xc2, 0xd2, 0xd4, 0xe2, 0x12, 0x21, 0x21, 0x2e, 0x96, 0xbc, 0xc4, 0xdc, 0x54, 0x09,
	0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x30, 0x5b, 0x49, 0x87, 0x4b, 0x00, 0xa1, 0xac, 0xb8, 0x20,
	0x3f, 0xaf, 0x38, 0x55, 0x48, 0x82, 0x8b, 0x3d, 0x37, 0xb5, 0xb8, 0x38, 0x31, 0x1d, 0xa6, 0x14,
	0xc6, 0x35, 0x8a, 0xe1, 0xe2, 0x85, 0xa9, 0xf6, 0x4d, 0xcc, 0x4e, 0x2d, 0x12, 0xf2, 0xe6, 0x12,
	0x08, 0x4a, 0x2d, 0x29, 0xca, 0x4c, 0x2d, 0x4b, 0x85, 0x49, 0x08, 0x49, 0xe9, 0xc1, 0x5d, 0xa1,
	0x87, 0xe6, 0x04, 0x29, 0x69, 0xac, 0x72, 0x10, 0x7b, 0x9d, 0x04, 0x57, 0x31, 0xf1, 0xc1, 0x04,
	0x83, 0x53, 0x8b, 0xca, 0x52, 0x8b, 0x92, 0xd8, 0xc0, 0xfe, 0x32, 0x06, 0x04, 0x00, 0x00, 0xff,
	0xff, 0x67, 0xde, 0x62, 0x82, 0xf0, 0x00, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// TemplateMakerClient is the client API for TemplateMaker service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type TemplateMakerClient interface {
	RetrieveTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error)
}

type templateMakerClient struct {
	cc grpc.ClientConnInterface
}

func NewTemplateMakerClient(cc grpc.ClientConnInterface) TemplateMakerClient {
	return &templateMakerClient{cc}
}

func (c *templateMakerClient) RetrieveTemplate(ctx context.Context, in *TemplateRequest, opts ...grpc.CallOption) (*TemplateResponse, error) {
	out := new(TemplateResponse)
	err := c.cc.Invoke(ctx, "/templates.TemplateMaker/RetrieveTemplate", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// TemplateMakerServer is the server API for TemplateMaker service.
type TemplateMakerServer interface {
	RetrieveTemplate(context.Context, *TemplateRequest) (*TemplateResponse, error)
}

// UnimplementedTemplateMakerServer can be embedded to have forward compatible implementations.
type UnimplementedTemplateMakerServer struct {
}

func (*UnimplementedTemplateMakerServer) RetrieveTemplate(ctx context.Context, req *TemplateRequest) (*TemplateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RetrieveTemplate not implemented")
}

func RegisterTemplateMakerServer(s *grpc.Server, srv TemplateMakerServer) {
	s.RegisterService(&_TemplateMaker_serviceDesc, srv)
}

func _TemplateMaker_RetrieveTemplate_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(TemplateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(TemplateMakerServer).RetrieveTemplate(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/templates.TemplateMaker/RetrieveTemplate",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(TemplateMakerServer).RetrieveTemplate(ctx, req.(*TemplateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _TemplateMaker_serviceDesc = grpc.ServiceDesc{
	ServiceName: "templates.TemplateMaker",
	HandlerType: (*TemplateMakerServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RetrieveTemplate",
			Handler:    _TemplateMaker_RetrieveTemplate_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ppt-proto/greet.proto",
}
