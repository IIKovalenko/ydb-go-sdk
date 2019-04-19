// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_operation.proto

package Ydb_Operations

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb"
	Ydb_Issue "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb_Issue"
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	any "github.com/golang/protobuf/ptypes/any"
	duration "github.com/golang/protobuf/ptypes/duration"
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

type OperationParams_OperationMode int32

const (
	OperationParams_OPERATION_MODE_UNSPECIFIED OperationParams_OperationMode = 0
	// Server will only reply once operation is finished (ready=true), and operation object won't be
	// accessible after the reply. This is a basic request-response mode.
	OperationParams_SYNC OperationParams_OperationMode = 1
)

var OperationParams_OperationMode_name = map[int32]string{
	0: "OPERATION_MODE_UNSPECIFIED",
	1: "SYNC",
}

var OperationParams_OperationMode_value = map[string]int32{
	"OPERATION_MODE_UNSPECIFIED": 0,
	"SYNC":                       1,
}

func (x OperationParams_OperationMode) String() string {
	return proto.EnumName(OperationParams_OperationMode_name, int32(x))
}

func (OperationParams_OperationMode) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_589bdee9fb97c360, []int{0, 0}
}

type OperationParams struct {
	OperationMode OperationParams_OperationMode `protobuf:"varint,1,opt,name=operation_mode,json=operationMode,proto3,enum=Ydb.Operations.OperationParams_OperationMode" json:"operation_mode,omitempty"`
	// Indicates that client is no longer interested in the result of operation after the specified duration
	// starting from the time operation arrives at the server.
	// Server will try to stop the execution of operation and if no result is currently available the operation
	// will receive TIMEOUT status code, which will be sent back to client if it was waiting for the operation result.
	// Timeout of operation does not tell anything about its result, it might be completed successfully
	// or cancelled on server.
	OperationTimeout *duration.Duration `protobuf:"bytes,2,opt,name=operation_timeout,json=operationTimeout,proto3" json:"operation_timeout,omitempty"`
	// Server will try to cancel the operation after the specified duration starting from the time
	// the operation arrives at server.
	// In case of successful cancellation operation will receive CANCELLED status code, which will be
	// sent back to client if it was waiting for the operation result.
	// In case when cancellation isn't possible, no action will be performed.
	CancelAfter          *duration.Duration `protobuf:"bytes,3,opt,name=cancel_after,json=cancelAfter,proto3" json:"cancel_after,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *OperationParams) Reset()         { *m = OperationParams{} }
func (m *OperationParams) String() string { return proto.CompactTextString(m) }
func (*OperationParams) ProtoMessage()    {}
func (*OperationParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_589bdee9fb97c360, []int{0}
}

func (m *OperationParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_OperationParams.Unmarshal(m, b)
}
func (m *OperationParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_OperationParams.Marshal(b, m, deterministic)
}
func (m *OperationParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_OperationParams.Merge(m, src)
}
func (m *OperationParams) XXX_Size() int {
	return xxx_messageInfo_OperationParams.Size(m)
}
func (m *OperationParams) XXX_DiscardUnknown() {
	xxx_messageInfo_OperationParams.DiscardUnknown(m)
}

var xxx_messageInfo_OperationParams proto.InternalMessageInfo

func (m *OperationParams) GetOperationMode() OperationParams_OperationMode {
	if m != nil {
		return m.OperationMode
	}
	return OperationParams_OPERATION_MODE_UNSPECIFIED
}

func (m *OperationParams) GetOperationTimeout() *duration.Duration {
	if m != nil {
		return m.OperationTimeout
	}
	return nil
}

func (m *OperationParams) GetCancelAfter() *duration.Duration {
	if m != nil {
		return m.CancelAfter
	}
	return nil
}

type GetOperationRequest struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *GetOperationRequest) Reset()         { *m = GetOperationRequest{} }
func (m *GetOperationRequest) String() string { return proto.CompactTextString(m) }
func (*GetOperationRequest) ProtoMessage()    {}
func (*GetOperationRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_589bdee9fb97c360, []int{1}
}

func (m *GetOperationRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationRequest.Unmarshal(m, b)
}
func (m *GetOperationRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationRequest.Marshal(b, m, deterministic)
}
func (m *GetOperationRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationRequest.Merge(m, src)
}
func (m *GetOperationRequest) XXX_Size() int {
	return xxx_messageInfo_GetOperationRequest.Size(m)
}
func (m *GetOperationRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationRequest.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationRequest proto.InternalMessageInfo

func (m *GetOperationRequest) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

type GetOperationResponse struct {
	Operation            *Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}   `json:"-"`
	XXX_unrecognized     []byte     `json:"-"`
	XXX_sizecache        int32      `json:"-"`
}

func (m *GetOperationResponse) Reset()         { *m = GetOperationResponse{} }
func (m *GetOperationResponse) String() string { return proto.CompactTextString(m) }
func (*GetOperationResponse) ProtoMessage()    {}
func (*GetOperationResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_589bdee9fb97c360, []int{2}
}

func (m *GetOperationResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_GetOperationResponse.Unmarshal(m, b)
}
func (m *GetOperationResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_GetOperationResponse.Marshal(b, m, deterministic)
}
func (m *GetOperationResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_GetOperationResponse.Merge(m, src)
}
func (m *GetOperationResponse) XXX_Size() int {
	return xxx_messageInfo_GetOperationResponse.Size(m)
}
func (m *GetOperationResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_GetOperationResponse.DiscardUnknown(m)
}

var xxx_messageInfo_GetOperationResponse proto.InternalMessageInfo

func (m *GetOperationResponse) GetOperation() *Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type Operation struct {
	// Identifier of the operation, empty value means no active operation object is present (it was forgotten or
	// not created in the first place, as in SYNC operation mode).
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// true - this operation has beed finished (doesn't matter successful or not),
	// so Status field has status code, and Result field can contains result data.
	// false - this operation still running. You can repeat request using operation Id.
	Ready  bool                      `protobuf:"varint,2,opt,name=ready,proto3" json:"ready,omitempty"`
	Status Ydb.StatusIds_StatusCode  `protobuf:"varint,3,opt,name=status,proto3,enum=Ydb.StatusIds_StatusCode" json:"status,omitempty"`
	Issues []*Ydb_Issue.IssueMessage `protobuf:"bytes,4,rep,name=issues,proto3" json:"issues,omitempty"`
	// Result data
	Result               *any.Any `protobuf:"bytes,5,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Operation) Reset()         { *m = Operation{} }
func (m *Operation) String() string { return proto.CompactTextString(m) }
func (*Operation) ProtoMessage()    {}
func (*Operation) Descriptor() ([]byte, []int) {
	return fileDescriptor_589bdee9fb97c360, []int{3}
}

func (m *Operation) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Operation.Unmarshal(m, b)
}
func (m *Operation) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Operation.Marshal(b, m, deterministic)
}
func (m *Operation) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Operation.Merge(m, src)
}
func (m *Operation) XXX_Size() int {
	return xxx_messageInfo_Operation.Size(m)
}
func (m *Operation) XXX_DiscardUnknown() {
	xxx_messageInfo_Operation.DiscardUnknown(m)
}

