// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: publicawesome/stargaze/cron/v1/proposal.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-proto"
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

// PromoteToPrivilegedContractProposal gov proposal content type to add
// "privileges" to a contract
type PromoteToPrivilegedContractProposal struct {
	// Title is a short summary
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	// Description is a human readable text
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	// Contract is the bech32 address of the smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
}

func (m *PromoteToPrivilegedContractProposal) Reset()         { *m = PromoteToPrivilegedContractProposal{} }
func (m *PromoteToPrivilegedContractProposal) String() string { return proto.CompactTextString(m) }
func (*PromoteToPrivilegedContractProposal) ProtoMessage()    {}
func (*PromoteToPrivilegedContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17d2b53c25d70a5, []int{0}
}
func (m *PromoteToPrivilegedContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *PromoteToPrivilegedContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_PromoteToPrivilegedContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *PromoteToPrivilegedContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PromoteToPrivilegedContractProposal.Merge(m, src)
}
func (m *PromoteToPrivilegedContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *PromoteToPrivilegedContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_PromoteToPrivilegedContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_PromoteToPrivilegedContractProposal proto.InternalMessageInfo

func (m *PromoteToPrivilegedContractProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *PromoteToPrivilegedContractProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *PromoteToPrivilegedContractProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

// DemotePrivilegedContractProposal gov proposal content type to remove
// "privileges" from a contract
type DemotePrivilegedContractProposal struct {
	// Title is a short summary
	Title string `protobuf:"bytes,1,opt,name=title,proto3" json:"title,omitempty" yaml:"title"`
	// Description is a human readable text
	Description string `protobuf:"bytes,2,opt,name=description,proto3" json:"description,omitempty" yaml:"description"`
	// Contract is the bech32 address of the smart contract
	Contract string `protobuf:"bytes,3,opt,name=contract,proto3" json:"contract,omitempty" yaml:"contract"`
}

func (m *DemotePrivilegedContractProposal) Reset()         { *m = DemotePrivilegedContractProposal{} }
func (m *DemotePrivilegedContractProposal) String() string { return proto.CompactTextString(m) }
func (*DemotePrivilegedContractProposal) ProtoMessage()    {}
func (*DemotePrivilegedContractProposal) Descriptor() ([]byte, []int) {
	return fileDescriptor_d17d2b53c25d70a5, []int{1}
}
func (m *DemotePrivilegedContractProposal) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DemotePrivilegedContractProposal) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DemotePrivilegedContractProposal.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DemotePrivilegedContractProposal) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DemotePrivilegedContractProposal.Merge(m, src)
}
func (m *DemotePrivilegedContractProposal) XXX_Size() int {
	return m.Size()
}
func (m *DemotePrivilegedContractProposal) XXX_DiscardUnknown() {
	xxx_messageInfo_DemotePrivilegedContractProposal.DiscardUnknown(m)
}

var xxx_messageInfo_DemotePrivilegedContractProposal proto.InternalMessageInfo

func (m *DemotePrivilegedContractProposal) GetTitle() string {
	if m != nil {
		return m.Title
	}
	return ""
}

func (m *DemotePrivilegedContractProposal) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *DemotePrivilegedContractProposal) GetContract() string {
	if m != nil {
		return m.Contract
	}
	return ""
}

func init() {
	proto.RegisterType((*PromoteToPrivilegedContractProposal)(nil), "publicawesome.stargaze.cron.v1.PromoteToPrivilegedContractProposal")
	proto.RegisterType((*DemotePrivilegedContractProposal)(nil), "publicawesome.stargaze.cron.v1.DemotePrivilegedContractProposal")
}

func init() {
	proto.RegisterFile("publicawesome/stargaze/cron/v1/proposal.proto", fileDescriptor_d17d2b53c25d70a5)
}

