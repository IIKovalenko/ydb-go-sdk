// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_status_codes.proto

package Ydb

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
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

// reserved range [400000, 400999]
type StatusIds_StatusCode int32

const (
	StatusIds_STATUS_CODE_UNSPECIFIED StatusIds_StatusCode = 0
	StatusIds_SUCCESS                 StatusIds_StatusCode = 400000
	StatusIds_BAD_REQUEST             StatusIds_StatusCode = 400010
	StatusIds_UNAUTHORIZED            StatusIds_StatusCode = 400020
	StatusIds_INTERNAL_ERROR          StatusIds_StatusCode = 400030
	StatusIds_ABORTED                 StatusIds_StatusCode = 400040
	StatusIds_UNAVAILABLE             StatusIds_StatusCode = 400050
	StatusIds_OVERLOADED              StatusIds_StatusCode = 400060
	StatusIds_SCHEME_ERROR            StatusIds_StatusCode = 400070
	StatusIds_GENERIC_ERROR           StatusIds_StatusCode = 400080
	StatusIds_TIMEOUT                 StatusIds_StatusCode = 400090
	StatusIds_BAD_SESSION             StatusIds_StatusCode = 400100
	StatusIds_PRECONDITION_FAILED     StatusIds_StatusCode = 400120
	StatusIds_ALREADY_EXISTS          StatusIds_StatusCode = 400130
	StatusIds_NOT_FOUND               StatusIds_StatusCode = 400140
	StatusIds_SESSION_EXPIRED         StatusIds_StatusCode = 400150
	StatusIds_CANCELLED               StatusIds_StatusCode = 400160
)

var StatusIds_StatusCode_name = map[int32]string{
	0:      "STATUS_CODE_UNSPECIFIED",
	400000: "SUCCESS",
	400010: "BAD_REQUEST",
	400020: "UNAUTHORIZED",
	400030: "INTERNAL_ERROR",
	400040: "ABORTED",
	400050: "UNAVAILABLE",
	400060: "OVERLOADED",
	400070: "SCHEME_ERROR",
	400080: "GENERIC_ERROR",
	400090: "TIMEOUT",
	400100: "BAD_SESSION",
	400120: "PRECONDITION_FAILED",
	400130: "ALREADY_EXISTS",
	400140: "NOT_FOUND",
	400150: "SESSION_EXPIRED",
	400160: "CANCELLED",
}

var StatusIds_StatusCode_value = map[string]int32{
	"STATUS_CODE_UNSPECIFIED": 0,
	"SUCCESS":                 400000,
	"BAD_REQUEST":             400010,
	"UNAUTHORIZED":            400020,
	"INTERNAL_ERROR":          400030,
	"ABORTED":                 400040,
	"UNAVAILABLE":             400050,
	"OVERLOADED":              400060,
	"SCHEME_ERROR":            400070,
	"GENERIC_ERROR":           400080,
	"TIMEOUT":                 400090,
	"BAD_SESSION":             400100,
	"PRECONDITION_FAILED":     400120,
	"ALREADY_EXISTS":          400130,
	"NOT_FOUND":               400140,
	"SESSION_EXPIRED":         400150,
	"CANCELLED":               400160,
}

func (x StatusIds_StatusCode) String() string {
	return proto.EnumName(StatusIds_StatusCode_name, int32(x))
}

func (StatusIds_StatusCode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_f81e45819472f2bf, []int{0, 0}
}

type StatusIds struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *StatusIds) Reset()         { *m = StatusIds{} }
func (m *StatusIds) String() string { return proto.CompactTextString(m) }
func (*StatusIds) ProtoMessage()    {}
func (*StatusIds) Descriptor() ([]byte, []int) {
	return fileDescriptor_f81e45819472f2bf, []int{0}
}

func (m *StatusIds) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_StatusIds.Unmarshal(m, b)
}
func (m *StatusIds) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_StatusIds.Marshal(b, m, deterministic)
}
func (m *StatusIds) XXX_Merge(src proto.Message) {
	xxx_messageInfo_StatusIds.Merge(m, src)
}
func (m *StatusIds) XXX_Size() int {
	return xxx_messageInfo_StatusIds.Size(m)
}
func (m *StatusIds) XXX_DiscardUnknown() {
	xxx_messageInfo_StatusIds.DiscardUnknown(m)
}

var xxx_messageInfo_StatusIds proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("Ydb.StatusIds_StatusCode", StatusIds_StatusCode_name, StatusIds_StatusCode_value)
	proto.RegisterType((*StatusIds)(nil), "Ydb.StatusIds")
}

func init() { proto.RegisterFile("ydb_status_codes.proto", fileDescriptor_f81e45819472f2bf) }

