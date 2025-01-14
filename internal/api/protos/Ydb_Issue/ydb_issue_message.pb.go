// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_issue_message.proto

package Ydb_Issue

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

// IssueMessage is a transport format for yql/public/issue library
type IssueMessage struct {
	Position    *IssueMessage_Position `protobuf:"bytes,1,opt,name=position" json:"position,omitempty"`
	Message     *string                `protobuf:"bytes,2,opt,name=message" json:"message,omitempty"`
	EndPosition *IssueMessage_Position `protobuf:"bytes,3,opt,name=end_position,json=endPosition" json:"end_position,omitempty"`
	IssueCode   *uint32                `protobuf:"varint,4,opt,name=issue_code,json=issueCode" json:"issue_code,omitempty"`
	// Severity values from yql/public/issue/protos/issue_severity.proto
	// FATAL = 0;
	// ERROR = 1;
	// WARNING = 2;
	// INFO = 3;
	Severity             *uint32         `protobuf:"varint,5,opt,name=severity" json:"severity,omitempty"`
	Issues               []*IssueMessage `protobuf:"bytes,6,rep,name=issues" json:"issues,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *IssueMessage) Reset()         { *m = IssueMessage{} }
func (m *IssueMessage) String() string { return proto.CompactTextString(m) }
func (*IssueMessage) ProtoMessage()    {}
func (*IssueMessage) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b472e41df8bb0e5, []int{0}
}

func (m *IssueMessage) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueMessage.Unmarshal(m, b)
}
func (m *IssueMessage) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueMessage.Marshal(b, m, deterministic)
}
func (m *IssueMessage) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueMessage.Merge(m, src)
}
func (m *IssueMessage) XXX_Size() int {
	return xxx_messageInfo_IssueMessage.Size(m)
}
func (m *IssueMessage) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueMessage.DiscardUnknown(m)
}

var xxx_messageInfo_IssueMessage proto.InternalMessageInfo

func (m *IssueMessage) GetPosition() *IssueMessage_Position {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *IssueMessage) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

func (m *IssueMessage) GetEndPosition() *IssueMessage_Position {
	if m != nil {
		return m.EndPosition
	}
	return nil
}

func (m *IssueMessage) GetIssueCode() uint32 {
	if m != nil && m.IssueCode != nil {
		return *m.IssueCode
	}
	return 0
}

func (m *IssueMessage) GetSeverity() uint32 {
	if m != nil && m.Severity != nil {
		return *m.Severity
	}
	return 0
}

func (m *IssueMessage) GetIssues() []*IssueMessage {
	if m != nil {
		return m.Issues
	}
	return nil
}

type IssueMessage_Position struct {
	Row                  *uint32  `protobuf:"varint,1,opt,name=row" json:"row,omitempty"`
	Column               *uint32  `protobuf:"varint,2,opt,name=column" json:"column,omitempty"`
	File                 *string  `protobuf:"bytes,3,opt,name=file" json:"file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *IssueMessage_Position) Reset()         { *m = IssueMessage_Position{} }
func (m *IssueMessage_Position) String() string { return proto.CompactTextString(m) }
func (*IssueMessage_Position) ProtoMessage()    {}
func (*IssueMessage_Position) Descriptor() ([]byte, []int) {
	return fileDescriptor_0b472e41df8bb0e5, []int{0, 0}
}

func (m *IssueMessage_Position) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_IssueMessage_Position.Unmarshal(m, b)
}
func (m *IssueMessage_Position) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_IssueMessage_Position.Marshal(b, m, deterministic)
}
func (m *IssueMessage_Position) XXX_Merge(src proto.Message) {
	xxx_messageInfo_IssueMessage_Position.Merge(m, src)
}
func (m *IssueMessage_Position) XXX_Size() int {
	return xxx_messageInfo_IssueMessage_Position.Size(m)
}
func (m *IssueMessage_Position) XXX_DiscardUnknown() {
	xxx_messageInfo_IssueMessage_Position.DiscardUnknown(m)
}

var xxx_messageInfo_IssueMessage_Position proto.InternalMessageInfo

func (m *IssueMessage_Position) GetRow() uint32 {
	if m != nil && m.Row != nil {
		return *m.Row
	}
	return 0
}

