// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: verifiable-credential-service/verifiable-credential.proto

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

type VerifiableCredential struct {
	Context           []string           `protobuf:"bytes,1,rep,name=context,proto3" json:"context,omitempty"`
	Id                string             `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Type              []string           `protobuf:"bytes,3,rep,name=type,proto3" json:"type,omitempty"`
	Issuer            string             `protobuf:"bytes,4,opt,name=issuer,proto3" json:"issuer,omitempty"`
	IssuanceDate      string             `protobuf:"bytes,5,opt,name=issuanceDate,proto3" json:"issuanceDate,omitempty"`
	CredentialSubject *CredentialSubject `protobuf:"bytes,6,opt,name=credential_subject,json=credentialSubject,proto3" json:"credential_subject,omitempty"`
	Proof             *Proof             `protobuf:"bytes,7,opt,name=proof,proto3" json:"proof,omitempty"`
}

func (m *VerifiableCredential) Reset()         { *m = VerifiableCredential{} }
func (m *VerifiableCredential) String() string { return proto.CompactTextString(m) }
func (*VerifiableCredential) ProtoMessage()    {}
func (*VerifiableCredential) Descriptor() ([]byte, []int) {
	return fileDescriptor_7396f4d0b066eb58, []int{0}
}
func (m *VerifiableCredential) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VerifiableCredential) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_VerifiableCredential.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *VerifiableCredential) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VerifiableCredential.Merge(m, src)
}
func (m *VerifiableCredential) XXX_Size() int {
	return m.Size()
}
func (m *VerifiableCredential) XXX_DiscardUnknown() {
	xxx_messageInfo_VerifiableCredential.DiscardUnknown(m)
}

var xxx_messageInfo_VerifiableCredential proto.InternalMessageInfo

func (m *VerifiableCredential) GetContext() []string {
	if m != nil {
		return m.Context
	}
	return nil
}

func (m *VerifiableCredential) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *VerifiableCredential) GetType() []string {
	if m != nil {
		return m.Type
	}
	return nil
}

func (m *VerifiableCredential) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *VerifiableCredential) GetIssuanceDate() string {
	if m != nil {
		return m.IssuanceDate
	}
	return ""
}

func (m *VerifiableCredential) GetCredentialSubject() *CredentialSubject {
	if m != nil {
		return m.CredentialSubject
	}
	return nil
}

func (m *VerifiableCredential) GetProof() *Proof {
	if m != nil {
		return m.Proof
	}
	return nil
}

type CredentialSubject struct {
	Id     string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	HasKyc bool   `protobuf:"varint,2,opt,name=has_kyc,json=hasKyc,proto3" json:"has_kyc,omitempty"`
}

func (m *CredentialSubject) Reset()         { *m = CredentialSubject{} }
func (m *CredentialSubject) String() string { return proto.CompactTextString(m) }
func (*CredentialSubject) ProtoMessage()    {}
func (*CredentialSubject) Descriptor() ([]byte, []int) {
	return fileDescriptor_7396f4d0b066eb58, []int{1}
}
func (m *CredentialSubject) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CredentialSubject) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CredentialSubject.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CredentialSubject) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CredentialSubject.Merge(m, src)
}
func (m *CredentialSubject) XXX_Size() int {
	return m.Size()
}
func (m *CredentialSubject) XXX_DiscardUnknown() {
	xxx_messageInfo_CredentialSubject.DiscardUnknown(m)
}

var xxx_messageInfo_CredentialSubject proto.InternalMessageInfo

func (m *CredentialSubject) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *CredentialSubject) GetHasKyc() bool {
	if m != nil {
		return m.HasKyc
	}
	return false
}

type Proof struct {
	Type               string `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Created            string `protobuf:"bytes,2,opt,name=created,proto3" json:"created,omitempty"`
	ProofPurpose       string `protobuf:"bytes,3,opt,name=proof_purpose,json=proofPurpose,proto3" json:"proof_purpose,omitempty"`
	VerificationMethod string `protobuf:"bytes,4,opt,name=verification_method,json=verificationMethod,proto3" json:"verification_method,omitempty"`
	Signature          string `protobuf:"bytes,5,opt,name=signature,proto3" json:"signature,omitempty"`
}

