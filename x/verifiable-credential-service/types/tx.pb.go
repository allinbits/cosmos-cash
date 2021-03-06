// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifiable-credential-service/tx.proto

package types

import (
	context "context"
	fmt "fmt"
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

// MsgCreateVerifiableCredential defines a SDK message for creating a new identifer.
type MsgCreateVerifiableCredential struct {
	VerifiableCredential *VerifiableCredential `protobuf:"bytes,1,opt,name=verifiable_credential,json=verifiableCredential,proto3" json:"verifiable_credential,omitempty"`
	// owner represents the user creating the message
	Owner string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
}

func (m *MsgCreateVerifiableCredential) Reset()         { *m = MsgCreateVerifiableCredential{} }
func (m *MsgCreateVerifiableCredential) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVerifiableCredential) ProtoMessage()    {}
func (*MsgCreateVerifiableCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e22bfe9514fc1a, []int{0}
}
func (m *MsgCreateVerifiableCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVerifiableCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVerifiableCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVerifiableCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVerifiableCredential.Merge(m, src)
}
func (m *MsgCreateVerifiableCredential) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVerifiableCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVerifiableCredential.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVerifiableCredential proto.InternalMessageInfo

type MsgCreateVerifiableCredentialResponse struct {
}

func (m *MsgCreateVerifiableCredentialResponse) Reset()         { *m = MsgCreateVerifiableCredentialResponse{} }
func (m *MsgCreateVerifiableCredentialResponse) String() string { return proto.CompactTextString(m) }
func (*MsgCreateVerifiableCredentialResponse) ProtoMessage()    {}
func (*MsgCreateVerifiableCredentialResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_c6e22bfe9514fc1a, []int{1}
}
func (m *MsgCreateVerifiableCredentialResponse) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *MsgCreateVerifiableCredentialResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_MsgCreateVerifiableCredentialResponse.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *MsgCreateVerifiableCredentialResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MsgCreateVerifiableCredentialResponse.Merge(m, src)
}
func (m *MsgCreateVerifiableCredentialResponse) XXX_Size() int {
	return m.Size()
}
func (m *MsgCreateVerifiableCredentialResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_MsgCreateVerifiableCredentialResponse.DiscardUnknown(m)
}

var xxx_messageInfo_MsgCreateVerifiableCredentialResponse proto.InternalMessageInfo

func init() {
	proto.RegisterType((*MsgCreateVerifiableCredential)(nil), "allinbits.cosmoscash.verifiablecredentialservice.MsgCreateVerifiableCredential")
	proto.RegisterType((*MsgCreateVerifiableCredentialResponse)(nil), "allinbits.cosmoscash.verifiablecredentialservice.MsgCreateVerifiableCredentialResponse")
}

func init() {
	proto.RegisterFile("verifiable-credential-service/tx.proto", fileDescriptor_c6e22bfe9514fc1a)
}

var fileDescriptor_c6e22bfe9514fc1a = []byte{
	// 299 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x2b, 0x4b, 0x2d, 0xca,
	0x4c, 0xcb, 0x4c, 0x4c, 0xca, 0x49, 0xd5, 0x4d, 0x2e, 0x4a, 0x4d, 0x49, 0xcd, 0x2b, 0xc9, 0x4c,
	0xcc, 0xd1, 0x2d, 0x4e, 0x2d, 0x2a, 0xcb, 0x4c, 0x4e, 0xd5, 0x2f, 0xa9, 0xd0, 0x2b, 0x28, 0xca,
	0x2f, 0xc9, 0x17, 0x32, 0x48, 0xcc, 0xc9, 0xc9, 0xcc, 0x4b, 0xca, 0x2c, 0x29, 0xd6, 0x4b, 0xce,
	0x2f, 0xce, 0xcd, 0x2f, 0x4e, 0x4e, 0x2c, 0xce, 0xd0, 0x43, 0x68, 0x46, 0xe8, 0x85, 0x6a, 0x95,
	0x12, 0x49, 0xcf, 0x4f, 0xcf, 0x07, 0x6b, 0xd6, 0x07, 0xb1, 0x20, 0xe6, 0x48, 0x59, 0xe2, 0xb7,
	0x0f, 0xab, 0x2c, 0x44, 0xab, 0xd2, 0x1e, 0x46, 0x2e, 0x59, 0xdf, 0xe2, 0x74, 0xe7, 0xa2, 0xd4,
	0xc4, 0x92, 0xd4, 0x30, 0xb8, 0x42, 0x67, 0xb8, 0x3a, 0xa1, 0x6a, 0x2e, 0x51, 0x84, 0x01, 0xf1,
	0x08, 0x03, 0x24, 0x18, 0x15, 0x18, 0x35, 0xb8, 0x8d, 0xdc, 0xf4, 0x48, 0xf5, 0x84, 0x1e, 0x36,
	0x6b, 0x82, 0x44, 0xca, 0xb0, 0x59, 0x2e, 0xc2, 0xc5, 0x9a, 0x5f, 0x9e, 0x97, 0x5a, 0x24, 0xc1,
	0xa4, 0xc0, 0xa8, 0xc1, 0x19, 0x04, 0xe1, 0x58, 0x71, 0x74, 0x2c, 0x90, 0x67, 0x78, 0xb1, 0x40,
	0x9e, 0x41, 0x49, 0x9d, 0x4b, 0x15, 0xaf, 0xeb, 0x83, 0x52, 0x8b, 0x0b, 0xf2, 0xf3, 0x8a, 0x53,
	0x8d, 0xce, 0x31, 0x72, 0x31, 0xfb, 0x16, 0xa7, 0x0b, 0x1d, 0x63, 0xe4, 0x92, 0xc2, 0xe3, 0x59,
	0x7f, 0xd2, 0x7d, 0x83, 0xd7, 0x7e, 0xa9, 0x70, 0x2a, 0x1b, 0x08, 0xf3, 0x90, 0x53, 0xc2, 0x89,
	0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38, 0xe1, 0xb1, 0x1c, 0xc3,
	0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0xb9, 0xa5, 0x67, 0x96, 0x64, 0x94, 0x26,
	0xe9, 0x25, 0xe7, 0xe7, 0xea, 0xc3, 0x2d, 0xd7, 0x87, 0x58, 0xae, 0x0b, 0xb2, 0x5d, 0xbf, 0x42,
	0x9f, 0x40, 0x02, 0xad, 0x2c, 0x48, 0x2d, 0x4e, 0x62, 0x03, 0xa7, 0x10, 0x63, 0x40, 0x00, 0x00,
	0x00, 0xff, 0xff, 0x7e, 0xb7, 0xa3, 0x74, 0xce, 0x02, 0x00, 0x00,
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
	CreateVerifiableCredential(ctx context.Context, in *MsgCreateVerifiableCredential, opts ...grpc.CallOption) (*MsgCreateVerifiableCredentialResponse, error)
}

