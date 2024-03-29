// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carauction/genesis.proto

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

// GenesisState defines the carauction module's genesis state.
type GenesisState struct {
	Params         Params       `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	LotsList       []Lots       `protobuf:"bytes,2,rep,name=lotsList,proto3" json:"lotsList"`
	LotsCount      uint64       `protobuf:"varint,3,opt,name=lotsCount,proto3" json:"lotsCount,omitempty"`
	AssetsList     []Assets     `protobuf:"bytes,4,rep,name=assetsList,proto3" json:"assetsList"`
	AssetsCount    uint64       `protobuf:"varint,5,opt,name=assetsCount,proto3" json:"assetsCount,omitempty"`
	AuctionInfo    *AuctionInfo `protobuf:"bytes,6,opt,name=auctionInfo,proto3" json:"auctionInfo,omitempty"`
	LotsQueueList  []LotsQueue  `protobuf:"bytes,7,rep,name=lotsQueueList,proto3" json:"lotsQueueList"`
	LotsQueueCount uint64       `protobuf:"varint,8,opt,name=lotsQueueCount,proto3" json:"lotsQueueCount,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_7bb09a60c909d61b, []int{0}
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

func (m *GenesisState) GetParams() Params {
	if m != nil {
		return m.Params
	}
	return Params{}
}

func (m *GenesisState) GetLotsList() []Lots {
	if m != nil {
		return m.LotsList
	}
	return nil
}

func (m *GenesisState) GetLotsCount() uint64 {
	if m != nil {
		return m.LotsCount
	}
	return 0
}

func (m *GenesisState) GetAssetsList() []Assets {
	if m != nil {
		return m.AssetsList
	}
	return nil
}

func (m *GenesisState) GetAssetsCount() uint64 {
	if m != nil {
		return m.AssetsCount
	}
	return 0
}

func (m *GenesisState) GetAuctionInfo() *AuctionInfo {
	if m != nil {
		return m.AuctionInfo
	}
	return nil
}

func (m *GenesisState) GetLotsQueueList() []LotsQueue {
	if m != nil {
		return m.LotsQueueList
	}
	return nil
}

func (m *GenesisState) GetLotsQueueCount() uint64 {
	if m != nil {
		return m.LotsQueueCount
	}
	return 0
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "nasa70.carauction.carauction.GenesisState")
}

func init() { proto.RegisterFile("carauction/genesis.proto", fileDescriptor_7bb09a60c909d61b) }

