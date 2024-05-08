// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: babylonchain/babylon/v1beta1/scheduler.proto

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

// ScheduledWork is XX
type ScheduledWork struct {
	Repeat bool `protobuf:"varint,1,opt,name=repeat,proto3" json:"repeat,omitempty"`
}

func (m *ScheduledWork) Reset()         { *m = ScheduledWork{} }
func (m *ScheduledWork) String() string { return proto.CompactTextString(m) }
func (*ScheduledWork) ProtoMessage()    {}
func (*ScheduledWork) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdcaeea24db3b88c, []int{0}
}
func (m *ScheduledWork) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ScheduledWork) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ScheduledWork.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ScheduledWork) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScheduledWork.Merge(m, src)
}
func (m *ScheduledWork) XXX_Size() int {
	return m.Size()
}
func (m *ScheduledWork) XXX_DiscardUnknown() {
	xxx_messageInfo_ScheduledWork.DiscardUnknown(m)
}

var xxx_messageInfo_ScheduledWork proto.InternalMessageInfo

// ValidatorAddress payload data to be used with the scheduler
type ValidatorAddress struct {
	// Address is the ValAddress bech32 string
	Address string `protobuf:"bytes,1,opt,name=address,proto3" json:"address,omitempty"`
}

func (m *ValidatorAddress) Reset()         { *m = ValidatorAddress{} }
func (m *ValidatorAddress) String() string { return proto.CompactTextString(m) }
func (*ValidatorAddress) ProtoMessage()    {}
func (*ValidatorAddress) Descriptor() ([]byte, []int) {
	return fileDescriptor_cdcaeea24db3b88c, []int{1}
}
func (m *ValidatorAddress) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ValidatorAddress) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ValidatorAddress.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ValidatorAddress) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ValidatorAddress.Merge(m, src)
}
func (m *ValidatorAddress) XXX_Size() int {
	return m.Size()
}
func (m *ValidatorAddress) XXX_DiscardUnknown() {
	xxx_messageInfo_ValidatorAddress.DiscardUnknown(m)
}

var xxx_messageInfo_ValidatorAddress proto.InternalMessageInfo

func init() {
	proto.RegisterType((*ScheduledWork)(nil), "babylonchain.babylon.v1beta1.ScheduledWork")
	proto.RegisterType((*ValidatorAddress)(nil), "babylonchain.babylon.v1beta1.ValidatorAddress")
}

func init() {
	proto.RegisterFile("babylonchain/babylon/v1beta1/scheduler.proto", fileDescriptor_cdcaeea24db3b88c)
}

