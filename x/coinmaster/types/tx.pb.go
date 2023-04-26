// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: coinmaster/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
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

type MsgCoinmasterMint struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount  types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgCoinmasterMint) Reset()         { *m = MsgCoinmasterMint{} }
func (m *MsgCoinmasterMint) String() string { return proto.CompactTextString(m) }
func (*MsgCoinmasterMint) ProtoMessage()    {}
func (*MsgCoinmasterMint) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f68783d02bba24b, []int{0}
}
func (m *MsgCoinmasterMint) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinmasterMint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinmasterMint.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinmasterMint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinmasterMint.Merge(m, src)
}
func (m *MsgCoinmasterMint) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinmasterMint) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinmasterMint.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinmasterMint proto.InternalMessageInfo

func (m *MsgCoinmasterMint) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCoinmasterMint) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgCoinmasterMintResponse struct {
}

func (m *MsgCoinmasterMintResponse) Reset()         { *m = MsgCoinmasterMintResponse{} }
func (m *MsgCoinmasterMintResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCoinmasterMintResponse) ProtoMessage()    {}
func (*MsgCoinmasterMintResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f68783d02bba24b, []int{1}
}
func (m *MsgCoinmasterMintResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinmasterMintResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinmasterMintResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinmasterMintResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinmasterMintResponse.Merge(m, src)
}
func (m *MsgCoinmasterMintResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinmasterMintResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinmasterMintResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinmasterMintResponse proto.InternalMessageInfo

type MsgCoinmasterBurn struct {
	Creator string     `protobuf:"bytes,1,opt,name=creator,proto3" json:"creator,omitempty"`
	Amount  types.Coin `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount"`
}

func (m *MsgCoinmasterBurn) Reset()         { *m = MsgCoinmasterBurn{} }
func (m *MsgCoinmasterBurn) String() string { return proto.CompactTextString(m) }
func (*MsgCoinmasterBurn) ProtoMessage()    {}
func (*MsgCoinmasterBurn) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f68783d02bba24b, []int{2}
}
func (m *MsgCoinmasterBurn) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinmasterBurn) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinmasterBurn.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinmasterBurn) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinmasterBurn.Merge(m, src)
}
func (m *MsgCoinmasterBurn) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinmasterBurn) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinmasterBurn.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinmasterBurn proto.InternalMessageInfo

func (m *MsgCoinmasterBurn) GetCreator() string {
	if m != nil {
		return m.Creator
	}
	return ""
}

func (m *MsgCoinmasterBurn) GetAmount() types.Coin {
	if m != nil {
		return m.Amount
	}
	return types.Coin{}
}

type MsgCoinmasterBurnResponse struct {
}

func (m *MsgCoinmasterBurnResponse) Reset()         { *m = MsgCoinmasterBurnResponse{} }
func (m *MsgCoinmasterBurnResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCoinmasterBurnResponse) ProtoMessage()    {}
func (*MsgCoinmasterBurnResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_2f68783d02bba24b, []int{3}
}
func (m *MsgCoinmasterBurnResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCoinmasterBurnResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCoinmasterBurnResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCoinmasterBurnResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCoinmasterBurnResponse.Merge(m, src)
}
func (m *MsgCoinmasterBurnResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCoinmasterBurnResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCoinmasterBurnResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCoinmasterBurnResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCoinmasterMint)(nil), "norianet.noria.coinmaster.MsgCoinmasterMint")
	proto.RegisterType((*MsgCoinmasterMintResponse)(nil), "norianet.noria.coinmaster.MsgCoinmasterMintResponse")
	proto.RegisterType((*MsgCoinmasterBurn)(nil), "norianet.noria.coinmaster.MsgCoinmasterBurn")
	proto.RegisterType((*MsgCoinmasterBurnResponse)(nil), "norianet.noria.coinmaster.MsgCoinmasterBurnResponse")
}

func init() { proto.RegisterFile("coinmaster/tx.proto", fileDescriptor_2f68783d02bba24b) }

