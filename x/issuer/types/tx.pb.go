// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issuer/tx.proto

package types

import (
	context "context"
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
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

// MsgCreateIdentifier defines a SDK message for creating a new identifer.
type MsgCreateIssuer struct {
	Name    string `protobuf:"bytes,1,opt,name=Name,proto3" json:"Name,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=Token,proto3" json:"Token,omitempty"`
	Fee     int32  `protobuf:"varint,3,opt,name=Fee,proto3" json:"Fee,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=State,proto3" json:"State,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=Address,proto3" json:"Address,omitempty"`
}

func (m *MsgCreateIssuer) Reset()         { *m = MsgCreateIssuer{} }
func (m *MsgCreateIssuer) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIssuer) ProtoMessage()    {}
func (*MsgCreateIssuer) Descriptor() ([]byte, []int) {
	return fileDescriptor_8167a165536ee83e, []int{0}
}
func (m *MsgCreateIssuer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIssuer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIssuer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIssuer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIssuer.Merge(m, src)
}
func (m *MsgCreateIssuer) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIssuer) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIssuer.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIssuer proto.InternalMessageInfo

type MsgCreateIssuerResponse struct {
}

func (m *MsgCreateIssuerResponse) Reset()         { *m = MsgCreateIssuerResponse{} }
func (m *MsgCreateIssuerResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIssuerResponse) ProtoMessage()    {}
func (*MsgCreateIssuerResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_8167a165536ee83e, []int{1}
}
func (m *MsgCreateIssuerResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIssuerResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIssuerResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIssuerResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIssuerResponse.Merge(m, src)
}
func (m *MsgCreateIssuerResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIssuerResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIssuerResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIssuerResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateIssuer)(nil), "allinbits.cosmoscash.identifier.MsgCreateIssuer")
	proto.RegisterType((*MsgCreateIssuerResponse)(nil), "allinbits.cosmoscash.identifier.MsgCreateIssuerResponse")
}

func init() { proto.RegisterFile("issuer/tx.proto", fileDescriptor_8167a165536ee83e) }

