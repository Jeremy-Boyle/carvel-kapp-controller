// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: github.com/vmware-tanzu/carvel-vendir/pkg/vendir/versions/v1alpha1/generated.proto

package v1alpha1

import (
	fmt "fmt"

	io "io"
	math "math"
	math_bits "math/bits"
	reflect "reflect"
	strings "strings"

	proto "github.com/gogo/protobuf/proto"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

func (m *VersionSelection) Reset()      { *m = VersionSelection{} }
func (*VersionSelection) ProtoMessage() {}
func (*VersionSelection) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d3ef80c738fc456, []int{0}
}
func (m *VersionSelection) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionSelection) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VersionSelection) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionSelection.Merge(m, src)
}
func (m *VersionSelection) XXX_Size() int {
	return m.Size()
}
func (m *VersionSelection) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionSelection.DiscardUnknown(m)
}

var xxx_messageInfo_VersionSelection proto.InternalMessageInfo

func (m *VersionSelectionSemver) Reset()      { *m = VersionSelectionSemver{} }
func (*VersionSelectionSemver) ProtoMessage() {}
func (*VersionSelectionSemver) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d3ef80c738fc456, []int{1}
}
func (m *VersionSelectionSemver) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionSelectionSemver) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VersionSelectionSemver) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionSelectionSemver.Merge(m, src)
}
func (m *VersionSelectionSemver) XXX_Size() int {
	return m.Size()
}
func (m *VersionSelectionSemver) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionSelectionSemver.DiscardUnknown(m)
}

var xxx_messageInfo_VersionSelectionSemver proto.InternalMessageInfo

func (m *VersionSelectionSemverPrereleases) Reset()      { *m = VersionSelectionSemverPrereleases{} }
func (*VersionSelectionSemverPrereleases) ProtoMessage() {}
func (*VersionSelectionSemverPrereleases) Descriptor() ([]byte, []int) {
	return fileDescriptor_8d3ef80c738fc456, []int{2}
}
func (m *VersionSelectionSemverPrereleases) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *VersionSelectionSemverPrereleases) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	b = b[:cap(b)]
	n, err := m.MarshalToSizedBuffer(b)
	if err != nil {
		return nil, err
	}
	return b[:n], nil
}
func (m *VersionSelectionSemverPrereleases) XXX_Merge(src proto.Message) {
	xxx_messageInfo_VersionSelectionSemverPrereleases.Merge(m, src)
}
func (m *VersionSelectionSemverPrereleases) XXX_Size() int {
	return m.Size()
}
func (m *VersionSelectionSemverPrereleases) XXX_DiscardUnknown() {
	xxx_messageInfo_VersionSelectionSemverPrereleases.DiscardUnknown(m)
}

var xxx_messageInfo_VersionSelectionSemverPrereleases proto.InternalMessageInfo

func init() {
	proto.RegisterType((*VersionSelection)(nil), "github.com.vmware_tanzu.carvel_vendir.pkg.vendir.versions.v1alpha1.VersionSelection")
	proto.RegisterType((*VersionSelectionSemver)(nil), "github.com.vmware_tanzu.carvel_vendir.pkg.vendir.versions.v1alpha1.VersionSelectionSemver")
	proto.RegisterType((*VersionSelectionSemverPrereleases)(nil), "github.com.vmware_tanzu.carvel_vendir.pkg.vendir.versions.v1alpha1.VersionSelectionSemverPrereleases")
}

func init() {
	proto.RegisterFile("github.com/vmware-tanzu/carvel-vendir/pkg/vendir/versions/v1alpha1/generated.proto", fileDescriptor_8d3ef80c738fc456)
}

