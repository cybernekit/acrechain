// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: acrechain/mint/v1beta1/mint.proto

package types

import (
	fmt "fmt"
	_ "github.com/cosmos/cosmos-sdk/codec/types"
	github_com_cosmos_cosmos_sdk_types "github.com/cosmos/cosmos-sdk/types"
	_ "github.com/gogo/protobuf/gogoproto"
	proto "github.com/gogo/protobuf/proto"
	_ "google.golang.org/protobuf/types/known/durationpb"
	_ "google.golang.org/protobuf/types/known/timestamppb"
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

// Minter represents the minting state.
type Minter struct {
	// last mint time
	LastMintTime int64 `protobuf:"varint,1,opt,name=last_mint_time,json=lastMintTime,proto3" json:"last_mint_time,omitempty"`
	// current daily provisions
	DailyProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=daily_provisions,json=dailyProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"daily_provisions"`
}

func (m *Minter) Reset()         { *m = Minter{} }
func (m *Minter) String() string { return proto.CompactTextString(m) }
func (*Minter) ProtoMessage()    {}
func (*Minter) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa6c02acf2a0105, []int{0}
}
func (m *Minter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Minter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Minter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Minter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Minter.Merge(m, src)
}
func (m *Minter) XXX_Size() int {
	return m.Size()
}
func (m *Minter) XXX_DiscardUnknown() {
	xxx_messageInfo_Minter.DiscardUnknown(m)
}

var xxx_messageInfo_Minter proto.InternalMessageInfo

func (m *Minter) GetLastMintTime() int64 {
	if m != nil {
		return m.LastMintTime
	}
	return 0
}

type DistributionProportions struct {
	// staking defines the proportion of the minted minted_denom that is to be
	// allocated as staking rewards.
	Staking github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,1,opt,name=staking,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"staking"`
}

func (m *DistributionProportions) Reset()         { *m = DistributionProportions{} }
func (m *DistributionProportions) String() string { return proto.CompactTextString(m) }
func (*DistributionProportions) ProtoMessage()    {}
func (*DistributionProportions) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa6c02acf2a0105, []int{1}
}
func (m *DistributionProportions) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *DistributionProportions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_DistributionProportions.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *DistributionProportions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DistributionProportions.Merge(m, src)
}
func (m *DistributionProportions) XXX_Size() int {
	return m.Size()
}
func (m *DistributionProportions) XXX_DiscardUnknown() {
	xxx_messageInfo_DistributionProportions.DiscardUnknown(m)
}

var xxx_messageInfo_DistributionProportions proto.InternalMessageInfo