func (m *Proof) Reset()         { *m = Proof{} }
func (m *Proof) String() string { return proto.CompactTextString(m) }
func (*Proof) ProtoMessage()    {}
func (*Proof) Descriptor() ([]byte, []int) {
	return fileDescriptor_7396f4d0b066eb58, []int{2}
}
func (m *Proof) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Proof) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Proof.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Proof) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Proof.Merge(m, src)
}
func (m *Proof) XXX_Size() int {
	return m.Size()
}
func (m *Proof) XXX_DiscardUnknown() {
	xxx_messageInfo_Proof.DiscardUnknown(m)
}

var xxx_messageInfo_Proof proto.InternalMessageInfo

func (m *Proof) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Proof) GetCreated() string {
	if m != nil {
		return m.Created
	}
	return ""
}

func (m *Proof) GetProofPurpose() string {
	if m != nil {
		return m.ProofPurpose
	}
	return ""
}

func (m *Proof) GetVerificationMethod() string {
	if m != nil {
		return m.VerificationMethod
	}
	return ""
}

func (m *Proof) GetSignature() string {
	if m != nil {
		return m.Signature
	}
	return ""
}

func init() {
	proto.RegisterType((*VerifiableCredential)(nil), "allinbits.cosmoscash.verifiablecredentialservice.VerifiableCredential")
	proto.RegisterType((*CredentialSubject)(nil), "allinbits.cosmoscash.verifiablecredentialservice.CredentialSubject")
	proto.RegisterType((*Proof)(nil), "allinbits.cosmoscash.verifiablecredentialservice.Proof")
}

func init() {
	proto.RegisterFile("verifiable-credential-service/verifiable-credential.proto", fileDescriptor_7396f4d0b066eb58)
}

