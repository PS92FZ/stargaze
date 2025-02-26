// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/globalfee/v1/globalfee.proto

package types

import (
	fmt "fmt"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/cosmos/gogoproto/gogoproto"
	proto "github.com/cosmos/gogoproto/proto"
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

// Params holds parameters for the globalfee module.
type Params struct {
	// Addresses which are whitelisted to modify the gas free operations
	PrivilegedAddresses []string `protobuf:"bytes,1,rep,name=privileged_addresses,json=privilegedAddresses,proto3" json:"privileged_addresses,omitempty"`
	// Minimum stores the minimum gas price(s) for all TX on the chain.
	MinimumGasPrices github_com_cosmos_cosmos_sdk_types.DecCoins `protobuf:"bytes,2,rep,name=minimum_gas_prices,json=minimumGasPrices,proto3,castrepeated=github.com/cosmos/cosmos-sdk/types.DecCoins" json:"minimum_gas_prices,omitempty" yaml:"minimum_gas_prices"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b74e65c141ff46e, []int{0}
}
func (m *Params) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Params) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Params.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Params) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Params.Merge(m, src)
}
func (m *Params) XXX_Size() int {
	return m.Size()
}
func (m *Params) XXX_DiscardUnknown() {
	xxx_messageInfo_Params.DiscardUnknown(m)
}

var xxx_messageInfo_Params proto.InternalMessageInfo

func (m *Params) GetPrivilegedAddresses() []string {
	if m != nil {
		return m.PrivilegedAddresses
	}
	return nil
}

func (m *Params) GetMinimumGasPrices() github_com_cosmos_cosmos_sdk_types.DecCoins {
	if m != nil {
		return m.MinimumGasPrices
	}
	return nil
}

// Configuration for code Ids which can have zero gas operations
type CodeAuthorization struct {
	// authorized code ids
	CodeID uint64 `protobuf:"varint,1,opt,name=code_id,json=codeId,proto3" json:"code_id,omitempty" yaml:"code_id"`
	// authorized contract operation methods
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty" yaml:"methods"`
}

func (m *CodeAuthorization) Reset()         { *m = CodeAuthorization{} }
func (m *CodeAuthorization) String() string { return proto.CompactTextString(m) }
func (*CodeAuthorization) ProtoMessage()    {}
func (*CodeAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b74e65c141ff46e, []int{1}
}
func (m *CodeAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *CodeAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_CodeAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *CodeAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CodeAuthorization.Merge(m, src)
}
func (m *CodeAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *CodeAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_CodeAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_CodeAuthorization proto.InternalMessageInfo

func (m *CodeAuthorization) GetCodeID() uint64 {
	if m != nil {
		return m.CodeID
	}
	return 0
}

func (m *CodeAuthorization) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

// Configuration for contract addresses which can have zero gas operations
type ContractAuthorization struct {
	// authorized contract addresses
	ContractAddress string `protobuf:"bytes,1,opt,name=contract_address,json=contractAddress,proto3" json:"contract_address,omitempty" yaml:"contract_address"`
	// authorized contract operation methods
	Methods []string `protobuf:"bytes,2,rep,name=methods,proto3" json:"methods,omitempty" yaml:"methods"`
}

func (m *ContractAuthorization) Reset()         { *m = ContractAuthorization{} }
func (m *ContractAuthorization) String() string { return proto.CompactTextString(m) }
func (*ContractAuthorization) ProtoMessage()    {}
func (*ContractAuthorization) Descriptor() ([]byte, []int) {
	return fileDescriptor_5b74e65c141ff46e, []int{2}
}
func (m *ContractAuthorization) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ContractAuthorization) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ContractAuthorization.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ContractAuthorization) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContractAuthorization.Merge(m, src)
}
func (m *ContractAuthorization) XXX_Size() int {
	return m.Size()
}
func (m *ContractAuthorization) XXX_DiscardUnknown() {
	xxx_messageInfo_ContractAuthorization.DiscardUnknown(m)
}

var xxx_messageInfo_ContractAuthorization proto.InternalMessageInfo

func (m *ContractAuthorization) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

func (m *ContractAuthorization) GetMethods() []string {
	if m != nil {
		return m.Methods
	}
	return nil
}

func init() {
	proto.RegisterType((*Params)(nil), "publicawesome.stargaze.globalfee.v1.Params")
	proto.RegisterType((*CodeAuthorization)(nil), "publicawesome.stargaze.globalfee.v1.CodeAuthorization")
	proto.RegisterType((*ContractAuthorization)(nil), "publicawesome.stargaze.globalfee.v1.ContractAuthorization")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/globalfee/v1/globalfee.proto", fileDescriptor_5b74e65c141ff46e)
}

