// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: carauction/assets.proto

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

type Assets struct {
	AssetId     uint64 `protobuf:"varint,1,opt,name=assetId,proto3" json:"assetId,omitempty"`
	Owner       string `protobuf:"bytes,2,opt,name=owner,proto3" json:"owner,omitempty"`
	Description string `protobuf:"bytes,3,opt,name=description,proto3" json:"description,omitempty"`
	LotId       uint64 `protobuf:"varint,4,opt,name=lotId,proto3" json:"lotId,omitempty"`
}

func (m *Assets) Reset()         { *m = Assets{} }
func (m *Assets) String() string { return proto.CompactTextString(m) }
func (*Assets) ProtoMessage()    {}
func (*Assets) Descriptor() ([]byte, []int) {
	return fileDescriptor_ec205576d4a8bd19, []int{0}
}
func (m *Assets) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Assets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Assets.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Assets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Assets.Merge(m, src)
}
func (m *Assets) XXX_Size() int {
	return m.Size()
}
func (m *Assets) XXX_DiscardUnknown() {
	xxx_messageInfo_Assets.DiscardUnknown(m)
}

var xxx_messageInfo_Assets proto.InternalMessageInfo

func (m *Assets) GetAssetId() uint64 {
	if m != nil {
		return m.AssetId
	}
	return 0
}

func (m *Assets) GetOwner() string {
	if m != nil {
		return m.Owner
	}
	return ""
}

func (m *Assets) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Assets) GetLotId() uint64 {
	if m != nil {
		return m.LotId
	}
	return 0
}

func init() {
	proto.RegisterType((*Assets)(nil), "nasa70.carauction.carauction.Assets")
}

func init() { proto.RegisterFile("carauction/assets.proto", fileDescriptor_ec205576d4a8bd19) }

var fileDescriptor_ec205576d4a8bd19 = []byte{
	// 198 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0x4f, 0x4e, 0x2c, 0x4a,
	0x2c, 0x4d, 0x2e, 0xc9, 0xcc, 0xcf, 0xd3, 0x4f, 0x2c, 0x2e, 0x4e, 0x2d, 0x29, 0xd6, 0x2b, 0x28,
	0xca, 0x2f, 0xc9, 0x17, 0x92, 0xc9, 0x4b, 0x2c, 0x4e, 0x34, 0x37, 0xd0, 0x43, 0xc8, 0x23, 0x31,
	0x95, 0x0a, 0xb8, 0xd8, 0x1c, 0xc1, 0xaa, 0x85, 0x24, 0xb8, 0xd8, 0xc1, 0xfa, 0x3c, 0x53, 0x24,
	0x18, 0x15, 0x18, 0x35, 0x58, 0x82, 0x60, 0x5c, 0x21, 0x11, 0x2e, 0xd6, 0xfc, 0xf2, 0xbc, 0xd4,
	0x22, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x08, 0x47, 0x48, 0x81, 0x8b, 0x3b, 0x25, 0xb5,
	0x38, 0xb9, 0x28, 0xb3, 0x00, 0x64, 0x90, 0x04, 0x33, 0x58, 0x0e, 0x59, 0x08, 0xa4, 0x2f, 0x27,
	0x1f, 0x64, 0x1e, 0x0b, 0xd8, 0x3c, 0x08, 0xc7, 0xc9, 0xeb, 0xc4, 0x23, 0x39, 0xc6, 0x0b, 0x8f,
	0xe4, 0x18, 0x1f, 0x3c, 0x92, 0x63, 0x9c, 0xf0, 0x58, 0x8e, 0xe1, 0xc2, 0x63, 0x39, 0x86, 0x1b,
	0x8f, 0xe5, 0x18, 0xa2, 0x0c, 0xd2, 0x33, 0x4b, 0x32, 0x4a, 0x93, 0xf4, 0x92, 0xf3, 0x73, 0xf5,
	0x21, 0x8e, 0xd6, 0x4f, 0x4e, 0x2c, 0xd2, 0x85, 0xf9, 0xaa, 0x42, 0x1f, 0xc9, 0x8b, 0x25, 0x95,
	0x05, 0xa9, 0xc5, 0x49, 0x6c, 0x60, 0x2f, 0x1a, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff, 0x9b, 0xf0,
	0x6d, 0x63, 0xfd, 0x00, 0x00, 0x00,
}

func (m *Assets) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Assets) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Assets) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.LotId != 0 {
		i = encodeVarintAssets(dAtA, i, uint64(m.LotId))
		i--
		dAtA[i] = 0x20
	}
	if len(m.Description) > 0 {
		i -= len(m.Description)
		copy(dAtA[i:], m.Description)
		i = encodeVarintAssets(dAtA, i, uint64(len(m.Description)))
		i--
		dAtA[i] = 0x1a
	}
	if len(m.Owner) > 0 {
		i -= len(m.Owner)
		copy(dAtA[i:], m.Owner)
		i = encodeVarintAssets(dAtA, i, uint64(len(m.Owner)))
		i--
		dAtA[i] = 0x12
	}
	if m.AssetId != 0 {
		i = encodeVarintAssets(dAtA, i, uint64(m.AssetId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintAssets(dAtA []byte, offset int, v uint64) int {
	offset -= sovAssets(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Assets) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.AssetId != 0 {
		n += 1 + sovAssets(uint64(m.AssetId))
	}
	l = len(m.Owner)
	if l > 0 {
		n += 1 + l + sovAssets(uint64(l))
	}
	l = len(m.Description)
	if l > 0 {
		n += 1 + l + sovAssets(uint64(l))
	}
	if m.LotId != 0 {
		n += 1 + sovAssets(uint64(m.LotId))
	}
	return n
}

func sovAssets(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozAssets(x uint64) (n int) {
	return sovAssets(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Assets) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowAssets
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
			return fmt.Errorf("proto: Assets: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Assets: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field AssetId", wireType)
			}
			m.AssetId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.AssetId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Owner", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssets
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
				return ErrInvalidLengthAssets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Owner = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Description", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssets
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
				return ErrInvalidLengthAssets
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthAssets
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Description = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LotId", wireType)
			}
			m.LotId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowAssets
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LotId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipAssets(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthAssets
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
func skipAssets(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowAssets
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
					return 0, ErrIntOverflowAssets
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
					return 0, ErrIntOverflowAssets
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
				return 0, ErrInvalidLengthAssets
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupAssets
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthAssets
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthAssets        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowAssets          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupAssets = fmt.Errorf("proto: unexpected end of group")
)