var fileDescriptor_2f68783d02bba24b = []byte{
	// 304 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xbc, 0x92, 0xb1, 0x4e, 0xf3, 0x30,
	0x14, 0x85, 0xe3, 0xff, 0xaf, 0x8a, 0x30, 0x13, 0x81, 0xa1, 0x0d, 0x92, 0xa9, 0x3a, 0x75, 0xa0,
	0xb6, 0x5a, 0x90, 0xd8, 0xc3, 0xc0, 0x94, 0x25, 0x23, 0x9b, 0x13, 0x99, 0x10, 0xa4, 0xf8, 0x46,
	0xb6, 0x83, 0xca, 0x5b, 0xf0, 0x58, 0x1d, 0x3b, 0xb2, 0x80, 0x50, 0xf2, 0x22, 0xc8, 0x4e, 0x03,
	0x95, 0x22, 0xa4, 0xb2, 0xb0, 0x1d, 0xdb, 0xc7, 0xf7, 0xbb, 0xe7, 0xea, 0xe2, 0x93, 0x14, 0x72,
	0x59, 0x70, 0x6d, 0x84, 0x62, 0x66, 0x45, 0x4b, 0x05, 0x06, 0xfc, 0xb1, 0x04, 0x95, 0x73, 0x29,
	0x0c, 0x75, 0x82, 0x7e, 0x7b, 0x02, 0x92, 0x82, 0x2e, 0x40, 0xb3, 0x84, 0x6b, 0xc1, 0x9e, 0x16,
	0x89, 0x30, 0x7c, 0xc1, 0xec, 0x7b, 0xfb, 0x35, 0x38, 0xcd, 0x20, 0x03, 0x27, 0x99, 0x55, 0xed,
	0xed, 0xf4, 0x1e, 0x1f, 0x47, 0x3a, 0xbb, 0xf9, 0x2a, 0x13, 0xe5, 0xd2, 0xf8, 0x23, 0x7c, 0x90,
	0x2a, 0xc1, 0x0d, 0xa8, 0x11, 0x9a, 0xa0, 0xd9, 0x61, 0xdc, 0x1d, 0xfd, 0x6b, 0x3c, 0xe4, 0x05,
	0x54, 0xd2, 0x8c, 0xfe, 0x4d, 0xd0, 0xec, 0x68, 0x39, 0xa6, 0x2d, 0x95, 0x5a, 0x2a, 0xdd, 0x52,
	0xa9, 0x2d, 0x17, 0x0e, 0xd6, 0xef, 0xe7, 0x5e, 0xbc, 0xb5, 0x4f, 0xcf, 0xf0, 0xb8, 0xc7, 0x89,
	0x85, 0x2e, 0x41, 0x6a, 0xd1, 0x6b, 0x22, 0xac, 0x94, 0xfc, 0x8b, 0x26, 0x2c, 0xa7, 0x6b, 0x62,
	0xf9, 0x86, 0xf0, 0xff, 0x48, 0x67, 0xfe, 0x23, 0x1e, 0xb8, 0x21, 0x5c, 0xd0, 0x1f, 0x67, 0x4d,
	0x7b, 0x51, 0x82, 0xab, 0xdf, 0xb8, 0x3b, 0xa6, 0x65, 0xb9, 0xac, 0x7b, 0xb3, 0xac, 0x7b, 0x7f,
	0xd6, 0x6e, 0xbe, 0xf0, 0x76, 0x5d, 0x13, 0xb4, 0xa9, 0x09, 0xfa, 0xa8, 0x09, 0x7a, 0x69, 0x88,
	0xb7, 0x69, 0x88, 0xf7, 0xda, 0x10, 0xef, 0x6e, 0x9e, 0xe5, 0xe6, 0xa1, 0x4a, 0x68, 0x0a, 0x05,
	0x73, 0x05, 0xe7, 0x52, 0x98, 0x56, 0xb1, 0x15, 0xdb, 0x5d, 0xc3, 0xe7, 0x52, 0xe8, 0x64, 0xe8,
	0x36, 0xe7, 0xf2, 0x33, 0x00, 0x00, 0xff, 0xff, 0x7d, 0x48, 0x8f, 0x08, 0xa1, 0x02, 0x00, 0x00,
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion4

// MsgClient is the client API for Msg service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type MsgClient interface {
	Mint(ctx context.Context, in *MsgCoinmasterMint, opts ...grpc.CallOption) (*MsgCoinmasterMintResponse, error)
	Burn(ctx context.Context, in *MsgCoinmasterBurn, opts ...grpc.CallOption) (*MsgCoinmasterBurnResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) Mint(ctx context.Context, in *MsgCoinmasterMint, opts ...grpc.CallOption) (*MsgCoinmasterMintResponse, error) {
	out := new(MsgCoinmasterMintResponse)
	err := c.cc.Invoke(ctx, "/norianet.noria.coinmaster.Msg/Mint", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *msgClient) Burn(ctx context.Context, in *MsgCoinmasterBurn, opts ...grpc.CallOption) (*MsgCoinmasterBurnResponse, error) {
	out := new(MsgCoinmasterBurnResponse)
	err := c.cc.Invoke(ctx, "/norianet.noria.coinmaster.Msg/Burn", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	Mint(context.Context, *MsgCoinmasterMint) (*MsgCoinmasterMintResponse, error)
	Burn(context.Context, *MsgCoinmasterBurn) (*MsgCoinmasterBurnResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) Mint(ctx context.Context, req *MsgCoinmasterMint) (*MsgCoinmasterMintResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Mint not implemented")
}
func (*UnimplementedMsgServer) Burn(ctx context.Context, req *MsgCoinmasterBurn) (*MsgCoinmasterBurnResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Burn not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_Mint_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCoinmasterMint)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Mint(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/norianet.noria.coinmaster.Msg/Mint",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Mint(ctx, req.(*MsgCoinmasterMint))
	}
	return interceptor(ctx, in, info, handler)
}

func _Msg_Burn_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCoinmasterBurn)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).Burn(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/norianet.noria.coinmaster.Msg/Burn",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).Burn(ctx, req.(*MsgCoinmasterBurn))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "norianet.noria.coinmaster.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Mint",
			Handler:    _Msg_Mint_Handler,
		},
		{
			MethodName: "Burn",
			Handler:    _Msg_Burn_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "coinmaster/tx.proto",
}

func (m *MsgCoinmasterMint) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinmasterMint) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinmasterMint) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCoinmasterMintResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinmasterMintResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinmasterMintResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func (m *MsgCoinmasterBurn) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinmasterBurn) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinmasterBurn) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size, err := m.Amount.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.Creator) > 0 {
		i -= len(m.Creator)
		copy(dAtA[i:], m.Creator)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Creator)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCoinmasterBurnResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCoinmasterBurnResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCoinmasterBurnResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	return len(dAtA) - i, nil
}

func encodeVarintTx(dAtA []byte, offset int, v uint64) int {
	offset -= sovTx(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *MsgCoinmasterMint) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCoinmasterMintResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func (m *MsgCoinmasterBurn) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Creator)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.Amount.Size()
	n += 1 + l + sovTx(uint64(l))
	return n
}

func (m *MsgCoinmasterBurnResponse) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	return n
}

func sovTx(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozTx(x uint64) (n int) {
	return sovTx(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *MsgCoinmasterMint) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCoinmasterMint: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinmasterMint: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCoinmasterMintResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCoinmasterMintResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinmasterMintResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCoinmasterBurn) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCoinmasterBurn: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinmasterBurn: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Creator", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Creator = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Amount", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
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
				return ErrInvalidLengthTx
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthTx
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Amount.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func (m *MsgCoinmasterBurnResponse) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowTx
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
			return fmt.Errorf("proto: MsgCoinmasterBurnResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCoinmasterBurnResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthTx
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
func skipTx(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
					return 0, ErrIntOverflowTx
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
				return 0, ErrInvalidLengthTx
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupTx
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthTx
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthTx        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowTx          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupTx = fmt.Errorf("proto: unexpected end of group")
)