var xxx_messageInfo_Operation proto.InternalMessageInfo

func (m *Operation) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Operation) GetReady() bool {
	if m != nil {
		return m.Ready
	}
	return false
}

func (m *Operation) GetStatus() Ydb.StatusIds_StatusCode {
	if m != nil {
		return m.Status
	}
	return Ydb.StatusIds_STATUS_CODE_UNSPECIFIED
}

func (m *Operation) GetIssues() []*Ydb_Issue.IssueMessage {
	if m != nil {
		return m.Issues
	}
	return nil
}

func (m *Operation) GetResult() *any.Any {
	if m != nil {
		return m.Result
	}
	return nil
}

func init() {
	proto.RegisterEnum("Ydb.Operations.OperationParams_OperationMode", OperationParams_OperationMode_name, OperationParams_OperationMode_value)
	proto.RegisterType((*OperationParams)(nil), "Ydb.Operations.OperationParams")
	proto.RegisterType((*GetOperationRequest)(nil), "Ydb.Operations.GetOperationRequest")
	proto.RegisterType((*GetOperationResponse)(nil), "Ydb.Operations.GetOperationResponse")
	proto.RegisterType((*Operation)(nil), "Ydb.Operations.Operation")
}

func init() { proto.RegisterFile("ydb_operation.proto", fileDescriptor_589bdee9fb97c360) }

