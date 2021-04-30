// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: btcx/denom_cross_chain_info.proto

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

type DenomCrossChainInfo struct {
	DenomInfo   *DenomInfo `protobuf:"bytes,1,opt,name=denomInfo,proto3" json:"denomInfo,omitempty"`
	ToChainId   uint64     `protobuf:"varint,2,opt,name=toChainId,proto3" json:"toChainId,omitempty"`
	ToAssetHash string     `protobuf:"bytes,3,opt,name=toAssetHash,proto3" json:"toAssetHash,omitempty"`
}

func (m *DenomCrossChainInfo) Reset()         { *m = DenomCrossChainInfo{} }
func (m *DenomCrossChainInfo) String() string { return proto.CompactTextString(m) }
func (*DenomCrossChainInfo) ProtoMessage()    {}
func (*DenomCrossChainInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_25ff300fab0f746d, []int{0}
}
func (m *DenomCrossChainInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DenomCrossChainInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DenomCrossChainInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DenomCrossChainInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DenomCrossChainInfo.Merge(m, src)
}
func (m *DenomCrossChainInfo) XXX_Size() int {
	return m.Size()
}
func (m *DenomCrossChainInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DenomCrossChainInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DenomCrossChainInfo proto.InternalMessageInfo

func (m *DenomCrossChainInfo) GetDenomInfo() *DenomInfo {
	if m != nil {
		return m.DenomInfo
	}
	return nil
}

func (m *DenomCrossChainInfo) GetToChainId() uint64 {
	if m != nil {
		return m.ToChainId
	}
	return 0
}

func (m *DenomCrossChainInfo) GetToAssetHash() string {
	if m != nil {
		return m.ToAssetHash
	}
	return ""
}

func init() {
	proto.RegisterType((*DenomCrossChainInfo)(nil), "Switcheo.polynetworkcosmos.btcx.DenomCrossChainInfo")
}

func init() { proto.RegisterFile("btcx/denom_cross_chain_info.proto", fileDescriptor_25ff300fab0f746d) }

var fileDescriptor_25ff300fab0f746d = []byte{
	// 248 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x52, 0x4c, 0x2a, 0x49, 0xae,
	0xd0, 0x4f, 0x49, 0xcd, 0xcb, 0xcf, 0x8d, 0x4f, 0x2e, 0xca, 0x2f, 0x2e, 0x8e, 0x4f, 0xce, 0x48,
	0xcc, 0xcc, 0x8b, 0xcf, 0xcc, 0x4b, 0xcb, 0xd7, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x92, 0x0f,
	0x2e, 0xcf, 0x2c, 0x49, 0xce, 0x48, 0xcd, 0xd7, 0x2b, 0xc8, 0xcf, 0xa9, 0xcc, 0x4b, 0x2d, 0x29,
	0xcf, 0x2f, 0xca, 0x4e, 0xce, 0x2f, 0xce, 0xcd, 0x2f, 0xd6, 0x03, 0xe9, 0x96, 0x12, 0x45, 0x32,
	0x03, 0xa1, 0x4f, 0x69, 0x3e, 0x23, 0x97, 0xb0, 0x0b, 0x48, 0xd0, 0x19, 0x64, 0xae, 0x33, 0xc8,
	0x58, 0xcf, 0xbc, 0xb4, 0x7c, 0x21, 0x0f, 0x2e, 0x4e, 0xb0, 0x5a, 0x10, 0x47, 0x82, 0x51, 0x81,
	0x51, 0x83, 0xdb, 0x48, 0x4b, 0x8f, 0x80, 0x1d, 0x7a, 0x2e, 0x30, 0x1d, 0x41, 0x08, 0xcd, 0x42,
	0x32, 0x5c, 0x9c, 0x25, 0xf9, 0x10, 0x83, 0x53, 0x24, 0x98, 0x14, 0x18, 0x35, 0x58, 0x82, 0x10,
	0x02, 0x42, 0x0a, 0x5c, 0xdc, 0x25, 0xf9, 0x8e, 0xc5, 0xc5, 0xa9, 0x25, 0x1e, 0x89, 0xc5, 0x19,
	0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x9c, 0x41, 0xc8, 0x42, 0x4e, 0xbe, 0x27, 0x1e, 0xc9, 0x31, 0x5e,
	0x78, 0x24, 0xc7, 0xf8, 0xe0, 0x91, 0x1c, 0xe3, 0x84, 0xc7, 0x72, 0x0c, 0x17, 0x1e, 0xcb, 0x31,
	0xdc, 0x78, 0x2c, 0xc7, 0x10, 0x65, 0x9c, 0x9e, 0x59, 0x92, 0x51, 0x9a, 0xa4, 0x97, 0x9c, 0x9f,
	0xab, 0x0f, 0x73, 0x9a, 0x3e, 0x92, 0xd3, 0x74, 0x21, 0x6e, 0xd3, 0xaf, 0xd0, 0x07, 0xfb, 0xbd,
	0xa4, 0xb2, 0x20, 0xb5, 0x38, 0x89, 0x0d, 0xec, 0x6f, 0x63, 0x40, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x4e, 0x74, 0x64, 0x38, 0x54, 0x01, 0x00, 0x00,
}

func (m *DenomCrossChainInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DenomCrossChainInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DenomCrossChainInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.ToAssetHash) > 0 {
		i -= len(m.ToAssetHash)
		copy(dAtA[i:], m.ToAssetHash)
		i = encodeVarintDenomCrossChainInfo(dAtA, i, uint64(len(m.ToAssetHash)))
		i--
		dAtA[i] = 0x1a
	}
	if m.ToChainId != 0 {
		i = encodeVarintDenomCrossChainInfo(dAtA, i, uint64(m.ToChainId))
		i--
		dAtA[i] = 0x10
	}
	if m.DenomInfo != nil {
		{
			size, err := m.DenomInfo.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintDenomCrossChainInfo(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintDenomCrossChainInfo(dAtA []byte, offset int, v uint64) int {
	offset -= sovDenomCrossChainInfo(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *DenomCrossChainInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.DenomInfo != nil {
		l = m.DenomInfo.Size()
		n += 1 + l + sovDenomCrossChainInfo(uint64(l))
	}
	if m.ToChainId != 0 {
		n += 1 + sovDenomCrossChainInfo(uint64(m.ToChainId))
	}
	l = len(m.ToAssetHash)
	if l > 0 {
		n += 1 + l + sovDenomCrossChainInfo(uint64(l))
	}
	return n
}

func sovDenomCrossChainInfo(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozDenomCrossChainInfo(x uint64) (n int) {
	return sovDenomCrossChainInfo(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *DenomCrossChainInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowDenomCrossChainInfo
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
			return fmt.Errorf("proto: DenomCrossChainInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DenomCrossChainInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomInfo", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomCrossChainInfo
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
				return ErrInvalidLengthDenomCrossChainInfo
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthDenomCrossChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.DenomInfo == nil {
				m.DenomInfo = &DenomInfo{}
			}
			if err := m.DenomInfo.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToChainId", wireType)
			}
			m.ToChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomCrossChainInfo
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ToChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ToAssetHash", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowDenomCrossChainInfo
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
				return ErrInvalidLengthDenomCrossChainInfo
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthDenomCrossChainInfo
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ToAssetHash = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipDenomCrossChainInfo(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthDenomCrossChainInfo
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
func skipDenomCrossChainInfo(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowDenomCrossChainInfo
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
					return 0, ErrIntOverflowDenomCrossChainInfo
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
					return 0, ErrIntOverflowDenomCrossChainInfo
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
				return 0, ErrInvalidLengthDenomCrossChainInfo
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupDenomCrossChainInfo
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthDenomCrossChainInfo
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthDenomCrossChainInfo        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowDenomCrossChainInfo          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupDenomCrossChainInfo = fmt.Errorf("proto: unexpected end of group")
)