// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: lorenzo/btclightclient/v1/btclightclient.proto

package types

import (
	cosmossdk_io_math "cosmossdk.io/math"
	fmt "fmt"
	github_com_Lorenzo_Protocol_lorenzo_types "github.com/Lorenzo-Protocol/lorenzo/types"
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

// BTCHeaderInfo is a structure that contains all relevant information about a
// BTC header
//   - Full header bytes
//   - Header hash for easy retrieval
//   - Height of the header in the BTC chain
//   - Total work spent on the header. This is the sum of the work corresponding
//     to the header Bits field
//     and the total work of the header.
type BTCHeaderInfo struct {
	Header *github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderBytes     `protobuf:"bytes,1,opt,name=header,proto3,customtype=github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderBytes" json:"header,omitempty"`
	Hash   *github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderHashBytes `protobuf:"bytes,2,opt,name=hash,proto3,customtype=github.com/Lorenzo-Protocol/lorenzo/types.BTCHeaderHashBytes" json:"hash,omitempty"`
	Height uint64                                                        `protobuf:"varint,3,opt,name=height,proto3" json:"height,omitempty"`
	Work   *cosmossdk_io_math.Uint                                       `protobuf:"bytes,4,opt,name=work,proto3,customtype=cosmossdk.io/math.Uint" json:"work,omitempty"`
}

func (m *BTCHeaderInfo) Reset()         { *m = BTCHeaderInfo{} }
func (m *BTCHeaderInfo) String() string { return proto.CompactTextString(m) }
func (*BTCHeaderInfo) ProtoMessage()    {}
func (*BTCHeaderInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_2107ea75e21373d0, []int{0}
}
func (m *BTCHeaderInfo) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCHeaderInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCHeaderInfo.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCHeaderInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCHeaderInfo.Merge(m, src)
}
func (m *BTCHeaderInfo) XXX_Size() int {
	return m.Size()
}
func (m *BTCHeaderInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCHeaderInfo.DiscardUnknown(m)
}

var xxx_messageInfo_BTCHeaderInfo proto.InternalMessageInfo

func (m *BTCHeaderInfo) GetHeight() uint64 {
	if m != nil {
		return m.Height
	}
	return 0
}

type BTCFeeRate struct {
	FeeRate uint64 `protobuf:"varint,1,opt,name=fee_rate,json=feeRate,proto3" json:"fee_rate,omitempty"`
}

func (m *BTCFeeRate) Reset()         { *m = BTCFeeRate{} }
func (m *BTCFeeRate) String() string { return proto.CompactTextString(m) }
func (*BTCFeeRate) ProtoMessage()    {}
func (*BTCFeeRate) Descriptor() ([]byte, []int) {
	return fileDescriptor_2107ea75e21373d0, []int{1}
}
func (m *BTCFeeRate) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *BTCFeeRate) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_BTCFeeRate.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *BTCFeeRate) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BTCFeeRate.Merge(m, src)
}
func (m *BTCFeeRate) XXX_Size() int {
	return m.Size()
}
func (m *BTCFeeRate) XXX_DiscardUnknown() {
	xxx_messageInfo_BTCFeeRate.DiscardUnknown(m)
}

var xxx_messageInfo_BTCFeeRate proto.InternalMessageInfo

func (m *BTCFeeRate) GetFeeRate() uint64 {
	if m != nil {
		return m.FeeRate
	}
	return 0
}

func init() {
	proto.RegisterType((*BTCHeaderInfo)(nil), "lorenzo.btclightclient.v1.BTCHeaderInfo")
	proto.RegisterType((*BTCFeeRate)(nil), "lorenzo.btclightclient.v1.BTCFeeRate")
}

func init() {
	proto.RegisterFile("lorenzo/btclightclient/v1/btclightclient.proto", fileDescriptor_2107ea75e21373d0)
}

