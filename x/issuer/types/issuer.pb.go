// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issuer/issuer.proto

package types

import (
	fmt "fmt"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
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

// Issuer represents a dencentralised issuer.
type Issuer struct {
	Name    string `protobuf:"bytes,1,opt,name=Name,json=name,proto3" json:"Name,omitempty"`
	Token   string `protobuf:"bytes,2,opt,name=Token,json=token,proto3" json:"Token,omitempty"`
	Fee     int32  `protobuf:"varint,3,opt,name=Fee,json=fee,proto3" json:"Fee,omitempty"`
	State   string `protobuf:"bytes,4,opt,name=State,json=state,proto3" json:"State,omitempty"`
	Address string `protobuf:"bytes,5,opt,name=Address,json=address,proto3" json:"Address,omitempty"`
}

func (m *Issuer) Reset()         { *m = Issuer{} }
func (m *Issuer) String() string { return proto.CompactTextString(m) }
func (*Issuer) ProtoMessage()    {}
func (*Issuer) Descriptor() ([]byte, []int) {
	return fileDescriptor_c4fd01cb6e3441b4, []int{0}
}
func (m *Issuer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Issuer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Issuer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Issuer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Issuer.Merge(m, src)
}
func (m *Issuer) XXX_Size() int {
	return m.Size()
}
func (m *Issuer) XXX_DiscardUnknown() {
	xxx_messageInfo_Issuer.DiscardUnknown(m)
}

var xxx_messageInfo_Issuer proto.InternalMessageInfo

func (m *Issuer) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Issuer) GetToken() string {
	if m != nil {
		return m.Token
	}
	return ""
}

func (m *Issuer) GetFee() int32 {
	if m != nil {
		return m.Fee
	}
	return 0
}

func (m *Issuer) GetState() string {
	if m != nil {
		return m.State
	}
	return ""
}

func (m *Issuer) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

func init() {
	proto.RegisterType((*Issuer)(nil), "allinbits.cosmoscash.issuer.Issuer")
}

func init() { proto.RegisterFile("issuer/issuer.proto", fileDescriptor_c4fd01cb6e3441b4) }

var fileDescriptor_c4fd01cb6e3441b4 = []byte{
	// 235 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x3c, 0x8f, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x63, 0x92, 0xb4, 0xc2, 0x13, 0x32, 0x1d, 0x2c, 0x90, 0xac, 0x8a, 0xa9, 0x0b, 0xf1,
	0xc0, 0x13, 0xc0, 0x80, 0xd4, 0x85, 0xa1, 0x30, 0xb1, 0x39, 0xe9, 0x91, 0x46, 0x34, 0xb9, 0x2a,
	0x77, 0x95, 0xe0, 0x2d, 0x78, 0x2c, 0xc6, 0x8c, 0x8c, 0x28, 0x79, 0x11, 0x64, 0x1b, 0x75, 0xba,
	0xfb, 0xff, 0xfb, 0x6e, 0xf8, 0xe4, 0x65, 0x43, 0x74, 0x84, 0xde, 0xc6, 0x51, 0x1c, 0x7a, 0x64,
	0x54, 0xd7, 0x6e, 0xbf, 0x6f, 0xba, 0xb2, 0x61, 0x2a, 0x2a, 0xa4, 0x16, 0xa9, 0x72, 0xb4, 0x2b,
	0x22, 0x72, 0xb5, 0xa8, 0xb1, 0xc6, 0xc0, 0x59, 0xbf, 0xc5, 0x97, 0x1b, 0x96, 0xb3, 0x75, 0xb8,
	0x2b, 0x25, 0xb3, 0x27, 0xd7, 0x82, 0x16, 0x4b, 0xb1, 0x3a, 0xdf, 0x64, 0x9d, 0x6b, 0x41, 0x2d,
	0x64, 0xfe, 0x82, 0xef, 0xd0, 0xe9, 0xb3, 0x50, 0xe6, 0xec, 0x83, 0xba, 0x90, 0xe9, 0x23, 0x80,
	0x4e, 0x97, 0x62, 0x95, 0x6f, 0xd2, 0x37, 0x08, 0xdc, 0x33, 0x3b, 0x06, 0x9d, 0x45, 0x8e, 0x7c,
	0x50, 0x5a, 0xce, 0xef, 0xb7, 0xdb, 0x1e, 0x88, 0x74, 0x1e, 0xfa, 0xb9, 0x8b, 0xf1, 0x61, 0xfd,
	0x3d, 0x1a, 0x31, 0x8c, 0x46, 0xfc, 0x8e, 0x46, 0x7c, 0x4d, 0x26, 0x19, 0x26, 0x93, 0xfc, 0x4c,
	0x26, 0x79, 0xb5, 0x75, 0xc3, 0xbb, 0x63, 0x59, 0x54, 0xd8, 0xda, 0x93, 0x8d, 0x8d, 0x36, 0xb7,
	0x5e, 0xc7, 0x7e, 0xfc, 0x3b, 0x5b, 0xfe, 0x3c, 0x00, 0x95, 0xb3, 0xe0, 0x71, 0xf7, 0x17, 0x00,
	0x00, 0xff, 0xff, 0x0c, 0xa9, 0x27, 0xa9, 0x11, 0x01, 0x00, 0x00,
}

func (m *Issuer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Issuer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Issuer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.State) > 0 {
		i -= len(m.State)
		copy(dAtA[i:], m.State)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.State)))
		i--
		dAtA[i] = 0x22
	}
	if m.Fee != 0 {
		i = encodeVarintIssuer(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x18
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.Token)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Name) > 0 {
		i -= len(m.Name)
		copy(dAtA[i:], m.Name)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.Name)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintIssuer(dAtA []byte, offset int, v uint64) int {
	offset -= sovIssuer(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Issuer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Name)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	if m.Fee != 0 {
		n += 1 + sovIssuer(uint64(m.Fee))
	}
	l = len(m.State)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	return n
}

func sovIssuer(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozIssuer(x uint64) (n int) {
	return sovIssuer(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Issuer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowIssuer
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
			return fmt.Errorf("proto: Issuer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Issuer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Name", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
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
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
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
					return ErrIntOverflowIssuer
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
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
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
					return ErrIntOverflowIssuer
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
				return ErrInvalidLengthIssuer
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthIssuer
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipIssuer(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthIssuer
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthIssuer
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
func skipIssuer(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowIssuer
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
					return 0, ErrIntOverflowIssuer
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
					return 0, ErrIntOverflowIssuer
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
				return 0, ErrInvalidLengthIssuer
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupIssuer
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthIssuer
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthIssuer        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowIssuer          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupIssuer = fmt.Errorf("proto: unexpected end of group")
)
