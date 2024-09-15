// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/amm/genesis.proto

package types

import (
	fmt "fmt"
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

// GenesisState defines the amm module's genesis state.
type GenesisState struct {
	Params             Params                    `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	PoolList           []Pool                    `protobuf:"bytes,2,rep,name=pool_list,json=poolList,proto3" json:"pool_list"`
	DenomLiquidityList []DenomLiquidity          `protobuf:"bytes,3,rep,name=denom_liquidity_list,json=denomLiquidityList,proto3" json:"denom_liquidity_list"`
	SlippageTracks     []OraclePoolSlippageTrack `protobuf:"bytes,4,rep,name=slippage_tracks,json=slippageTracks,proto3" json:"slippage_tracks"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_836f20eb8daba51a, []int{0}
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

func (m *GenesisState) GetPoolList() []Pool {
	if m != nil {
		return m.PoolList
	}
	return nil
}

func (m *GenesisState) GetDenomLiquidityList() []DenomLiquidity {
	if m != nil {
		return m.DenomLiquidityList
	}
	return nil
}

func (m *GenesisState) GetSlippageTracks() []OraclePoolSlippageTrack {
	if m != nil {
		return m.SlippageTracks
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "elys.amm.GenesisState")
}

func init() { proto.RegisterFile("elys/amm/genesis.proto", fileDescriptor_836f20eb8daba51a) }

var fileDescriptor_836f20eb8daba51a = []byte{
	// 320 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x5c, 0x90, 0xb1, 0x4e, 0xeb, 0x30,
	0x14, 0x86, 0x93, 0xb6, 0xaa, 0x7a, 0xd3, 0xab, 0x82, 0x42, 0x41, 0x51, 0x07, 0x53, 0x98, 0xba,
	0xe0, 0x88, 0xf2, 0x06, 0x11, 0x12, 0x4b, 0x25, 0x2a, 0xca, 0xc4, 0x12, 0xb9, 0x89, 0x15, 0xac,
	0xda, 0x71, 0x88, 0x5d, 0x41, 0xde, 0x82, 0x85, 0x77, 0xea, 0xd8, 0x91, 0x09, 0xa1, 0xe4, 0x45,
	0x90, 0x1d, 0xb7, 0x29, 0x6c, 0xc9, 0xe7, 0xff, 0xff, 0xce, 0xd1, 0x71, 0xce, 0x30, 0x2d, 0x84,
	0x8f, 0x18, 0xf3, 0x13, 0x9c, 0x62, 0x41, 0x04, 0xcc, 0x72, 0x2e, 0xb9, 0xdb, 0x53, 0x1c, 0x22,
	0xc6, 0x46, 0xc3, 0x84, 0x27, 0x5c, 0x43, 0x5f, 0x7d, 0xd5, 0xef, 0xa3, 0xd3, 0x7d, 0x2f, 0x43,
	0x39, 0x62, 0xa6, 0x36, 0x3a, 0x69, 0x30, 0xe7, 0xd4, 0x40, 0xb0, 0x87, 0x31, 0x4e, 0x39, 0x0b,
	0x29, 0x79, 0x59, 0x93, 0x98, 0xc8, 0xa2, 0x7e, 0xbf, 0xfc, 0x68, 0x39, 0xff, 0xef, 0xea, 0xe9,
	0x0b, 0x89, 0x24, 0x76, 0xa1, 0xd3, 0xad, 0xad, 0x9e, 0x3d, 0xb6, 0x27, 0xfd, 0xe9, 0x31, 0xdc,
	0x6d, 0x03, 0xe7, 0x9a, 0x07, 0x9d, 0xcd, 0xd7, 0xb9, 0xf5, 0x60, 0x52, 0xee, 0xb5, 0xf3, 0x4f,
	0x8d, 0x0b, 0x29, 0x11, 0xd2, 0x6b, 0x8d, 0xdb, 0x93, 0xfe, 0x74, 0x70, 0x50, 0xe1, 0x9c, 0x9a,
	0x42, 0x4f, 0xc5, 0x66, 0x44, 0x48, 0x77, 0xee, 0x0c, 0xff, 0x2c, 0x53, 0xb7, 0xdb, 0xba, 0xed,
	0x35, 0xed, 0x5b, 0x95, 0x9a, 0xed, 0x42, 0xc6, 0xe3, 0xc6, 0xbf, 0xa8, 0x31, 0x1e, 0x09, 0x4a,
	0xb2, 0x0c, 0x25, 0x38, 0x94, 0x39, 0x8a, 0x56, 0xc2, 0xeb, 0x68, 0xd9, 0x45, 0x23, 0xbb, 0xcf,
	0x51, 0x44, 0xb1, 0x5a, 0x68, 0x61, 0xa2, 0x8f, 0x2a, 0x69, 0xac, 0x03, 0x71, 0x08, 0x45, 0x10,
	0x6c, 0x4a, 0x60, 0x6f, 0x4b, 0x60, 0x7f, 0x97, 0xc0, 0x7e, 0xaf, 0x80, 0xb5, 0xad, 0x80, 0xf5,
	0x59, 0x01, 0xeb, 0x69, 0x92, 0x10, 0xf9, 0xbc, 0x5e, 0xc2, 0x88, 0x33, 0x5f, 0xc9, 0xaf, 0x52,
	0x2c, 0x5f, 0x79, 0xbe, 0xd2, 0x3f, 0xfe, 0x9b, 0xbe, 0xb5, 0x2c, 0x32, 0x2c, 0x96, 0x5d, 0x7d,
	0xe2, 0x9b, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x69, 0x0b, 0x0a, 0x47, 0xe8, 0x01, 0x00, 0x00,
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
	if len(m.SlippageTracks) > 0 {
		for iNdEx := len(m.SlippageTracks) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.SlippageTracks[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.DenomLiquidityList) > 0 {
		for iNdEx := len(m.DenomLiquidityList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.DenomLiquidityList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PoolList) > 0 {
		for iNdEx := len(m.PoolList) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolList[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.PoolList) > 0 {
		for _, e := range m.PoolList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.DenomLiquidityList) > 0 {
		for _, e := range m.DenomLiquidityList {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.SlippageTracks) > 0 {
		for _, e := range m.SlippageTracks {
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
				return fmt.Errorf("proto: wrong wireType = %d for field PoolList", wireType)
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
			m.PoolList = append(m.PoolList, Pool{})
			if err := m.PoolList[len(m.PoolList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DenomLiquidityList", wireType)
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
			m.DenomLiquidityList = append(m.DenomLiquidityList, DenomLiquidity{})
			if err := m.DenomLiquidityList[len(m.DenomLiquidityList)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field SlippageTracks", wireType)
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
			m.SlippageTracks = append(m.SlippageTracks, OraclePoolSlippageTrack{})
			if err := m.SlippageTracks[len(m.SlippageTracks)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
