// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_common.proto

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

type FeatureFlag_Status int32

const (
	FeatureFlag_STATUS_UNSPECIFIED FeatureFlag_Status = 0
	FeatureFlag_ENABLED            FeatureFlag_Status = 1
	FeatureFlag_DISABLED           FeatureFlag_Status = 2
)

var FeatureFlag_Status_name = map[int32]string{
	0: "STATUS_UNSPECIFIED",
	1: "ENABLED",
	2: "DISABLED",
}

var FeatureFlag_Status_value = map[string]int32{
	"STATUS_UNSPECIFIED": 0,
	"ENABLED":            1,
	"DISABLED":           2,
}

func (x FeatureFlag_Status) String() string {
	return proto.EnumName(FeatureFlag_Status_name, int32(x))
}

func (FeatureFlag_Status) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_ac23b1f8425834a0, []int{0, 0}
}

type FeatureFlag struct {
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *FeatureFlag) Reset()         { *m = FeatureFlag{} }
func (m *FeatureFlag) String() string { return proto.CompactTextString(m) }
func (*FeatureFlag) ProtoMessage()    {}
func (*FeatureFlag) Descriptor() ([]byte, []int) {
	return fileDescriptor_ac23b1f8425834a0, []int{0}
}

func (m *FeatureFlag) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_FeatureFlag.Unmarshal(m, b)
}
func (m *FeatureFlag) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_FeatureFlag.Marshal(b, m, deterministic)
}
func (m *FeatureFlag) XXX_Merge(src proto.Message) {
	xxx_messageInfo_FeatureFlag.Merge(m, src)
}
func (m *FeatureFlag) XXX_Size() int {
	return xxx_messageInfo_FeatureFlag.Size(m)
}
func (m *FeatureFlag) XXX_DiscardUnknown() {
	xxx_messageInfo_FeatureFlag.DiscardUnknown(m)
}

var xxx_messageInfo_FeatureFlag proto.InternalMessageInfo

func init() {
	proto.RegisterEnum("Ydb.FeatureFlag_Status", FeatureFlag_Status_name, FeatureFlag_Status_value)
	proto.RegisterType((*FeatureFlag)(nil), "Ydb.FeatureFlag")
}

func init() { proto.RegisterFile("ydb_common.proto", fileDescriptor_ac23b1f8425834a0) }

var fileDescriptor_ac23b1f8425834a0 = []byte{
	// 157 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x12, 0xa8, 0x4c, 0x49, 0x8a,
	0x4f, 0xce, 0xcf, 0xcd, 0xcd, 0xcf, 0xd3, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x8e, 0x4c,
	0x49, 0x52, 0xf2, 0xe2, 0xe2, 0x76, 0x4b, 0x4d, 0x2c, 0x29, 0x2d, 0x4a, 0x75, 0xcb, 0x49, 0x4c,
	0x57, 0xb2, 0xe6, 0x62, 0x0b, 0x2e, 0x49, 0x2c, 0x29, 0x2d, 0x16, 0x12, 0xe3, 0x12, 0x0a, 0x0e,
	0x71, 0x0c, 0x09, 0x0d, 0x8e, 0x0f, 0xf5, 0x0b, 0x0e, 0x70, 0x75, 0xf6, 0x74, 0xf3, 0x74, 0x75,
	0x11, 0x60, 0x10, 0xe2, 0xe6, 0x62, 0x77, 0xf5, 0x73, 0x74, 0xf2, 0x71, 0x75, 0x11, 0x60, 0x14,
	0xe2, 0xe1, 0xe2, 0x70, 0xf1, 0x0c, 0x86, 0xf0, 0x98, 0x9c, 0x34, 0xb8, 0x44, 0x93, 0xf3, 0x73,
	0xf5, 0x2a, 0x13, 0xf3, 0x52, 0x52, 0x2b, 0xf4, 0x2a, 0x53, 0x92, 0xf4, 0x20, 0xf6, 0x39, 0xf1,
	0x38, 0x83, 0xe9, 0x00, 0x90, 0xb5, 0xc5, 0x3f, 0x18, 0x19, 0x93, 0xd8, 0xc0, 0x2e, 0x30, 0x06,
	0x04, 0x00, 0x00, 0xff, 0xff, 0x76, 0x4d, 0x56, 0x50, 0x95, 0x00, 0x00, 0x00,
}

const ()
