// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: nft/v1beta1/genesis.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/gogoproto/gogoproto"
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

// GenesisState defines the nft module's genesis state.
type GenesisState struct {
	Collections []Collection  `protobuf:"bytes,1,rep,name=collections,proto3" json:"collections"`
	Orders      []MarketPlace `protobuf:"bytes,2,rep,name=orders,proto3" json:"orders"`
	Communities []Community   `protobuf:"bytes,3,rep,name=communities,proto3" json:"communities"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_52737c725dd1928d, []int{0}
}
func (m *GenesisState) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GenesisState) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GenesisState.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GenesisState) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GenesisState.Merge(m, src)
}
func (m *GenesisState) XXX_Size() int {
	return m.Size()
}
func (m *GenesisState) XXX_DiscardUnknown() {
	xxx_messageInfo_GenesisState.DiscardUnknown(m)
}

var xxx_messageInfo_GenesisState proto.InternalMessageInfo

func (m *GenesisState) GetCollections() []Collection {
	if m != nil {
		return m.Collections
	}
	return nil
}

func (m *GenesisState) GetOrders() []MarketPlace {
	if m != nil {
		return m.Orders
	}
	return nil
}

func (m *GenesisState) GetCommunities() []Community {
	if m != nil {
		return m.Communities
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nft.v1beta1.GenesisState")
}

func init() { proto.RegisterFile("nft/v1beta1/genesis.proto", fileDescriptor_52737c725dd1928d) }

var fileDescriptor_52737c725dd1928d = []byte{
	// 282 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xbd, 0x4a, 0xc4, 0x40,
	0x14, 0x85, 0x13, 0x57, 0xb6, 0x48, 0xac, 0x82, 0x3f, 0x71, 0x85, 0x51, 0xc4, 0xc2, 0x2a, 0x61,
	0x15, 0x2c, 0x57, 0x5c, 0x0b, 0x2b, 0x45, 0xb4, 0xb3, 0x91, 0x24, 0xde, 0x8d, 0x61, 0x93, 0xb9,
	0x21, 0x73, 0xa3, 0xe4, 0x2d, 0x7c, 0x2b, 0xb7, 0xdc, 0xd2, 0x4a, 0x24, 0x79, 0x11, 0xc9, 0xcc,
	0xb0, 0x8e, 0xdb, 0x0d, 0xf7, 0x9c, 0x8f, 0xf3, 0x31, 0xce, 0x3e, 0x9f, 0x51, 0xf8, 0x36, 0x8e,
	0x81, 0xa2, 0x71, 0x98, 0x02, 0x07, 0x91, 0x89, 0xa0, 0xac, 0x90, 0xd0, 0x73, 0xf9, 0x8c, 0x02,
	0x1d, 0x8d, 0xb6, 0x53, 0x4c, 0x51, 0xde, 0xc3, 0xfe, 0xa5, 0x2a, 0xa3, 0x1d, 0x93, 0xee, 0xeb,
	0xea, 0xcc, 0xcc, 0x73, 0x11, 0x55, 0x73, 0xa0, 0xe7, 0x32, 0x8f, 0x12, 0xd0, 0xf9, 0x81, 0x99,
	0x27, 0x58, 0x14, 0x35, 0xcf, 0xa8, 0x51, 0xe1, 0xf1, 0xa7, 0xed, 0x6c, 0xdd, 0x28, 0x91, 0x47,
	0x8a, 0x08, 0xbc, 0x4b, 0xc7, 0x4d, 0x30, 0xcf, 0x21, 0xa1, 0x0c, 0xb9, 0xf0, 0xed, 0xa3, 0xc1,
	0xa9, 0x7b, 0xb6, 0x17, 0x18, 0x76, 0xc1, 0xf5, 0x2a, 0x9f, 0x6e, 0x2e, 0xbe, 0x0f, 0xad, 0x07,
	0x93, 0xf0, 0x2e, 0x9c, 0x21, 0x56, 0x2f, 0x50, 0x09, 0x7f, 0x43, 0xb2, 0xfe, 0x3f, 0xf6, 0x56,
	0xfa, 0xdd, 0xf7, 0x7a, 0x1a, 0xd6, 0x6d, 0x6f, 0xd2, 0x0f, 0x2b, 0xb9, 0x0c, 0x84, 0x3f, 0x90,
	0xf0, 0xee, 0xda, 0xb0, 0x96, 0xff, 0xdb, 0x5d, 0x01, 0xd3, 0xc9, 0xa2, 0x65, 0xf6, 0xb2, 0x65,
	0xf6, 0x4f, 0xcb, 0xec, 0x8f, 0x8e, 0x59, 0xcb, 0x8e, 0x59, 0x5f, 0x1d, 0xb3, 0x9e, 0x4e, 0xd2,
	0x8c, 0x5e, 0xeb, 0x38, 0x48, 0xb0, 0x08, 0xaf, 0x6a, 0x42, 0x8e, 0x45, 0x73, 0x07, 0xf4, 0x8e,
	0xd5, 0xbc, 0xff, 0xc6, 0x90, 0x9a, 0x12, 0x44, 0x3c, 0x94, 0x1f, 0x72, 0xfe, 0x1b, 0x00, 0x00,
	0xff, 0xff, 0xd6, 0xda, 0xe9, 0x15, 0xa4, 0x01, 0x00, 0x00,
}

func (m *GenesisState) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GenesisState) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GenesisState) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Communities) > 0 {
		for iNdEx := len(m.Communities) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Communities[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x1a
		}
	}
	if len(m.Orders) > 0 {
		for iNdEx := len(m.Orders) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Orders[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x12
		}
	}
	if len(m.Collections) > 0 {
		for iNdEx := len(m.Collections) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.Collections[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenesis(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenesis(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GenesisState) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Collections) > 0 {
		for _, e := range m.Collections {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Orders) > 0 {
		for _, e := range m.Orders {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.Communities) > 0 {
		for _, e := range m.Communities {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	return n
}

func sovGenesis(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenesis(x uint64) (n int) {
	return sovGenesis(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GenesisState) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenesis
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
			return fmt.Errorf("proto: GenesisState: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GenesisState: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Collections", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Collections = append(m.Collections, Collection{})
			if err := m.Collections[len(m.Collections)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Orders", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Orders = append(m.Orders, MarketPlace{})
			if err := m.Orders[len(m.Orders)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Communities", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
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
				return ErrInvalidLengthGenesis
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenesis
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Communities = append(m.Communities, Community{})
			if err := m.Communities[len(m.Communities)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenesis(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthGenesis
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
func skipGenesis(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
					return 0, ErrIntOverflowGenesis
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
				return 0, ErrInvalidLengthGenesis
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGenesis
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGenesis
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGenesis        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenesis          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGenesis = fmt.Errorf("proto: unexpected end of group")
)