var fileDescriptor_5b74e65c141ff46e = []byte{
	// 471 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x92, 0x3f, 0x6f, 0xd3, 0x40,
	0x18, 0xc6, 0xed, 0xb6, 0x4a, 0xd5, 0x43, 0x82, 0x62, 0x8a, 0x08, 0x25, 0xb2, 0x23, 0xb3, 0x44,
	0xa2, 0xf5, 0xc9, 0xad, 0x90, 0x50, 0xb7, 0x3a, 0x15, 0xa8, 0x5b, 0xe5, 0x91, 0x25, 0x3a, 0xdf,
	0x1d, 0xce, 0x09, 0x9f, 0x5f, 0xcb, 0x77, 0x09, 0x4d, 0x27, 0xbe, 0x00, 0x12, 0x23, 0x23, 0x33,
	0x9f, 0x81, 0x0f, 0xd0, 0xb1, 0x23, 0x93, 0x41, 0xc9, 0xc6, 0x98, 0x4f, 0x80, 0xfc, 0x4f, 0x69,
	0x94, 0x89, 0xc9, 0x67, 0x3f, 0xbf, 0xf7, 0x7d, 0x5e, 0xbf, 0xf7, 0xa0, 0xd3, 0x6c, 0x12, 0x25,
	0x82, 0x92, 0x4f, 0x5c, 0x81, 0xe4, 0x58, 0x69, 0x92, 0xc7, 0xe4, 0x86, 0xe3, 0x38, 0x81, 0x88,
	0x24, 0x1f, 0x38, 0xc7, 0x53, 0x7f, 0xf5, 0xe2, 0x65, 0x39, 0x68, 0xb0, 0x5e, 0xae, 0x15, 0x79,
	0x6d, 0x91, 0xb7, 0xe2, 0xa6, 0xfe, 0xe1, 0x41, 0x0c, 0x31, 0x54, 0x3c, 0x2e, 0x4f, 0x75, 0xe9,
	0xa1, 0x4d, 0x41, 0x49, 0x50, 0x38, 0x22, 0xaa, 0x6c, 0x1d, 0x71, 0x4d, 0x7c, 0x4c, 0x41, 0xa4,
	0xb5, 0xee, 0x7e, 0xde, 0x42, 0x9d, 0x2b, 0x92, 0x13, 0xa9, 0x2c, 0x1f, 0x1d, 0x64, 0xb9, 0x98,
	0x8a, 0x84, 0xc7, 0x9c, 0x8d, 0x08, 0x63, 0x39, 0x57, 0x8a, 0xab, 0xae, 0xd9, 0xdf, 0x1e, 0xec,
	0x85, 0x4f, 0x56, 0xda, 0x79, 0x2b, 0x59, 0x3f, 0x4d, 0x64, 0x49, 0x91, 0x0a, 0x39, 0x91, 0xa3,
	0x98, 0xa8, 0x51, 0x96, 0x0b, 0xca, 0x55, 0x77, 0xab, 0xbf, 0x3d, 0x78, 0x70, 0xd2, 0xf3, 0x6a,
	0x6f, 0xaf, 0xf4, 0xf6, 0x1a, 0x6f, 0xef, 0x82, 0xd3, 0x21, 0x88, 0x34, 0xc8, 0x6e, 0x0b, 0xc7,
	0xf8, 0x5b, 0x38, 0xbd, 0xcd, 0xfa, 0x23, 0x90, 0x42, 0x73, 0x99, 0xe9, 0xd9, 0xb2, 0x70, 0x9e,
	0xcf, 0x88, 0x4c, 0xce, 0xdc, 0x4d, 0xca, 0xfd, 0xf1, 0xdb, 0x79, 0x15, 0x0b, 0x3d, 0x9e, 0x44,
	0x1e, 0x05, 0x89, 0x9b, 0x1f, 0xad, 0x1f, 0xc7, 0x8a, 0x7d, 0xc4, 0x7a, 0x96, 0x71, 0xd5, 0x1a,
	0xaa, 0x70, 0xbf, 0xe9, 0xf1, 0x8e, 0xa8, 0xab, 0xaa, 0xc3, 0xd9, 0xce, 0xb7, 0xef, 0x8e, 0xe1,
	0x5e, 0xa3, 0xc7, 0x43, 0x60, 0xfc, 0x7c, 0xa2, 0xc7, 0x90, 0x8b, 0x1b, 0xa2, 0x05, 0xa4, 0xd6,
	0x6b, 0xb4, 0x4b, 0x81, 0xf1, 0x91, 0x60, 0x5d, 0xb3, 0x6f, 0x0e, 0x76, 0x82, 0xde, 0xbc, 0x70,
	0x3a, 0x25, 0x77, 0x79, 0xb1, 0x2c, 0x9c, 0x87, 0xf5, 0x54, 0x0d, 0xe2, 0x86, 0x9d, 0xf2, 0x74,
	0xc9, 0xac, 0x23, 0xb4, 0x2b, 0xb9, 0x1e, 0x03, 0xab, 0x97, 0xb0, 0x17, 0x58, 0x2b, 0xb8, 0x11,
	0xdc, 0xb0, 0x45, 0xdc, 0x2f, 0x26, 0x7a, 0x3a, 0x84, 0x54, 0xe7, 0x84, 0xea, 0x75, 0xfb, 0xb7,
	0x68, 0x9f, 0x36, 0x42, 0x7b, 0x13, 0xd5, 0x1c, 0x7b, 0xc1, 0x8b, 0x65, 0xe1, 0x3c, 0x6b, 0xdd,
	0xd7, 0x09, 0x37, 0x7c, 0xd4, 0x7e, 0x6a, 0xae, 0xe8, 0xff, 0xe6, 0x09, 0xc2, 0xdb, 0xb9, 0x6d,
	0xde, 0xcd, 0x6d, 0xf3, 0xcf, 0xdc, 0x36, 0xbf, 0x2e, 0x6c, 0xe3, 0x6e, 0x61, 0x1b, 0xbf, 0x16,
	0xb6, 0xf1, 0xfe, 0xcd, 0xbd, 0x45, 0xd7, 0x61, 0x3c, 0xde, 0x88, 0xf0, 0xd4, 0x3f, 0xc1, 0xd7,
	0xf7, 0x82, 0x5c, 0xad, 0x3f, 0xea, 0x54, 0x39, 0x3b, 0xfd, 0x17, 0x00, 0x00, 0xff, 0xff, 0x4b,
	0xfa, 0xa4, 0x42, 0xf9, 0x02, 0x00, 0x00,
}