var fileDescriptor_7396f4d0b066eb58 = []byte{
	// 421 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0xbd, 0xae, 0xd3, 0x30,
	0x14, 0xc7, 0xeb, 0xb4, 0x4d, 0xa9, 0xf9, 0x90, 0x6a, 0x10, 0x78, 0x40, 0x51, 0x15, 0x96, 0x2e,
	0x4d, 0x10, 0x0c, 0x08, 0x89, 0x89, 0x22, 0x16, 0x54, 0xa9, 0x0a, 0x12, 0x03, 0x4b, 0x70, 0x1c,
	0xb7, 0x31, 0xa4, 0x71, 0x64, 0x3b, 0x55, 0xfb, 0x16, 0xbc, 0x06, 0x6f, 0xc2, 0x58, 0x89, 0x85,
	0x11, 0xb5, 0x2f, 0x82, 0xe2, 0x7c, 0xc1, 0xbd, 0x57, 0x57, 0xea, 0xe6, 0x73, 0xfe, 0x3e, 0xe7,
	0xf8, 0xfc, 0xfe, 0x86, 0xaf, 0x77, 0x4c, 0xf2, 0x35, 0x27, 0x51, 0xca, 0xe6, 0x54, 0xb2, 0x98,
	0x65, 0x9a, 0x93, 0x74, 0xae, 0x98, 0xdc, 0x71, 0xca, 0xfc, 0x1b, 0x55, 0x2f, 0x97, 0x42, 0x0b,
	0xf4, 0x9c, 0xa4, 0x29, 0xcf, 0x22, 0xae, 0x95, 0x47, 0x85, 0xda, 0x0a, 0x45, 0x89, 0x4a, 0xbc,
	0xae, 0xa2, 0x2b, 0xa8, 0xbb, 0xb9, 0xbf, 0x2c, 0xf8, 0xe8, 0x53, 0xab, 0x2f, 0x5a, 0x1d, 0x61,
	0x38, 0xa2, 0x22, 0xd3, 0x6c, 0xaf, 0x31, 0x98, 0xf6, 0x67, 0xe3, 0xa0, 0x09, 0xd1, 0x03, 0x68,
	0xf1, 0x18, 0x5b, 0x53, 0x30, 0x1b, 0x07, 0x16, 0x8f, 0x11, 0x82, 0x03, 0x7d, 0xc8, 0x19, 0xee,
	0x9b, 0x6b, 0xe6, 0x8c, 0x1e, 0x43, 0x9b, 0x2b, 0x55, 0x30, 0x89, 0x07, 0xe6, 0x5e, 0x1d, 0x21,
	0x17, 0xde, 0x2b, 0x4f, 0x24, 0xa3, 0xec, 0x1d, 0xd1, 0x0c, 0x0f, 0x8d, 0xfa, 0x5f, 0x0e, 0x49,
	0x88, 0xba, 0x77, 0x86, 0xaa, 0x88, 0xbe, 0x32, 0xaa, 0xb1, 0x3d, 0x05, 0xb3, 0xbb, 0x2f, 0x16,
	0xde, 0xa5, 0x1b, 0x7a, 0xdd, 0x4e, 0x1f, 0xab, 0x56, 0xc1, 0x84, 0x5e, 0x4d, 0xa1, 0x25, 0x1c,
	0xe6, 0x52, 0x88, 0x35, 0x1e, 0x99, 0x31, 0xaf, 0x2e, 0x1f, 0xb3, 0x2a, 0xcb, 0x83, 0xaa, 0x8b,
	0xfb, 0x06, 0x4e, 0xae, 0x8d, 0xad, 0xb9, 0x81, 0x96, 0xdb, 0x13, 0x38, 0x4a, 0x88, 0x0a, 0xbf,
	0x1d, 0xa8, 0x81, 0x79, 0x27, 0xb0, 0x13, 0xa2, 0x3e, 0x1c, 0xa8, 0xfb, 0x03, 0xc0, 0xa1, 0x69,
	0xd7, 0xa2, 0xad, 0x8a, 0x2a, 0xb4, 0xa5, 0x31, 0x92, 0x11, 0xcd, 0x1a, 0x0f, 0x9a, 0x10, 0x3d,
	0x83, 0xf7, 0xcd, 0xf8, 0x30, 0x2f, 0x64, 0x2e, 0x54, 0xe9, 0x88, 0xa1, 0x6b, 0x92, 0xab, 0x2a,
	0x87, 0x7c, 0xf8, 0xb0, 0x5a, 0x83, 0x12, 0xcd, 0x45, 0x16, 0x6e, 0x99, 0x4e, 0x44, 0x5c, 0xdb,
	0x84, 0xfe, 0x95, 0x96, 0x46, 0x41, 0x4f, 0xe1, 0x58, 0xf1, 0x4d, 0x46, 0x74, 0x21, 0x1b, 0xbf,
	0xba, 0xc4, 0xdb, 0x2f, 0x3f, 0x4f, 0x0e, 0x38, 0x9e, 0x1c, 0xf0, 0xe7, 0xe4, 0x80, 0xef, 0x67,
	0xa7, 0x77, 0x3c, 0x3b, 0xbd, 0xdf, 0x67, 0xa7, 0xf7, 0xf9, 0xfd, 0x86, 0xeb, 0xa4, 0x88, 0x3c,
	0x2a, 0xb6, 0x7e, 0x4b, 0xd3, 0xaf, 0x68, 0xce, 0x4b, 0x9c, 0xfe, 0xde, 0xbf, 0xfd, 0xa7, 0x97,
	0xeb, 0xaa, 0xc8, 0x36, 0x5f, 0xfb, 0xe5, 0xdf, 0x00, 0x00, 0x00, 0xff, 0xff, 0xf5, 0x2b, 0xe7,
	0x8d, 0x17, 0x03, 0x00, 0x00,
}