var fileDescriptor_f81e45819472f2bf = []byte{
	// 372 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x44, 0xd1, 0xbf, 0x8f, 0x12, 0x41,
	0x14, 0x07, 0x70, 0x83, 0x89, 0x86, 0xa7, 0xc0, 0x30, 0xf8, 0x63, 0x8c, 0x1d, 0xf6, 0x34, 0xfe,
	0x05, 0xb3, 0x33, 0x0f, 0x99, 0x64, 0x99, 0x59, 0xe7, 0x07, 0x01, 0x9b, 0x09, 0xcb, 0x52, 0xea,
	0x1a, 0x17, 0x13, 0xe9, 0x8c, 0xa5, 0xb1, 0x34, 0x96, 0xc6, 0xd2, 0xfa, 0x92, 0xa5, 0xba, 0xfa,
	0xea, 0xab, 0xaf, 0xb8, 0xea, 0xca, 0xfb, 0x03, 0xae, 0xbc, 0xc0, 0x92, 0x50, 0xce, 0xbc, 0x97,
	0x4f, 0xbe, 0x79, 0x5f, 0x78, 0xb1, 0x2d, 0xf2, 0x58, 0x6d, 0x96, 0x9b, 0xaf, 0x55, 0x5c, 0x95,
	0xc5, 0xba, 0x1a, 0x7d, 0xfe, 0x52, 0x6e, 0x4a, 0xfa, 0x70, 0x51, 0xe4, 0xc3, 0xdb, 0x16, 0xb4,
	0xdd, 0x61, 0xa6, 0x8a, 0x6a, 0x78, 0xdd, 0x02, 0x68, 0x5e, 0xa2, 0x2c, 0xd6, 0xf4, 0x35, 0xbc,
	0x74, 0x9e, 0xfb, 0xe0, 0xa2, 0x30, 0x12, 0x63, 0xd0, 0x2e, 0x43, 0xa1, 0xc6, 0x0a, 0x25, 0x79,
	0x40, 0x3b, 0xf0, 0xd8, 0x05, 0x21, 0xd0, 0x39, 0xf2, 0xbd, 0x66, 0xb4, 0x0f, 0x4f, 0x12, 0x2e,
	0xa3, 0xc5, 0xf7, 0x01, 0x9d, 0x27, 0x3f, 0x6b, 0x46, 0x29, 0x3c, 0x0d, 0x9a, 0x07, 0x3f, 0x31,
	0x56, 0x7d, 0x40, 0x49, 0x7e, 0xd7, 0x8c, 0x3e, 0x83, 0xae, 0xd2, 0x1e, 0xad, 0xe6, 0x69, 0x44,
	0x6b, 0x8d, 0x25, 0x7f, 0x6b, 0xb6, 0xb7, 0x78, 0x62, 0xac, 0x47, 0x49, 0xfe, 0x37, 0x56, 0xd0,
	0x7c, 0xc6, 0x55, 0xca, 0x93, 0x14, 0xc9, 0x59, 0xcd, 0x28, 0x01, 0x30, 0x33, 0xb4, 0xa9, 0xe1,
	0x12, 0x25, 0x39, 0x6f, 0x74, 0x27, 0x26, 0x38, 0xc5, 0xa3, 0x73, 0x51, 0x33, 0x3a, 0x80, 0xce,
	0x3b, 0xd4, 0x68, 0x95, 0x38, 0x7e, 0x5e, 0x36, 0xb8, 0x57, 0x53, 0x34, 0xc1, 0x93, 0xab, 0x53,
	0x50, 0x87, 0xce, 0x29, 0xa3, 0xc9, 0x4d, 0xcd, 0xe8, 0x2b, 0x18, 0x64, 0x16, 0x85, 0xd1, 0x52,
	0x79, 0x65, 0x74, 0x1c, 0x73, 0x95, 0xa2, 0x24, 0x77, 0x4d, 0x5e, 0x9e, 0x5a, 0xe4, 0x72, 0x11,
	0x71, 0xae, 0x9c, 0x77, 0xe4, 0xc7, 0x8e, 0xd1, 0x1e, 0xb4, 0xb5, 0xf1, 0x71, 0x6c, 0x82, 0x96,
	0xe4, 0xd7, 0x8e, 0xd1, 0xe7, 0xd0, 0x3b, 0x82, 0x11, 0xe7, 0x99, 0xb2, 0x28, 0xc9, 0x9f, 0x66,
	0x4f, 0x70, 0x2d, 0x30, 0xdd, 0x73, 0xff, 0x76, 0x2c, 0x79, 0x03, 0xdd, 0x55, 0xf9, 0x71, 0xb4,
	0x5d, 0x7e, 0x2a, 0xd6, 0xdf, 0x46, 0xdb, 0x22, 0x4f, 0xfa, 0xa7, 0x7b, 0x57, 0xd9, 0xbe, 0x97,
	0x2a, 0x7f, 0x74, 0xe8, 0xe7, 0xed, 0x7d, 0x00, 0x00, 0x00, 0xff, 0xff, 0xe7, 0xfd, 0xac, 0x60,
	0xb9, 0x01, 0x00, 0x00,
}

const ()