// Params holds parameters for the mint module.
type Params struct {
	// type of coin to mint
	MintDenom string `protobuf:"bytes,1,opt,name=mint_denom,json=mintDenom,proto3" json:"mint_denom,omitempty"`
	// daily provisions from the genesis
	GenesisDailyProvisions github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,2,opt,name=genesis_daily_provisions,json=genesisDailyProvisions,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"genesis_daily_provisions"`
	// number of seconds take to reduce rewards
	ReductionPeriodInSeconds int64 `protobuf:"varint,3,opt,name=reduction_period_in_seconds,json=reductionPeriodInSeconds,proto3" json:"reduction_period_in_seconds,omitempty"`
	// reduction multiplier to execute on each period
	ReductionFactor github_com_cosmos_cosmos_sdk_types.Dec `protobuf:"bytes,4,opt,name=reduction_factor,json=reductionFactor,proto3,customtype=github.com/cosmos/cosmos-sdk/types.Dec" json:"reduction_factor"`
	// distribution_proportions defines the proportion of the minted denom
	DistributionProportions DistributionProportions `protobuf:"bytes,5,opt,name=distribution_proportions,json=distributionProportions,proto3" json:"distribution_proportions"`
	// next block to reduce minting rewards
	NextRewardsReductionTime int64 `protobuf:"varint,6,opt,name=next_rewards_reduction_time,json=nextRewardsReductionTime,proto3" json:"next_rewards_reduction_time,omitempty"`
}

func (m *Params) Reset()      { *m = Params{} }
func (*Params) ProtoMessage() {}
func (*Params) Descriptor() ([]byte, []int) {
	return fileDescriptor_2fa6c02acf2a0105, []int{2}
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

func (m *Params) GetMintDenom() string {
	if m != nil {
		return m.MintDenom
	}
	return ""
}

func (m *Params) GetReductionPeriodInSeconds() int64 {
	if m != nil {
		return m.ReductionPeriodInSeconds
	}
	return 0
}

func (m *Params) GetDistributionProportions() DistributionProportions {
	if m != nil {
		return m.DistributionProportions
	}
	return DistributionProportions{}
}

func (m *Params) GetNextRewardsReductionTime() int64 {
	if m != nil {
		return m.NextRewardsReductionTime
	}
	return 0
}

func init() {
	proto.RegisterType((*Minter)(nil), "acrechain.mint.v1beta1.Minter")
	proto.RegisterType((*DistributionProportions)(nil), "acrechain.mint.v1beta1.DistributionProportions")
	proto.RegisterType((*Params)(nil), "acrechain.mint.v1beta1.Params")
}

func init() { proto.RegisterFile("acrechain/mint/v1beta1/mint.proto", fileDescriptor_2fa6c02acf2a0105) }

var fileDescriptor_2fa6c02acf2a0105 = []byte{
	// 507 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xac, 0x93, 0x41, 0x6f, 0xd3, 0x30,
	0x14, 0xc7, 0x13, 0x5a, 0x8a, 0x66, 0x10, 0xa0, 0x08, 0x6d, 0x61, 0x13, 0x69, 0xa9, 0x10, 0xea,
	0x85, 0x84, 0xc1, 0x0d, 0x69, 0x07, 0xaa, 0x0a, 0x01, 0x12, 0x52, 0x14, 0xb8, 0xc0, 0x25, 0x72,
	0x6c, 0x2f, 0xb5, 0x96, 0xf8, 0x45, 0xb6, 0x33, 0xd6, 0x8f, 0xc0, 0x8d, 0x23, 0x47, 0x3e, 0xce,
	0x8e, 0x3b, 0x22, 0x0e, 0x13, 0x6a, 0xbf, 0x08, 0xb2, 0x93, 0xb6, 0x63, 0xa2, 0x07, 0xd0, 0x4e,
	0x89, 0xdf, 0xfb, 0xf9, 0xf9, 0xbd, 0xf7, 0x7f, 0x0f, 0x3d, 0xc4, 0x44, 0x32, 0x32, 0xc5, 0x5c,
	0x44, 0x25, 0x17, 0x3a, 0x3a, 0xde, 0xcf, 0x98, 0xc6, 0xfb, 0xf6, 0x10, 0x56, 0x12, 0x34, 0x78,
	0xdb, 0x2b, 0x24, 0xb4, 0xd6, 0x16, 0xd9, 0xbd, 0x97, 0x43, 0x0e, 0x16, 0x89, 0xcc, 0x5f, 0x43,
	0xef, 0xf6, 0x73, 0x80, 0xbc, 0x60, 0x91, 0x3d, 0x65, 0xf5, 0x61, 0xa4, 0x79, 0xc9, 0x94, 0xc6,
	0x65, 0xd5, 0x02, 0xf7, 0x2f, 0x03, 0x58, 0xcc, 0x5a, 0x57, 0x70, 0xd9, 0x45, 0x6b, 0x89, 0x35,
	0x07, 0xd1, 0xf8, 0x87, 0x5f, 0x5c, 0xd4, 0x7b, 0xc7, 0x85, 0x66, 0xd2, 0x7b, 0x84, 0x6e, 0x17,
	0x58, 0xe9, 0xd4, 0x64, 0x94, 0x9a, 0x27, 0x7c, 0x77, 0xe0, 0x8e, 0x3a, 0xc9, 0x2d, 0x63, 0x35,
	0xcc, 0x07, 0x5e, 0x32, 0xef, 0x23, 0xba, 0x4b, 0x31, 0x2f, 0x66, 0x69, 0x25, 0xe1, 0x98, 0x2b,
	0x0e, 0x42, 0xf9, 0xd7, 0x06, 0xee, 0x68, 0x6b, 0x1c, 0x9e, 0x9e, 0xf7, 0x9d, 0x9f, 0xe7, 0xfd,
	0xc7, 0x39, 0xd7, 0xd3, 0x3a, 0x0b, 0x09, 0x94, 0x11, 0x01, 0x55, 0x82, 0x6a, 0x3f, 0x4f, 0x14,
	0x3d, 0x8a, 0xf4, 0xac, 0x62, 0x2a, 0x9c, 0x30, 0x92, 0xdc, 0xb1, 0x71, 0xe2, 0x55, 0x98, 0x21,
	0x41, 0x3b, 0x13, 0xae, 0xb4, 0xe4, 0x59, 0x6d, 0x32, 0x8c, 0x25, 0x54, 0x20, 0xcd, 0x9f, 0xf2,
	0x5e, 0xa3, 0x1b, 0x4a, 0xe3, 0x23, 0x2e, 0x72, 0x9b, 0xd4, 0xbf, 0x3f, 0xb6, 0xbc, 0x3e, 0x9c,
	0x77, 0x50, 0x2f, 0xc6, 0x12, 0x97, 0xca, 0x7b, 0x80, 0x90, 0xad, 0x95, 0x32, 0x01, 0x65, 0x13,
	0x37, 0xd9, 0x32, 0x96, 0x89, 0x31, 0x78, 0x53, 0xe4, 0xe7, 0x4c, 0x30, 0xc5, 0x55, 0x7a, 0x45,
	0x15, 0x6f, 0xb7, 0xf1, 0x26, 0x7f, 0x16, 0xee, 0x1d, 0xa0, 0x3d, 0xc9, 0x68, 0x4d, 0x4c, 0xad,
	0x69, 0xc5, 0x24, 0x07, 0x9a, 0x72, 0x91, 0x2a, 0x46, 0x40, 0x50, 0xe5, 0x77, 0xac, 0x0c, 0xfe,
	0x0a, 0x89, 0x2d, 0xf1, 0x46, 0xbc, 0x6f, 0xfc, 0x46, 0x92, 0xf5, 0xf5, 0x43, 0x4c, 0x34, 0x48,
	0xbf, 0xfb, 0x7f, 0x92, 0xac, 0xe2, 0xbc, 0xb2, 0x61, 0xbc, 0x0a, 0xf9, 0xf4, 0x82, 0x24, 0xa6,
	0x05, 0x4b, 0x4d, 0xfc, 0xeb, 0x03, 0x77, 0x74, 0xf3, 0x59, 0x14, 0xfe, 0x7d, 0x96, 0xc3, 0x0d,
	0x52, 0x8e, 0xbb, 0x26, 0xa7, 0x64, 0x87, 0x6e, 0x50, 0xfa, 0x00, 0xed, 0x09, 0x76, 0xa2, 0x53,
	0xc9, 0x3e, 0x63, 0x49, 0x55, 0xba, 0xae, 0xcc, 0x8e, 0x64, 0xaf, 0xe9, 0x85, 0x41, 0x92, 0x86,
	0x48, 0x96, 0x80, 0x19, 0xcf, 0x17, 0xdd, 0x6f, 0xdf, 0xfb, 0xce, 0xf8, 0xed, 0xe9, 0x3c, 0x70,
	0xcf, 0xe6, 0x81, 0xfb, 0x6b, 0x1e, 0xb8, 0x5f, 0x17, 0x81, 0x73, 0xb6, 0x08, 0x9c, 0x1f, 0x8b,
	0xc0, 0xf9, 0xf4, 0xf4, 0x42, 0x27, 0x5e, 0x4a, 0x9c, 0x15, 0x2c, 0x36, 0x7b, 0x40, 0xa0, 0x88,
	0xd6, 0x6b, 0x7b, 0xd2, 0x2c, 0xae, 0xed, 0x4b, 0xd6, 0xb3, 0x8b, 0xf2, 0xfc, 0x77, 0x00, 0x00,
	0x00, 0xff, 0xff, 0xbc, 0xa6, 0xff, 0xca, 0xd7, 0x03, 0x00, 0x00,
}

func (m *Minter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Minter) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *Minter) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.DailyProvisions.Size()
		i -= size
		if _, err := m.DailyProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if m.LastMintTime != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.LastMintTime))
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *DistributionProportions) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *DistributionProportions) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *DistributionProportions) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	{
		size := m.Staking.Size()
		i -= size
		if _, err := m.Staking.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
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
	if m.NextRewardsReductionTime != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.NextRewardsReductionTime))
		i--
		dAtA[i] = 0x30
	}
	{
		size, err := m.DistributionProportions.MarshalToSizedBuffer(dAtA[:i])
		if err != nil {
			return 0, err
		}
		i -= size
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x2a
	{
		size := m.ReductionFactor.Size()
		i -= size
		if _, err := m.ReductionFactor.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x22
	if m.ReductionPeriodInSeconds != 0 {
		i = encodeVarintMint(dAtA, i, uint64(m.ReductionPeriodInSeconds))
		i--
		dAtA[i] = 0x18
	}
	{
		size := m.GenesisDailyProvisions.Size()
		i -= size
		if _, err := m.GenesisDailyProvisions.MarshalTo(dAtA[i:]); err != nil {
			return 0, err
		}
		i = encodeVarintMint(dAtA, i, uint64(size))
	}
	i--
	dAtA[i] = 0x12
	if len(m.MintDenom) > 0 {
		i -= len(m.MintDenom)
		copy(dAtA[i:], m.MintDenom)
		i = encodeVarintMint(dAtA, i, uint64(len(m.MintDenom)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintMint(dAtA []byte, offset int, v uint64) int {
	offset -= sovMint(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *Minter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.LastMintTime != 0 {
		n += 1 + sovMint(uint64(m.LastMintTime))
	}
	l = m.DailyProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *DistributionProportions) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = m.Staking.Size()
	n += 1 + l + sovMint(uint64(l))
	return n
}

func (m *Params) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.MintDenom)
	if l > 0 {
		n += 1 + l + sovMint(uint64(l))
	}
	l = m.GenesisDailyProvisions.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.ReductionPeriodInSeconds != 0 {
		n += 1 + sovMint(uint64(m.ReductionPeriodInSeconds))
	}
	l = m.ReductionFactor.Size()
	n += 1 + l + sovMint(uint64(l))
	l = m.DistributionProportions.Size()
	n += 1 + l + sovMint(uint64(l))
	if m.NextRewardsReductionTime != 0 {
		n += 1 + sovMint(uint64(m.NextRewardsReductionTime))
	}
	return n
}

func sovMint(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozMint(x uint64) (n int) {
	return sovMint(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *Minter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: Minter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Minter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field LastMintTime", wireType)
			}
			m.LastMintTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.LastMintTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DailyProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DailyProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *DistributionProportions) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
			return fmt.Errorf("proto: DistributionProportions: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: DistributionProportions: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Staking", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.Staking.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func (m *Params) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowMint
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
				return fmt.Errorf("proto: wrong wireType = %d for field MintDenom", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.MintDenom = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field GenesisDailyProvisions", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.GenesisDailyProvisions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionPeriodInSeconds", wireType)
			}
			m.ReductionPeriodInSeconds = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ReductionPeriodInSeconds |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ReductionFactor", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.ReductionFactor.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field DistributionProportions", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
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
				return ErrInvalidLengthMint
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthMint
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if err := m.DistributionProportions.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NextRewardsReductionTime", wireType)
			}
			m.NextRewardsReductionTime = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowMint
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NextRewardsReductionTime |= int64(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		default:
			iNdEx = preIndex
			skippy, err := skipMint(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthMint
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
func skipMint(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
					return 0, ErrIntOverflowMint
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
				return 0, ErrInvalidLengthMint
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupMint
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthMint
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthMint        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowMint          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupMint = fmt.Errorf("proto: unexpected end of group")
)