func (m *VerifiableCredential) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VerifiableCredential) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VerifiableCredential) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Proof != nil {
		{
			size, err := m.Proof.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVerifiableCredential(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x3a
	}
	if m.CredentialSubject != nil {
		{
			size, err := m.CredentialSubject.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintVerifiableCredential(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if len(m.IssuanceDate) > 0 {
		i -= len(m.IssuanceDate)
		copy(dAtA[i:], m.IssuanceDate)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.IssuanceDate)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.Issuer) > 0 {
		i -= len(m.Issuer)
		copy(dAtA[i:], m.Issuer)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Issuer)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.Type) > 0 {
		for iNdEx := len(m.Type) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Type[iNdEx])
			copy(dAtA[i:], m.Type[iNdEx])
			i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Type[iNdEx])))
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Context) > 0 {
		for iNdEx := len(m.Context) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Context[iNdEx])
			copy(dAtA[i:], m.Context[iNdEx])
			i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Context[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CredentialSubject) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CredentialSubject) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CredentialSubject) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.HasKyc {
		i--
		if m.HasKyc {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x10
	}
	if len(m.Id) > 0 {
		i -= len(m.Id)
		copy(dAtA[i:], m.Id)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Id)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *Proof) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Proof) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Proof) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Signature) > 0 {
		i -= len(m.Signature)
		copy(dAtA[i:], m.Signature)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Signature)))
		i--
		dAtA[i] = 0x2a
	}
	if len(m.VerificationMethod) > 0 {
		i -= len(m.VerificationMethod)
		copy(dAtA[i:], m.VerificationMethod)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.VerificationMethod)))
		i--
		dAtA[i] = 0x22
	}
	if len(m.ProofPurpose) > 0 {
		i -= len(m.ProofPurpose)
		copy(dAtA[i:], m.ProofPurpose)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.ProofPurpose)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Created) > 0 {
		i -= len(m.Created)
		copy(dAtA[i:], m.Created)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Created)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Type) > 0 {
		i -= len(m.Type)
		copy(dAtA[i:], m.Type)
		i = encodeVarintVerifiableCredential(dAtA, i, uint64(len(m.Type)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintVerifiableCredential(dAtA []byte, offset int, v uint64) int {
	offset -= sovVerifiableCredential(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VerifiableCredential) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Context) > 0 {
		for _, s := range m.Context {
			l = len(s)
			n += 1 + l + sovVerifiableCredential(uint64(l))
		}
	}
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	if len(m.Type) > 0 {
		for _, s := range m.Type {
			l = len(s)
			n += 1 + l + sovVerifiableCredential(uint64(l))
		}
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	l = len(m.IssuanceDate)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	if m.CredentialSubject != nil {
		l = m.CredentialSubject.Size()
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	if m.Proof != nil {
		l = m.Proof.Size()
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	return n
}

func (m *CredentialSubject) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Id)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	if m.HasKyc {
		n += 2
	}
	return n
}

func (m *Proof) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Type)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	l = len(m.Created)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	l = len(m.ProofPurpose)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	l = len(m.VerificationMethod)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	l = len(m.Signature)
	if l > 0 {
		n += 1 + l + sovVerifiableCredential(uint64(l))
	}
	return n
}

func sovVerifiableCredential(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozVerifiableCredential(x uint64) (n int) {
	return sovVerifiableCredential(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *VerifiableCredential) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifiableCredential
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
			return fmt.Errorf("proto: VerifiableCredential: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VerifiableCredential: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Context", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Context = append(m.Context, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = append(m.Type, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuanceDate", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.IssuanceDate = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field CredentialSubject", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.CredentialSubject == nil {
				m.CredentialSubject = &CredentialSubject{}
			}
			if err := m.CredentialSubject.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Proof", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Proof == nil {
				m.Proof = &Proof{}
			}
			if err := m.Proof.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerifiableCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerifiableCredential
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
func (m *CredentialSubject) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifiableCredential
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
			return fmt.Errorf("proto: CredentialSubject: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CredentialSubject: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Id", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Id = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field HasKyc", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
			m.HasKyc = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipVerifiableCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerifiableCredential
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
func (m *Proof) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowVerifiableCredential
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
			return fmt.Errorf("proto: Proof: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Proof: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Type", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Type = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Created", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Created = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ProofPurpose", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ProofPurpose = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field VerificationMethod", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.VerificationMethod = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Signature", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowVerifiableCredential
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
				return ErrInvalidLengthVerifiableCredential
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthVerifiableCredential
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Signature = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipVerifiableCredential(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthVerifiableCredential
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
func skipVerifiableCredential(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowVerifiableCredential
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
					return 0, ErrIntOverflowVerifiableCredential
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
					return 0, ErrIntOverflowVerifiableCredential
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
				return 0, ErrInvalidLengthVerifiableCredential
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupVerifiableCredential
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthVerifiableCredential
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthVerifiableCredential        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowVerifiableCredential          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupVerifiableCredential = fmt.Errorf("proto: unexpected end of group")
)