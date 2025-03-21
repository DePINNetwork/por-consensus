// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: cometbft/rpc/grpc/v1beta1/types.proto

package v1beta1

import (
	context "context"
	fmt "fmt"
	v1beta1 "github.com/depinnetwork/por-consensus/api/cometbft/abci/v1beta1"
	grpc1 "github.com/cosmos/gogoproto/grpc"
	proto "github.com/cosmos/gogoproto/proto"
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

// RequestPing is a request to confirm that the connection is alive.
type RequestPing struct {
}

func (m *RequestPing) Reset()         { *m = RequestPing{} }
func (m *RequestPing) String() string { return proto.CompactTextString(m) }
func (*RequestPing) ProtoMessage()    {}
func (*RequestPing) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd787770c28cdca, []int{0}
}
func (m *RequestPing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestPing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestPing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestPing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestPing.Merge(m, src)
}
func (m *RequestPing) XXX_Size() int {
	return m.Size()
}
func (m *RequestPing) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestPing.DiscardUnknown(m)
}

var xxx_messageInfo_RequestPing proto.InternalMessageInfo

// RequestBroadcastTx is a request to broadcast the transaction.
type RequestBroadcastTx struct {
	Tx []byte `protobuf:"bytes,1,opt,name=tx,proto3" json:"tx,omitempty"`
}

func (m *RequestBroadcastTx) Reset()         { *m = RequestBroadcastTx{} }
func (m *RequestBroadcastTx) String() string { return proto.CompactTextString(m) }
func (*RequestBroadcastTx) ProtoMessage()    {}
func (*RequestBroadcastTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd787770c28cdca, []int{1}
}
func (m *RequestBroadcastTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *RequestBroadcastTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_RequestBroadcastTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *RequestBroadcastTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RequestBroadcastTx.Merge(m, src)
}
func (m *RequestBroadcastTx) XXX_Size() int {
	return m.Size()
}
func (m *RequestBroadcastTx) XXX_DiscardUnknown() {
	xxx_messageInfo_RequestBroadcastTx.DiscardUnknown(m)
}

var xxx_messageInfo_RequestBroadcastTx proto.InternalMessageInfo

func (m *RequestBroadcastTx) GetTx() []byte {
	if m != nil {
		return m.Tx
	}
	return nil
}

// ResponsePing is a response to confirm that the connection is alive.
type ResponsePing struct {
}

func (m *ResponsePing) Reset()         { *m = ResponsePing{} }
func (m *ResponsePing) String() string { return proto.CompactTextString(m) }
func (*ResponsePing) ProtoMessage()    {}
func (*ResponsePing) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd787770c28cdca, []int{2}
}
func (m *ResponsePing) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponsePing) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponsePing.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponsePing) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponsePing.Merge(m, src)
}
func (m *ResponsePing) XXX_Size() int {
	return m.Size()
}
func (m *ResponsePing) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponsePing.DiscardUnknown(m)
}

var xxx_messageInfo_ResponsePing proto.InternalMessageInfo

// ResponseBroadcastTx is a response of broadcasting the transaction.
type ResponseBroadcastTx struct {
	CheckTx   *v1beta1.ResponseCheckTx   `protobuf:"bytes,1,opt,name=check_tx,json=checkTx,proto3" json:"check_tx,omitempty"`
	DeliverTx *v1beta1.ResponseDeliverTx `protobuf:"bytes,2,opt,name=deliver_tx,json=deliverTx,proto3" json:"deliver_tx,omitempty"`
}

func (m *ResponseBroadcastTx) Reset()         { *m = ResponseBroadcastTx{} }
func (m *ResponseBroadcastTx) String() string { return proto.CompactTextString(m) }
func (*ResponseBroadcastTx) ProtoMessage()    {}
func (*ResponseBroadcastTx) Descriptor() ([]byte, []int) {
	return fileDescriptor_2bd787770c28cdca, []int{3}
}
func (m *ResponseBroadcastTx) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ResponseBroadcastTx) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ResponseBroadcastTx.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ResponseBroadcastTx) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponseBroadcastTx.Merge(m, src)
}
func (m *ResponseBroadcastTx) XXX_Size() int {
	return m.Size()
}
func (m *ResponseBroadcastTx) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponseBroadcastTx.DiscardUnknown(m)
}

var xxx_messageInfo_ResponseBroadcastTx proto.InternalMessageInfo

func (m *ResponseBroadcastTx) GetCheckTx() *v1beta1.ResponseCheckTx {
	if m != nil {
		return m.CheckTx
	}
	return nil
}

func (m *ResponseBroadcastTx) GetDeliverTx() *v1beta1.ResponseDeliverTx {
	if m != nil {
		return m.DeliverTx
	}
	return nil
}

func init() {
	proto.RegisterType((*RequestPing)(nil), "cometbft.rpc.grpc.v1beta1.RequestPing")
	proto.RegisterType((*RequestBroadcastTx)(nil), "cometbft.rpc.grpc.v1beta1.RequestBroadcastTx")
	proto.RegisterType((*ResponsePing)(nil), "cometbft.rpc.grpc.v1beta1.ResponsePing")
	proto.RegisterType((*ResponseBroadcastTx)(nil), "cometbft.rpc.grpc.v1beta1.ResponseBroadcastTx")
}