var fileDescriptor_8167a165536ee83e = []byte{
	// 323 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x51, 0x3f, 0x4f, 0x02, 0x31,
	0x1c, 0xbd, 0x0a, 0xf8, 0xa7, 0x31, 0xc1, 0x34, 0x24, 0x1e, 0x0c, 0x07, 0x61, 0x62, 0xb1, 0x35,
	0xba, 0x18, 0x37, 0x35, 0x31, 0x61, 0xc0, 0x01, 0x9d, 0x5c, 0x4c, 0x0f, 0x7e, 0x94, 0x46, 0xb8,
	0x92, 0xfb, 0x95, 0x04, 0xdc, 0x9c, 0x74, 0xf4, 0x23, 0xf0, 0x71, 0x1c, 0x19, 0x1d, 0x0d, 0x2c,
	0x7e, 0x0c, 0x73, 0x2d, 0x98, 0xc8, 0x62, 0xdc, 0xde, 0x7b, 0x7d, 0xbf, 0x97, 0x97, 0x57, 0x5a,
	0xd4, 0x88, 0x63, 0x48, 0x85, 0x9d, 0xf0, 0x51, 0x6a, 0xac, 0x61, 0x55, 0x39, 0x18, 0xe8, 0x24,
	0xd6, 0x16, 0x79, 0xc7, 0xe0, 0xd0, 0x60, 0x47, 0x62, 0x9f, 0xeb, 0x2e, 0x24, 0x56, 0xf7, 0x34,
	0xa4, 0x95, 0xb2, 0x32, 0x46, 0x0d, 0x40, 0x38, 0x7b, 0x3c, 0xee, 0x09, 0x99, 0x4c, 0xfd, 0x6d,
	0xa5, 0xa4, 0x8c, 0x32, 0x0e, 0x8a, 0x0c, 0xad, 0xd4, 0xb2, 0xcf, 0x79, 0xf0, 0x0f, 0x9e, 0xf8,
	0xa7, 0xfa, 0x0b, 0xa1, 0xc5, 0x16, 0xaa, 0xab, 0x14, 0xa4, 0x85, 0xa6, 0x6b, 0xc2, 0x18, 0xcd,
	0xdf, 0xc8, 0x21, 0x84, 0xa4, 0x46, 0x1a, 0x7b, 0x6d, 0x87, 0x59, 0x89, 0x16, 0xee, 0xcc, 0x23,
	0x24, 0xe1, 0x96, 0x13, 0x3d, 0x61, 0x07, 0x34, 0x77, 0x0d, 0x10, 0xe6, 0x6a, 0xa4, 0x51, 0x68,
	0x67, 0x30, 0xf3, 0xdd, 0x5a, 0x69, 0x21, 0xcc, 0x7b, 0x9f, 0x23, 0x2c, 0xa4, 0x3b, 0x17, 0xdd,
	0x6e, 0x0a, 0x88, 0x61, 0xc1, 0xe9, 0x6b, 0x7a, 0xbe, 0xfb, 0x3a, 0xab, 0x06, 0x5f, 0xb3, 0x6a,
	0x50, 0x2f, 0xd3, 0xc3, 0x8d, 0x22, 0x6d, 0xc0, 0x91, 0x49, 0x10, 0x4e, 0x9e, 0x09, 0xcd, 0xb5,
	0x50, 0xb1, 0x27, 0xba, 0xff, 0xab, 0xe8, 0x31, 0xff, 0x63, 0x2a, 0xbe, 0x91, 0x58, 0x39, 0xfb,
	0xef, 0xc5, 0xba, 0xc3, 0x65, 0xf3, 0x7d, 0x11, 0x91, 0xf9, 0x22, 0x22, 0x9f, 0x8b, 0x88, 0xbc,
	0x2d, 0xa3, 0x60, 0xbe, 0x8c, 0x82, 0x8f, 0x65, 0x14, 0xdc, 0x0b, 0xa5, 0x6d, 0x7f, 0x1c, 0xf3,
	0x8e, 0x19, 0x8a, 0x9f, 0xf4, 0xd5, 0xca, 0x47, 0x59, 0xbc, 0x98, 0x88, 0xf5, 0x1f, 0x4f, 0x47,
	0x80, 0xf1, 0xb6, 0x9b, 0xfe, 0xf4, 0x3b, 0x00, 0x00, 0xff, 0xff, 0xfd, 0xfd, 0x99, 0xad, 0xfa,
	0x01, 0x00, 0x00,
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
	// CreateDidDocument defines a method for creating a new identity.
	CreateIssuer(ctx context.Context, in *MsgCreateIssuer, opts ...grpc.CallOption) (*MsgCreateIssuerResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateIssuer(ctx context.Context, in *MsgCreateIssuer, opts ...grpc.CallOption) (*MsgCreateIssuerResponse, error) {
	out := new(MsgCreateIssuerResponse)
	err := c.cc.Invoke(ctx, "/allinbits.cosmoscash.identifier.Msg/CreateIssuer", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateDidDocument defines a method for creating a new identity.
	CreateIssuer(context.Context, *MsgCreateIssuer) (*MsgCreateIssuerResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateIssuer(ctx context.Context, req *MsgCreateIssuer) (*MsgCreateIssuerResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIssuer not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateIssuer_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateIssuer)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateIssuer(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allinbits.cosmoscash.identifier.Msg/CreateIssuer",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateIssuer(ctx, req.(*MsgCreateIssuer))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allinbits.cosmoscash.identifier.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIssuer",
			Handler:    _Msg_CreateIssuer_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "issuer/tx.proto",
}

func (m *MsgCreateIssuer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIssuer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIssuer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintTx(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if m.Fee != 0 {
		i = encodeVarintTx(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateIssuerResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIssuerResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIssuerResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateIssuer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if m.Fee != 0 {
		n += 1 + sovTx(uint64(m.Fee))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateIssuerResponse) Size() (n int) {
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
func (m *MsgCreateIssuer) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateIssuer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIssuer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
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
			m.Name = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Token", wireType)
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
			m.Token = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Fee", wireType)
			}
			m.Fee = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowTx
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Fee |= int32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field State", wireType)
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
			m.State = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
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
			m.Address = string(dAtA[iNdEx:postIndex])
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
func (m *MsgCreateIssuerResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateIssuerResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIssuerResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