var fileDescriptor_8d3ef80c738fc456 = []byte{
	// 354 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x8f, 0x31, 0x4b, 0xf3, 0x40,
	0x1c, 0xc6, 0x73, 0xef, 0x0b, 0xc5, 0x5e, 0x06, 0x25, 0x82, 0x14, 0x87, 0xab, 0x76, 0xea, 0xd2,
	0x0b, 0x15, 0xfc, 0x02, 0x71, 0x72, 0x93, 0x14, 0x3a, 0x74, 0x29, 0xd7, 0xf4, 0xdf, 0xf4, 0x68,
	0x7a, 0x17, 0x2e, 0xd7, 0x08, 0x0e, 0xe2, 0x47, 0x10, 0x27, 0x3f, 0x52, 0xc7, 0x8e, 0x9d, 0x8a,
	0x8d, 0x9b, 0x9f, 0x42, 0x7a, 0xd7, 0x92, 0x20, 0x82, 0x83, 0xb8, 0x3d, 0xf7, 0xbf, 0x87, 0xdf,
	0xf3, 0x3c, 0x38, 0x8c, 0xb9, 0x9e, 0x2e, 0x46, 0x34, 0x92, 0x73, 0x3f, 0x9f, 0xdf, 0x33, 0x05,
	0x1d, 0xcd, 0xc4, 0xc3, 0xc2, 0x8f, 0x98, 0xca, 0x21, 0xe9, 0xe4, 0x20, 0xc6, 0x5c, 0xf9, 0xe9,
	0x2c, 0xf6, 0xf7, 0x32, 0x07, 0x95, 0x71, 0x29, 0x32, 0x3f, 0xef, 0xb2, 0x24, 0x9d, 0xb2, 0xae,
	0x1f, 0x83, 0x00, 0xc5, 0x34, 0x8c, 0x69, 0xaa, 0xa4, 0x96, 0x5e, 0x50, 0x32, 0xa9, 0x65, 0x0e,
	0x0d, 0x93, 0x5a, 0xe6, 0xd0, 0x82, 0x68, 0x3a, 0x8b, 0xe9, 0x5e, 0x1e, 0x98, 0xf4, 0xc0, 0x3c,
	0xef, 0x54, 0x7a, 0xc5, 0x32, 0x96, 0xbe, 0x41, 0x8f, 0x16, 0x13, 0xf3, 0x32, 0x0f, 0xa3, 0x6c,
	0x64, 0xeb, 0x05, 0xe1, 0x93, 0xbe, 0x85, 0xf4, 0x20, 0x81, 0x48, 0x73, 0x29, 0xbc, 0x47, 0x5c,
	0xcb, 0x60, 0x9e, 0x83, 0x6a, 0xa0, 0x0b, 0xd4, 0x76, 0xaf, 0x06, 0xf4, 0xf7, 0xc5, 0xe8, 0xd7,
	0x94, 0x9e, 0x49, 0x08, 0x70, 0xb1, 0x69, 0xd6, 0xac, 0x0e, 0xf7, 0xa9, 0xad, 0x0f, 0x84, 0xcf,
	0xbe, 0xb7, 0x7b, 0xd7, 0xd8, 0x8d, 0xa4, 0xc8, 0xb4, 0x62, 0x5c, 0xe8, 0xcc, 0xf4, 0xab, 0x07,
	0xa7, 0xcb, 0x4d, 0xd3, 0x29, 0x36, 0x4d, 0xf7, 0xa6, 0xfc, 0x0a, 0xab, 0x3e, 0xef, 0x15, 0x61,
	0x37, 0x55, 0xa0, 0x20, 0x01, 0x96, 0x41, 0xd6, 0xf8, 0x67, 0x76, 0xc1, 0xdf, 0xed, 0xba, 0x2b,
	0xc3, 0x82, 0xe3, 0x5d, 0xb5, 0xca, 0x21, 0xac, 0x56, 0x69, 0xf5, 0xf1, 0xe5, 0x8f, 0x08, 0xaf,
	0x8b, 0x5d, 0x3e, 0x06, 0xa1, 0xf9, 0x84, 0x83, 0xda, 0xcd, 0xfe, 0xdf, 0xae, 0x5b, 0xee, 0x6d,
	0x79, 0x0e, 0xab, 0x9e, 0x80, 0x2e, 0xb7, 0xc4, 0x59, 0x6d, 0x89, 0xb3, 0xde, 0x12, 0xe7, 0xa9,
	0x20, 0x68, 0x59, 0x10, 0xb4, 0x2a, 0x08, 0x5a, 0x17, 0x04, 0xbd, 0x15, 0x04, 0x3d, 0xbf, 0x13,
	0x67, 0x70, 0x74, 0xd8, 0xf1, 0x19, 0x00, 0x00, 0xff, 0xff, 0x97, 0x30, 0x55, 0x40, 0xd1, 0x02,
	0x00, 0x00,
}

func (m *VersionSelection) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionSelection) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionSelection) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Semver != nil {
		{
			size, err := m.Semver.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0xa
	}
	return len(dAtA) - i, nil
}

func (m *VersionSelectionSemver) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionSelectionSemver) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionSelectionSemver) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if m.Prereleases != nil {
		{
			size, err := m.Prereleases.MarshalToSizedBuffer(dAtA[:i])
			if err != nil {
				return 0, err
			}
			i -= size
			i = encodeVarintGenerated(dAtA, i, uint64(size))
		}
		i--
		dAtA[i] = 0x12
	}
	i -= len(m.Constraints)
	copy(dAtA[i:], m.Constraints)
	i = encodeVarintGenerated(dAtA, i, uint64(len(m.Constraints)))
	i--
	dAtA[i] = 0xa
	return len(dAtA) - i, nil
}