func init() {
	proto.RegisterFile("cometbft/rpc/grpc/v1beta1/types.proto", fileDescriptor_2bd787770c28cdca)
}

var fileDescriptor_2bd787770c28cdca = []byte{
	// 325 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4d, 0xce, 0xcf, 0x4d,
	0x2d, 0x49, 0x4a, 0x2b, 0xd1, 0x2f, 0x2a, 0x48, 0xd6, 0x4f, 0x07, 0x11, 0x65, 0x86, 0x49, 0xa9,
	0x25, 0x89, 0x86, 0xfa, 0x25, 0x95, 0x05, 0xa9, 0xc5, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42,
	0x92, 0x30, 0x65, 0x7a, 0x45, 0x05, 0xc9, 0x7a, 0x20, 0x65, 0x7a, 0x50, 0x65, 0x52, 0x8a, 0x70,
	0x13, 0x12, 0x93, 0x92, 0x33, 0xb1, 0xe9, 0x56, 0xe2, 0xe5, 0xe2, 0x0e, 0x4a, 0x2d, 0x2c, 0x4d,
	0x2d, 0x2e, 0x09, 0xc8, 0xcc, 0x4b, 0x57, 0x52, 0xe1, 0x12, 0x82, 0x72, 0x9d, 0x8a, 0xf2, 0x13,
	0x53, 0x92, 0x13, 0x8b, 0x4b, 0x42, 0x2a, 0x84, 0xf8, 0xb8, 0x98, 0x4a, 0x2a, 0x24, 0x18, 0x15,
	0x18, 0x35, 0x78, 0x82, 0x98, 0x4a, 0x2a, 0x94, 0xf8, 0xb8, 0x78, 0x82, 0x52, 0x8b, 0x0b, 0xf2,
	0xf3, 0x8a, 0x53, 0xc1, 0xba, 0x16, 0x32, 0x72, 0x09, 0xc3, 0x04, 0x90, 0xf5, 0x39, 0x72, 0x71,
	0x24, 0x67, 0xa4, 0x26, 0x67, 0xc7, 0x43, 0x75, 0x73, 0x1b, 0xa9, 0xe9, 0xc1, 0x5d, 0x0b, 0x72,
	0x12, 0xcc, 0xa5, 0x7a, 0x30, 0xdd, 0xce, 0x20, 0xe5, 0x21, 0x15, 0x41, 0xec, 0xc9, 0x10, 0x86,
	0x90, 0x3b, 0x17, 0x57, 0x4a, 0x6a, 0x4e, 0x66, 0x59, 0x6a, 0x11, 0xc8, 0x10, 0x26, 0xb0, 0x21,
	0x1a, 0x04, 0x0c, 0x71, 0x81, 0x68, 0x08, 0xa9, 0x08, 0xe2, 0x4c, 0x81, 0x31, 0x8d, 0xae, 0x32,
	0x72, 0xf1, 0xc0, 0xdd, 0xe6, 0x18, 0xe0, 0x29, 0x14, 0xce, 0xc5, 0x02, 0x72, 0xbc, 0x10, 0x92,
	0x93, 0xd0, 0x03, 0x50, 0x0f, 0x29, 0x68, 0xa4, 0xd4, 0xf1, 0xaa, 0x43, 0x84, 0x86, 0x50, 0x0e,
	0x17, 0x37, 0x72, 0x20, 0xe8, 0x12, 0x36, 0x1f, 0x49, 0xb9, 0x94, 0x1e, 0x11, 0xd6, 0x20, 0xa9,
	0x77, 0x0a, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39, 0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27,
	0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63, 0x39, 0x86, 0x28, 0xab, 0xf4, 0xcc,
	0x92, 0x8c, 0xd2, 0x24, 0x90, 0x79, 0xfa, 0xf0, 0x84, 0x80, 0x48, 0x11, 0x05, 0x99, 0xfa, 0x38,
	0x13, 0x58, 0x12, 0x1b, 0x38, 0x75, 0x18, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0xc7, 0x56, 0x2e,
	0x4a, 0x84, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// BroadcastAPIClient is the client API for BroadcastAPI service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type BroadcastAPIClient interface {
	// Ping the connection.
	Ping(ctx context.Context, in *RequestPing, opts ...grpc.CallOption) (*ResponsePing, error)
	// BroadcastTx broadcasts the transaction.
	BroadcastTx(ctx context.Context, in *RequestBroadcastTx, opts ...grpc.CallOption) (*ResponseBroadcastTx, error)
}

type broadcastAPIClient struct {
	cc grpc1.ClientConn
}

func NewBroadcastAPIClient(cc grpc1.ClientConn) BroadcastAPIClient {
	return &broadcastAPIClient{cc}
}