func (m *IssueMessage_Position) GetColumn() uint32 {
	if m != nil && m.Column != nil {
		return *m.Column
	}
	return 0
}

func (m *IssueMessage_Position) GetFile() string {
	if m != nil && m.File != nil {
		return *m.File
	}
	return ""
}

func init() {
	proto.RegisterType((*IssueMessage)(nil), "Ydb.Issue.IssueMessage")
	proto.RegisterType((*IssueMessage_Position)(nil), "Ydb.Issue.IssueMessage.Position")
}

func init() { proto.RegisterFile("ydb_issue_message.proto", fileDescriptor_0b472e41df8bb0e5) }

var fileDescriptor_0b472e41df8bb0e5 = []byte{
	// 256 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x90, 0xc1, 0x4b, 0xc3, 0x30,
	0x14, 0xc6, 0xe9, 0x3a, 0x6b, 0xfb, 0xba, 0xca, 0x78, 0x07, 0x17, 0x06, 0x42, 0xf0, 0xd4, 0x53,
	0x84, 0x9d, 0x3d, 0xb9, 0x8b, 0x1e, 0x04, 0xc9, 0xcd, 0x53, 0x59, 0x9b, 0xa7, 0x04, 0xd6, 0x64,
	0x34, 0x9d, 0xda, 0x3f, 0xd7, 0xff, 0x44, 0x96, 0xb5, 0xc1, 0x8b, 0xe0, 0xed, 0xfb, 0xf2, 0x7e,
	0x79, 0x1f, 0xdf, 0x83, 0xd5, 0xa0, 0xea, 0x4a, 0x3b, 0x77, 0xa4, 0xaa, 0x25, 0xe7, 0x76, 0xef,
	0x24, 0x0e, 0x9d, 0xed, 0x2d, 0x66, 0xaf, 0xaa, 0x16, 0x4f, 0xa7, 0xc1, 0xed, 0xf7, 0x0c, 0x16,
	0x5e, 0x3d, 0x9f, 0x09, 0xbc, 0x87, 0xf4, 0x60, 0x9d, 0xee, 0xb5, 0x35, 0x2c, 0xe2, 0x51, 0x99,
	0x6f, 0xb8, 0x08, 0xb8, 0xf8, 0x8d, 0x8a, 0x97, 0x91, 0x93, 0xe1, 0x07, 0x32, 0xb8, 0x1c, 0xa3,
	0xd8, 0x8c, 0x47, 0x65, 0x26, 0x27, 0x8b, 0x5b, 0x58, 0x90, 0x51, 0x55, 0xd8, 0x1d, 0xff, 0x73,
	0x77, 0x4e, 0x46, 0x4d, 0x06, 0x6f, 0x00, 0xce, 0x7d, 0x1a, 0xab, 0x88, 0xcd, 0x79, 0x54, 0x16,
	0x32, 0xf3, 0x2f, 0x5b, 0xab, 0x08, 0xd7, 0x90, 0x3a, 0xfa, 0xa0, 0x4e, 0xf7, 0x03, 0xbb, 0xf0,
	0xc3, 0xe0, 0xf1, 0x0e, 0x12, 0x0f, 0x3a, 0x96, 0xf0, 0xb8, 0xcc, 0x37, 0xab, 0x3f, 0x92, 0xe5,
	0x88, 0xad, 0x1f, 0x21, 0x0d, 0xb9, 0x4b, 0x88, 0x3b, 0xfb, 0xe9, 0xef, 0x51, 0xc8, 0x93, 0xc4,
	0x6b, 0x48, 0x1a, 0xbb, 0x3f, 0xb6, 0xc6, 0xf7, 0x2c, 0xe4, 0xe8, 0x10, 0x61, 0xfe, 0xa6, 0xf7,
	0xe4, 0xeb, 0x65, 0xd2, 0xeb, 0x87, 0x25, 0x5c, 0x35, 0xb6, 0x15, 0xc3, 0xce, 0x28, 0xfa, 0x12,
	0x83, 0xaa, 0x7f, 0x02, 0x00, 0x00, 0xff, 0xff, 0xfe, 0x9d, 0x35, 0x6b, 0x9a, 0x01, 0x00, 0x00,
}

const ()
