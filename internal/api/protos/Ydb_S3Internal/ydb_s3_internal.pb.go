// Code generated by protoc-gen-go. DO NOT EDIT.
// source: ydb_s3_internal.proto

package Ydb_S3Internal

import (
	Ydb "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb"
	Ydb_Operations "github.com/yandex-cloud/ydb-go-sdk/internal/api/protos/Ydb_Operations"
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

type S3ListingRequest struct {
	TableName            string                          `protobuf:"bytes,1,opt,name=table_name,json=tableName,proto3" json:"table_name,omitempty"`
	KeyPrefix            *Ydb.TypedValue                 `protobuf:"bytes,2,opt,name=key_prefix,json=keyPrefix,proto3" json:"key_prefix,omitempty"`
	PathColumnPrefix     string                          `protobuf:"bytes,3,opt,name=path_column_prefix,json=pathColumnPrefix,proto3" json:"path_column_prefix,omitempty"`
	PathColumnDelimiter  string                          `protobuf:"bytes,4,opt,name=path_column_delimiter,json=pathColumnDelimiter,proto3" json:"path_column_delimiter,omitempty"`
	StartAfterKeySuffix  *Ydb.TypedValue                 `protobuf:"bytes,5,opt,name=start_after_key_suffix,json=startAfterKeySuffix,proto3" json:"start_after_key_suffix,omitempty"`
	MaxKeys              uint32                          `protobuf:"varint,6,opt,name=max_keys,json=maxKeys,proto3" json:"max_keys,omitempty"`
	ColumnsToReturn      []string                        `protobuf:"bytes,7,rep,name=columns_to_return,json=columnsToReturn,proto3" json:"columns_to_return,omitempty"`
	OperationParams      *Ydb_Operations.OperationParams `protobuf:"bytes,8,opt,name=operation_params,json=operationParams,proto3" json:"operation_params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                        `json:"-"`
	XXX_unrecognized     []byte                          `json:"-"`
	XXX_sizecache        int32                           `json:"-"`
}

func (m *S3ListingRequest) Reset()         { *m = S3ListingRequest{} }
func (m *S3ListingRequest) String() string { return proto.CompactTextString(m) }
func (*S3ListingRequest) ProtoMessage()    {}
func (*S3ListingRequest) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a251a79ecaf5049, []int{0}
}

func (m *S3ListingRequest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3ListingRequest.Unmarshal(m, b)
}
func (m *S3ListingRequest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3ListingRequest.Marshal(b, m, deterministic)
}
func (m *S3ListingRequest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3ListingRequest.Merge(m, src)
}
func (m *S3ListingRequest) XXX_Size() int {
	return xxx_messageInfo_S3ListingRequest.Size(m)
}
func (m *S3ListingRequest) XXX_DiscardUnknown() {
	xxx_messageInfo_S3ListingRequest.DiscardUnknown(m)
}

var xxx_messageInfo_S3ListingRequest proto.InternalMessageInfo

func (m *S3ListingRequest) GetTableName() string {
	if m != nil {
		return m.TableName
	}
	return ""
}

func (m *S3ListingRequest) GetKeyPrefix() *Ydb.TypedValue {
	if m != nil {
		return m.KeyPrefix
	}
	return nil
}

func (m *S3ListingRequest) GetPathColumnPrefix() string {
	if m != nil {
		return m.PathColumnPrefix
	}
	return ""
}

func (m *S3ListingRequest) GetPathColumnDelimiter() string {
	if m != nil {
		return m.PathColumnDelimiter
	}
	return ""
}

func (m *S3ListingRequest) GetStartAfterKeySuffix() *Ydb.TypedValue {
	if m != nil {
		return m.StartAfterKeySuffix
	}
	return nil
}

func (m *S3ListingRequest) GetMaxKeys() uint32 {
	if m != nil {
		return m.MaxKeys
	}
	return 0
}

func (m *S3ListingRequest) GetColumnsToReturn() []string {
	if m != nil {
		return m.ColumnsToReturn
	}
	return nil
}

func (m *S3ListingRequest) GetOperationParams() *Ydb_Operations.OperationParams {
	if m != nil {
		return m.OperationParams
	}
	return nil
}

type S3ListingResponse struct {
	Operation            *Ydb_Operations.Operation `protobuf:"bytes,1,opt,name=operation,proto3" json:"operation,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                  `json:"-"`
	XXX_unrecognized     []byte                    `json:"-"`
	XXX_sizecache        int32                     `json:"-"`
}

func (m *S3ListingResponse) Reset()         { *m = S3ListingResponse{} }
func (m *S3ListingResponse) String() string { return proto.CompactTextString(m) }
func (*S3ListingResponse) ProtoMessage()    {}
func (*S3ListingResponse) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a251a79ecaf5049, []int{1}
}