var fileDescriptor_d17d2b53c25d70a5 = []byte{
	// 332 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x92, 0xc1, 0x4a, 0xc3, 0x30,
	0x18, 0xc7, 0x17, 0x45, 0xd1, 0x2a, 0x28, 0x55, 0x64, 0xee, 0x90, 0x8d, 0x0a, 0xe2, 0x65, 0x09,
	0x75, 0x17, 0xd9, 0x71, 0x7a, 0x15, 0xc6, 0xf0, 0xe4, 0x45, 0xd2, 0x2c, 0xd4, 0x40, 0xdb, 0xaf,
	0x24, 0x59, 0x74, 0x3e, 0x85, 0x0f, 0xe3, 0x43, 0x88, 0xa7, 0x5d, 0x84, 0x9d, 0x86, 0x6c, 0x6f,
	0xb0, 0x27, 0x90, 0x36, 0x9d, 0x4c, 0x7c, 0x03, 0x6f, 0x49, 0x7e, 0xbf, 0x8f, 0x2f, 0x7f, 0xf8,
	0x7b, 0xed, 0x7c, 0x14, 0x25, 0x92, 0xb3, 0x27, 0xa1, 0x21, 0x15, 0x54, 0x1b, 0xa6, 0x62, 0xf6,
	0x22, 0x28, 0x57, 0x90, 0x51, 0x1b, 0xd2, 0x5c, 0x41, 0x0e, 0x9a, 0x25, 0x24, 0x57, 0x60, 0xc0,
	0xc7, 0xbf, 0x74, 0xb2, 0xd2, 0x49, 0xa1, 0x13, 0x1b, 0x36, 0x8e, 0x63, 0x88, 0xa1, 0x54, 0x69,
	0x71, 0x72, 0x53, 0x8d, 0x53, 0x0e, 0x3a, 0x05, 0xfd, 0xe0, 0x80, 0xbb, 0x38, 0x14, 0x4c, 0x91,
	0x77, 0xd6, 0x57, 0x90, 0x82, 0x11, 0x77, 0xd0, 0x57, 0xd2, 0xca, 0x44, 0xc4, 0x62, 0x78, 0x0d,
	0x99, 0x51, 0x8c, 0x9b, 0x7e, 0xb5, 0xde, 0x3f, 0xf7, 0xb6, 0x8c, 0x34, 0x89, 0xa8, 0xa3, 0x16,
	0xba, 0xd8, 0xed, 0x1d, 0x2e, 0x67, 0xcd, 0xfd, 0x31, 0x4b, 0x93, 0x6e, 0x50, 0x3e, 0x07, 0x03,
	0x87, 0xfd, 0x2b, 0x6f, 0x6f, 0x28, 0x34, 0x57, 0x32, 0x37, 0x12, 0xb2, 0xfa, 0x46, 0x69, 0x9f,
	0x2c, 0x67, 0x4d, 0xdf, 0xd9, 0x6b, 0x30, 0x18, 0xac, 0xab, 0x3e, 0xf5, 0x76, 0x78, 0xb5, 0xb5,
	0xbe, 0x59, 0x8e, 0x1d, 0x2d, 0x67, 0xcd, 0x03, 0x37, 0xb6, 0x22, 0xc1, 0xe0, 0x47, 0xea, 0xe2,
	0x8f, 0xb7, 0x76, 0xa3, 0x0a, 0x13, 0x83, 0x25, 0x36, 0x8c, 0x84, 0x61, 0x21, 0x29, 0xfe, 0x2e,
	0x32, 0x13, 0x7c, 0x22, 0xaf, 0x75, 0x23, 0x8a, 0x64, 0xff, 0x2a, 0x57, 0xef, 0xf6, 0x7d, 0x8e,
	0xd1, 0x64, 0x8e, 0xd1, 0xd7, 0x1c, 0xa3, 0xd7, 0x05, 0xae, 0x4d, 0x16, 0xb8, 0x36, 0x5d, 0xe0,
	0xda, 0x7d, 0x27, 0x96, 0xe6, 0x71, 0x14, 0x11, 0x0e, 0x29, 0x75, 0x45, 0x69, 0xff, 0x29, 0x96,
	0x0d, 0x2f, 0xe9, 0xb3, 0xab, 0x97, 0x19, 0xe7, 0x42, 0x47, 0xdb, 0x65, 0x11, 0x3a, 0xdf, 0x01,
	0x00, 0x00, 0xff, 0xff, 0xcd, 0xad, 0x2a, 0xc1, 0x8a, 0x02, 0x00, 0x00,
}

func (m *PromoteToPrivilegedContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *PromoteToPrivilegedContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *PromoteToPrivilegedContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *DemotePrivilegedContractProposal) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DemotePrivilegedContractProposal) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DemotePrivilegedContractProposal) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Contract) > 0 {
		i -= len(m.Contract)
		copy(dAtA[i:], m.Contract)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Contract)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.Title) > 0 {
		i -= len(m.Title)
		copy(dAtA[i:], m.Title)
		i = encodeVarintProposal(dAtA, i, uint64(len(m.Title)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintProposal(dAtA []byte, offset int, v uint64) int {
	offset -= sovProposal(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *PromoteToPrivilegedContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func (m *DemotePrivilegedContractProposal) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Title)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	l = len(m.Contract)
	if l > 0 {
		n += 1 + l + sovProposal(uint64(l))
	}
	return n
}

func sovProposal(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozProposal(x uint64) (n int) {
	return sovProposal(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *PromoteToPrivilegedContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: PromoteToPrivilegedContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: PromoteToPrivilegedContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func (m *DemotePrivilegedContractProposal) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowProposal
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
			return fmt.Errorf("proto: DemotePrivilegedContractProposal: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DemotePrivilegedContractProposal: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Title", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Title = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Contract", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowProposal
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
				return ErrInvalidLengthProposal
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthProposal
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Contract = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipProposal(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthProposal
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
func skipProposal(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
					return 0, ErrIntOverflowProposal
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
				return 0, ErrInvalidLengthProposal
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupProposal
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthProposal
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthProposal        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowProposal          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupProposal = fmt.Errorf("proto: unexpected end of group")
)
