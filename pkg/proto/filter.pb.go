// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: filter.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type LogicOperator int32

const (
	LogicOperator_UNKNOWN LogicOperator = 0
	LogicOperator_AND     LogicOperator = 1
	LogicOperator_OR      LogicOperator = 2
)

var LogicOperator_name = map[int32]string{
	0: "UNKNOWN",
	1: "AND",
	2: "OR",
}

var LogicOperator_value = map[string]int32{
	"UNKNOWN": 0,
	"AND":     1,
	"OR":      2,
}

func (x LogicOperator) String() string {
	return proto.EnumName(LogicOperator_name, int32(x))
}

func (LogicOperator) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0}
}

type TagFilter struct {
	TagKey               string   `protobuf:"bytes,1,opt,name=tagKey,proto3" json:"tagKey,omitempty"`
	Op                   Operator `protobuf:"varint,2,opt,name=op,proto3,enum=proto.Operator" json:"op,omitempty"`
	TagValue             string   `protobuf:"bytes,3,opt,name=tagValue,proto3" json:"tagValue,omitempty"`
	TagValueItems        []string `protobuf:"bytes,4,rep,name=tagValueItems,proto3" json:"tagValueItems,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *TagFilter) Reset()         { *m = TagFilter{} }
func (m *TagFilter) String() string { return proto.CompactTextString(m) }
func (*TagFilter) ProtoMessage()    {}
func (*TagFilter) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{0}
}
func (m *TagFilter) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *TagFilter) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_TagFilter.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *TagFilter) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TagFilter.Merge(m, src)
}
func (m *TagFilter) XXX_Size() int {
	return m.Size()
}
func (m *TagFilter) XXX_DiscardUnknown() {
	xxx_messageInfo_TagFilter.DiscardUnknown(m)
}

var xxx_messageInfo_TagFilter proto.InternalMessageInfo

func (m *TagFilter) GetTagKey() string {
	if m != nil {
		return m.TagKey
	}
	return ""
}

func (m *TagFilter) GetOp() Operator {
	if m != nil {
		return m.Op
	}
	return Operator_EQUAL
}

func (m *TagFilter) GetTagValue() string {
	if m != nil {
		return m.TagValue
	}
	return ""
}

func (m *TagFilter) GetTagValueItems() []string {
	if m != nil {
		return m.TagValueItems
	}
	return nil
}

type Filters struct {
	Operator             LogicOperator `protobuf:"varint,1,opt,name=operator,proto3,enum=proto.LogicOperator" json:"operator,omitempty"`
	TagFilters           []*TagFilter  `protobuf:"bytes,2,rep,name=tagFilters,proto3" json:"tagFilters,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Filters) Reset()         { *m = Filters{} }
func (m *Filters) String() string { return proto.CompactTextString(m) }
func (*Filters) ProtoMessage()    {}
func (*Filters) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{1}
}
func (m *Filters) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Filters) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Filters.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Filters) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Filters.Merge(m, src)
}
func (m *Filters) XXX_Size() int {
	return m.Size()
}
func (m *Filters) XXX_DiscardUnknown() {
	xxx_messageInfo_Filters.DiscardUnknown(m)
}

var xxx_messageInfo_Filters proto.InternalMessageInfo

func (m *Filters) GetOperator() LogicOperator {
	if m != nil {
		return m.Operator
	}
	return LogicOperator_UNKNOWN
}

