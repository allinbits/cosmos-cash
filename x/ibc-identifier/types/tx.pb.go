// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: ibc-identifier/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	types "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	_ "github.com/gogo/protobuf/gogoproto"
	grpc1 "github.com/gogo/protobuf/grpc"
	proto "github.com/gogo/protobuf/proto"
	_ "github.com/regen-network/cosmos-proto"
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

// MsgTransferIdentifierIBC defines a SDK message for creating a new identifer.
type MsgTransferIdentifierIBC struct {
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// the port by which the packet will be sent
	SourcePort string `protobuf:"bytes,2,opt,name=source_port,json=sourcePort,proto3" json:"source_port,omitempty" yaml:"source_port"`
	// the channel by which the packet will be sent
	SourceChannel string `protobuf:"bytes,3,opt,name=source_channel,json=sourceChannel,proto3" json:"source_channel,omitempty" yaml:"source_channel"`
	// Timeout height relative to the current block height.
	// The timeout is disabled when set to 0.
	TimeoutHeight types.Height `protobuf:"bytes,4,opt,name=timeout_height,json=timeoutHeight,proto3" json:"timeout_height" yaml:"timeout_height"`
	// Timeout timestamp (in nanoseconds) relative to the current block timestamp.
	// The timeout is disabled when set to 0.
	TimeoutTimestamp uint64 `protobuf:"varint,5,opt,name=timeout_timestamp,json=timeoutTimestamp,proto3" json:"timeout_timestamp,omitempty" yaml:"timeout_timestamp"`
	Owner            string `protobuf:"bytes,6,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgTransferIdentifierIBC) Reset()         { *m = MsgTransferIdentifierIBC{} }
func (m *MsgTransferIdentifierIBC) String() string { return proto.CompactTextString(m) }
func (*MsgTransferIdentifierIBC) ProtoMessage()    {}
func (*MsgTransferIdentifierIBC) Descriptor() ([]byte, []int) {
	return fileDescriptor_7651f29490ef7140, []int{0}
}
func (m *MsgTransferIdentifierIBC) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferIdentifierIBC) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferIdentifierIBC.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferIdentifierIBC) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferIdentifierIBC.Merge(m, src)
}
func (m *MsgTransferIdentifierIBC) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferIdentifierIBC) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferIdentifierIBC.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferIdentifierIBC proto.InternalMessageInfo

type MsgTransferIdentifierIBCResponse struct {
}

func (m *MsgTransferIdentifierIBCResponse) Reset()         { *m = MsgTransferIdentifierIBCResponse{} }
func (m *MsgTransferIdentifierIBCResponse) String() string { return proto.CompactTextString(m) }
func (*MsgTransferIdentifierIBCResponse) ProtoMessage()    {}
func (*MsgTransferIdentifierIBCResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_7651f29490ef7140, []int{1}
}
func (m *MsgTransferIdentifierIBCResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgTransferIdentifierIBCResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgTransferIdentifierIBCResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgTransferIdentifierIBCResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgTransferIdentifierIBCResponse.Merge(m, src)
}
func (m *MsgTransferIdentifierIBCResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgTransferIdentifierIBCResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgTransferIdentifierIBCResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgTransferIdentifierIBCResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgTransferIdentifierIBC)(nil), "allinbits.cosmoscash.ibcidentifier.MsgTransferIdentifierIBC")
	proto.RegisterType((*MsgTransferIdentifierIBCResponse)(nil), "allinbits.cosmoscash.ibcidentifier.MsgTransferIdentifierIBCResponse")
}

func init() { proto.RegisterFile("ibc-identifier/tx.proto", fileDescriptor_7651f29490ef7140) }

var fileDescriptor_7651f29490ef7140 = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xa4, 0x53, 0xbd, 0x6e, 0xdb, 0x3e,
	0x10, 0x97, 0x9c, 0x0f, 0xfc, 0xff, 0x0c, 0x62, 0xb4, 0x42, 0xd2, 0xca, 0x46, 0x2b, 0x19, 0x9c,
	0xbc, 0x84, 0x44, 0xd2, 0x21, 0x40, 0xd0, 0xa1, 0x50, 0x3a, 0xd4, 0x43, 0x80, 0x56, 0xc8, 0xd4,
	0xc5, 0x95, 0x68, 0x5a, 0x22, 0x20, 0x91, 0x02, 0x49, 0xa7, 0xf1, 0x1b, 0x74, 0xec, 0xde, 0x25,
	0x40, 0x5f, 0x26, 0x63, 0xc6, 0x4e, 0x46, 0x61, 0x2f, 0x9d, 0xfd, 0x04, 0x85, 0x44, 0xda, 0xa9,
	0x8b, 0x06, 0x1d, 0x3a, 0xe9, 0xee, 0xf7, 0xc1, 0xe3, 0x9d, 0x8e, 0xe0, 0x29, 0x4b, 0xc9, 0x11,
	0x1b, 0x51, 0xae, 0xd9, 0x98, 0x51, 0x89, 0xf5, 0x35, 0xaa, 0xa4, 0xd0, 0xc2, 0x83, 0x49, 0x51,
	0x30, 0x9e, 0x32, 0xad, 0x10, 0x11, 0xaa, 0x14, 0x8a, 0x24, 0x2a, 0x47, 0x2c, 0x25, 0xf7, 0xe2,
	0x6e, 0x27, 0x13, 0x22, 0x2b, 0x28, 0x6e, 0x1c, 0xe9, 0x64, 0x8c, 0x13, 0x3e, 0x35, 0xf6, 0xee,
	0x41, 0x26, 0x32, 0xd1, 0x84, 0xb8, 0x8e, 0x2c, 0xda, 0x31, 0x47, 0x0d, 0x0d, 0x61, 0x12, 0x4b,
	0x85, 0x2c, 0x25, 0x98, 0x08, 0x49, 0x31, 0x29, 0x18, 0xe5, 0x1a, 0x5f, 0x1d, 0xdb, 0xc8, 0x08,
	0xe0, 0xb2, 0x05, 0xfc, 0x0b, 0x95, 0x5d, 0xca, 0x84, 0xab, 0x31, 0x95, 0x83, 0xf5, 0x35, 0x06,
	0xd1, 0xb9, 0xd7, 0x06, 0x2d, 0x36, 0xf2, 0xdd, 0x9e, 0xdb, 0xff, 0x3f, 0x6e, 0xb1, 0x91, 0x77,
	0x0a, 0xf6, 0x94, 0x98, 0x48, 0x42, 0x87, 0x95, 0x90, 0xda, 0x6f, 0xd5, 0x44, 0xf4, 0x64, 0x39,
	0x0b, 0xbd, 0x69, 0x52, 0x16, 0x67, 0xf0, 0x17, 0x12, 0xc6, 0xc0, 0x64, 0x6f, 0x85, 0xd4, 0xde,
	0x2b, 0xd0, 0xb6, 0x1c, 0xc9, 0x13, 0xce, 0x69, 0xe1, 0x6f, 0x35, 0xde, 0xce, 0x72, 0x16, 0x1e,
	0x6e, 0x78, 0x2d, 0x0f, 0xe3, 0x7d, 0x03, 0x9c, 0x9b, 0xdc, 0xfb, 0x00, 0xda, 0x9a, 0x95, 0x54,
	0x4c, 0xf4, 0x30, 0xa7, 0x2c, 0xcb, 0xb5, 0xbf, 0xdd, 0x73, 0xfb, 0x7b, 0x27, 0xdd, 0x7a, 0x78,
	0xa8, 0xee, 0x10, 0xd9, 0xbe, 0xae, 0x8e, 0xd1, 0x9b, 0x46, 0x11, 0x3d, 0xbf, 0x9d, 0x85, 0xce,
	0x7d, 0x85, 0x4d, 0x3f, 0x8c, 0xf7, 0x2d, 0x60, 0xd4, 0xde, 0x00, 0x3c, 0x5e, 0x29, 0xea, 0xaf,
	0xd2, 0x49, 0x59, 0xf9, 0x3b, 0x3d, 0xb7, 0xbf, 0x1d, 0x3d, 0x5b, 0xce, 0x42, 0x7f, 0xf3, 0x90,
	0xb5, 0x04, 0xc6, 0x8f, 0x2c, 0x76, 0xb9, 0x82, 0xbc, 0x03, 0xb0, 0x23, 0x3e, 0x72, 0x2a, 0xfd,
	0xdd, 0x66, 0x74, 0x26, 0x39, 0xfb, 0xef, 0xd3, 0x4d, 0xe8, 0xfc, 0xb8, 0x09, 0x1d, 0x08, 0x41,
	0xef, 0xa1, 0x99, 0xc7, 0x54, 0x55, 0x82, 0x2b, 0x7a, 0xf2, 0xd5, 0x05, 0x5b, 0x17, 0x2a, 0xf3,
	0xbe, 0xb8, 0xe0, 0xf0, 0xcf, 0x7f, 0xe7, 0x25, 0xfa, 0xfb, 0x32, 0xa1, 0x87, 0xea, 0x74, 0x5f,
	0xff, 0x8b, 0x7b, 0x75, 0xcb, 0xe8, 0xdd, 0xed, 0x3c, 0x70, 0xef, 0xe6, 0x81, 0xfb, 0x7d, 0x1e,
	0xb8, 0x9f, 0x17, 0x81, 0x73, 0xb7, 0x08, 0x9c, 0x6f, 0x8b, 0xc0, 0x79, 0x7f, 0x9a, 0x31, 0x9d,
	0x4f, 0x52, 0x44, 0x44, 0x89, 0xd7, 0x95, 0xec, 0x72, 0x1e, 0xd5, 0xa5, 0xf0, 0x35, 0xfe, 0xfd,
	0x95, 0x4c, 0x2b, 0xaa, 0xd2, 0xdd, 0x66, 0x31, 0x5f, 0xfc, 0x0c, 0x00, 0x00, 0xff, 0xff, 0x9d,
	0x24, 0xe6, 0x00, 0x44, 0x03, 0x00, 0x00,
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
	// TransferIdentifierIBC defines a method for transfering an identifier to another chain.
	TransferIdentifierIBC(ctx context.Context, in *MsgTransferIdentifierIBC, opts ...grpc.CallOption) (*MsgTransferIdentifierIBCResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) TransferIdentifierIBC(ctx context.Context, in *MsgTransferIdentifierIBC, opts ...grpc.CallOption) (*MsgTransferIdentifierIBCResponse, error) {
	out := new(MsgTransferIdentifierIBCResponse)
	err := c.cc.Invoke(ctx, "/allinbits.cosmoscash.ibcidentifier.Msg/TransferIdentifierIBC", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// TransferIdentifierIBC defines a method for transfering an identifier to another chain.
	TransferIdentifierIBC(context.Context, *MsgTransferIdentifierIBC) (*MsgTransferIdentifierIBCResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) TransferIdentifierIBC(ctx context.Context, req *MsgTransferIdentifierIBC) (*MsgTransferIdentifierIBCResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method TransferIdentifierIBC not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_TransferIdentifierIBC_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgTransferIdentifierIBC)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).TransferIdentifierIBC(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allinbits.cosmoscash.ibcidentifier.Msg/TransferIdentifierIBC",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).TransferIdentifierIBC(ctx, req.(*MsgTransferIdentifierIBC))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allinbits.cosmoscash.ibcidentifier.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "TransferIdentifierIBC",
			Handler:    _Msg_TransferIdentifierIBC_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "ibc-identifier/tx.proto",
}

func (m *MsgTransferIdentifierIBC) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferIdentifierIBC) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferIdentifierIBC) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x32
	}
	if m.TimeoutTimestamp != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.TimeoutTimestamp))
		i--
		dAtA[i] = 0x28
	}
	{
		size, err := m.TimeoutHeight.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintTx(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if len(m.SourceChannel) > 0 {
		i -= len(m.SourceChannel)
		copy(dAtA[i:], m.SourceChannel)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourceChannel)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.SourcePort) > 0 {
		i -= len(m.SourcePort)
		copy(dAtA[i:], m.SourcePort)
		i = encodeVarintTx(dAtA, i, uint64(len(m.SourcePort)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgTransferIdentifierIBCResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgTransferIdentifierIBCResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgTransferIdentifierIBCResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgTransferIdentifierIBC) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourcePort)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.SourceChannel)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = m.TimeoutHeight.Size()
	n += 1 + l + sovTx(uint64(l))
	if m.TimeoutTimestamp != 0 {
		n += 1 + sovTx(uint64(m.TimeoutTimestamp))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgTransferIdentifierIBCResponse) Size() (n int) {
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
func (m *MsgTransferIdentifierIBC) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgTransferIdentifierIBC: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferIdentifierIBC: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
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
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourcePort", wireType)
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
			m.SourcePort = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SourceChannel", wireType)
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
			m.SourceChannel = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutHeight", wireType)
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
			if err := m.TimeoutHeight.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field TimeoutTimestamp", wireType)
			}
			m.TimeoutTimestamp = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.TimeoutTimestamp |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
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
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
func (m *MsgTransferIdentifierIBCResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgTransferIdentifierIBCResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgTransferIdentifierIBCResponse: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		default:
			iNdEx = preIndex
			skippy, err := skipTx(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthTx
			}
			if (iNdEx + skippy) < 0 {
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