func (m *Params) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Params) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Params) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.MinimumGasPrices) > 0 {
		for iNdEx := len(m.MinimumGasPrices) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.MinimumGasPrices[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGlobalfee(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.PrivilegedAddresses) > 0 {
		for iNdEx := len(m.PrivilegedAddresses) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.PrivilegedAddresses[iNdEx])
			copy(dAtA[i:], m.PrivilegedAddresses[iNdEx])
			i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.PrivilegedAddresses[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func (m *CodeAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *CodeAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *CodeAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Methods[iNdEx])
			copy(dAtA[i:], m.Methods[iNdEx])
			i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.Methods[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if m.CodeID != 0 {
		i = encodeVarintGlobalfee(dAtA, i, uint64(m.CodeID))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ContractAuthorization) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ContractAuthorization) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ContractAuthorization) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Methods) > 0 {
		for iNdEx := len(m.Methods) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Methods[iNdEx])
			copy(dAtA[i:], m.Methods[iNdEx])
			i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.Methods[iNdEx])))
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.ContractAddress) > 0 {
		i -= len(m.ContractAddress)
		copy(dAtA[i:], m.ContractAddress)
		i = encodeVarintGlobalfee(dAtA, i, uint64(len(m.ContractAddress)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintGlobalfee(dAtA []byte, offset int, v uint64) int {
	offset -= sovGlobalfee(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.PrivilegedAddresses) > 0 {
		for _, s := range m.PrivilegedAddresses {
			l = len(s)
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	if len(m.MinimumGasPrices) > 0 {
		for _, e := range m.MinimumGasPrices {
			l = e.Size()
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	return n
}

func (m *CodeAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.CodeID != 0 {
		n += 1 + sovGlobalfee(uint64(m.CodeID))
	}
	if len(m.Methods) > 0 {
		for _, s := range m.Methods {
			l = len(s)
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	return n
}

func (m *ContractAuthorization) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ContractAddress)
	if l > 0 {
		n += 1 + l + sovGlobalfee(uint64(l))
	}
	if len(m.Methods) > 0 {
		for _, s := range m.Methods {
			l = len(s)
			n += 1 + l + sovGlobalfee(uint64(l))
		}
	}
	return n
}

func sovGlobalfee(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGlobalfee(x uint64) (n int) {
	return sovGlobalfee(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalfee
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
			return fmt.Errorf("proto: Params: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Params: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PrivilegedAddresses", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PrivilegedAddresses = append(m.PrivilegedAddresses, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field MinimumGasPrices", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MinimumGasPrices = append(m.MinimumGasPrices, types.DecCoin{})
			if err := m.MinimumGasPrices[len(m.MinimumGasPrices)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalfee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalfee
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
func (m *CodeAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalfee
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
			return fmt.Errorf("proto: CodeAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: CodeAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field CodeID", wireType)
			}
			m.CodeID = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.CodeID |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Methods = append(m.Methods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalfee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalfee
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
func (m *ContractAuthorization) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGlobalfee
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
			return fmt.Errorf("proto: ContractAuthorization: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ContractAuthorization: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ContractAddress", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ContractAddress = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Methods", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGlobalfee
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
				return ErrInvalidLengthGlobalfee
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGlobalfee
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Methods = append(m.Methods, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGlobalfee(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGlobalfee
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
func skipGlobalfee(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGlobalfee
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
					return 0, ErrIntOverflowGlobalfee
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
					return 0, ErrIntOverflowGlobalfee
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
				return 0, ErrInvalidLengthGlobalfee
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGlobalfee
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGlobalfee
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGlobalfee        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGlobalfee          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGlobalfee = fmt.Errorf("proto: unexpected end of group")
)
