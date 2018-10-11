// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: services/common/proto/jwt.proto

package jwt // import "pixielabs.ai/pixielabs/services/common/proto"

import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "github.com/gogo/protobuf/gogoproto"

import strings "strings"
import reflect "reflect"

import io "io"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.GoGoProtoPackageIsVersion2 // please upgrade the proto package

type JWTClaims struct {
	Audience             string   `protobuf:"bytes,1,opt,name=audience,proto3" json:"audience,omitempty"`
	ExpiresAt            int64    `protobuf:"varint,2,opt,name=expires_at,json=expiresAt,proto3" json:"expires_at,omitempty"`
	ID                   string   `protobuf:"bytes,3,opt,name=id,proto3" json:"id,omitempty"`
	IssuedAt             int64    `protobuf:"varint,4,opt,name=issued_at,json=issuedAt,proto3" json:"issued_at,omitempty"`
	Issuer               string   `protobuf:"bytes,5,opt,name=issuer,proto3" json:"issuer,omitempty"`
	NotBefore            int64    `protobuf:"varint,6,opt,name=not_before,json=notBefore,proto3" json:"not_before,omitempty"`
	Subject              string   `protobuf:"bytes,7,opt,name=subject,proto3" json:"subject,omitempty"`
	UserID               string   `protobuf:"bytes,8,opt,name=user_id,json=userId,proto3" json:"userID"`
	Email                string   `protobuf:"bytes,9,opt,name=email,proto3" json:"email,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JWTClaims) Reset()      { *m = JWTClaims{} }
func (*JWTClaims) ProtoMessage() {}
func (*JWTClaims) Descriptor() ([]byte, []int) {
	return fileDescriptor_jwt_cd8269b5744a05e9, []int{0}
}
func (m *JWTClaims) XXX_Unmarshal(b []byte) error {
	return m.Unmarshal(b)
}
func (m *JWTClaims) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	if deterministic {
		return xxx_messageInfo_JWTClaims.Marshal(b, m, deterministic)
	} else {
		b = b[:cap(b)]
		n, err := m.MarshalTo(b)
		if err != nil {
			return nil, err
		}
		return b[:n], nil
	}
}
func (dst *JWTClaims) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JWTClaims.Merge(dst, src)
}
func (m *JWTClaims) XXX_Size() int {
	return m.Size()
}
func (m *JWTClaims) XXX_DiscardUnknown() {
	xxx_messageInfo_JWTClaims.DiscardUnknown(m)
}

var xxx_messageInfo_JWTClaims proto.InternalMessageInfo

func (m *JWTClaims) GetAudience() string {
	if m != nil {
		return m.Audience
	}
	return ""
}

func (m *JWTClaims) GetExpiresAt() int64 {
	if m != nil {
		return m.ExpiresAt
	}
	return 0
}

func (m *JWTClaims) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *JWTClaims) GetIssuedAt() int64 {
	if m != nil {
		return m.IssuedAt
	}
	return 0
}

func (m *JWTClaims) GetIssuer() string {
	if m != nil {
		return m.Issuer
	}
	return ""
}

func (m *JWTClaims) GetNotBefore() int64 {
	if m != nil {
		return m.NotBefore
	}
	return 0
}

func (m *JWTClaims) GetSubject() string {
	if m != nil {
		return m.Subject
	}
	return ""
}

func (m *JWTClaims) GetUserID() string {
	if m != nil {
		return m.UserID
	}
	return ""
}

func (m *JWTClaims) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func init() {
	proto.RegisterType((*JWTClaims)(nil), "pl.common.JWTClaims")
}
func (this *JWTClaims) Equal(that interface{}) bool {
	if that == nil {
		return this == nil
	}

	that1, ok := that.(*JWTClaims)
	if !ok {
		that2, ok := that.(JWTClaims)
		if ok {
			that1 = &that2
		} else {
			return false
		}
	}
	if that1 == nil {
		return this == nil
	} else if this == nil {
		return false
	}
	if this.Audience != that1.Audience {
		return false
	}
	if this.ExpiresAt != that1.ExpiresAt {
		return false
	}
	if this.ID != that1.ID {
		return false
	}
	if this.IssuedAt != that1.IssuedAt {
		return false
	}
	if this.Issuer != that1.Issuer {
		return false
	}
	if this.NotBefore != that1.NotBefore {
		return false
	}
	if this.Subject != that1.Subject {
		return false
	}
	if this.UserID != that1.UserID {
		return false
	}
	if this.Email != that1.Email {
		return false
	}
	return true
}
func (this *JWTClaims) GoString() string {
	if this == nil {
		return "nil"
	}
	s := make([]string, 0, 13)
	s = append(s, "&jwt.JWTClaims{")
	s = append(s, "Audience: "+fmt.Sprintf("%#v", this.Audience)+",\n")
	s = append(s, "ExpiresAt: "+fmt.Sprintf("%#v", this.ExpiresAt)+",\n")
	s = append(s, "ID: "+fmt.Sprintf("%#v", this.ID)+",\n")
	s = append(s, "IssuedAt: "+fmt.Sprintf("%#v", this.IssuedAt)+",\n")
	s = append(s, "Issuer: "+fmt.Sprintf("%#v", this.Issuer)+",\n")
	s = append(s, "NotBefore: "+fmt.Sprintf("%#v", this.NotBefore)+",\n")
	s = append(s, "Subject: "+fmt.Sprintf("%#v", this.Subject)+",\n")
	s = append(s, "UserID: "+fmt.Sprintf("%#v", this.UserID)+",\n")
	s = append(s, "Email: "+fmt.Sprintf("%#v", this.Email)+",\n")
	s = append(s, "}")
	return strings.Join(s, "")
}
func valueToGoStringJwt(v interface{}, typ string) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("func(v %v) *%v { return &v } ( %#v )", typ, typ, pv)
}
func (m *JWTClaims) Marshal() (dAtA []byte, err error) {
	size := m.Size()
	dAtA = make([]byte, size)
	n, err := m.MarshalTo(dAtA)
	if err != nil {
		return nil, err
	}
	return dAtA[:n], nil
}

func (m *JWTClaims) MarshalTo(dAtA []byte) (int, error) {
	var i int
	_ = i
	var l int
	_ = l
	if len(m.Audience) > 0 {
		dAtA[i] = 0xa
		i++
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Audience)))
		i += copy(dAtA[i:], m.Audience)
	}
	if m.ExpiresAt != 0 {
		dAtA[i] = 0x10
		i++
		i = encodeVarintJwt(dAtA, i, uint64(m.ExpiresAt))
	}
	if len(m.ID) > 0 {
		dAtA[i] = 0x1a
		i++
		i = encodeVarintJwt(dAtA, i, uint64(len(m.ID)))
		i += copy(dAtA[i:], m.ID)
	}
	if m.IssuedAt != 0 {
		dAtA[i] = 0x20
		i++
		i = encodeVarintJwt(dAtA, i, uint64(m.IssuedAt))
	}
	if len(m.Issuer) > 0 {
		dAtA[i] = 0x2a
		i++
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Issuer)))
		i += copy(dAtA[i:], m.Issuer)
	}
	if m.NotBefore != 0 {
		dAtA[i] = 0x30
		i++
		i = encodeVarintJwt(dAtA, i, uint64(m.NotBefore))
	}
	if len(m.Subject) > 0 {
		dAtA[i] = 0x3a
		i++
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Subject)))
		i += copy(dAtA[i:], m.Subject)
	}
	if len(m.UserID) > 0 {
		dAtA[i] = 0x42
		i++
		i = encodeVarintJwt(dAtA, i, uint64(len(m.UserID)))
		i += copy(dAtA[i:], m.UserID)
	}
	if len(m.Email) > 0 {
		dAtA[i] = 0x4a
		i++
		i = encodeVarintJwt(dAtA, i, uint64(len(m.Email)))
		i += copy(dAtA[i:], m.Email)
	}
	return i, nil
}

func encodeVarintJwt(dAtA []byte, offset int, v uint64) int {
	for v >= 1<<7 {
		dAtA[offset] = uint8(v&0x7f | 0x80)
		v >>= 7
		offset++
	}
	dAtA[offset] = uint8(v)
	return offset + 1
}
func (m *JWTClaims) Size() (n int) {
	var l int
	_ = l
	l = len(m.Audience)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.ExpiresAt != 0 {
		n += 1 + sovJwt(uint64(m.ExpiresAt))
	}
	l = len(m.ID)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.IssuedAt != 0 {
		n += 1 + sovJwt(uint64(m.IssuedAt))
	}
	l = len(m.Issuer)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	if m.NotBefore != 0 {
		n += 1 + sovJwt(uint64(m.NotBefore))
	}
	l = len(m.Subject)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	l = len(m.UserID)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	l = len(m.Email)
	if l > 0 {
		n += 1 + l + sovJwt(uint64(l))
	}
	return n
}

func sovJwt(x uint64) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}
func sozJwt(x uint64) (n int) {
	return sovJwt(uint64((x << 1) ^ uint64((int64(x) >> 63))))
}
func (this *JWTClaims) String() string {
	if this == nil {
		return "nil"
	}
	s := strings.Join([]string{`&JWTClaims{`,
		`Audience:` + fmt.Sprintf("%v", this.Audience) + `,`,
		`ExpiresAt:` + fmt.Sprintf("%v", this.ExpiresAt) + `,`,
		`ID:` + fmt.Sprintf("%v", this.ID) + `,`,
		`IssuedAt:` + fmt.Sprintf("%v", this.IssuedAt) + `,`,
		`Issuer:` + fmt.Sprintf("%v", this.Issuer) + `,`,
		`NotBefore:` + fmt.Sprintf("%v", this.NotBefore) + `,`,
		`Subject:` + fmt.Sprintf("%v", this.Subject) + `,`,
		`UserID:` + fmt.Sprintf("%v", this.UserID) + `,`,
		`Email:` + fmt.Sprintf("%v", this.Email) + `,`,
		`}`,
	}, "")
	return s
}
func valueToStringJwt(v interface{}) string {
	rv := reflect.ValueOf(v)
	if rv.IsNil() {
		return "nil"
	}
	pv := reflect.Indirect(rv).Interface()
	return fmt.Sprintf("*%v", pv)
}
func (m *JWTClaims) Unmarshal(dAtA []byte) error {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		preIndex := iNdEx
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return ErrIntOverflowJwt
			}
			if iNdEx >= l {
				return io.ErrUnexpectedEOF
			}
			b := dAtA[iNdEx]
			iNdEx++
			wire |= (uint64(b) & 0x7F) << shift
			if b < 0x80 {
				break
			}
		}
		fieldNum := int32(wire >> 3)
		wireType := int(wire & 0x7)
		if wireType == 4 {
			return fmt.Errorf("proto: JWTClaims: wiretype end group for non-group")
		}
		if fieldNum <= 0 {
			return fmt.Errorf("proto: JWTClaims: illegal tag %d (wire type %d)", fieldNum, wire)
		}
		switch fieldNum {
		case 1:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Audience", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Audience = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 2:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field ExpiresAt", wireType)
			}
			m.ExpiresAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.ExpiresAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 3:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field ID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.ID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 4:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field IssuedAt", wireType)
			}
			m.IssuedAt = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.IssuedAt |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 5:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Issuer", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Issuer = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 6:
			if wireType != 0 {
				return fmt.Errorf("proto: wrong wireType = %d for field NotBefore", wireType)
			}
			m.NotBefore = 0
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				m.NotBefore |= (int64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
		case 7:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Subject", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Subject = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 8:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field UserID", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.UserID = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		case 9:
			if wireType != 2 {
				return fmt.Errorf("proto: wrong wireType = %d for field Email", wireType)
			}
			var stringLen uint64
			for shift := uint(0); ; shift += 7 {
				if shift >= 64 {
					return ErrIntOverflowJwt
				}
				if iNdEx >= l {
					return io.ErrUnexpectedEOF
				}
				b := dAtA[iNdEx]
				iNdEx++
				stringLen |= (uint64(b) & 0x7F) << shift
				if b < 0x80 {
					break
				}
			}
			intStringLen := int(stringLen)
			if intStringLen < 0 {
				return ErrInvalidLengthJwt
			}
			postIndex := iNdEx + intStringLen
			if postIndex > l {
				return io.ErrUnexpectedEOF
			}
			m.Email = string(dAtA[iNdEx:postIndex])
			iNdEx = postIndex
		default:
			iNdEx = preIndex
			skippy, err := skipJwt(dAtA[iNdEx:])
			if err != nil {
				return err
			}
			if skippy < 0 {
				return ErrInvalidLengthJwt
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
func skipJwt(dAtA []byte) (n int, err error) {
	l := len(dAtA)
	iNdEx := 0
	for iNdEx < l {
		var wire uint64
		for shift := uint(0); ; shift += 7 {
			if shift >= 64 {
				return 0, ErrIntOverflowJwt
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
					return 0, ErrIntOverflowJwt
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
					return 0, ErrIntOverflowJwt
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
			iNdEx += length
			if length < 0 {
				return 0, ErrInvalidLengthJwt
			}
			return iNdEx, nil
		case 3:
			for {
				var innerWire uint64
				var start int = iNdEx
				for shift := uint(0); ; shift += 7 {
					if shift >= 64 {
						return 0, ErrIntOverflowJwt
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
				next, err := skipJwt(dAtA[start:])
				if err != nil {
					return 0, err
				}
				iNdEx = start + next
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
	ErrInvalidLengthJwt = fmt.Errorf("proto: negative length found during unmarshaling")
	ErrIntOverflowJwt   = fmt.Errorf("proto: integer overflow")
)

func init() {
	proto.RegisterFile("services/common/proto/jwt.proto", fileDescriptor_jwt_cd8269b5744a05e9)
}

var fileDescriptor_jwt_cd8269b5744a05e9 = []byte{
	// 351 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0x3f, 0x4e, 0xf3, 0x30,
	0x1c, 0x86, 0xe3, 0xf4, 0x6b, 0xda, 0x78, 0xb4, 0xaa, 0x4f, 0x56, 0x11, 0x4e, 0xc5, 0xd4, 0xa5,
	0x0d, 0x12, 0x23, 0x13, 0xa5, 0x4b, 0x19, 0x23, 0x10, 0x12, 0x03, 0x55, 0xfe, 0xb8, 0xc5, 0x55,
	0x12, 0x47, 0xb1, 0x43, 0x3b, 0x72, 0x04, 0xb8, 0x05, 0x47, 0x61, 0xec, 0xc8, 0x54, 0x51, 0xb3,
	0x20, 0xa6, 0x1e, 0x01, 0xc5, 0x2e, 0x9d, 0x98, 0xfc, 0x3e, 0x3f, 0xfb, 0xf1, 0x6b, 0xc9, 0xd0,
	0x13, 0xb4, 0x7c, 0x64, 0x31, 0x15, 0x7e, 0xcc, 0xb3, 0x8c, 0xe7, 0x7e, 0x51, 0x72, 0xc9, 0xfd,
	0xc5, 0x52, 0x0e, 0x75, 0x42, 0x6e, 0x91, 0x0e, 0xcd, 0x56, 0x77, 0x30, 0x67, 0xf2, 0xa1, 0x8a,
	0x6a, 0xf4, 0xe7, 0x7c, 0xce, 0xcd, 0xd9, 0xa8, 0x9a, 0x69, 0x32, 0x62, 0x9d, 0x8c, 0x79, 0xf2,
	0x62, 0x43, 0xf7, 0xea, 0xf6, 0xfa, 0x32, 0x0d, 0x59, 0x26, 0x50, 0x17, 0xb6, 0xc3, 0x2a, 0x61,
	0x34, 0x8f, 0x29, 0x06, 0x3d, 0xd0, 0x77, 0x83, 0x03, 0xa3, 0x63, 0x08, 0xe9, 0xaa, 0x60, 0x25,
	0x15, 0xd3, 0x50, 0x62, 0xbb, 0x07, 0xfa, 0x8d, 0xc0, 0xdd, 0x4f, 0x2e, 0x24, 0xfa, 0x0f, 0x6d,
	0x96, 0xe0, 0x46, 0x2d, 0x8d, 0x1c, 0xb5, 0xf1, 0xec, 0xc9, 0x38, 0xb0, 0x59, 0x82, 0x8e, 0xa0,
	0xcb, 0x84, 0xa8, 0x68, 0x52, 0x5b, 0xff, 0xb4, 0xd5, 0x36, 0x03, 0x2d, 0x39, 0x3a, 0x97, 0xb8,
	0xa9, 0xdb, 0xf6, 0x54, 0x77, 0xe5, 0x5c, 0x4e, 0x23, 0x3a, 0xe3, 0x25, 0xc5, 0x8e, 0xe9, 0xca,
	0xb9, 0x1c, 0xe9, 0x01, 0xc2, 0xb0, 0x25, 0xaa, 0x68, 0x41, 0x63, 0x89, 0x5b, 0xda, 0xfb, 0x45,
	0x34, 0x80, 0xad, 0x4a, 0xd0, 0x72, 0xca, 0x12, 0xdc, 0xd6, 0x4f, 0xe9, 0xa8, 0x8d, 0xe7, 0xdc,
	0x08, 0x5a, 0x4e, 0xc6, 0xdf, 0x1b, 0xcf, 0xa9, 0x74, 0x0a, 0xcc, 0x9a, 0xa0, 0x0e, 0x6c, 0xd2,
	0x2c, 0x64, 0x29, 0x76, 0xf5, 0x35, 0x06, 0x46, 0xf7, 0xeb, 0x2d, 0xb1, 0xde, 0xb7, 0xc4, 0xda,
	0x6d, 0x09, 0x78, 0x52, 0x04, 0xbc, 0x2a, 0x02, 0xde, 0x14, 0x01, 0x6b, 0x45, 0xc0, 0x87, 0x22,
	0xe0, 0x4b, 0x11, 0x6b, 0xa7, 0x08, 0x78, 0xfe, 0x24, 0xd6, 0xdd, 0x69, 0xc1, 0x56, 0x8c, 0xa6,
	0x61, 0x24, 0x86, 0x21, 0xf3, 0x0f, 0xe0, 0xff, 0xf9, 0x67, 0xe7, 0x8b, 0xa5, 0x8c, 0x1c, 0x1d,
	0xcf, 0x7e, 0x02, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x55, 0x23, 0xd8, 0xd7, 0x01, 0x00, 0x00,
}