func (m *Filters) GetTagFilters() []*TagFilter {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

type Condition struct {
	Operator             LogicOperator `protobuf:"varint,1,opt,name=operator,proto3,enum=proto.LogicOperator" json:"operator,omitempty"`
	TagFilters           []*TagFilter  `protobuf:"bytes,2,rep,name=tagFilters,proto3" json:"tagFilters,omitempty"`
	Condition            []*Condition  `protobuf:"bytes,3,rep,name=condition,proto3" json:"condition,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Condition) Reset()         { *m = Condition{} }
func (m *Condition) String() string { return proto.CompactTextString(m) }
func (*Condition) ProtoMessage()    {}
func (*Condition) Descriptor() ([]byte, []int) {
	return fileDescriptor_1f5303cab7a20d6f, []int{2}
}
func (m *Condition) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *Condition) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_Condition.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *Condition) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Condition.Merge(m, src)
}
func (m *Condition) XXX_Size() int {
	return m.Size()
}
func (m *Condition) XXX_DiscardUnknown() {
	xxx_messageInfo_Condition.DiscardUnknown(m)
}

var xxx_messageInfo_Condition proto.InternalMessageInfo

func (m *Condition) GetOperator() LogicOperator {
	if m != nil {
		return m.Operator
	}
	return LogicOperator_UNKNOWN
}

func (m *Condition) GetTagFilters() []*TagFilter {
	if m != nil {
		return m.TagFilters
	}
	return nil
}

func (m *Condition) GetCondition() []*Condition {
	if m != nil {
		return m.Condition
	}
	return nil
}

func init() {
	proto.RegisterEnum("proto.LogicOperator", LogicOperator_name, LogicOperator_value)
	proto.RegisterType((*TagFilter)(nil), "proto.TagFilter")
	proto.RegisterType((*Filters)(nil), "proto.Filters")
	proto.RegisterType((*Condition)(nil), "proto.Condition")
}

func init() { proto.RegisterFile("filter.proto", fileDescriptor_1f5303cab7a20d6f) }

var fileDescriptor_1f5303cab7a20d6f = []byte{
	// 296 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0xe2, 0x49, 0xcb, 0xcc, 0x29,
	0x49, 0x2d, 0xd2, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x05, 0x53, 0x52, 0x3c, 0xc9, 0xf9,
	0xb9, 0xb9, 0xf9, 0x79, 0x10, 0x41, 0xa5, 0x36, 0x46, 0x2e, 0xce, 0x90, 0xc4, 0x74, 0x37, 0xb0,
	0x42, 0x21, 0x31, 0x2e, 0xb6, 0x92, 0xc4, 0x74, 0xef, 0xd4, 0x4a, 0x09, 0x46, 0x05, 0x46, 0x0d,
	0xce, 0x20, 0x28, 0x4f, 0x48, 0x9e, 0x8b, 0x29, 0xbf, 0x40, 0x82, 0x49, 0x81, 0x51, 0x83, 0xcf,
	0x88, 0x1f, 0xa2, 0x53, 0xcf, 0xbf, 0x20, 0xb5, 0x28, 0xb1, 0x24, 0xbf, 0x28, 0x88, 0x29, 0xbf,
	0x40, 0x48, 0x8a, 0x8b, 0xa3, 0x24, 0x31, 0x3d, 0x2c, 0x31, 0xa7, 0x34, 0x55, 0x82, 0x19, 0xac,
	0x15, 0xce, 0x17, 0x52, 0xe1, 0xe2, 0x85, 0xb1, 0x3d, 0x4b, 0x52, 0x73, 0x8b, 0x25, 0x58, 0x14,
	0x98, 0x35, 0x38, 0x83, 0x50, 0x05, 0x95, 0x72, 0xb9, 0xd8, 0x21, 0x8e, 0x28, 0x16, 0x32, 0xe0,
	0xe2, 0xc8, 0x87, 0x1a, 0x0e, 0x76, 0x07, 0x9f, 0x91, 0x08, 0xd4, 0x4e, 0x9f, 0xfc, 0xf4, 0xcc,
	0x64, 0xb8, 0xc5, 0x70, 0x55, 0x42, 0x06, 0x5c, 0x5c, 0x25, 0x30, 0x4f, 0x14, 0x4b, 0x30, 0x29,
	0x30, 0x6b, 0x70, 0x1b, 0x09, 0x40, 0xf5, 0xc0, 0x7d, 0x17, 0x84, 0xa4, 0x46, 0x69, 0x3e, 0x23,
	0x17, 0xa7, 0x73, 0x7e, 0x5e, 0x4a, 0x66, 0x49, 0x66, 0x7e, 0x1e, 0x3d, 0x6c, 0x14, 0xd2, 0xe3,
	0xe2, 0x4c, 0x86, 0x59, 0x28, 0xc1, 0x8c, 0xa2, 0x01, 0xee, 0x90, 0x20, 0x84, 0x12, 0x2d, 0x5d,
	0x2e, 0x5e, 0x14, 0xcb, 0x85, 0xb8, 0xb9, 0xd8, 0x43, 0xfd, 0xbc, 0xfd, 0xfc, 0xc3, 0xfd, 0x04,
	0x18, 0x84, 0xd8, 0xb9, 0x98, 0x1d, 0xfd, 0x5c, 0x04, 0x18, 0x85, 0xd8, 0xb8, 0x98, 0xfc, 0x83,
	0x04, 0x98, 0x9c, 0x04, 0x4e, 0x3c, 0x92, 0x63, 0xbc, 0xf0, 0x48, 0x8e, 0xf1, 0xc1, 0x23, 0x39,
	0xc6, 0x19, 0x8f, 0xe5, 0x18, 0x92, 0xd8, 0xc0, 0x86, 0x1b, 0x03, 0x02, 0x00, 0x00, 0xff, 0xff,
	0x05, 0x29, 0xd2, 0x25, 0x06, 0x02, 0x00, 0x00,
}

func (m *TagFilter) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *TagFilter) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.TagKey) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.TagKey)))
		i += copy(dAtA[i:], m.TagKey)
	}
	if m.Op != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintFilter(dAtA, i, uint64(m.Op))
	}
	if len(m.TagValue) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintFilter(dAtA, i, uint64(len(m.TagValue)))
		i += copy(dAtA[i:], m.TagValue)
	}
	if len(m.TagValueItems) > 0 {
		for _, s := range m.TagValueItems {
			dAtA[i] = 0x22
			i++
			l = len(s)
			for l >= 1<<7 {
				dAtA[i] = uint8(uint64(l)&0x7f | 0x80)
				l >>= 7
				i++
			}
			dAtA[i] = uint8(l)
			i++
			i += copy(dAtA[i:], s)
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Filters) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Filters) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Operator != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFilter(dAtA, i, uint64(m.Operator))
	}
	if len(m.TagFilters) > 0 {
		for _, msg := range m.TagFilters {
			dAtA[i] = 0x12
			i++
			i = encodeVarintFilter(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func (m *Condition) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *Condition) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if m.Operator != 0 {
		dAtA[i] = 0x8
		i++
		i = encodeVarintFilter(dAtA, i, uint64(m.Operator))
	}
	if len(m.TagFilters) > 0 {
		for _, msg := range m.TagFilters {
			dAtA[i] = 0x12
			i++
			i = encodeVarintFilter(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if len(m.Condition) > 0 {
		for _, msg := range m.Condition {
			dAtA[i] = 0x1a
			i++
			i = encodeVarintFilter(dAtA, i, uint64(msg.Size()))
			n, err := msg.MarshalTo(dAtA[i:])
			if err != nil {
				return 0, err
			}
			i += n
		}
	}
	if m.XXX_unrecognized != nil {
		i += copy(dAtA[i:], m.XXX_unrecognized)
	}
	return i, nil
}

func encodeVarintFilter(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *TagFilter) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.TagKey)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	if m.Op != 0 {
		n += 1 + sovFilter(uint64(m.Op))
	}
	l = len(m.TagValue)
	if l > 0 {
		n += 1 + l + sovFilter(uint64(l))
	}
	if len(m.TagValueItems) > 0 {
		for _, s := range m.TagValueItems {
			l = len(s)
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Filters) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Operator != 0 {
		n += 1 + sovFilter(uint64(m.Operator))
	}
	if len(m.TagFilters) > 0 {
		for _, e := range m.TagFilters {
			l = e.Size()
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *Condition) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Operator != 0 {
		n += 1 + sovFilter(uint64(m.Operator))
	}
	if len(m.TagFilters) > 0 {
		for _, e := range m.TagFilters {
			l = e.Size()
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if len(m.Condition) > 0 {
		for _, e := range m.Condition {
			l = e.Size()
			n += 1 + l + sovFilter(uint64(l))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovFilter(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozFilter(x uint64) (n int) {
	return sovFilter(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *TagFilter) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilter
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
			return fmt.Errorf("proto: TagFilter: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: TagFilter: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagKey", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
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
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagKey = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Op", wireType)
			}
			m.Op = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Op |= Operator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagValue", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
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
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagValue = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagValueItems", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
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
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagValueItems = append(m.TagValueItems, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Filters) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilter
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
			return fmt.Errorf("proto: Filters: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Filters: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			m.Operator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operator |= LogicOperator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagFilters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
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
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagFilters = append(m.TagFilters, &TagFilter{})
			if err := m.TagFilters[len(m.TagFilters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func (m *Condition) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowFilter
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
			return fmt.Errorf("proto: Condition: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: Condition: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Operator", wireType)
			}
			m.Operator = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.Operator |= LogicOperator(b&0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field TagFilters", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
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
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.TagFilters = append(m.TagFilters, &TagFilter{})
			if err := m.TagFilters[len(m.TagFilters)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Condition", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowFilter
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
				return ErrInvalidLengthFilter
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthFilter
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Condition = append(m.Condition, &Condition{})
			if err := m.Condition[len(m.Condition)-1].Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipFilter(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthFilter
			}
			if (iNdEx + skippy) > l {
				return io.ErrUnexpectedEOF
			}
			m.XXX_unrecognized = append(m.XXX_unrecognized, dAtA[iNdEx:iNdEx+skippy]...)
			iNdEx += skippy
		}
	}

	if iNdEx > l {
		return io.ErrUnexpectedEOF
	}
	return nil
}
func skipFilter(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowFilter
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
					return 0, ErrIntOverflowFilter
				}
				if iNdEx >= l {
					return 0, io.ErrUnexpectedEOF
				}
				iNdEx++
				if dAtA[iNdEx-1] < 0x80 {
					break
				}
			}
			return iNdEx, nil
		case 1:
			iNdEx += 8
			return iNdEx, nil
		case 2:
			var length int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return 0, ErrIntOverflowFilter
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
				return 0, ErrInvalidLengthFilter
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthFilter
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowFilter
					}
					if iNdEx >= l {
						return 0, io.ErrUnexpectedEOF
					}
					b := dAtA[iNdEx]
					iNdEx++
					innerWire |= (uint64(b) & 0x7F) << shift
					if b < 0x80 {
						break
					}
				}
				innerWireType := int(innerWire & 0x7)
				if innerWireType == 4 {
					break
				}
				next, err := skipFilter(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthFilter
				}
			}
			return iNdEx, nil
		case 4:
			return iNdEx, nil
		case 5:
			iNdEx += 4
			return iNdEx, nil
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
	}
	panic("unreachable")
}

var (
	ErrInvalidLengthFilter = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowFilter   = fmt.Errorf("proto: integer overflow")
)
