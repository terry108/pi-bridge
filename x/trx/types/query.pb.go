// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: trx/query.proto

package types

import (
	context "context"
	fmt "fmt"
	query "github.com/cosmos/cosmos-sdk/types/query"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/genproto/googleapis/api/annotations"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	io "io"
	math "math"
	math_bits "math/bits"
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

// this line is used by starport scaffolding # 3
type QueryGetTrxRequest struct {
	Index string `protobuf:"bytes,1,opt,name=index,proto3" json:"index,omitempty"`
}

func (m *QueryGetTrxRequest) Reset()         { *m = QueryGetTrxRequest{} }
func (m *QueryGetTrxRequest) String() string { return proto.CompactTextString(m) }
func (*QueryGetTrxRequest) ProtoMessage()    {}
func (*QueryGetTrxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2c204dc277b005, []int{0}
}
func (m *QueryGetTrxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTrxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTrxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTrxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTrxRequest.Merge(m, src)
}
func (m *QueryGetTrxRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTrxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTrxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTrxRequest proto.InternalMessageInfo

func (m *QueryGetTrxRequest) GetIndex() string {
	if m != nil {
		return m.Index
	}
	return ""
}

type QueryGetTrxResponse struct {
	Trx *Trx `protobuf:"bytes,1,opt,name=Trx,proto3" json:"Trx,omitempty"`
}

func (m *QueryGetTrxResponse) Reset()         { *m = QueryGetTrxResponse{} }
func (m *QueryGetTrxResponse) String() string { return proto.CompactTextString(m) }
func (*QueryGetTrxResponse) ProtoMessage()    {}
func (*QueryGetTrxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2c204dc277b005, []int{1}
}
func (m *QueryGetTrxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryGetTrxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryGetTrxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryGetTrxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryGetTrxResponse.Merge(m, src)
}
func (m *QueryGetTrxResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryGetTrxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryGetTrxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryGetTrxResponse proto.InternalMessageInfo

func (m *QueryGetTrxResponse) GetTrx() *Trx {
	if m != nil {
		return m.Trx
	}
	return nil
}

type QueryAllTrxRequest struct {
	Pagination *query.PageRequest `protobuf:"bytes,1,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTrxRequest) Reset()         { *m = QueryAllTrxRequest{} }
func (m *QueryAllTrxRequest) String() string { return proto.CompactTextString(m) }
func (*QueryAllTrxRequest) ProtoMessage()    {}
func (*QueryAllTrxRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2c204dc277b005, []int{2}
}
func (m *QueryAllTrxRequest) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTrxRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTrxRequest.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTrxRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTrxRequest.Merge(m, src)
}
func (m *QueryAllTrxRequest) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTrxRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTrxRequest.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTrxRequest proto.InternalMessageInfo

func (m *QueryAllTrxRequest) GetPagination() *query.PageRequest {
	if m != nil {
		return m.Pagination
	}
	return nil
}

type QueryAllTrxResponse struct {
	Trx        []*Trx              `protobuf:"bytes,1,rep,name=Trx,proto3" json:"Trx,omitempty"`
	Pagination *query.PageResponse `protobuf:"bytes,2,opt,name=pagination,proto3" json:"pagination,omitempty"`
}

func (m *QueryAllTrxResponse) Reset()         { *m = QueryAllTrxResponse{} }
func (m *QueryAllTrxResponse) String() string { return proto.CompactTextString(m) }
func (*QueryAllTrxResponse) ProtoMessage()    {}
func (*QueryAllTrxResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_0f2c204dc277b005, []int{3}
}
func (m *QueryAllTrxResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *QueryAllTrxResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_QueryAllTrxResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *QueryAllTrxResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_QueryAllTrxResponse.Merge(m, src)
}
func (m *QueryAllTrxResponse) XXX_Size() int {
	return m.Size()
}
func (m *QueryAllTrxResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_QueryAllTrxResponse.DiscardUnknown(m)
}

var xxx_messageInfo_QueryAllTrxResponse proto.InternalMessageInfo

func (m *QueryAllTrxResponse) GetTrx() []*Trx {
	if m != nil {
		return m.Trx
	}
	return nil
}

func (m *QueryAllTrxResponse) GetPagination() *query.PageResponse {
	if m != nil {
		return m.Pagination
	}
	return nil
}

func init() {
	proto.RegisterType((*QueryGetTrxRequest)(nil), "pchainorg.pibridge.trx.QueryGetTrxRequest")
	proto.RegisterType((*QueryGetTrxResponse)(nil), "pchainorg.pibridge.trx.QueryGetTrxResponse")
	proto.RegisterType((*QueryAllTrxRequest)(nil), "pchainorg.pibridge.trx.QueryAllTrxRequest")
	proto.RegisterType((*QueryAllTrxResponse)(nil), "pchainorg.pibridge.trx.QueryAllTrxResponse")
}

func init() { proto.RegisterFile("trx/query.proto", fileDescriptor_0f2c204dc277b005) }

var fileDescriptor_0f2c204dc277b005 = []byte{
	// 405 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x92, 0xcf, 0x6a, 0x1a, 0x41,
	0x1c, 0xc7, 0x1d, 0x45, 0xa1, 0x53, 0x4a, 0x61, 0x5a, 0x4a, 0xb1, 0xb2, 0x94, 0x45, 0x6c, 0xd1,
	0x3a, 0x83, 0xf6, 0x09, 0x2c, 0x6d, 0xbd, 0xb6, 0xe2, 0xa9, 0xf4, 0x32, 0xab, 0xc3, 0x38, 0xb0,
	0xee, 0xac, 0x33, 0x63, 0x59, 0x29, 0xb9, 0xe4, 0x92, 0x4b, 0x02, 0x81, 0xbc, 0x40, 0x1e, 0x27,
	0x47, 0x21, 0x97, 0x1c, 0x83, 0xe6, 0x41, 0xc2, 0xee, 0x8c, 0xc4, 0xcd, 0x3f, 0x3d, 0x2e, 0x7c,
	0xff, 0x7c, 0xbe, 0xf3, 0x5b, 0xf8, 0xda, 0xa8, 0x84, 0xcc, 0xe6, 0x4c, 0x2d, 0x70, 0xac, 0xa4,
	0x91, 0xe8, 0x5d, 0x3c, 0x9a, 0x50, 0x11, 0x49, 0xc5, 0x71, 0x2c, 0x02, 0x25, 0xc6, 0x9c, 0x61,
	0xa3, 0x92, 0x6a, 0x8d, 0x4b, 0xc9, 0x43, 0x46, 0x68, 0x2c, 0x08, 0x8d, 0x22, 0x69, 0xa8, 0x11,
	0x32, 0xd2, 0xd6, 0x55, 0x6d, 0x8e, 0xa4, 0x9e, 0x4a, 0x4d, 0x02, 0xaa, 0x99, 0x8d, 0x23, 0xff,
	0x3a, 0x01, 0x33, 0xb4, 0x43, 0x62, 0xca, 0x45, 0x94, 0x89, 0x9d, 0xf6, 0x55, 0x5a, 0x69, 0x54,
	0x62, 0x3f, 0xfd, 0x26, 0x44, 0xbf, 0x53, 0x43, 0x9f, 0x99, 0xa1, 0x4a, 0x06, 0x6c, 0x36, 0x67,
	0xda, 0xa0, 0xb7, 0xb0, 0x2c, 0xa2, 0x31, 0x4b, 0xde, 0x83, 0x8f, 0xe0, 0xf3, 0x8b, 0x81, 0xfd,
	0xf0, 0xbf, 0xc3, 0x37, 0x39, 0xad, 0x8e, 0x65, 0xa4, 0x19, 0x6a, 0xc3, 0xd2, 0x50, 0x59, 0xe9,
	0xcb, 0xee, 0x07, 0xfc, 0xf8, 0x02, 0x9c, 0x3a, 0x52, 0x9d, 0xff, 0xd7, 0x35, 0xf6, 0xc2, 0x70,
	0xab, 0xf1, 0x27, 0x84, 0x77, 0xa8, 0x2e, 0xab, 0x81, 0xed, 0x2e, 0x9c, 0xee, 0xc2, 0xf6, 0x99,
	0xdc, 0x2e, 0xfc, 0x8b, 0x72, 0xe6, 0xbc, 0x83, 0x2d, 0xa7, 0x7f, 0x02, 0x1c, 0xe4, 0x26, 0xfe,
	0x3e, 0x64, 0x69, 0x1f, 0x48, 0xd4, 0xcf, 0xe1, 0x14, 0x33, 0x9c, 0x4f, 0x3b, 0x71, 0x6c, 0xd7,
	0x36, 0x4f, 0xf7, 0xbc, 0x08, 0xcb, 0x19, 0x0f, 0x3a, 0x06, 0x19, 0x02, 0x6a, 0x3e, 0x55, 0xfe,
	0xf0, 0x0e, 0xd5, 0xd6, 0x5e, 0x5a, 0x5b, 0xeb, 0x7f, 0x39, 0xbc, 0xbc, 0x39, 0x2b, 0x36, 0x50,
	0x9d, 0x58, 0x53, 0x5b, 0x2a, 0x4e, 0x36, 0x2e, 0xe2, 0xce, 0x4e, 0xfe, 0x67, 0xb7, 0x3c, 0x40,
	0x47, 0x00, 0x56, 0x86, 0x2a, 0xe9, 0x85, 0xe1, 0x0e, 0xa2, 0xdc, 0x9d, 0x76, 0x10, 0xe5, 0x1f,
	0xdd, 0xaf, 0x67, 0x44, 0x1e, 0xaa, 0x3d, 0x47, 0xf4, 0xed, 0xc7, 0xc5, 0xca, 0x03, 0xcb, 0x95,
	0x07, 0xae, 0x57, 0x1e, 0x38, 0x5d, 0x7b, 0x85, 0xe5, 0xda, 0x2b, 0x5c, 0xad, 0xbd, 0xc2, 0x9f,
	0x16, 0x17, 0x66, 0x32, 0x0f, 0xf0, 0x48, 0x4e, 0xf3, 0x09, 0x6d, 0x17, 0x61, 0x27, 0x99, 0x45,
	0xcc, 0x74, 0x50, 0xc9, 0x7e, 0xe8, 0xaf, 0xb7, 0x01, 0x00, 0x00, 0xff, 0xff, 0xf9, 0x2a, 0x7a,
	0x48, 0x54, 0x03, 0x00, 0x00,
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
	// Queries a trx by index.
	Trx(ctx context.Context, in *QueryGetTrxRequest, opts ...grpc.CallOption) (*QueryGetTrxResponse, error)
	// Queries a list of trx items.
	TrxAll(ctx context.Context, in *QueryAllTrxRequest, opts ...grpc.CallOption) (*QueryAllTrxResponse, error)
}

type queryClient struct {
	cc grpc1.ClientConn
}

func NewQueryClient(cc grpc1.ClientConn) QueryClient {
	return &queryClient{cc}
}

func (c *queryClient) Trx(ctx context.Context, in *QueryGetTrxRequest, opts ...grpc.CallOption) (*QueryGetTrxResponse, error) {
	out := new(QueryGetTrxResponse)
	err := c.cc.Invoke(ctx, "/pchainorg.pibridge.trx.Query/Trx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *queryClient) TrxAll(ctx context.Context, in *QueryAllTrxRequest, opts ...grpc.CallOption) (*QueryAllTrxResponse, error) {
	out := new(QueryAllTrxResponse)
	err := c.cc.Invoke(ctx, "/pchainorg.pibridge.trx.Query/TrxAll", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// QueryServer is the server API for Query service.
type QueryServer interface {
	// Queries a trx by index.
	Trx(context.Context, *QueryGetTrxRequest) (*QueryGetTrxResponse, error)
	// Queries a list of trx items.
	TrxAll(context.Context, *QueryAllTrxRequest) (*QueryAllTrxResponse, error)
}

// UnimplementedQueryServer can be embedded to have forward compatible implementations.
type UnimplementedQueryServer struct {
}

func (*UnimplementedQueryServer) Trx(ctx context.Context, req *QueryGetTrxRequest) (*QueryGetTrxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Trx not implemented")
}
func (*UnimplementedQueryServer) TrxAll(ctx context.Context, req *QueryAllTrxRequest) (*QueryAllTrxResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TrxAll not implemented")
}

func RegisterQueryServer(s grpc1.Server, srv QueryServer) {
	s.RegisterService(&_Query_serviceDesc, srv)
}

func _Query_Trx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryGetTrxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).Trx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pchainorg.pibridge.trx.Query/Trx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).Trx(ctx, req.(*QueryGetTrxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _Query_TrxAll_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(QueryAllTrxRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(QueryServer).TrxAll(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/pchainorg.pibridge.trx.Query/TrxAll",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(QueryServer).TrxAll(ctx, req.(*QueryAllTrxRequest))
	}
	return interceptor(ctx, in, info, handler)
}

var _Query_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pchainorg.pibridge.trx.Query",
	HandlerType: (*QueryServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Trx",
			Handler:    _Query_Trx_Handler,
		},
		{
			MethodName: "TrxAll",
			Handler:    _Query_TrxAll_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "trx/query.proto",
}

func (m *QueryGetTrxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTrxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTrxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Index) > 0 {
		i -= len(m.Index)
		copy(dAtA[i:], m.Index)
		i = encodeVarintQuery(dAtA, i, uint64(len(m.Index)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryGetTrxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryGetTrxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryGetTrxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Trx != nil {
		{
			size, err := m.Trx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllTrxRequest) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTrxRequest) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTrxRequest) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *QueryAllTrxResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *QueryAllTrxResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *QueryAllTrxResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Pagination != nil {
		{
			size, err := m.Pagination.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintQuery(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if len(m.Trx) > 0 {
		for iNdEx := len(m.Trx) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Trx[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintQuery(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintQuery(dAtA []byte, offset int, v uint64) int {
	offset -= sovQuery(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *QueryGetTrxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Index)
	if l > 0 {
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryGetTrxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Trx != nil {
		l = m.Trx.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTrxRequest) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func (m *QueryAllTrxResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Trx) > 0 {
		for _, e := range m.Trx {
			l = e.Size()
			n += 1 + l + sovQuery(uint64(l))
		}
	}
	if m.Pagination != nil {
		l = m.Pagination.Size()
		n += 1 + l + sovQuery(uint64(l))
	}
	return n
}

func sovQuery(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozQuery(x uint64) (n int) {
	return sovQuery(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *QueryGetTrxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetTrxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTrxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Index = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryGetTrxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryGetTrxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryGetTrxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Trx == nil {
				m.Trx = &Trx{}
			}
			if err := m.Trx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllTrxRequest) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllTrxRequest: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTrxRequest: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageRequest{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *QueryAllTrxResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= uint64(b&0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: QueryAllTrxResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: QueryAllTrxResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Trx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Trx = append(m.Trx, &Trx{})
			if err := m.Trx[len(m.Trx)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Pagination", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				msglen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if msglen < 0 {
				return ErrInvalidLengthQuery
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthQuery
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Pagination == nil {
				m.Pagination = &query.PageResponse{}
			}
			if err := m.Pagination.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipQuery(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthQuery
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipQuery(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowQuery
			}
			if iNdEx >= l {
				return 0, io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		wireType := int(wire & 0x7)
		switch wireType {
		case 0:
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
		case 1:
			iNdEx += 8
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowQuery
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				length |= (int(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if length < 0 {
				return 0, ErrInvalidLengthQuery
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupQuery
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthQuery
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthQuery        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowQuery          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupQuery = fmt.Errorf("proto: unexpected end of group")
)