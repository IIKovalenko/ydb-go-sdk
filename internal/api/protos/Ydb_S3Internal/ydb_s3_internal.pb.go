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
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x93, 0x51, 0x6f, 0xd3, 0x3e,
	0x14, 0xc5, 0x95, 0x7f, 0xff, 0x6c, 0xad, 0xa7, 0x36, 0x9d, 0xa7, 0xa1, 0xac, 0x12, 0xa2, 0xea,
	0x03, 0xaa, 0xa6, 0x29, 0x15, 0xed, 0xc3, 0x9e, 0x19, 0x7b, 0x81, 0x4d, 0x50, 0xb9, 0x13, 0x12,
	0x4f, 0x96, 0x93, 0xdc, 0x82, 0x95, 0xd8, 0x0e, 0xb6, 0x83, 0x9a, 0x7d, 0x16, 0xc4, 0xe7, 0xe4,
	0x11, 0xc5, 0x49, 0x9a, 0x82, 0x80, 0xb7, 0xea, 0xde, 0xdf, 0x39, 0x3d, 0xca, 0xb9, 0x46, 0xe7,
	0x65, 0x12, 0x51, 0xb3, 0xa2, 0x5c, 0x5a, 0xd0, 0x92, 0x65, 0x61, 0xae, 0x95, 0x55, 0x78, 0xf4,
	0x31, 0x89, 0xc2, 0xcd, 0xea, 0x4d, 0x33, 0x9d, 0x5c, 0xa5, 0x3c, 0xe5, 0x42, 0x2f, 0xf2, 0x22,
	0xca, 0x78, 0xbc, 0x60, 0x39, 0x5f, 0x38, 0xd0, 0x2c, 0x2a, 0xbd, 0xca, 0x41, 0x33, 0xcb, 0x95,
	0xac, 0xd5, 0x93, 0xf9, 0x3f, 0xe9, 0xaf, 0x2c, 0x2b, 0xa0, 0x26, 0x67, 0xdf, 0x7a, 0x68, 0xbc,
	0x59, 0xdd, 0x73, 0x63, 0xb9, 0xfc, 0x44, 0xe0, 0x4b, 0x01, 0xc6, 0xe2, 0x67, 0x08, 0x59, 0x16,
	0x65, 0x40, 0x25, 0x13, 0x10, 0x78, 0x53, 0x6f, 0x3e, 0x20, 0x03, 0x37, 0x79, 0xc7, 0x04, 0xe0,
	0x10, 0xa1, 0x14, 0x4a, 0x9a, 0x6b, 0xd8, 0xf2, 0x5d, 0xf0, 0xdf, 0xd4, 0x9b, 0x9f, 0x2c, 0xfd,
	0xb0, 0x0a, 0xfc, 0x50, 0xe6, 0x90, 0x7c, 0xa8, 0xec, 0xc9, 0x20, 0x85, 0x72, 0xed, 0x08, 0x7c,
	0x85, 0x70, 0xce, 0xec, 0x67, 0x1a, 0xab, 0xac, 0x10, 0xb2, 0xd5, 0xf5, 0x9c, 0xed, 0xb8, 0xda,
	0xbc, 0x76, 0x8b, 0x86, 0x5e, 0xa2, 0xf3, 0x43, 0x3a, 0x81, 0x8c, 0x0b, 0x6e, 0x41, 0x07, 0xff,
	0x3b, 0xc1, 0x59, 0x27, 0xb8, 0x6d, 0x57, 0xf8, 0x16, 0x3d, 0x35, 0x96, 0x69, 0x4b, 0xd9, 0xd6,
	0x82, 0xa6, 0x55, 0x3a, 0x53, 0x6c, 0xab, 0x7f, 0x79, 0xf2, 0xe7, 0x74, 0x67, 0x0e, 0x7f, 0x55,
	0xd1, 0x77, 0x50, 0x6e, 0x1c, 0x8b, 0x2f, 0x50, 0x5f, 0xb0, 0x5d, 0xa5, 0x36, 0xc1, 0xd1, 0xd4,
	0x9b, 0x0f, 0xc9, 0xb1, 0x60, 0xbb, 0x3b, 0x28, 0x0d, 0xbe, 0x44, 0xa7, 0x75, 0x1e, 0x43, 0xad,
	0xa2, 0x1a, 0x6c, 0xa1, 0x65, 0x70, 0x3c, 0xed, 0xcd, 0x07, 0xc4, 0x6f, 0x16, 0x0f, 0x8a, 0xb8,
	0x31, 0x7e, 0x8b, 0xc6, 0xfb, 0x3e, 0x68, 0xce, 0x34, 0x13, 0x26, 0xe8, 0xbb, 0x18, 0xcf, 0x5d,
	0x8c, 0xf7, 0xed, 0xd2, 0x74, 0x3f, 0xd7, 0x0e, 0x23, 0xbe, 0xfa, 0x75, 0x30, 0xbb, 0x47, 0xa7,
	0x07, 0xed, 0x98, 0x5c, 0x49, 0x03, 0xf8, 0x1a, 0x0d, 0xf6, 0x9c, 0x6b, 0xe7, 0x64, 0x79, 0xf1,
	0x57, 0x67, 0xd2, 0xb1, 0xb3, 0xef, 0x1e, 0xf2, 0x0f, 0xed, 0x8a, 0xcc, 0xe2, 0x6b, 0xe4, 0xc7,
	0x4a, 0x08, 0xd5, 0xf6, 0x02, 0xa6, 0xb1, 0x1c, 0x39, 0xcb, 0x9a, 0xda, 0x80, 0x25, 0xa3, 0x1a,
	0x5b, 0x37, 0x14, 0xbe, 0x44, 0xfd, 0x58, 0x49, 0x0b, 0xd2, 0x9a, 0xe6, 0x06, 0x7e, 0x57, 0xec,
	0xf7, 0xf8, 0x05, 0xf2, 0xbb, 0x4e, 0xa8, 0xe1, 0x8f, 0xe0, 0xea, 0x1f, 0x92, 0x61, 0xda, 0x7e,
	0xfd, 0x0d, 0x7f, 0x84, 0x9b, 0x97, 0x68, 0x12, 0x2b, 0x11, 0x96, 0x4c, 0x26, 0xb0, 0x0b, 0xcb,
	0x24, 0x0a, 0x0f, 0x5e, 0xc6, 0xcd, 0xb8, 0x7b, 0x0f, 0x6b, 0x77, 0xcd, 0x3f, 0x3c, 0x2f, 0x3a,
	0x72, 0x77, 0xbc, 0xfa, 0x19, 0x00, 0x00, 0xff, 0xff, 0x38, 0x85, 0x8d, 0xf8, 0x48, 0x03, 0x00,
	0x00,
}

const ()
