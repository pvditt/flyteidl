// Code generated by protoc-gen-go. DO NOT EDIT.
// source: flyteidl/service/cache.proto

package service

import (
	context "context"
	fmt "fmt"
	core "github.com/flyteorg/flyteidl/gen/pb-go/flyteidl/core"
	proto "github.com/golang/protobuf/proto"
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

type EvictCacheRequest struct {
	// Identifier of resource cached data should be evicted for.
	//
	// Types that are valid to be assigned to Id:
	//	*EvictCacheRequest_WorkflowExecutionId
	//	*EvictCacheRequest_TaskExecutionId
	Id                   isEvictCacheRequest_Id `protobuf_oneof:"id"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *EvictCacheRequest) Reset()         { *m = EvictCacheRequest{} }
func (m *EvictCacheRequest) String() string { return proto.CompactTextString(m) }
func (*EvictCacheRequest) ProtoMessage()    {}
func (*EvictCacheRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff5da69b96fa44, []int{0}
}

func (m *EvictCacheRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvictCacheRequest.Unmarshal(m, b)
}
func (m *EvictCacheRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvictCacheRequest.Marshal(b, m, deterministic)
}
func (m *EvictCacheRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvictCacheRequest.Merge(m, src)
}
func (m *EvictCacheRequest) XXX_Size() int {
	return xxx_messageInfo_EvictCacheRequest.Size(m)
}
func (m *EvictCacheRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_EvictCacheRequest.DiscardUnknown(m)
}

var xxx_messageInfo_EvictCacheRequest proto.InternalMessageInfo

type isEvictCacheRequest_Id interface {
	isEvictCacheRequest_Id()
}

type EvictCacheRequest_WorkflowExecutionId struct {
	WorkflowExecutionId *core.WorkflowExecutionIdentifier `protobuf:"bytes,1,opt,name=workflow_execution_id,json=workflowExecutionId,proto3,oneof"`
}

type EvictCacheRequest_TaskExecutionId struct {
	TaskExecutionId *core.TaskExecutionIdentifier `protobuf:"bytes,2,opt,name=task_execution_id,json=taskExecutionId,proto3,oneof"`
}

func (*EvictCacheRequest_WorkflowExecutionId) isEvictCacheRequest_Id() {}

func (*EvictCacheRequest_TaskExecutionId) isEvictCacheRequest_Id() {}

func (m *EvictCacheRequest) GetId() isEvictCacheRequest_Id {
	if m != nil {
		return m.Id
	}
	return nil
}

func (m *EvictCacheRequest) GetWorkflowExecutionId() *core.WorkflowExecutionIdentifier {
	if x, ok := m.GetId().(*EvictCacheRequest_WorkflowExecutionId); ok {
		return x.WorkflowExecutionId
	}
	return nil
}

func (m *EvictCacheRequest) GetTaskExecutionId() *core.TaskExecutionIdentifier {
	if x, ok := m.GetId().(*EvictCacheRequest_TaskExecutionId); ok {
		return x.TaskExecutionId
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*EvictCacheRequest) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*EvictCacheRequest_WorkflowExecutionId)(nil),
		(*EvictCacheRequest_TaskExecutionId)(nil),
	}
}

type EvictCacheResponse struct {
	// List of errors encountered during cache eviction (if any).
	Errors               *core.CacheEvictionErrorList `protobuf:"bytes,1,opt,name=errors,proto3" json:"errors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                     `json:"-"`
	XXX_unrecognized     []byte                       `json:"-"`
	XXX_sizecache        int32                        `json:"-"`
}

func (m *EvictCacheResponse) Reset()         { *m = EvictCacheResponse{} }
func (m *EvictCacheResponse) String() string { return proto.CompactTextString(m) }
func (*EvictCacheResponse) ProtoMessage()    {}
func (*EvictCacheResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c5ff5da69b96fa44, []int{1}
}

func (m *EvictCacheResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EvictCacheResponse.Unmarshal(m, b)
}
func (m *EvictCacheResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EvictCacheResponse.Marshal(b, m, deterministic)
}
func (m *EvictCacheResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EvictCacheResponse.Merge(m, src)
}
func (m *EvictCacheResponse) XXX_Size() int {
	return xxx_messageInfo_EvictCacheResponse.Size(m)
}
func (m *EvictCacheResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_EvictCacheResponse.DiscardUnknown(m)
}

var xxx_messageInfo_EvictCacheResponse proto.InternalMessageInfo

func (m *EvictCacheResponse) GetErrors() *core.CacheEvictionErrorList {
	if m != nil {
		return m.Errors
	}
	return nil
}

func init() {
	proto.RegisterType((*EvictCacheRequest)(nil), "flyteidl.service.EvictCacheRequest")
	proto.RegisterType((*EvictCacheResponse)(nil), "flyteidl.service.EvictCacheResponse")
}

func init() { proto.RegisterFile("flyteidl/service/cache.proto", fileDescriptor_c5ff5da69b96fa44) }

var fileDescriptor_c5ff5da69b96fa44 = []byte{
	// 463 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x54, 0xcd, 0x8a, 0x13, 0x41,
	0x10, 0x36, 0xd9, 0xec, 0x1e, 0x5a, 0x41, 0x77, 0x44, 0x90, 0xb0, 0x88, 0xc4, 0x1f, 0x24, 0xe0,
	0x34, 0xae, 0x07, 0x51, 0xf0, 0xb2, 0x12, 0x50, 0xf0, 0x94, 0x5d, 0x10, 0xf6, 0x12, 0x3b, 0x3d,
	0x95, 0xd9, 0x32, 0x49, 0xd7, 0xd8, 0x5d, 0x49, 0x5c, 0x96, 0x5c, 0x3c, 0xf8, 0x02, 0x1e, 0x7c,
	0x02, 0x05, 0xc1, 0x97, 0x11, 0x5f, 0xc1, 0x07, 0x91, 0xe9, 0xc9, 0x8c, 0x4e, 0x92, 0x59, 0x37,
	0xd7, 0xfa, 0xbe, 0xfa, 0xea, 0xab, 0xfa, 0x66, 0x5a, 0xec, 0x0d, 0x46, 0xa7, 0x0c, 0x18, 0x8d,
	0xa4, 0x03, 0x3b, 0x45, 0x0d, 0x52, 0x2b, 0x7d, 0x02, 0x61, 0x62, 0x89, 0x29, 0xb8, 0x96, 0xa3,
	0xe1, 0x02, 0x6d, 0xee, 0xc5, 0x44, 0xf1, 0x08, 0xa4, 0x4a, 0x50, 0x2a, 0x63, 0x88, 0x15, 0x23,
	0x19, 0x97, 0xf1, 0x9b, 0xcd, 0x42, 0x4d, 0x93, 0x05, 0x09, 0xd6, 0x92, 0xcd, 0xb1, 0x5b, 0x65,
	0x0c, 0x23, 0x30, 0x8c, 0x03, 0x04, 0x9b, 0xe1, 0xad, 0x9f, 0x35, 0xb1, 0xdb, 0x99, 0xa2, 0xe6,
	0x17, 0xa9, 0x81, 0x2e, 0xbc, 0x9f, 0x80, 0xe3, 0xe0, 0xad, 0xb8, 0x31, 0x23, 0x3b, 0x1c, 0x8c,
	0x68, 0xd6, 0x83, 0x0f, 0xa0, 0x27, 0xe9, 0xb8, 0x1e, 0x46, 0x37, 0x6b, 0xb7, 0x6b, 0x0f, 0x2e,
	0xef, 0xb7, 0xc3, 0xc2, 0x61, 0xaa, 0x1a, 0xbe, 0x59, 0x70, 0x3b, 0x39, 0xf5, 0x55, 0x31, 0xe6,
	0xe5, 0xa5, 0xee, 0xf5, 0xd9, 0x2a, 0x1c, 0x1c, 0x89, 0x5d, 0x56, 0x6e, 0x58, 0x56, 0xaf, 0x7b,
	0xf5, 0xfb, 0x4b, 0xea, 0x47, 0xca, 0x0d, 0xd7, 0x2b, 0x5f, 0xe5, 0x32, 0x74, 0xd0, 0x10, 0x75,
	0x8c, 0x5a, 0x87, 0x22, 0xf8, 0x77, 0x25, 0x97, 0x90, 0x71, 0x10, 0x3c, 0x17, 0x3b, 0xd9, 0x65,
	0x16, 0x4b, 0xdc, 0x5b, 0x1a, 0xe3, 0xd9, 0xbe, 0x0f, 0xc9, 0x74, 0x52, 0xe6, 0x6b, 0x74, 0xdc,
	0x5d, 0x34, 0xed, 0x7f, 0xd9, 0x16, 0x57, 0x3c, 0xe5, 0x30, 0xcb, 0x24, 0xf8, 0xb4, 0x2d, 0xc4,
	0xdf, 0x31, 0xc1, 0x9d, 0x70, 0x39, 0xb5, 0x70, 0xe5, 0xae, 0xcd, 0xbb, 0xe7, 0x93, 0x32, 0xa7,
	0xad, 0x1f, 0x8d, 0x8f, 0xbf, 0x7e, 0x7f, 0xae, 0x7f, 0x6d, 0xb4, 0xd9, 0x27, 0x3e, 0x7d, 0x94,
	0x7d, 0x1e, 0xb2, 0x38, 0x96, 0x93, 0x67, 0x6b, 0xf3, 0x49, 0x43, 0x7d, 0x07, 0x9a, 0xe7, 0x55,
	0x78, 0x44, 0x63, 0x85, 0xa6, 0x12, 0x36, 0x6a, 0x0c, 0xf3, 0x67, 0xb5, 0xf6, 0xf1, 0xf7, 0xad,
	0xf6, 0xb7, 0xad, 0xf2, 0xf0, 0x72, 0x5c, 0x4e, 0x9e, 0xad, 0xe4, 0x17, 0x1a, 0x8a, 0xa0, 0x5c,
	0xa9, 0x30, 0xb7, 0x71, 0x6b, 0xe1, 0x7b, 0xe3, 0x4e, 0xbf, 0xd2, 0xc5, 0xfa, 0x7c, 0x05, 0xa3,
	0xb5, 0x6c, 0x5f, 0xf9, 0xcf, 0x0e, 0x39, 0xe7, 0x1c, 0xb3, 0x39, 0xa5, 0xd2, 0x55, 0x4e, 0x98,
	0x82, 0x75, 0x48, 0xeb, 0x45, 0x2c, 0xb0, 0x3d, 0xed, 0x29, 0x66, 0x18, 0x27, 0x3c, 0x3f, 0x78,
	0x7a, 0xfc, 0x24, 0x46, 0x3e, 0x99, 0xf4, 0x43, 0x4d, 0x63, 0xe9, 0x3f, 0x30, 0xb2, 0xb1, 0x2c,
	0x7e, 0xfc, 0x18, 0x8c, 0x4c, 0xfa, 0x0f, 0x63, 0x92, 0xcb, 0xaf, 0x4e, 0x7f, 0xc7, 0x3f, 0x02,
	0x8f, 0xff, 0x04, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x73, 0x21, 0xfd, 0x90, 0x04, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// CacheServiceClient is the client API for CacheService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type CacheServiceClient interface {
	// Evicts all cached data for the referenced :ref:`ref_flyteidl.admin.Execution` or :ref:`ref_flyteidl.admin.TaskExecution`.
	EvictCache(ctx context.Context, in *EvictCacheRequest, opts ...grpc.CallOption) (*EvictCacheResponse, error)
}

type cacheServiceClient struct {
	cc *grpc.ClientConn
}

func NewCacheServiceClient(cc *grpc.ClientConn) CacheServiceClient {
	return &cacheServiceClient{cc}
}

func (c *cacheServiceClient) EvictCache(ctx context.Context, in *EvictCacheRequest, opts ...grpc.CallOption) (*EvictCacheResponse, error) {
	out := new(EvictCacheResponse)
	err := c.cc.Invoke(ctx, "/flyteidl.service.CacheService/EvictCache", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// CacheServiceServer is the server API for CacheService service.
type CacheServiceServer interface {
	// Evicts all cached data for the referenced :ref:`ref_flyteidl.admin.Execution` or :ref:`ref_flyteidl.admin.TaskExecution`.
	EvictCache(context.Context, *EvictCacheRequest) (*EvictCacheResponse, error)
}

// UnimplementedCacheServiceServer can be embedded to have forward compatible implementations.
type UnimplementedCacheServiceServer struct {
}

func (*UnimplementedCacheServiceServer) EvictCache(ctx context.Context, req *EvictCacheRequest) (*EvictCacheResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method EvictCache not implemented")
}

func RegisterCacheServiceServer(s *grpc.Server, srv CacheServiceServer) {
	s.RegisterService(&_CacheService_serviceDesc, srv)
}

func _CacheService_EvictCache_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(EvictCacheRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(CacheServiceServer).EvictCache(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/flyteidl.service.CacheService/EvictCache",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(CacheServiceServer).EvictCache(ctx, req.(*EvictCacheRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _CacheService_serviceDesc = grpc.ServiceDesc{
	ServiceName: "flyteidl.service.CacheService",
	HandlerType: (*CacheServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "EvictCache",
			Handler:    _CacheService_EvictCache_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "flyteidl/service/cache.proto",
}