var fileDescriptor_589bdee9fb97c360 = []byte{
	// 468 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x52, 0x4d, 0x6f, 0xd3, 0x4c,
	0x10, 0x7e, 0x9d, 0xb6, 0x51, 0xb3, 0x79, 0x63, 0xc2, 0x36, 0x12, 0x6e, 0x0e, 0x55, 0x64, 0x84,
	0x94, 0x03, 0xd8, 0x10, 0x0e, 0x08, 0x89, 0x4b, 0x9a, 0xa4, 0xc8, 0x87, 0xc4, 0x91, 0x13, 0x0e,
	0x3d, 0x59, 0xeb, 0xec, 0x34, 0xb2, 0x1a, 0x7b, 0x8d, 0x77, 0x57, 0xc2, 0xff, 0x88, 0xff, 0xc1,
	0x9f, 0xe2, 0x88, 0xbc, 0xeb, 0x3a, 0x1f, 0x08, 0xb8, 0xac, 0x76, 0x66, 0x9e, 0x67, 0x3e, 0x9e,
	0x19, 0x74, 0x55, 0xd0, 0x28, 0x64, 0x19, 0xe4, 0x44, 0xc4, 0x2c, 0x75, 0xb2, 0x9c, 0x09, 0x86,
	0xcd, 0x7b, 0x1a, 0x39, 0xfe, 0x93, 0x93, 0xf7, 0xaf, 0xb7, 0x8c, 0x6d, 0x77, 0xe0, 0xaa, 0x68,
	0x24, 0x1f, 0x5c, 0x92, 0x16, 0x1a, 0xda, 0xbf, 0x39, 0x0d, 0x51, 0x79, 0x98, 0xaa, 0xff, 0xf6,
	0x31, 0x7e, 0x8c, 0x93, 0xdc, 0xcd, 0x64, 0xb4, 0x8b, 0x37, 0x2e, 0xc9, 0x62, 0x0d, 0xe5, 0x6e,
	0x59, 0x38, 0xe6, 0x5c, 0x42, 0x98, 0x00, 0xe7, 0x64, 0x0b, 0x15, 0xc3, 0xfd, 0x2b, 0x83, 0x0b,
	0x22, 0x24, 0x0f, 0x37, 0x8c, 0x02, 0xd7, 0x04, 0xfb, 0x7b, 0x03, 0x3d, 0xab, 0x9b, 0x5d, 0x92,
	0x9c, 0x24, 0x1c, 0xaf, 0x91, 0x59, 0x0f, 0x15, 0x26, 0x8c, 0x82, 0x65, 0x0c, 0x8c, 0xa1, 0x39,
	0x7a, 0xe3, 0x1c, 0x8f, 0xe6, 0x9c, 0x10, 0xf7, 0xf6, 0x9c, 0x51, 0x08, 0x3a, 0xec, 0xd0, 0xc4,
	0x77, 0xe8, 0xf9, 0x3e, 0xab, 0x88, 0x13, 0x60, 0x52, 0x58, 0x8d, 0x81, 0x31, 0x6c, 0x8f, 0xae,
	0x1d, 0x2d, 0x84, 0xf3, 0x24, 0x84, 0x33, 0xad, 0x84, 0x08, 0xba, 0x35, 0x67, 0xad, 0x29, 0xf8,
	0x13, 0xfa, 0x7f, 0x43, 0xd2, 0x0d, 0xec, 0x42, 0xf2, 0x20, 0x20, 0xb7, 0xce, 0xfe, 0x95, 0xa2,
	0xad, 0xe1, 0xe3, 0x12, 0x6d, 0x7f, 0x44, 0x9d, 0xa3, 0x2e, 0xf1, 0x0d, 0xea, 0xfb, 0xcb, 0x59,
	0x30, 0x5e, 0x7b, 0xfe, 0x22, 0x9c, 0xfb, 0xd3, 0x59, 0xf8, 0x65, 0xb1, 0x5a, 0xce, 0x26, 0xde,
	0x9d, 0x37, 0x9b, 0x76, 0xff, 0xc3, 0x97, 0xe8, 0x7c, 0x75, 0xbf, 0x98, 0x74, 0x0d, 0xfb, 0x15,
	0xba, 0xfa, 0x0c, 0xa2, 0x66, 0x07, 0xf0, 0x55, 0x02, 0x17, 0xd8, 0x44, 0x8d, 0x98, 0x2a, 0x85,
	0x5a, 0x41, 0x23, 0xa6, 0xb6, 0x8f, 0x7a, 0xc7, 0x30, 0x9e, 0xb1, 0x94, 0x03, 0xfe, 0x80, 0x5a,
	0xf5, 0x2c, 0x0a, 0x5e, 0x36, 0xfd, 0x27, 0x41, 0x83, 0x3d, 0xd6, 0xfe, 0x61, 0xa0, 0x56, 0x1d,
	0x38, 0x2d, 0x87, 0x7b, 0xe8, 0x22, 0x07, 0x42, 0x0b, 0x25, 0xe5, 0x65, 0xa0, 0x0d, 0xfc, 0x0e,
	0x35, 0xf5, 0xb2, 0x95, 0x3c, 0x66, 0x55, 0x69, 0xa5, 0x5c, 0x1e, 0xe5, 0xd5, 0x6f, 0x52, 0xae,
	0xa9, 0x02, 0x62, 0x17, 0x35, 0xd5, 0x45, 0x71, 0xeb, 0x7c, 0x70, 0x36, 0x6c, 0x8f, 0x5e, 0x28,
	0x8a, 0x57, 0xba, 0xf4, 0x3b, 0xd7, 0x97, 0x16, 0x54, 0x30, 0xfc, 0x1a, 0x35, 0x73, 0xe0, 0x72,
	0x27, 0xac, 0x0b, 0x35, 0x4d, 0xef, 0xb7, 0x15, 0x8c, 0xd3, 0x22, 0xa8, 0x30, 0xb7, 0x2f, 0x51,
	0x27, 0x97, 0x4e, 0x41, 0x52, 0x0a, 0xdf, 0x9c, 0x82, 0x46, 0xb7, 0x07, 0x67, 0xa7, 0x2e, 0xf4,
	0xa7, 0x61, 0x44, 0x4d, 0x45, 0x7d, 0xff, 0x2b, 0x00, 0x00, 0xff, 0xff, 0x65, 0xe7, 0xf1, 0x43,
	0x59, 0x03, 0x00, 0x00,
}

const ()