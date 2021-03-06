// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: GrpcCommon.proto

package sagagrpc

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type GrpcServiceConfig struct {
	ServiceName          string   `protobuf:"bytes,1,opt,name=serviceName,proto3" json:"serviceName,omitempty"`
	InstanceId           string   `protobuf:"bytes,2,opt,name=instanceId,proto3" json:"instanceId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcServiceConfig) Reset()         { *m = GrpcServiceConfig{} }
func (m *GrpcServiceConfig) String() string { return proto.CompactTextString(m) }
func (*GrpcServiceConfig) ProtoMessage()    {}
func (*GrpcServiceConfig) Descriptor() ([]byte, []int) {
	return fileDescriptor_3141a2505e5333d5, []int{0}
}
func (m *GrpcServiceConfig) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcServiceConfig) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcServiceConfig.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcServiceConfig) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcServiceConfig.Merge(m, src)
}
func (m *GrpcServiceConfig) XXX_Size() int {
	return m.Size()
}
func (m *GrpcServiceConfig) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcServiceConfig.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcServiceConfig proto.InternalMessageInfo

func (m *GrpcServiceConfig) GetServiceName() string {
	if m != nil {
		return m.ServiceName
	}
	return ""
}

func (m *GrpcServiceConfig) GetInstanceId() string {
	if m != nil {
		return m.InstanceId
	}
	return ""
}

type GrpcAck struct {
	Aborted              bool     `protobuf:"varint,1,opt,name=aborted,proto3" json:"aborted,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GrpcAck) Reset()         { *m = GrpcAck{} }
func (m *GrpcAck) String() string { return proto.CompactTextString(m) }
func (*GrpcAck) ProtoMessage()    {}
func (*GrpcAck) Descriptor() ([]byte, []int) {
	return fileDescriptor_3141a2505e5333d5, []int{1}
}
func (m *GrpcAck) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *GrpcAck) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_GrpcAck.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *GrpcAck) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GrpcAck.Merge(m, src)
}
func (m *GrpcAck) XXX_Size() int {
	return m.Size()
}
func (m *GrpcAck) XXX_DiscardUnknown() {
	xxx_messageInfo_GrpcAck.DiscardUnknown(m)
}

var xxx_messageInfo_GrpcAck proto.InternalMessageInfo

func (m *GrpcAck) GetAborted() bool {
	if m != nil {
		return m.Aborted
	}
	return false
}