var fileDescriptor_2107ea75e21373d0 = []byte{
	// 319 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0xcb, 0xc9, 0x2f, 0x4a,
	0xcd, 0xab, 0xca, 0xd7, 0x4f, 0x2a, 0x49, 0xce, 0xc9, 0x4c, 0xcf, 0x00, 0x91, 0xa9, 0x79, 0x25,
	0xfa, 0x65, 0x86, 0x68, 0x22, 0x7a, 0x05, 0x45, 0xf9, 0x25, 0xf9, 0x42, 0x92, 0x50, 0xf5, 0x7a,
	0x68, 0xb2, 0x65, 0x86, 0x52, 0x22, 0xe9, 0xf9, 0xe9, 0xf9, 0x60, 0x55, 0xfa, 0x20, 0x16, 0x44,
	0x83, 0x52, 0x33, 0x13, 0x17, 0xaf, 0x53, 0x88, 0xb3, 0x47, 0x6a, 0x62, 0x4a, 0x6a, 0x91, 0x67,
	0x5e, 0x5a, 0xbe, 0x50, 0x08, 0x17, 0x5b, 0x06, 0x98, 0x27, 0xc1, 0xa8, 0xc0, 0xa8, 0xc1, 0xe3,
	0x64, 0x73, 0xeb, 0x9e, 0xbc, 0x45, 0x7a, 0x66, 0x49, 0x46, 0x69, 0x92, 0x5e, 0x72, 0x7e, 0xae,
	0xbe, 0x0f, 0xc4, 0x06, 0xdd, 0x00, 0x90, 0xfe, 0xe4, 0xfc, 0x1c, 0x7d, 0x98, 0x13, 0x4b, 0x2a,
	0x0b, 0x52, 0x8b, 0xf5, 0xe0, 0x86, 0x39, 0x55, 0x96, 0xa4, 0x16, 0x07, 0x41, 0xcd, 0x12, 0x0a,
	0xe1, 0x62, 0xc9, 0x48, 0x2c, 0xce, 0x90, 0x60, 0x02, 0x9b, 0xe9, 0x70, 0xeb, 0x9e, 0xbc, 0x0d,
	0x19, 0x66, 0x7a, 0x24, 0x16, 0x67, 0x40, 0xcc, 0x05, 0x9b, 0x26, 0x24, 0x06, 0x72, 0x2b, 0xc8,
	0x9b, 0x12, 0xcc, 0x0a, 0x8c, 0x1a, 0x2c, 0x41, 0x50, 0x9e, 0x90, 0x1e, 0x17, 0x4b, 0x79, 0x7e,
	0x51, 0xb6, 0x04, 0x0b, 0xd8, 0x36, 0xa9, 0x5b, 0xf7, 0xe4, 0xc5, 0x92, 0xf3, 0x8b, 0x73, 0xf3,
	0x8b, 0x8b, 0x53, 0xb2, 0xf5, 0x32, 0xf3, 0xf5, 0x73, 0x13, 0x4b, 0x32, 0xf4, 0x42, 0x33, 0xf3,
	0x4a, 0x82, 0xc0, 0xea, 0x94, 0xd4, 0xb9, 0xb8, 0x9c, 0x42, 0x9c, 0xdd, 0x52, 0x53, 0x83, 0x12,
	0x4b, 0x52, 0x85, 0x24, 0xb9, 0x38, 0xd2, 0x52, 0x53, 0xe3, 0x8b, 0x12, 0x4b, 0x52, 0xc1, 0x61,
	0xc0, 0x12, 0xc4, 0x9e, 0x06, 0x91, 0x72, 0x0a, 0x39, 0xf1, 0x48, 0x8e, 0xf1, 0xc2, 0x23, 0x39,
	0xc6, 0x07, 0x8f, 0xe4, 0x18, 0x27, 0x3c, 0x96, 0x63, 0xb8, 0xf0, 0x58, 0x8e, 0xe1, 0xc6, 0x63,
	0x39, 0x86, 0x28, 0x2b, 0x62, 0xbc, 0x53, 0x81, 0x1e, 0x8f, 0x60, 0xff, 0x25, 0xb1, 0x81, 0xe3,
	0xc2, 0x18, 0x10, 0x00, 0x00, 0xff, 0xff, 0x32, 0x3b, 0xaa, 0x1d, 0xee, 0x01, 0x00, 0x00,
}

