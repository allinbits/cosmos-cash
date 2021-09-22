// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: issuer/issuer.proto

package types

import (
	fmt "fmt"
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

// Issuer represents an e-money token issuer
type Issuer struct {
	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`
	Fee       int32  `protobuf:"varint,2,opt,name=fee,proto3" json:"fee,omitempty"`
	IssuerDid string `protobuf:"bytes,3,opt,name=issuer_did,json=issuerDid,proto3" json:"issuer_did,omitempty"`
	Paused    bool   `protobuf:"varint,4,opt,name=paused,proto3" json:"paused,omitempty"`
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

func (m *Issuer) GetIssuerDid() string {
	if m != nil {
		return m.IssuerDid
	}
	return ""
}

func (m *Issuer) GetPaused() bool {
	if m != nil {
		return m.Paused
	}
	return false
}

func init() {
	proto.RegisterType((*Issuer)(nil), "allinbits.cosmoscash.issuer.Issuer")
}

func init() { proto.RegisterFile("issuer/issuer.proto", fileDescriptor_c4fd01cb6e3441b4) }

var fileDescriptor_c4fd01cb6e3441b4 = []byte{
	// 209 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xce, 0x2c, 0x2e, 0x2e,
	0x4d, 0x2d, 0xd2, 0x87, 0x50, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0xd2, 0x89, 0x39, 0x39,
	0x99, 0x79, 0x49, 0x99, 0x25, 0xc5, 0x7a, 0xc9, 0xf9, 0xc5, 0xb9, 0xf9, 0xc5, 0xc9, 0x89, 0xc5,
	0x19, 0x7a, 0x10, 0x25, 0x4a, 0xe9, 0x5c, 0x6c, 0x9e, 0x60, 0x96, 0x90, 0x08, 0x17, 0x6b, 0x49,
	0x7e, 0x76, 0x6a, 0x9e, 0x04, 0xa3, 0x02, 0xa3, 0x06, 0x67, 0x10, 0x84, 0x23, 0x24, 0xc0, 0xc5,
	0x9c, 0x96, 0x9a, 0x2a, 0xc1, 0xa4, 0xc0, 0xa8, 0xc1, 0x1a, 0x04, 0x62, 0x0a, 0xc9, 0x72, 0x71,
	0x41, 0xf4, 0xc6, 0xa7, 0x64, 0xa6, 0x48, 0x30, 0x83, 0x15, 0x73, 0x42, 0x44, 0x5c, 0x32, 0x53,
	0x84, 0xc4, 0xb8, 0xd8, 0x0a, 0x12, 0x4b, 0x8b, 0x53, 0x53, 0x24, 0x58, 0x14, 0x18, 0x35, 0x38,
	0x82, 0xa0, 0x3c, 0x27, 0xcf, 0x13, 0x8f, 0xe4, 0x18, 0x2f, 0x3c, 0x92, 0x63, 0x7c, 0xf0, 0x48,
	0x8e, 0x71, 0xc2, 0x63, 0x39, 0x86, 0x0b, 0x8f, 0xe5, 0x18, 0x6e, 0x3c, 0x96, 0x63, 0x88, 0xd2,
	0x4f, 0xcf, 0x2c, 0xc9, 0x28, 0x4d, 0xd2, 0x4b, 0xce, 0xcf, 0xd5, 0x87, 0x3b, 0x55, 0x1f, 0xe2,
	0x54, 0x5d, 0x90, 0x5b, 0xf5, 0x2b, 0xa0, 0x1e, 0xd2, 0x2f, 0xa9, 0x2c, 0x48, 0x2d, 0x4e, 0x62,
	0x03, 0xfb, 0xcb, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x36, 0x3d, 0x0d, 0xee, 0x00, 0x00,
	0x00,
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
	if m.Paused {
		i--
		if m.Paused {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x20
	}
	if len(m.IssuerDid) > 0 {
		i -= len(m.IssuerDid)
		copy(dAtA[i:], m.IssuerDid)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.IssuerDid)))
		i--
		dAtA[i] = 0x1a
	}
	if m.Fee != 0 {
		i = encodeVarintIssuer(dAtA, i, uint64(m.Fee))
		i--
		dAtA[i] = 0x10
	}
	if len(m.Token) > 0 {
		i -= len(m.Token)
		copy(dAtA[i:], m.Token)
		i = encodeVarintIssuer(dAtA, i, uint64(len(m.Token)))
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
	l = len(m.Token)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	if m.Fee != 0 {
		n += 1 + sovIssuer(uint64(m.Fee))
	}
	l = len(m.IssuerDid)
	if l > 0 {
		n += 1 + l + sovIssuer(uint64(l))
	}
	if m.Paused {
		n += 2
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
		case 2:
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
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuerDid", wireType)
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
			m.IssuerDid = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Paused", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowIssuer
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Paused = bool(v != 0)
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