type ServerMeta struct {
	Meta                 map[string]string `protobuf:"bytes,1,rep,name=meta,proto3" json:"meta,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ServerMeta) Reset()         { *m = ServerMeta{} }
func (m *ServerMeta) String() string { return proto.CompactTextString(m) }
func (*ServerMeta) ProtoMessage()    {}
func (*ServerMeta) Descriptor() ([]byte, []int) {
	return fileDescriptor_3141a2505e5333d5, []int{2}
}
func (m *ServerMeta) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *ServerMeta) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_ServerMeta.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalToSizedBuffer(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (m *ServerMeta) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ServerMeta.Merge(m, src)
}
func (m *ServerMeta) XXX_Size() int {
	return m.Size()
}
func (m *ServerMeta) XXX_DiscardUnknown() {
	xxx_messageInfo_ServerMeta.DiscardUnknown(m)
}

var xxx_messageInfo_ServerMeta proto.InternalMessageInfo

func (m *ServerMeta) GetMeta() map[string]string {
	if m != nil {
		return m.Meta
	}
	return nil
}

func init() {
	proto.RegisterType((*GrpcServiceConfig)(nil), "GrpcServiceConfig")
	proto.RegisterType((*GrpcAck)(nil), "GrpcAck")
	proto.RegisterType((*ServerMeta)(nil), "ServerMeta")
	proto.RegisterMapType((map[string]string)(nil), "ServerMeta.MetaEntry")
}

func init() { proto.RegisterFile("GrpcCommon.proto", fileDescriptor_3141a2505e5333d5) }

var fileDescriptor_3141a2505e5333d5 = []byte{
	// 274 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x4c, 0x90, 0xbf, 0x4e, 0xf3, 0x30,
	0x14, 0xc5, 0x3f, 0xb7, 0x1f, 0x94, 0xde, 0x2e, 0xc5, 0x02, 0x29, 0x62, 0x88, 0xa2, 0xb0, 0xb4,
	0x8b, 0x07, 0x18, 0x40, 0x30, 0x41, 0x85, 0x50, 0x07, 0x10, 0x0a, 0x62, 0x61, 0xbb, 0x71, 0x2f,
	0x21, 0x0a, 0xfe, 0x23, 0xc7, 0x54, 0xea, 0x9b, 0xf0, 0x48, 0x8c, 0x3c, 0x02, 0x0a, 0x2f, 0x82,
	0x9c, 0x06, 0xe8, 0x62, 0xf9, 0xfc, 0x7c, 0x74, 0x7c, 0xee, 0x85, 0xf1, 0xb5, 0xb3, 0x72, 0x66,
	0x94, 0x32, 0x5a, 0x58, 0x67, 0xbc, 0x49, 0x1f, 0x60, 0x37, 0xb0, 0x7b, 0x72, 0xcb, 0x52, 0xd2,
	0xcc, 0xe8, 0xa7, 0xb2, 0xe0, 0x09, 0x8c, 0xea, 0x35, 0xb8, 0x45, 0x45, 0x11, 0x4b, 0xd8, 0x64,
	0x98, 0x6d, 0x22, 0x1e, 0x03, 0x94, 0xba, 0xf6, 0xa8, 0x25, 0xcd, 0x17, 0x51, 0xaf, 0x35, 0x6c,
	0x90, 0xf4, 0x10, 0x06, 0x21, 0xf6, 0x42, 0x56, 0x3c, 0x82, 0x01, 0xe6, 0xc6, 0x79, 0x5a, 0xb4,
	0x41, 0x3b, 0xd9, 0x8f, 0x4c, 0x2d, 0x40, 0xf8, 0x97, 0xdc, 0x0d, 0x79, 0xe4, 0x53, 0xf8, 0xaf,
	0xc8, 0x63, 0xc4, 0x92, 0xfe, 0x64, 0x74, 0xb4, 0x2f, 0xfe, 0x9e, 0x44, 0x38, 0xae, 0xb4, 0x77,
	0xab, 0xac, 0xb5, 0x1c, 0x9c, 0xc0, 0xf0, 0x17, 0xf1, 0x31, 0xf4, 0x2b, 0x5a, 0x75, 0x25, 0xc3,
	0x95, 0xef, 0xc1, 0xd6, 0x12, 0x5f, 0x5e, 0xa9, 0xeb, 0xb5, 0x16, 0x67, 0xbd, 0x53, 0x76, 0x39,
	0x7f, 0x6f, 0x62, 0xf6, 0xd1, 0xc4, 0xec, 0xb3, 0x89, 0xd9, 0xdb, 0x57, 0xfc, 0x0f, 0xa6, 0xc6,
	0x15, 0x02, 0x2d, 0xca, 0x67, 0x12, 0xdd, 0x80, 0xd2, 0xa8, 0x5c, 0x58, 0x94, 0x95, 0x90, 0x46,
	0x7b, 0x87, 0xd2, 0x8b, 0xc2, 0x59, 0x79, 0xc7, 0x1e, 0x41, 0x9c, 0xd7, 0x58, 0x60, 0x50, 0xf9,
	0x76, 0xbb, 0xbf, 0xe3, 0xef, 0x00, 0x00, 0x00, 0xff, 0xff, 0x39, 0xf8, 0x48, 0x6e, 0x53, 0x01,
	0x00, 0x00,
}

func (m *GrpcServiceConfig) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcServiceConfig) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcServiceConfig) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.InstanceId) > 0 {
		i -= len(m.InstanceId)
		copy(dAtA[i:], m.InstanceId)
		i = encodeVarintGrpcCommon(dAtA, i, uint64(len(m.InstanceId)))
		i--
		dAtA[i] = 0x12
	}
	if len(m.ServiceName) > 0 {
		i -= len(m.ServiceName)
		copy(dAtA[i:], m.ServiceName)
		i = encodeVarintGrpcCommon(dAtA, i, uint64(len(m.ServiceName)))
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *GrpcAck) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *GrpcAck) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *GrpcAck) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if m.Aborted {
		i--
		if m.Aborted {
			dAtA[i] = 1
		} else {
			dAtA[i] = 0
		}
		i--
		dAtA[i] = 0x8
	}
	return len(dAtA) - i, nil
}

func (m *ServerMeta) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *ServerMeta) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *ServerMeta) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.XXX_unrecognized != nil {
		i -= len(m.XXX_unrecognized)
		copy(dAtA[i:], m.XXX_unrecognized)
	}
	if len(m.Meta) > 0 {
		for k := range m.Meta {
			v := m.Meta[k]
			baseI := i
			i -= len(v)
			copy(dAtA[i:], v)
			i = encodeVarintGrpcCommon(dAtA, i, uint64(len(v)))
			i--
			dAtA[i] = 0x12
			i -= len(k)
			copy(dAtA[i:], k)
			i = encodeVarintGrpcCommon(dAtA, i, uint64(len(k)))
			i--
			dAtA[i] = 0xa
			i = encodeVarintGrpcCommon(dAtA, i, uint64(baseI-i))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGrpcCommon(dAtA []byte, offset int, v uint64) int {
	offset -= sovGrpcCommon(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *GrpcServiceConfig) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.ServiceName)
	if l > 0 {
		n += 1 + l + sovGrpcCommon(uint64(l))
	}
	l = len(m.InstanceId)
	if l > 0 {
		n += 1 + l + sovGrpcCommon(uint64(l))
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *GrpcAck) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Aborted {
		n += 2
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func (m *ServerMeta) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Meta) > 0 {
		for k, v := range m.Meta {
			_ = k
			_ = v
			mapEntrySize := 1 + len(k) + sovGrpcCommon(uint64(len(k))) + 1 + len(v) + sovGrpcCommon(uint64(len(v)))
			n += mapEntrySize + 1 + sovGrpcCommon(uint64(mapEntrySize))
		}
	}
	if m.XXX_unrecognized != nil {
		n += len(m.XXX_unrecognized)
	}
	return n
}

func sovGrpcCommon(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGrpcCommon(x uint64) (n int) {
	return sovGrpcCommon(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (m *GrpcServiceConfig) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcCommon
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
			return fmt.Errorf("proto: GrpcServiceConfig: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcServiceConfig: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ServiceName", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcCommon
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
				return ErrInvalidLengthGrpcCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpcCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ServiceName = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field InstanceId", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcCommon
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
				return ErrInvalidLengthGrpcCommon
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGrpcCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.InstanceId = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpcCommon
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
func (m *GrpcAck) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcCommon
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
			return fmt.Errorf("proto: GrpcAck: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: GrpcAck: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field Aborted", wireType)
			}
			var v int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcCommon
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
			m.Aborted = bool(v != 0)
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpcCommon
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
func (m *ServerMeta) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGrpcCommon
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
			return fmt.Errorf("proto: ServerMeta: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: ServerMeta: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Meta", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGrpcCommon
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
				return ErrInvalidLengthGrpcCommon
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGrpcCommon
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Meta == nil {
				m.Meta = make(map[string]string)
			}
			var mapkey string
			var mapvalue string
			for iNdEx < postIndex {
				entryPreIndex := iNdEx
				var wire uint64
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return ErrIntOverflowGrpcCommon
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
							return ErrIntOverflowGrpcCommon
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
						return ErrInvalidLengthGrpcCommon
					}
					postStringIndexmapkey := iNdEx + intStringLenmapkey
					if postStringIndexmapkey < 0 {
						return ErrInvalidLengthGrpcCommon
					}
					if postStringIndexmapkey > l {
						return io.ErrUnexpectedEOF
					}
					mapkey = string(dAtA[iNdEx:postStringIndexmapkey])
					iNdEx = postStringIndexmapkey
				} else if fieldNum == 2 {
					var stringLenmapvalue uint64
					for shift := uint(0); ; shift += 7 {
						if shift >= 64 {
							return ErrIntOverflowGrpcCommon
						}
						if iNdEx >= l {
							return io.ErrUnexpectedEOF
						}
						b := dAtA[iNdEx]
						iNdEx++
						stringLenmapvalue |= uint64(b&0x7F) << shift
						if b < 0x80 {
							break
						}
					}
					intStringLenmapvalue := int(stringLenmapvalue)
					if intStringLenmapvalue < 0 {
						return ErrInvalidLengthGrpcCommon
					}
					postStringIndexmapvalue := iNdEx + intStringLenmapvalue
					if postStringIndexmapvalue < 0 {
						return ErrInvalidLengthGrpcCommon
					}
					if postStringIndexmapvalue > l {
						return io.ErrUnexpectedEOF
					}
					mapvalue = string(dAtA[iNdEx:postStringIndexmapvalue])
					iNdEx = postStringIndexmapvalue
				} else {
					iNdEx = entryPreIndex
					skippy, err := skipGrpcCommon(dAtA[iNdEx:])
					if err != nil {
						return err
					}
					if skippy < 0 {
						return ErrInvalidLengthGrpcCommon
					}
					if (iNdEx + skippy) > postIndex {
						return io.ErrUnexpectedEOF
					}
					iNdEx += skippy
				}
			}
			m.Meta[mapkey] = mapvalue
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGrpcCommon(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGrpcCommon
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGrpcCommon
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
func skipGrpcCommon(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	depth := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGrpcCommon
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
					return 0, ErrIntOverflowGrpcCommon
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
					return 0, ErrIntOverflowGrpcCommon
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
				return 0, ErrInvalidLengthGrpcCommon
			}
			iNdEx += length
		case 3:
			depth++
		case 4:
			if depth == 0 {
				return 0, ErrUnexpectedEndOfGroupGrpcCommon
			}
			depth--
		case 5:
			iNdEx += 4
		default:
			return 0, fmt.Errorf("proto: illegal wireType %d", wireType)
		}
		if iNdEx < 0 {
			return 0, ErrInvalidLengthGrpcCommon
		}
		if depth == 0 {
			return iNdEx, nil
		}
	}
	return 0, io.ErrUnexpectedEOF
}

var (
	ErrInvalidLengthGrpcCommon        = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGrpcCommon          = fmt.Errorf("proto: integer overflow")
	ErrUnexpectedEndOfGroupGrpcCommon = fmt.Errorf("proto: unexpected end of group")
)