func (m *S3ListingResponse) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3ListingResponse.Unmarshal(m, b)
}
func (m *S3ListingResponse) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3ListingResponse.Marshal(b, m, deterministic)
}
func (m *S3ListingResponse) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3ListingResponse.Merge(m, src)
}
func (m *S3ListingResponse) XXX_Size() int {
	return xxx_messageInfo_S3ListingResponse.Size(m)
}
func (m *S3ListingResponse) XXX_DiscardUnknown() {
	xxx_messageInfo_S3ListingResponse.DiscardUnknown(m)
}

var xxx_messageInfo_S3ListingResponse proto.InternalMessageInfo

func (m *S3ListingResponse) GetOperation() *Ydb_Operations.Operation {
	if m != nil {
		return m.Operation
	}
	return nil
}

type S3ListingResult struct {
	CommonPrefixes       *Ydb.ResultSet `protobuf:"bytes,1,opt,name=common_prefixes,json=commonPrefixes,proto3" json:"common_prefixes,omitempty"`
	Contents             *Ydb.ResultSet `protobuf:"bytes,2,opt,name=contents,proto3" json:"contents,omitempty"`
	KeySuffixSize        uint32         `protobuf:"varint,3,opt,name=key_suffix_size,json=keySuffixSize,proto3" json:"key_suffix_size,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *S3ListingResult) Reset()         { *m = S3ListingResult{} }
func (m *S3ListingResult) String() string { return proto.CompactTextString(m) }
func (*S3ListingResult) ProtoMessage()    {}
func (*S3ListingResult) Descriptor() ([]byte, []int) {
	return fileDescriptor_1a251a79ecaf5049, []int{2}
}

func (m *S3ListingResult) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_S3ListingResult.Unmarshal(m, b)
}
func (m *S3ListingResult) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_S3ListingResult.Marshal(b, m, deterministic)
}
func (m *S3ListingResult) XXX_Merge(src proto.Message) {
	xxx_messageInfo_S3ListingResult.Merge(m, src)
}
func (m *S3ListingResult) XXX_Size() int {
	return xxx_messageInfo_S3ListingResult.Size(m)
}
func (m *S3ListingResult) XXX_DiscardUnknown() {
	xxx_messageInfo_S3ListingResult.DiscardUnknown(m)
}

var xxx_messageInfo_S3ListingResult proto.InternalMessageInfo

func (m *S3ListingResult) GetCommonPrefixes() *Ydb.ResultSet {
	if m != nil {
		return m.CommonPrefixes
	}
	return nil
}

func (m *S3ListingResult) GetContents() *Ydb.ResultSet {
	if m != nil {
		return m.Contents
	}
	return nil
}

func (m *S3ListingResult) GetKeySuffixSize() uint32 {
	if m != nil {
		return m.KeySuffixSize
	}
	return 0
}

func init() {
	proto.RegisterType((*S3ListingRequest)(nil), "Ydb.S3Internal.S3ListingRequest")
	proto.RegisterType((*S3ListingResponse)(nil), "Ydb.S3Internal.S3ListingResponse")
	proto.RegisterType((*S3ListingResult)(nil), "Ydb.S3Internal.S3ListingResult")
}

func init() { proto.RegisterFile("ydb_s3_internal.proto", fileDescriptor_1a251a79ecaf5049) }

var fileDescriptor_1a251a79ecaf5049 = []byte{
	// 486 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0xc1, 0x6f, 0xd3, 0x30,
	0x14, 0xc6, 0x15, 0x0a, 0x5b, 0xeb, 0xa9, 0x4d, 0xe7, 0x69, 0x28, 0x9d, 0x84, 0xa8, 0x7a, 0x40,
	0x65, 0x9a, 0x52, 0xd4, 0x1e, 0x76, 0x66, 0xec, 0x02, 0x9b, 0xa0, 0x72, 0x26, 0x24, 0x4e, 0x96,
	0xd3, 0xbc, 0x82, 0x95, 0xd8, 0x0e, 0xb6, 0x83, 0x9a, 0xfd, 0x2d, 0x88, 0xbf, 0x93, 0x23, 0x8a,
	0x93, 0x36, 0x05, 0xb1, 0xdd, 0xaa, 0xf7, 0xfd, 0xbe, 0xaf, 0xaf, 0xfd, 0x9e, 0xd1, 0x69, 0x99,
	0xc4, 0xd4, 0x2c, 0x28, 0x97, 0x16, 0xb4, 0x64, 0x59, 0x98, 0x6b, 0x65, 0x15, 0x1e, 0x7c, 0x49,
	0xe2, 0x30, 0x5a, 0xbc, 0x6f, 0xa6, 0x67, 0xaf, 0x53, 0x9e, 0x72, 0xa1, 0x67, 0x79, 0x11, 0x67,
	0x7c, 0x35, 0x63, 0x39, 0x9f, 0x39, 0xd0, 0xcc, 0x2a, 0xff, 0x4a, 0x09, 0xa1, 0x64, 0x6d, 0x3d,
	0xbb, 0x78, 0x14, 0x55, 0x39, 0x68, 0x66, 0xf9, 0x8e, 0x9e, 0x3e, 0x4a, 0xff, 0x60, 0x59, 0x01,
	0x35, 0x39, 0xf9, 0xd9, 0x41, 0xc3, 0x68, 0x71, 0xcb, 0x8d, 0xe5, 0xf2, 0x2b, 0x81, 0xef, 0x05,
	0x18, 0x8b, 0x5f, 0x20, 0x64, 0x59, 0x9c, 0x01, 0x95, 0x4c, 0x40, 0xe0, 0x8d, 0xbd, 0x69, 0x8f,
	0xf4, 0xdc, 0xe4, 0x23, 0x13, 0x80, 0x43, 0x84, 0x52, 0x28, 0x69, 0xae, 0x61, 0xcd, 0x37, 0xc1,
	0x93, 0xb1, 0x37, 0x3d, 0x9a, 0xfb, 0x61, 0xf5, 0xdb, 0xee, 0xca, 0x1c, 0x92, 0xcf, 0x55, 0x3c,
	0xe9, 0xa5, 0x50, 0x2e, 0x1d, 0x81, 0x2f, 0x10, 0xce, 0x99, 0xfd, 0x46, 0x57, 0x2a, 0x2b, 0x84,
	0xdc, 0xfa, 0x3a, 0x2e, 0x76, 0x58, 0x29, 0xef, 0x9c, 0xd0, 0xd0, 0x73, 0x74, 0xba, 0x4f, 0x27,
	0x90, 0x71, 0xc1, 0x2d, 0xe8, 0xe0, 0xa9, 0x33, 0x9c, 0xb4, 0x86, 0xeb, 0xad, 0x84, 0xaf, 0xd1,
	0x73, 0x63, 0x99, 0xb6, 0x94, 0xad, 0x2d, 0x68, 0x5a, 0x6d, 0x67, 0x8a, 0x75, 0xf5, 0x2d, 0xcf,
	0xfe, 0xbf, 0xdd, 0x89, 0xc3, 0xdf, 0x56, 0xf4, 0x0d, 0x94, 0x91, 0x63, 0xf1, 0x08, 0x75, 0x05,
	0xdb, 0x54, 0x6e, 0x13, 0x1c, 0x8c, 0xbd, 0x69, 0x9f, 0x1c, 0x0a, 0xb6, 0xb9, 0x81, 0xd2, 0xe0,
	0x73, 0x74, 0x5c, 0xef, 0x63, 0xa8, 0x55, 0x54, 0x83, 0x2d, 0xb4, 0x0c, 0x0e, 0xc7, 0x9d, 0x69,
	0x8f, 0xf8, 0x8d, 0x70, 0xa7, 0x88, 0x1b, 0xe3, 0x0f, 0x68, 0xb8, 0xeb, 0x83, 0xe6, 0x4c, 0x33,
	0x61, 0x82, 0xae, 0x5b, 0xe3, 0xa5, 0x5b, 0xe3, 0xd3, 0x56, 0x34, 0xed, 0xc7, 0xa5, 0xc3, 0x88,
	0xaf, 0xfe, 0x1e, 0x4c, 0x6e, 0xd1, 0xf1, 0x5e, 0x3b, 0x26, 0x57, 0xd2, 0x00, 0xbe, 0x44, 0xbd,
	0x1d, 0xe7, 0xda, 0x39, 0x9a, 0x8f, 0x1e, 0x4c, 0x26, 0x2d, 0x3b, 0xf9, 0xe5, 0x21, 0x7f, 0x3f,
	0xae, 0xc8, 0x2c, 0xbe, 0x44, 0x7e, 0x7d, 0x68, 0x4d, 0x2f, 0x60, 0x9a, 0xc8, 0x81, 0x8b, 0xac,
	0xa9, 0x08, 0x2c, 0x19, 0xd4, 0xd8, 0xb2, 0xa1, 0xf0, 0x39, 0xea, 0xae, 0x94, 0xb4, 0x20, 0xad,
	0x69, 0x6e, 0xe0, 0x5f, 0xc7, 0x4e, 0xc7, 0xaf, 0x90, 0xdf, 0x76, 0x42, 0x0d, 0xbf, 0x07, 0x57,
	0x7f, 0x9f, 0xf4, 0xd3, 0xed, 0xbf, 0x1f, 0xf1, 0x7b, 0xb8, 0x7a, 0x83, 0x46, 0xba, 0x08, 0x4b,
	0x26, 0x13, 0xd8, 0x84, 0x65, 0x12, 0x87, 0x7b, 0x6f, 0xe8, 0x6a, 0xd8, 0xbe, 0x9c, 0xa5, 0x3b,
	0xe6, 0xdf, 0x9e, 0x17, 0x1f, 0xb8, 0x33, 0x5e, 0xfc, 0x09, 0x00, 0x00, 0xff, 0xff, 0x0a, 0xb0,
	0x58, 0x00, 0x72, 0x03, 0x00, 0x00,
}

const ()