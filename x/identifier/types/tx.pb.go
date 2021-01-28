// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: identifier/tx.proto

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
type MsgCreateIdentifier struct {
	Context string `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	Id      string `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	// authentication represents public key associated with the did document.
	Authentication []*Authentication `protobuf:"bytes,3,rep,name=authentication,proto3" json:"authentication,omitempty"`
	// services represents each service associated with a did
	Services []*Service `protobuf:"bytes,4,rep,name=services,proto3" json:"services,omitempty"`
	Owner    string     `protobuf:"bytes,5,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgCreateIdentifier) Reset()         { *m = MsgCreateIdentifier{} }
func (m *MsgCreateIdentifier) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIdentifier) ProtoMessage()    {}
func (*MsgCreateIdentifier) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7826d264b5e2237, []int{0}
}
func (m *MsgCreateIdentifier) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIdentifier) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIdentifier.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIdentifier) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIdentifier.Merge(m, src)
}
func (m *MsgCreateIdentifier) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIdentifier) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIdentifier.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIdentifier proto.InternalMessageInfo

type MsgCreateIdentifierResponse struct {
}

func (m *MsgCreateIdentifierResponse) Reset()         { *m = MsgCreateIdentifierResponse{} }
func (m *MsgCreateIdentifierResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateIdentifierResponse) ProtoMessage()    {}
func (*MsgCreateIdentifierResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c7826d264b5e2237, []int{1}
}
func (m *MsgCreateIdentifierResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateIdentifierResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateIdentifierResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateIdentifierResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateIdentifierResponse.Merge(m, src)
}
func (m *MsgCreateIdentifierResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateIdentifierResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateIdentifierResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateIdentifierResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateIdentifier)(nil), "allinbits.cosmoscash.identifier.MsgCreateIdentifier")
	proto.RegisterType((*MsgCreateIdentifierResponse)(nil), "allinbits.cosmoscash.identifier.MsgCreateIdentifierResponse")
}

func init() { proto.RegisterFile("identifier/tx.proto", fileDescriptor_c7826d264b5e2237) }

var fileDescriptor_c7826d264b5e2237 = []byte{
	// 360 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x92, 0xbd, 0x4e, 0x02, 0x41,
	0x10, 0xc7, 0x6f, 0x41, 0x14, 0xd7, 0x84, 0x98, 0x85, 0xe2, 0x80, 0x78, 0x10, 0x2a, 0x1a, 0x6f,
	0x13, 0xb0, 0x32, 0x36, 0x7e, 0x34, 0x16, 0x34, 0x67, 0x61, 0x62, 0x63, 0xee, 0x8e, 0x65, 0xd9,
	0x04, 0x76, 0xc8, 0xed, 0xa2, 0xf0, 0x02, 0xc6, 0xca, 0xf8, 0x08, 0x3c, 0x8e, 0x25, 0xa5, 0xa5,
	0x81, 0xc6, 0xa7, 0x30, 0x86, 0x5b, 0xbe, 0x54, 0x12, 0x12, 0xbb, 0x99, 0x9d, 0xf9, 0xfd, 0xff,
	0x33, 0x77, 0x83, 0xb3, 0xa2, 0xc9, 0xa4, 0x16, 0x2d, 0xc1, 0x22, 0xaa, 0x07, 0x6e, 0x2f, 0x02,
	0x0d, 0xa4, 0xe4, 0x77, 0x3a, 0x42, 0x06, 0x42, 0x2b, 0x37, 0x04, 0xd5, 0x05, 0x15, 0xfa, 0xaa,
	0xed, 0xae, 0x3a, 0x0b, 0x79, 0x0e, 0xc0, 0x3b, 0x8c, 0xc6, 0xed, 0x41, 0xbf, 0x45, 0x7d, 0x39,
	0x34, 0x6c, 0x21, 0xc7, 0x81, 0x43, 0x1c, 0xd2, 0x59, 0x34, 0x7f, 0x2d, 0xae, 0xd9, 0xac, 0xc2,
	0x79, 0x31, 0x6f, 0x4c, 0xee, 0x0d, 0x65, 0x12, 0x53, 0xaa, 0x7c, 0x21, 0x9c, 0x6d, 0x28, 0x7e,
	0x19, 0x31, 0x5f, 0xb3, 0xeb, 0x25, 0x48, 0x6c, 0xbc, 0x17, 0x82, 0xd4, 0x6c, 0xa0, 0x6d, 0x54,
	0x46, 0xd5, 0x7d, 0x6f, 0x91, 0x92, 0x0c, 0x4e, 0x88, 0xa6, 0x9d, 0x88, 0x1f, 0x13, 0xa2, 0x49,
	0x6e, 0x71, 0xc6, 0xef, 0xeb, 0xf6, 0x8c, 0x0c, 0x7d, 0x2d, 0x40, 0xda, 0xc9, 0x72, 0xb2, 0x7a,
	0x50, 0xa3, 0xee, 0x96, 0x25, 0xdd, 0xf3, 0x1f, 0x98, 0xf7, 0x4b, 0x86, 0x5c, 0xe1, 0xb4, 0x62,
	0xd1, 0x83, 0x08, 0x99, 0xb2, 0x77, 0x62, 0xc9, 0xea, 0x56, 0xc9, 0x1b, 0x03, 0x78, 0x4b, 0x92,
	0xe4, 0x70, 0x0a, 0x1e, 0x25, 0x8b, 0xec, 0x54, 0x3c, 0xb1, 0x49, 0x4e, 0xd3, 0xcf, 0xa3, 0x92,
	0xf5, 0x39, 0x2a, 0x59, 0x95, 0x23, 0x5c, 0xdc, 0xb0, 0xbf, 0xc7, 0x54, 0x0f, 0xa4, 0x62, 0xb5,
	0x17, 0x84, 0x93, 0x0d, 0xc5, 0xc9, 0x13, 0xc2, 0x87, 0x7f, 0x3e, 0xd2, 0xc9, 0xd6, 0x79, 0x36,
	0x48, 0x17, 0xce, 0xfe, 0x43, 0x2d, 0x06, 0xba, 0x68, 0xbc, 0x4d, 0x1c, 0x34, 0x9e, 0x38, 0xe8,
	0x63, 0xe2, 0xa0, 0xd7, 0xa9, 0x63, 0x8d, 0xa7, 0x8e, 0xf5, 0x3e, 0x75, 0xac, 0xbb, 0x3a, 0x17,
	0xba, 0xdd, 0x0f, 0xdc, 0x10, 0xba, 0x74, 0xe9, 0x30, 0xff, 0xdb, 0xc7, 0x33, 0x0b, 0x3a, 0xa0,
	0xeb, 0xc7, 0x38, 0xec, 0x31, 0x15, 0xec, 0xc6, 0x67, 0x50, 0xff, 0x0e, 0x00, 0x00, 0xff, 0xff,
	0x38, 0x7e, 0xcc, 0x35, 0xa7, 0x02, 0x00, 0x00,
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
	CreateIdentifier(ctx context.Context, in *MsgCreateIdentifier, opts ...grpc.CallOption) (*MsgCreateIdentifierResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateIdentifier(ctx context.Context, in *MsgCreateIdentifier, opts ...grpc.CallOption) (*MsgCreateIdentifierResponse, error) {
	out := new(MsgCreateIdentifierResponse)
	err := c.cc.Invoke(ctx, "/allinbits.cosmoscash.identifier.Msg/CreateIdentifier", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	// CreateDidDocument defines a method for creating a new identity.
	CreateIdentifier(context.Context, *MsgCreateIdentifier) (*MsgCreateIdentifierResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateIdentifier(ctx context.Context, req *MsgCreateIdentifier) (*MsgCreateIdentifierResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateIdentifier not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateIdentifier_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateIdentifier)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateIdentifier(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allinbits.cosmoscash.identifier.Msg/CreateIdentifier",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateIdentifier(ctx, req.(*MsgCreateIdentifier))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allinbits.cosmoscash.identifier.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateIdentifier",
			Handler:    _Msg_CreateIdentifier_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "identifier/tx.proto",
}

func (m *MsgCreateIdentifier) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIdentifier) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIdentifier) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Services) > 0 {
		for iNdEx := len(m.Services) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Services[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if len(m.Authentication) > 0 {
		for iNdEx := len(m.Authentication) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Authentication[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintTx(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		i -= len(m.Context)
		copy(dAtA[i:], m.Context)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Context)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateIdentifierResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateIdentifierResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateIdentifierResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateIdentifier) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Context)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	if len(m.Authentication) > 0 {
		for _, e := range m.Authentication {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	if len(m.Services) > 0 {
		for _, e := range m.Services {
			l = e.Size()
			n += 1 + l + sovTx(uint64(l))
		}
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateIdentifierResponse) Size() (n int) {
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
func (m *MsgCreateIdentifier) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateIdentifier: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIdentifier: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
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
			m.Context = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Authentication", wireType)
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
			m.Authentication = append(m.Authentication, &Authentication{})
			if err := m.Authentication[len(m.Authentication)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Services", wireType)
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
			m.Services = append(m.Services, &Service{})
			if err := m.Services[len(m.Services)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
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
func (m *MsgCreateIdentifierResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateIdentifierResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateIdentifierResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