var fileDescriptor_7bb09a60c909d61b = []byte{
	// 374 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0xc1, 0x4e, 0xc2, 0x40,
	0x10, 0x6d, 0x05, 0x11, 0xb7, 0xea, 0x61, 0xa3, 0xb1, 0x41, 0xac, 0x0d, 0x31, 0x8a, 0x07, 0x5b,
	0x82, 0x07, 0xcf, 0xa2, 0x89, 0x11, 0x39, 0x28, 0xdc, 0xbc, 0x90, 0xa5, 0x59, 0x6a, 0x13, 0xe9,
	0x62, 0x77, 0x9b, 0xe8, 0x5f, 0xf8, 0x59, 0x1c, 0x39, 0x7a, 0x32, 0x06, 0xee, 0x7e, 0x83, 0xe9,
	0x6c, 0x81, 0x95, 0x18, 0x7a, 0xea, 0x74, 0x66, 0xde, 0x7b, 0xf3, 0x66, 0x07, 0x99, 0x1e, 0x89,
	0x48, 0xec, 0x89, 0x80, 0x85, 0xae, 0x4f, 0x43, 0xca, 0x03, 0xee, 0x0c, 0x23, 0x26, 0x18, 0x2e,
	0x87, 0x84, 0x93, 0xcb, 0x9a, 0xb3, 0x68, 0x50, 0xc2, 0xd2, 0xae, 0xcf, 0x7c, 0x06, 0x8d, 0x6e,
	0x12, 0x49, 0x4c, 0x69, 0x5f, 0x61, 0x1b, 0x92, 0x88, 0x0c, 0x52, 0xb2, 0xd2, 0x9e, 0x52, 0x78,
	0x61, 0x82, 0xff, 0xd3, 0x4f, 0x38, 0xa7, 0xf3, 0xc2, 0xa1, 0x5a, 0x90, 0xdf, 0x6e, 0x10, 0xf6,
	0x67, 0x3a, 0x07, 0x4b, 0x74, 0xdd, 0xd7, 0x98, 0xc6, 0x54, 0x16, 0x2b, 0x3f, 0x39, 0xb4, 0x75,
	0x2b, 0xad, 0x74, 0x04, 0x11, 0x14, 0x37, 0x50, 0x41, 0x0e, 0x63, 0xea, 0xb6, 0x5e, 0x35, 0xea,
	0xc7, 0xce, 0x2a, 0x6b, 0xce, 0x03, 0xf4, 0x36, 0xf2, 0xa3, 0xaf, 0x23, 0xad, 0x9d, 0x22, 0xf1,
	0x0d, 0x2a, 0x26, 0x42, 0xad, 0x80, 0x0b, 0x73, 0xcd, 0xce, 0x55, 0x8d, 0x7a, 0x65, 0x35, 0x4b,
	0x8b, 0x89, 0x19, 0xc7, 0x1c, 0x89, 0xcb, 0x68, 0x33, 0x89, 0xaf, 0x59, 0x1c, 0x0a, 0x33, 0x67,
	0xeb, 0xd5, 0x7c, 0x7b, 0x91, 0xc0, 0x4d, 0x84, 0xe4, 0x12, 0x40, 0x25, 0x0f, 0x2a, 0x19, 0xb3,
	0x5e, 0x41, 0x7f, 0xaa, 0xa3, 0xa0, 0xb1, 0x8d, 0x0c, 0xf9, 0x27, 0xb5, 0xd6, 0x41, 0x4b, 0x4d,
	0xe1, 0x7b, 0x64, 0xa4, 0x2c, 0x77, 0x61, 0x9f, 0x99, 0x05, 0x58, 0xcd, 0x59, 0x86, 0xdc, 0x02,
	0xd0, 0x56, 0xd1, 0xb8, 0x83, 0xb6, 0x13, 0x1f, 0x8f, 0xc9, 0x33, 0xc0, 0xf4, 0x1b, 0x30, 0xfd,
	0x69, 0xf6, 0x8e, 0x00, 0x92, 0x1a, 0xf8, 0xcb, 0x81, 0x4f, 0xd0, 0xce, 0x3c, 0x21, 0x6d, 0x14,
	0xc1, 0xc6, 0x52, 0xb6, 0xd1, 0x1c, 0x4d, 0x2c, 0x7d, 0x3c, 0xb1, 0xf4, 0xef, 0x89, 0xa5, 0x7f,
	0x4c, 0x2d, 0x6d, 0x3c, 0xb5, 0xb4, 0xcf, 0xa9, 0xa5, 0x3d, 0xd5, 0xfc, 0x40, 0x3c, 0xc7, 0x3d,
	0xc7, 0x63, 0x03, 0x57, 0x4e, 0xe2, 0x7a, 0x24, 0x3a, 0x9f, 0x9d, 0xce, 0x9b, 0xab, 0xdc, 0x91,
	0x78, 0x1f, 0x52, 0xde, 0x2b, 0xc0, 0x0d, 0x5d, 0xfc, 0x06, 0x00, 0x00, 0xff, 0xff, 0x11, 0x16,
	0x9b, 0x78, 0x18, 0x03, 0x00, 0x00,
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
	if m.LotsQueueCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LotsQueueCount))
		i--
		dAtA[i] = 0x40
	}
	if len(m.LotsQueueList) > 0 {
		for iNdEx := len(m.LotsQueueList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LotsQueueList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x3a
		}
	}
	if m.AuctionInfo != nil {
		{
			size, err := m.AuctionInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenesis(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x32
	}
	if m.AssetsCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.AssetsCount))
		i--
		dAtA[i] = 0x28
	}
	if len(m.AssetsList) > 0 {
		for iNdEx := len(m.AssetsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.AssetsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x22
		}
	}
	if m.LotsCount != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.LotsCount))
		i--
		dAtA[i] = 0x18
	}
	if len(m.LotsList) > 0 {
		for iNdEx := len(m.LotsList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.LotsList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	{
		size, err := m.Params.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintGenesis(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
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
	l = m.Params.Size()
	n += 1 + l + sovGenesis(uint64(l))
	if len(m.LotsList) > 0 {
		for _, e := range m.LotsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LotsCount != 0 {
		n += 1 + sovGenesis(uint64(m.LotsCount))
	}
	if len(m.AssetsList) > 0 {
		for _, e := range m.AssetsList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.AssetsCount != 0 {
		n += 1 + sovGenesis(uint64(m.AssetsCount))
	}
	if m.AuctionInfo != nil {
		l = m.AuctionInfo.Size()
		n += 1 + l + sovGenesis(uint64(l))
	}
	if len(m.LotsQueueList) > 0 {
		for _, e := range m.LotsQueueList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.LotsQueueCount != 0 {
		n += 1 + sovGenesis(uint64(m.LotsQueueCount))
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
				return fmt.Errorf("proto: wrong wireType = %d for field Params", wireType)
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
			if err := m.Params.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotsList", wireType)
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
			m.LotsList = append(m.LotsList, Lots{})
			if err := m.LotsList[len(m.LotsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotsCount", wireType)
			}
			m.LotsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LotsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetsList", wireType)
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
			m.AssetsList = append(m.AssetsList, Assets{})
			if err := m.AssetsList[len(m.AssetsList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetsCount", wireType)
			}
			m.AssetsCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetsCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field AuctionInfo", wireType)
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
			if m.AuctionInfo == nil {
				m.AuctionInfo = &AuctionInfo{}
			}
			if err := m.AuctionInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotsQueueList", wireType)
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
			m.LotsQueueList = append(m.LotsQueueList, LotsQueue{})
			if err := m.LotsQueueList[len(m.LotsQueueList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 8:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotsQueueCount", wireType)
			}
			m.LotsQueueCount = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LotsQueueCount |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
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