func (m *BTCHeaderInfo) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCHeaderInfo) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCHeaderInfo) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Work != nil {
		{
			size := m.Work.Size()
			i -= size
			if _, err := m.Work.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtclightclient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x22
	}
	if m.Height != 0 {
		i = encodeVarintBtclightclient(dAtA, i, uint64(m.Height))
		i--
		dAtA[i] = 0x18
	}
	if m.Hash != nil {
		{
			size := m.Hash.Size()
			i -= size
			if _, err := m.Hash.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtclightclient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	if m.Header != nil {
		{
			size := m.Header.Size()
			i -= size
			if _, err := m.Header.MarshalTo(dAtA[i:]); err != nil {
				return 0, err
			}
			i = encodeVarintBtclightclient(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *BTCFeeRate) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *BTCFeeRate) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *BTCFeeRate) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.FeeRate != 0 {
		i = encodeVarintBtclightclient(dAtA, i, uint64(m.FeeRate))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func encodeVarintBtclightclient(dAtA []byte, offset int, v uint64) int {
	offset -= sovBtclightclient(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *BTCHeaderInfo) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Header != nil {
		l = m.Header.Size()
		n += 1 + l + sovBtclightclient(uint64(l))
	}
	if m.Hash != nil {
		l = m.Hash.Size()
		n += 1 + l + sovBtclightclient(uint64(l))
	}
	if m.Height != 0 {
		n += 1 + sovBtclightclient(uint64(m.Height))
	}
	if m.Work != nil {
		l = m.Work.Size()
		n += 1 + l + sovBtclightclient(uint64(l))
	}
	return n
}

func (m *BTCFeeRate) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.FeeRate != 0 {
		n += 1 + sovBtclightclient(uint64(m.FeeRate))
	}
	return n
}

func sovBtclightclient(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozBtclightclient(x uint64) (n int) {
	return sovBtclightclient(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *BTCHeaderInfo) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtclightclient
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
			return fmt.Errorf("proto: BTCHeaderInfo: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCHeaderInfo: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Header", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtclightclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtclightclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderBytes
			m.Header = &v
			if err := m.Header.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Hash", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtclightclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtclightclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v github_com_Lorenzo_Protocol_lorenzo_types.BTCHeaderHashBytes
			m.Hash = &v
			if err := m.Hash.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Height", wireType)
			}
			m.Height = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Height |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Work", wireType)
			}
			var byteLen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				byteLen |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			if byteLen < 0 {
				return ErrInvalidLengthBtclightclient
			}
			postIndex := iNdEx + byteLen
			if postIndex < 0 {
				return ErrInvalidLengthBtclightclient
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			var v cosmossdk_io_math.Uint
			m.Work = &v
			if err := m.Work.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipBtclightclient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtclightclient
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
func (m *BTCFeeRate) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowBtclightclient
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
			return fmt.Errorf("proto: BTCFeeRate: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: BTCFeeRate: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field FeeRate", wireType)
			}
			m.FeeRate = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowBtclightclient
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.FeeRate |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipBtclightclient(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthBtclightclient
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
func skipBtclightclient(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowBtclightclient
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
					return 0, ErrIntOverflowBtclightclient
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
					return 0, ErrIntOverflowBtclightclient
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
				return 0, ErrInvalidLengthBtclightclient
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupBtclightclient
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthBtclightclient
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthBtclightclient        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowBtclightclient          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupBtclightclient = fmt.Errorf("proto: unexpected end of group")
)