func (c *broadcastAPIClient) Ping(ctx context.Context, in *RequestPing, opts ...grpc.CallOption) (*ResponsePing, error) {
	out := new(ResponsePing)
	err := c.cc.Invoke(ctx, "/cometbft.rpc.grpc.v1beta1.BroadcastAPI/Ping", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *broadcastAPIClient) BroadcastTx(ctx context.Context, in *RequestBroadcastTx, opts ...grpc.CallOption) (*ResponseBroadcastTx, error) {
	out := new(ResponseBroadcastTx)
	err := c.cc.Invoke(ctx, "/cometbft.rpc.grpc.v1beta1.BroadcastAPI/BroadcastTx", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BroadcastAPIServer is the server API for BroadcastAPI service.
type BroadcastAPIServer interface {
	// Ping the connection.
	Ping(context.Context, *RequestPing) (*ResponsePing, error)
	// BroadcastTx broadcasts the transaction.
	BroadcastTx(context.Context, *RequestBroadcastTx) (*ResponseBroadcastTx, error)
}

// UnimplementedBroadcastAPIServer can be embedded to have forward compatible implementations.
type UnimplementedBroadcastAPIServer struct {
}

func (*UnimplementedBroadcastAPIServer) Ping(ctx context.Context, req *RequestPing) (*ResponsePing, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ping not implemented")
}
func (*UnimplementedBroadcastAPIServer) BroadcastTx(ctx context.Context, req *RequestBroadcastTx) (*ResponseBroadcastTx, error) {
	return nil, status.Errorf(codes.Unimplemented, "method BroadcastTx not implemented")
}

func RegisterBroadcastAPIServer(s grpc1.Server, srv BroadcastAPIServer) {
	s.RegisterService(&_BroadcastAPI_serviceDesc, srv)
}

func _BroadcastAPI_Ping_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestPing)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastAPIServer).Ping(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.rpc.grpc.v1beta1.BroadcastAPI/Ping",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastAPIServer).Ping(ctx, req.(*RequestPing))
	}
	return interceptor(ctx, in, info, handler)
}

func _BroadcastAPI_BroadcastTx_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestBroadcastTx)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BroadcastAPIServer).BroadcastTx(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/cometbft.rpc.grpc.v1beta1.BroadcastAPI/BroadcastTx",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BroadcastAPIServer).BroadcastTx(ctx, req.(*RequestBroadcastTx))
	}
	return interceptor(ctx, in, info, handler)
}

var BroadcastAPI_serviceDesc = _BroadcastAPI_serviceDesc
var _BroadcastAPI_serviceDesc = grpc.ServiceDesc{
	ServiceName: "cometbft.rpc.grpc.v1beta1.BroadcastAPI",
	HandlerType: (*BroadcastAPIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Ping",
			Handler:    _BroadcastAPI_Ping_Handler,
		},
		{
			MethodName: "BroadcastTx",
			Handler:    _BroadcastAPI_BroadcastTx_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "cometbft/rpc/grpc/v1beta1/types.proto",
}

func (m *RequestPing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestPing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestPing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *RequestBroadcastTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *RequestBroadcastTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *RequestBroadcastTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Tx) > 0 {
		i -= len(m.Tx)
		copy(dAtA[i:], m.Tx)
		i = encodeVarintTypes(dAtA, i, uint64(len(m.Tx)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *ResponsePing) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponsePing) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponsePing) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *ResponseBroadcastTx) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ResponseBroadcastTx) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ResponseBroadcastTx) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.DeliverTx != nil {
		{
			size, err := m.DeliverTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.CheckTx != nil {
		{
			size, err := m.CheckTx.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTypes(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintTypes(dAtA []byte, offset int, v uint64) int {
	offset -= sovTypes(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *RequestPing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *RequestBroadcastTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Tx)
	if l > 0 {
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func (m *ResponsePing) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *ResponseBroadcastTx) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CheckTx != nil {
		l = m.CheckTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	if m.DeliverTx != nil {
		l = m.DeliverTx.Size()
		n += 1 + l + sovTypes(uint64(l))
	}
	return n
}

func sovTypes(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTypes(x uint64) (n int) {
	return sovTypes(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *RequestPing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestPing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestPing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *RequestBroadcastTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: RequestBroadcastTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: RequestBroadcastTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Tx", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Tx = append(m.Tx[:0], dAtA[iNdEx:postIndex]...)
			if m.Tx == nil {
				m.Tx = []byte{}
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponsePing) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponsePing: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponsePing: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func (m *ResponseBroadcastTx) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTypes
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
			return fmt.Errorf("proto: ResponseBroadcastTx: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ResponseBroadcastTx: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CheckTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CheckTx == nil {
				m.CheckTx = &v1beta1.ResponseCheckTx{}
			}
			if err := m.CheckTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DeliverTx", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTypes
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
				return ErrInvalidLengthTypes
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTypes
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DeliverTx == nil {
				m.DeliverTx = &v1beta1.ResponseDeliverTx{}
			}
			if err := m.DeliverTx.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTypes(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTypes
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
func skipTypes(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
					return 0, ErrIntOverflowTypes
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
				return 0, ErrInvalidLengthTypes
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTypes
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTypes
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTypes        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTypes          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTypes = fmt.Errorf("proto: unexpected end of group")
)