type msgClient struct {
	cc grpc1.ClientConn
}

func NewMsgClient(cc grpc1.ClientConn) MsgClient {
	return &msgClient{cc}
}

func (c *msgClient) CreateVerifiableCredential(ctx context.Context, in *MsgCreateVerifiableCredential, opts ...grpc.CallOption) (*MsgCreateVerifiableCredentialResponse, error) {
	out := new(MsgCreateVerifiableCredentialResponse)
	err := c.cc.Invoke(ctx, "/allinbits.cosmoscash.verifiablecredentialservice.Msg/CreateVerifiableCredential", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MsgServer is the server API for Msg service.
type MsgServer interface {
	CreateVerifiableCredential(context.Context, *MsgCreateVerifiableCredential) (*MsgCreateVerifiableCredentialResponse, error)
}

// UnimplementedMsgServer can be embedded to have forward compatible implementations.
type UnimplementedMsgServer struct {
}

func (*UnimplementedMsgServer) CreateVerifiableCredential(ctx context.Context, req *MsgCreateVerifiableCredential) (*MsgCreateVerifiableCredentialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateVerifiableCredential not implemented")
}

func RegisterMsgServer(s grpc1.Server, srv MsgServer) {
	s.RegisterService(&_Msg_serviceDesc, srv)
}

func _Msg_CreateVerifiableCredential_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(MsgCreateVerifiableCredential)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MsgServer).CreateVerifiableCredential(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/allinbits.cosmoscash.verifiablecredentialservice.Msg/CreateVerifiableCredential",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MsgServer).CreateVerifiableCredential(ctx, req.(*MsgCreateVerifiableCredential))
	}
	return interceptor(ctx, in, info, handler)
}

var _Msg_serviceDesc = grpc.ServiceDesc{
	ServiceName: "allinbits.cosmoscash.verifiablecredentialservice.Msg",
	HandlerType: (*MsgServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreateVerifiableCredential",
			Handler:    _Msg_CreateVerifiableCredential_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "verifiable-credential-service/tx.proto",
}

func (m *MsgCreateVerifiableCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVerifiableCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVerifiableCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintTx(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.VerifiableCredential != nil {
		{
			size, err := m.VerifiableCredential.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintTx(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *MsgCreateVerifiableCredentialResponse) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *MsgCreateVerifiableCredentialResponse) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *MsgCreateVerifiableCredentialResponse) MarshalToSizedBuffer(dAtA []byte) (int, error) {
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
func (m *MsgCreateVerifiableCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.VerifiableCredential != nil {
		l = m.VerifiableCredential.Size()
		n += 1 + l + sovTx(uint64(l))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovTx(uint64(l))
	}
	return n
}

func (m *MsgCreateVerifiableCredentialResponse) Size() (n int) {
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
func (m *MsgCreateVerifiableCredential) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateVerifiableCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVerifiableCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerifiableCredential", wireType)
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
			if m.VerifiableCredential == nil {
				m.VerifiableCredential = &VerifiableCredential{}
			}
			if err := m.VerifiableCredential.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
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
func (m *MsgCreateVerifiableCredentialResponse) Unmarshal(dAtA []byte) error {
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
			return fmt.Errorf("proto: MsgCreateVerifiableCredentialResponse: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: MsgCreateVerifiableCredentialResponse: illegal tag %d (wire type %d)", fieldNum, wire)
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