func (m *VersionSelectionSemverPrereleases) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalToSizedBuffer(dAtA[:size])
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *VersionSelectionSemverPrereleases) MarshalTo(dAtA []byte) (int, error) {
	size := m.Size()
	return m.MarshalToSizedBuffer(dAtA[:size])
}

func (m *VersionSelectionSemverPrereleases) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	i := len(dAtA)
	_ = i
	var l int
	_ = l
	if len(m.Identifiers) > 0 {
		for iNdEx := len(m.Identifiers) - 1; iNdEx >= 0; iNdEx-- {
			i -= len(m.Identifiers[iNdEx])
			copy(dAtA[i:], m.Identifiers[iNdEx])
			i = encodeVarintGenerated(dAtA, i, uint64(len(m.Identifiers[iNdEx])))
			i--
			dAtA[i] = 0xa
		}
	}
	return len(dAtA) - i, nil
}

func encodeVarintGenerated(dAtA []byte, offset int, v uint64) int {
	offset -= sovGenerated(v)
	base := offset
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return base
}
func (m *VersionSelection) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if m.Semver != nil {
		l = m.Semver.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *VersionSelectionSemver) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	l = len(m.Constraints)
	n += 1 + l + sovGenerated(uint64(l))
	if m.Prereleases != nil {
		l = m.Prereleases.Size()
		n += 1 + l + sovGenerated(uint64(l))
	}
	return n
}

func (m *VersionSelectionSemverPrereleases) Size() (n int) {
	if m == nil {
		return 0
	}
	var l int
	_ = l
	if len(m.Identifiers) > 0 {
		for _, s := range m.Identifiers {
			l = len(s)
			n += 1 + l + sovGenerated(uint64(l))
		}
	}
	return n
}

func sovGenerated(x uint64) (n int) {
	return (math_bits.Len64(x|1) + 6) / 7
}
func sozGenerated(x uint64) (n int) {
	return sovGenerated(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *VersionSelection) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VersionSelection{`,
		`Semver:` + strings.Replace(this.Semver.String(), "VersionSelectionSemver", "VersionSelectionSemver", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VersionSelectionSemver) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VersionSelectionSemver{`,
		`Constraints:` + fmt.Sprintf("%v", this.Constraints) + `,`,
		`Prereleases:` + strings.Replace(this.Prereleases.String(), "VersionSelectionSemverPrereleases", "VersionSelectionSemverPrereleases", 1) + `,`,
		`}`,
	}, "")
	return s
}
func (this *VersionSelectionSemverPrereleases) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&VersionSelectionSemverPrereleases{`,
		`Identifiers:` + fmt.Sprintf("%v", this.Identifiers) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringGenerated(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *VersionSelection) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VersionSelection: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionSelection: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Semver", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Semver == nil {
				m.Semver = &VersionSelectionSemver{}
			}
			if err := m.Semver.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VersionSelectionSemver) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VersionSelectionSemver: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionSelectionSemver: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Constraints", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Constraints = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Prereleases", wireType)
			}
			var msglen int
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + msglen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			if m.Prereleases == nil {
				m.Prereleases = &VersionSelectionSemverPrereleases{}
			}
			if err := m.Prereleases.Unmarshal(dAtA[iNdEx:postIndex]); err != nil {
				return err
			}
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func (m *VersionSelectionSemverPrereleases) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowGenerated
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
			return fmt.Errorf("proto: VersionSelectionSemverPrereleases: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: VersionSelectionSemverPrereleases: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Identifiers", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowGenerated
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
				return ErrInvalidLengthGenerated
			}
			postIndex := iNdEx + intStringLen
			if postIndex < 0 {
				return ErrInvalidLengthGenerated
			}
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Identifiers = append(m.Identifiers, string(dAtA[iNdEx:postIndex]))
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipGenerated(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthGenerated
			}
			if (iNdEx + skippy) < 0 {
				return ErrInvalidLengthGenerated
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
func skipGenerated(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
					return 0, ErrIntOverflowGenerated
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
				return 0, ErrInvalidLengthGenerated
			}
			iNdEx += length
			if iNdEx < 0 {
				return 0, ErrInvalidLengthGenerated
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowGenerated
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
				next, err := skipGenerated(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
				if iNdEx < 0 {
					return 0, ErrInvalidLengthGenerated
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
	ErrInvalidLengthGenerated = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowGenerated   = fmt.Errorf("proto: integer overflow")
)