var fileDescriptor_cdcaeea24db3b88c = []byte{
	// 252 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xd2, 0x49, 0x4a, 0x4c, 0xaa,
	0xcc, 0xc9, 0xcf, 0x4b, 0xce, 0x48, 0xcc, 0xcc, 0xd3, 0x87, 0x72, 0xf4, 0xcb, 0x0c, 0x93, 0x52,
	0x4b, 0x12, 0x0d, 0xf5, 0x8b, 0x93, 0x33, 0x52, 0x53, 0x4a, 0x73, 0x52, 0x8b, 0xf4, 0x0a, 0x8a,
	0xf2, 0x4b, 0xf2, 0x85, 0x64, 0x90, 0x55, 0xeb, 0x41, 0x39, 0x7a, 0x50, 0xd5, 0x52, 0x22, 0xe9,
	0xf9, 0xe9, 0xf9, 0x60, 0x85, 0xfa, 0x20, 0x16, 0x44, 0x8f, 0x94, 0x64, 0x72, 0x7e, 0x71, 0x6e,
	0x7e, 0x71, 0x3c, 0x44, 0x02, 0xc2, 0x81, 0x48, 0x29, 0xa9, 0x73, 0xf1, 0x06, 0x43, 0x6d, 0x48,
	0x09, 0xcf, 0x2f, 0xca, 0x16, 0x12, 0xe3, 0x62, 0x2b, 0x4a, 0x2d, 0x48, 0x4d, 0x2c, 0x91, 0x60,
	0x54, 0x60, 0xd4, 0xe0, 0x08, 0x82, 0xf2, 0x94, 0xdc, 0xb8, 0x04, 0xc2, 0x12, 0x73, 0x32, 0x53,
	0x12, 0x4b, 0xf2, 0x8b, 0x1c, 0x53, 0x52, 0x8a, 0x52, 0x8b, 0x8b, 0x85, 0x8c, 0xb8, 0xd8, 0x13,
	0x21, 0x4c, 0xb0, 0x62, 0x4e, 0x27, 0x89, 0x4b, 0x5b, 0x74, 0x45, 0xa0, 0xe6, 0x43, 0x15, 0x05,
	0x97, 0x14, 0x65, 0xe6, 0xa5, 0x07, 0xc1, 0x14, 0x3a, 0x85, 0x9e, 0x78, 0x28, 0xc7, 0xb0, 0xe2,
	0x91, 0x1c, 0xc3, 0x89, 0x47, 0x72, 0x8c, 0x17, 0x1e, 0xc9, 0x31, 0x3e, 0x78, 0x24, 0xc7, 0x38,
	0xe1, 0xb1, 0x1c, 0xc3, 0x85, 0xc7, 0x72, 0x0c, 0x37, 0x1e, 0xcb, 0x31, 0x44, 0x19, 0xa7, 0x67,
	0x96, 0x64, 0x94, 0x26, 0xe9, 0x25, 0xe7, 0xe7, 0xea, 0x63, 0x0b, 0x1a, 0xdd, 0xe2, 0x94, 0x6c,
	0xfd, 0x0a, 0x78, 0x40, 0x95, 0x54, 0x16, 0xa4, 0x16, 0x27, 0xb1, 0x81, 0xbd, 0x63, 0x0c, 0x08,
	0x00, 0x00, 0xff, 0xff, 0x34, 0x7c, 0x99, 0x48, 0x4d, 0x01, 0x00, 0x00,
}

func (m *ScheduledWork) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ScheduledWork) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ScheduledWork) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Repeat {
		i--
		if m.Repeat {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ValidatorAddress) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ValidatorAddress) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ValidatorAddress) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Address) > 0 {
		i -= len(m.Address)
		copy(dAtA[i:], m.Address)
		i = encodeVarintScheduler(dAtA, i, uint64(len(m.Address)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func encodeVarintScheduler(dAtA []byte, offset int, v uint64) int {
	offset -= sovScheduler(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *ScheduledWork) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Repeat {
		n += 2
	}
	return n
}

func (m *ValidatorAddress) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Address)
	if l > 0 {
		n += 1 + l + sovScheduler(uint64(l))
	}
	return n
}

func sovScheduler(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozScheduler(x uint64) (n int) {
	return sovScheduler(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *ScheduledWork) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduler
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
			return fmt.Errorf("proto: ScheduledWork: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ScheduledWork: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Repeat", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduler
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				v |= int(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			m.Repeat = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipScheduler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScheduler
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
func (m *ValidatorAddress) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowScheduler
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
			return fmt.Errorf("proto: ValidatorAddress: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ValidatorAddress: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Address", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowScheduler
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
				return ErrInvalidLengthScheduler
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthScheduler
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Address = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipScheduler(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if (skippy < 0) || (iNdEx+skippy) < 0 {
				return ErrInvalidLengthScheduler
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
func skipScheduler(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowScheduler
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
					return 0, ErrIntOverflowScheduler
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
					return 0, ErrIntOverflowScheduler
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
				return 0, ErrInvalidLengthScheduler
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupScheduler
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthScheduler
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthScheduler        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowScheduler          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupScheduler = fmt.Errorf("proto: unexpected end of group")
)