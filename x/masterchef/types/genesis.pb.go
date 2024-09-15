// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: elys/masterchef/genesis.proto

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

// GenesisState defines the masterchef module's genesis state.
type GenesisState struct {
	Params                 Params              `protobuf:"bytes,1,opt,name=params,proto3" json:"params"`
	ExternalIncentives     []ExternalIncentive `protobuf:"bytes,2,rep,name=external_incentives,json=externalIncentives,proto3" json:"external_incentives"`
	ExternalIncentiveIndex uint64              `protobuf:"varint,3,opt,name=external_incentive_index,json=externalIncentiveIndex,proto3" json:"external_incentive_index,omitempty"`
	PoolInfos              []PoolInfo          `protobuf:"bytes,4,rep,name=pool_infos,json=poolInfos,proto3" json:"pool_infos"`
	PoolRewardInfos        []PoolRewardInfo    `protobuf:"bytes,5,rep,name=pool_reward_infos,json=poolRewardInfos,proto3" json:"pool_reward_infos"`
	UserRewardInfos        []UserRewardInfo    `protobuf:"bytes,6,rep,name=user_reward_infos,json=userRewardInfos,proto3" json:"user_reward_infos"`
	PoolRewardsAccum       []PoolRewardsAccum  `protobuf:"bytes,7,rep,name=pool_rewards_accum,json=poolRewardsAccum,proto3" json:"pool_rewards_accum"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}
func (*GenesisState) Descriptor() ([]byte, []int) {
	return fileDescriptor_b55d8c0403fbd8da, []int{0}
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

func (m *GenesisState) GetExternalIncentives() []ExternalIncentive {
	if m != nil {
		return m.ExternalIncentives
	}
	return nil
}

func (m *GenesisState) GetExternalIncentiveIndex() uint64 {
	if m != nil {
		return m.ExternalIncentiveIndex
	}
	return 0
}

func (m *GenesisState) GetPoolInfos() []PoolInfo {
	if m != nil {
		return m.PoolInfos
	}
	return nil
}

func (m *GenesisState) GetPoolRewardInfos() []PoolRewardInfo {
	if m != nil {
		return m.PoolRewardInfos
	}
	return nil
}

func (m *GenesisState) GetUserRewardInfos() []UserRewardInfo {
	if m != nil {
		return m.UserRewardInfos
	}
	return nil
}

func (m *GenesisState) GetPoolRewardsAccum() []PoolRewardsAccum {
	if m != nil {
		return m.PoolRewardsAccum
	}
	return nil
}

func init() {
	proto.RegisterType((*GenesisState)(nil), "elys.masterchef.GenesisState")
}

func init() { proto.RegisterFile("elys/masterchef/genesis.proto", fileDescriptor_b55d8c0403fbd8da) }

var fileDescriptor_b55d8c0403fbd8da = []byte{
	// 390 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x41, 0x6b, 0xdb, 0x30,
	0x18, 0x86, 0xed, 0xc5, 0xcb, 0x98, 0x32, 0xc8, 0xa6, 0x8d, 0xcd, 0x33, 0x9b, 0x93, 0xe5, 0xe4,
	0xcb, 0x6c, 0xc8, 0x18, 0xec, 0x34, 0xd8, 0x60, 0x0c, 0xdf, 0xda, 0x94, 0x1c, 0xda, 0x8b, 0x71,
	0x9c, 0x2f, 0x8e, 0xa9, 0x2d, 0x19, 0x49, 0x6e, 0x92, 0x7f, 0xd1, 0x5b, 0xff, 0x52, 0x8e, 0x39,
	0xf6, 0x54, 0x4a, 0xf2, 0x47, 0x8a, 0x65, 0x95, 0x26, 0x76, 0xdb, 0x9b, 0xf4, 0x3d, 0xaf, 0x1e,
	0xbd, 0x02, 0xa1, 0xaf, 0x90, 0xae, 0xb8, 0x97, 0x85, 0x5c, 0x00, 0x8b, 0xe6, 0x30, 0xf3, 0x62,
	0x20, 0xc0, 0x13, 0xee, 0xe6, 0x8c, 0x0a, 0x8a, 0xbb, 0x25, 0x76, 0x1f, 0xb0, 0xf5, 0x21, 0xa6,
	0x31, 0x95, 0xcc, 0x2b, 0x57, 0x55, 0xcc, 0xfa, 0x52, 0xb7, 0xe4, 0x21, 0x0b, 0x33, 0x25, 0xb1,
	0x9c, 0x3a, 0x85, 0xa5, 0x00, 0x46, 0xc2, 0x34, 0x48, 0x48, 0x04, 0x44, 0x24, 0x17, 0xa0, 0x92,
	0x56, 0xc3, 0x43, 0x69, 0x5a, 0xb1, 0xc1, 0x95, 0x81, 0xde, 0xfc, 0xaf, 0xca, 0x9d, 0x88, 0x50,
	0x00, 0xfe, 0x89, 0xda, 0xd5, 0x35, 0xa6, 0xde, 0xd7, 0x9d, 0xce, 0xf0, 0x93, 0x5b, 0x2b, 0xeb,
	0x1e, 0x49, 0xfc, 0xd7, 0x58, 0xdf, 0xf4, 0xb4, 0x91, 0x0a, 0xe3, 0x53, 0xf4, 0xbe, 0x79, 0x3f,
	0x37, 0x5f, 0xf4, 0x5b, 0x4e, 0x67, 0x38, 0x68, 0x38, 0xfe, 0xa9, 0xac, 0x7f, 0x1f, 0x55, 0x3a,
	0x0c, 0x75, 0xc0, 0xf1, 0x2f, 0x64, 0x36, 0xd5, 0x41, 0x42, 0xa6, 0xb0, 0x34, 0x5b, 0x7d, 0xdd,
	0x31, 0x46, 0x1f, 0x1b, 0xa7, 0xfc, 0x92, 0xe2, 0xdf, 0x08, 0x95, 0x4f, 0x0d, 0x12, 0x32, 0xa3,
	0xdc, 0x34, 0x64, 0x97, 0xcf, 0xcd, 0xf7, 0x50, 0x9a, 0xfa, 0x64, 0x46, 0x55, 0x85, 0xd7, 0xb9,
	0xda, 0x73, 0x7c, 0x8c, 0xde, 0xc9, 0xf3, 0x0c, 0x16, 0x21, 0x9b, 0x2a, 0xcd, 0x4b, 0xa9, 0xe9,
	0x3d, 0xaa, 0x19, 0xc9, 0xe0, 0x9e, 0xac, 0x9b, 0x1f, 0x4c, 0xa5, 0xb2, 0xe0, 0xc0, 0x0e, 0x95,
	0xed, 0x27, 0x94, 0x63, 0x0e, 0xac, 0xa9, 0x2c, 0x0e, 0xa6, 0x1c, 0x8f, 0x11, 0xde, 0x6b, 0xc9,
	0x83, 0x30, 0x8a, 0x8a, 0xcc, 0x7c, 0x25, 0x9d, 0xdf, 0x9e, 0xa9, 0xc9, 0xff, 0x94, 0x41, 0x65,
	0x7d, 0x9b, 0xd7, 0xe7, 0xfe, 0x7a, 0x6b, 0xeb, 0x9b, 0xad, 0xad, 0xdf, 0x6e, 0x6d, 0xfd, 0x72,
	0x67, 0x6b, 0x9b, 0x9d, 0xad, 0x5d, 0xef, 0x6c, 0xed, 0xcc, 0x8b, 0x13, 0x31, 0x2f, 0x26, 0x6e,
	0x44, 0x33, 0xaf, 0xd4, 0x7f, 0x27, 0x20, 0x16, 0x94, 0x9d, 0xcb, 0x8d, 0xb7, 0xdc, 0xff, 0x69,
	0x62, 0x95, 0x03, 0x9f, 0xb4, 0xe5, 0x5f, 0xfb, 0x71, 0x17, 0x00, 0x00, 0xff, 0xff, 0x63, 0x59,
	0xfe, 0xa5, 0x17, 0x03, 0x00, 0x00,
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
	if len(m.PoolRewardsAccum) > 0 {
		for iNdEx := len(m.PoolRewardsAccum) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolRewardsAccum[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.UserRewardInfos) > 0 {
		for iNdEx := len(m.UserRewardInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.UserRewardInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x32
		}
	}
	if len(m.PoolRewardInfos) > 0 {
		for iNdEx := len(m.PoolRewardInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolRewardInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
				if err != nil {
					return 0, err
				}
				i -= size
				i = encodeVarintGenesis(dAtA, i, uint64(size))
			}
			i--
			dAtA[i] = 0x2a
		}
	}
	if len(m.PoolInfos) > 0 {
		for iNdEx := len(m.PoolInfos) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.PoolInfos[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if m.ExternalIncentiveIndex != 0 {
		i = encodeVarintGenesis(dAtA, i, uint64(m.ExternalIncentiveIndex))
		i--
		dAtA[i] = 0x18
	}
	if len(m.ExternalIncentives) > 0 {
		for iNdEx := len(m.ExternalIncentives) - 1; iNdEx >= 0; iNdEx-- {
			{
				size, err := m.ExternalIncentives[iNdEx].MarshalToSizedBuffer(dAtA[:i])
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
	if len(m.ExternalIncentives) > 0 {
		for _, e := range m.ExternalIncentives {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if m.ExternalIncentiveIndex != 0 {
		n += 1 + sovGenesis(uint64(m.ExternalIncentiveIndex))
	}
	if len(m.PoolInfos) > 0 {
		for _, e := range m.PoolInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolRewardInfos) > 0 {
		for _, e := range m.PoolRewardInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.UserRewardInfos) > 0 {
		for _, e := range m.UserRewardInfos {
			l = e.Size()
			n += 1 + l + sovGenesis(uint64(l))
		}
	}
	if len(m.PoolRewardsAccum) > 0 {
		for _, e := range m.PoolRewardsAccum {
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
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalIncentives", wireType)
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
			m.ExternalIncentives = append(m.ExternalIncentives, ExternalIncentive{})
			if err := m.ExternalIncentives[len(m.ExternalIncentives)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExternalIncentiveIndex", wireType)
			}
			m.ExternalIncentiveIndex = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenesis
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExternalIncentiveIndex |= uint64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolInfos", wireType)
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
			m.PoolInfos = append(m.PoolInfos, PoolInfo{})
			if err := m.PoolInfos[len(m.PoolInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRewardInfos", wireType)
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
			m.PoolRewardInfos = append(m.PoolRewardInfos, PoolRewardInfo{})
			if err := m.PoolRewardInfos[len(m.PoolRewardInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserRewardInfos", wireType)
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
			m.UserRewardInfos = append(m.UserRewardInfos, UserRewardInfo{})
			if err := m.UserRewardInfos[len(m.UserRewardInfos)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field PoolRewardsAccum", wireType)
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
			m.PoolRewardsAccum = append(m.PoolRewardsAccum, PoolRewardsAccum{})
			if err := m.PoolRewardsAccum[len(m.PoolRewardsAccum)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
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
