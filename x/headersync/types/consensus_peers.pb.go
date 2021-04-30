// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: headersync/consensus_peers.proto

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

type Peer struct {
	Index      uint32 `protobuf:"varint,1,opt,name=index,proto3" json:"index,omitempty"`
	PeerPubKey string `protobuf:"bytes,2,opt,name=peerPubKey,proto3" json:"peerPubKey,omitempty"`
}

func (m *Peer) Reset()         { *m = Peer{} }
func (m *Peer) String() string { return proto.CompactTextString(m) }
func (*Peer) ProtoMessage()    {}
func (*Peer) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7b41d5847534ac, []int{0}
}
func (m *Peer) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Peer) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Peer.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Peer) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Peer.Merge(m, src)
}
func (m *Peer) XXX_Size() int {
	return m.Size()
}
func (m *Peer) XXX_DiscardUnknown() {
	xxx_messageInfo_Peer.DiscardUnknown(m)
}

var xxx_messageInfo_Peer proto.InternalMessageInfo

func (m *Peer) GetIndex() uint32 {
	if m != nil {
		return m.Index
	}
	return 0
}

func (m *Peer) GetPeerPubKey() string {
	if m != nil {
		return m.PeerPubKey
	}
	return ""
}

type ConsensusPeers struct {
	ChainId uint64           `protobuf:"varint,1,opt,name=chainId,proto3" json:"chainId,omitempty"`
	Height  uint32           `protobuf:"varint,2,opt,name=height,proto3" json:"height,omitempty"`
	Peers   map[string]*Peer `protobuf:"bytes,3,rep,name=peers,proto3" json:"peers,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (m *ConsensusPeers) Reset()         { *m = ConsensusPeers{} }
func (m *ConsensusPeers) String() string { return proto.CompactTextString(m) }
func (*ConsensusPeers) ProtoMessage()    {}
func (*ConsensusPeers) Descriptor() ([]byte, []int) {
	return fileDescriptor_5d7b41d5847534ac, []int{1}
}
func (m *ConsensusPeers) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ConsensusPeers) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ConsensusPeers.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ConsensusPeers) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ConsensusPeers.Merge(m, src)
}
func (m *ConsensusPeers) XXX_Size() int {
	return m.Size()
}
func (m *ConsensusPeers) XXX_DiscardUnknown() {
	xxx_messageInfo_ConsensusPeers.DiscardUnknown(m)
}

var xxx_messageInfo_ConsensusPeers proto.InternalMessageInfo

func (m *ConsensusPeers) GetChainId() uint64 {
	if m != nil {
		return m.ChainId
	}
	return 0
}

func (m *ConsensusPeers) GetHeight() uint32 {
	if m != nil {
		return m.Height
	}
	return 0
}

func (m *ConsensusPeers) GetPeers() map[string]*Peer {
	if m != nil {
		return m.Peers
	}
	return nil
}

func init() {
	proto.RegisterType((*Peer)(nil), "Switcheo.polynetworkcosmos.headersync.Peer")
	proto.RegisterType((*ConsensusPeers)(nil), "Switcheo.polynetworkcosmos.headersync.ConsensusPeers")
	proto.RegisterMapType((map[string]*Peer)(nil), "Switcheo.polynetworkcosmos.headersync.ConsensusPeers.PeersEntry")
}

func init() { proto.RegisterFile("headersync/consensus_peers.proto", fileDescriptor_5d7b41d5847534ac) }

var fileDescriptor_5d7b41d5847534ac = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x91, 0xbd, 0x4a, 0x33, 0x41,
	0x14, 0x86, 0x33, 0xf9, 0xfb, 0xc8, 0x09, 0xf9, 0x90, 0x41, 0x64, 0xb1, 0x18, 0x96, 0x80, 0x10,
	0x10, 0x67, 0x21, 0x36, 0x2a, 0x16, 0xfe, 0x60, 0x21, 0x36, 0x61, 0x02, 0x16, 0x36, 0x92, 0x6c,
	0x0e, 0xd9, 0x25, 0xc9, 0xcc, 0xb2, 0x33, 0x6b, 0x32, 0xa5, 0x77, 0xe0, 0x65, 0x59, 0xa6, 0xb4,
	0x94, 0xe4, 0x46, 0x64, 0x67, 0x0d, 0x89, 0x5d, 0x9a, 0x61, 0xde, 0x03, 0xe7, 0x79, 0x1f, 0x38,
	0xe0, 0x47, 0x38, 0x18, 0x61, 0xaa, 0xad, 0x0c, 0x83, 0x50, 0x49, 0x8d, 0x52, 0x67, 0xfa, 0x35,
	0x41, 0x4c, 0x35, 0x4f, 0x52, 0x65, 0x14, 0x3d, 0xe9, 0xcf, 0x63, 0x13, 0x46, 0xa8, 0x78, 0xa2,
	0xa6, 0x56, 0xa2, 0x99, 0xab, 0x74, 0x12, 0x2a, 0x3d, 0x53, 0x9a, 0x6f, 0x97, 0xdb, 0xd7, 0x50,
	0xed, 0x21, 0xa6, 0xf4, 0x10, 0x6a, 0xb1, 0x1c, 0xe1, 0xc2, 0x23, 0x3e, 0xe9, 0xb4, 0x44, 0x11,
	0x28, 0x03, 0xc8, 0x99, 0xbd, 0x6c, 0xf8, 0x84, 0xd6, 0x2b, 0xfb, 0xa4, 0xd3, 0x10, 0x3b, 0x93,
	0xf6, 0x7b, 0x19, 0xfe, 0xdf, 0x6f, 0xea, 0x73, 0x8e, 0xa6, 0x1e, 0xfc, 0x0b, 0xa3, 0x41, 0x2c,
	0x1f, 0x47, 0x0e, 0x55, 0x15, 0x9b, 0x48, 0x8f, 0xa0, 0x1e, 0x61, 0x3c, 0x8e, 0x8c, 0x03, 0xb5,
	0xc4, 0x6f, 0xa2, 0xcf, 0x50, 0x73, 0xe2, 0x5e, 0xc5, 0xaf, 0x74, 0x9a, 0xdd, 0x1b, 0xbe, 0x97,
	0x39, 0xff, 0xdb, 0xcb, 0xdd, 0xfb, 0x20, 0x4d, 0x6a, 0x45, 0x81, 0x3b, 0x46, 0x80, 0xed, 0x90,
	0x1e, 0x40, 0x65, 0x82, 0xd6, 0x39, 0x35, 0x44, 0xfe, 0xa5, 0xb7, 0x50, 0x7b, 0x1b, 0x4c, 0x33,
	0x74, 0x3a, 0xcd, 0xee, 0xe9, 0x9e, 0xbd, 0x39, 0x53, 0x14, 0x9b, 0x57, 0xe5, 0x0b, 0x72, 0xd7,
	0xff, 0x5c, 0x31, 0xb2, 0x5c, 0x31, 0xf2, 0xbd, 0x62, 0xe4, 0x63, 0xcd, 0x4a, 0xcb, 0x35, 0x2b,
	0x7d, 0xad, 0x59, 0xe9, 0xe5, 0x72, 0x1c, 0x9b, 0x28, 0x1b, 0xf2, 0x50, 0xcd, 0x82, 0x0d, 0x3b,
	0xd8, 0x61, 0x9f, 0x15, 0xf0, 0x60, 0x11, 0xec, 0x5c, 0xd3, 0xd8, 0x04, 0xf5, 0xb0, 0xee, 0x8e,
	0x78, 0xfe, 0x13, 0x00, 0x00, 0xff, 0xff, 0x9f, 0xf8, 0x78, 0x2e, 0xe8, 0x01, 0x00, 0x00,
}

func (m *Peer) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Peer) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Peer) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.PeerPubKey) > 0 {
		i -= len(m.PeerPubKey)
		copy(dAtA[i:], m.PeerPubKey)
		i = encodeVarintConsensusPeers(dAtA, i, uint64(len(m.PeerPubKey)))
		i--
		dAtA[i] = 0x12
	}
	if m.Index != 0 {
		i = encodeVarintConsensusPeers(dAtA, i, uint64(m.Index))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ConsensusPeers) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ConsensusPeers) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ConsensusPeers) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Peers) > 0 {
		for k := range m.Peers {
			v := m.Peers[k]
			baseI := i
			if v != nil {
				{
					size, err := v.MarshalToSizedBuffer(dAtA[:i])
					if err != nil {
						return 0, err
					}
					i -= size
					i = encodeVarintConsensusPeers(dAtA, i, uint64(size))
				}
				i--
				dAtA[i] = 0x12
			}
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintConsensusPeers(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintConsensusPeers(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0x1a
		}
	}
	if m.Height != 0 {
		i = encodeVarintConsensusPeers(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x10
	}
	if m.ChainId != 0 {
		i = encodeVarintConsensusPeers(dAtA, i, uint64(m.ChainId))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintConsensusPeers(dAtA []byte, offset int, v uint64) int {
	offset -= sovConsensusPeers(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Peer) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Index != 0 {
		n += 1 + sovConsensusPeers(uint64(m.Index))
	}
	l = len(m.PeerPubKey)
	if l > 0 {
		n += 1 + l + sovConsensusPeers(uint64(l))
	}
	return n
}

func (m *ConsensusPeers) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.ChainId != 0 {
		n += 1 + sovConsensusPeers(uint64(m.ChainId))
	}
	if m.Height != 0 {
		n += 1 + sovConsensusPeers(uint64(m.Height))
	}
	if len(m.Peers) > 0 {
		for k, v := range m.Peers {
			_ = k
			_ = v
			l = 0
			if v != nil {
				l = v.Size()
				l += 1 + sovConsensusPeers(uint64(l))
			}
			mapEntrySize := 1 + len(k) + sovConsensusPeers(uint64(len(k))) + l
			n += mapEntrySize + 1 + sovConsensusPeers(uint64(mapEntrySize))
		}
	}
	return n
}

func sovConsensusPeers(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozConsensusPeers(x uint64) (n int) {
	return sovConsensusPeers(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Peer) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusPeers
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
			return fmt.Errorf("proto: Peer: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Peer: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Index", wireType)
			}
			m.Index = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Index |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PeerPubKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusPeers
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
				return ErrInvalidLengthConsensusPeers
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusPeers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.PeerPubKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusPeers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusPeers
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
func (m *ConsensusPeers) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowConsensusPeers
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
			return fmt.Errorf("proto: ConsensusPeers: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ConsensusPeers: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ChainId", wireType)
			}
			m.ChainId = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ChainId |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusPeers
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint32(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Peers", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowConsensusPeers
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
				return ErrInvalidLengthConsensusPeers
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthConsensusPeers
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Peers == nil {
				m.Peers = make(map[string]*Peer)
			}
			var mapkey string
			var mapvalue *Peer
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowConsensusPeers
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
				if fieldNum == 1 {
					var stringLenmapkey uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConsensusPeers
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapkey |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapkey := int(stringLenmapkey)
					if intStringLenmapkey < 0 {
						return ErrInvalidLengthConsensusPeers
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthConsensusPeers
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var mapmsglen int
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowConsensusPeers
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						mapmsglen |= int(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					if mapmsglen < 0 {
						return ErrInvalidLengthConsensusPeers
					}
					postmsgIndex := iNdEx + mapmsglen
					if postmsgIndex < 0 {
						return ErrInvalidLengthConsensusPeers
					}
					if postmsgIndex > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = &Peer{}
					if err := mapvalue.Unmarshal(dAtA[iNdEx:postmsgIndex]); err != nil {
						return err
					}
					iNdEx = postmsgIndex
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipConsensusPeers(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if (skippy < 0) || (iNdEx+skippy) < 0 {
						return ErrInvalidLengthConsensusPeers
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Peers[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipConsensusPeers(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthConsensusPeers
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
func skipConsensusPeers(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowConsensusPeers
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
					return 0, ErrIntOverflowConsensusPeers
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
					return 0, ErrIntOverflowConsensusPeers
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
				return 0, ErrInvalidLengthConsensusPeers
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupConsensusPeers
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthConsensusPeers
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthConsensusPeers        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowConsensusPeers          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupConsensusPeers = fmt.Errorf("proto: unexpected end of group")
)
