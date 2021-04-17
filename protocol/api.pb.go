// Code generated by protoc-gen-go.
// source: api.proto
// DO NOT EDIT!

/*
Package protocol is a generated protocol buffer package.

It is generated from these files:
	api.proto

It has these top-level messages:
	OkResponse
	ErrorResponse
	OpenDBRequest
	OpenDBResponse
	CloseDBRequest
	ListDBRequest
	ListDBResponse
	DeleteDBRequest
	GetRequest
	GetResponse
	SetRequest
	DeleteRequest
	TouchRequest
	CacheRequest
	OptionRequest
*/
package protocol

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
const _ = proto.ProtoPackageIsVersion1

// empty response used to indicate OK.
type OkResponse struct {
}

func (m *OkResponse) Reset()                    { *m = OkResponse{} }
func (m *OkResponse) String() string            { return proto.CompactTextString(m) }
func (*OkResponse) ProtoMessage()               {}
func (*OkResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ErrorResponse struct {
	Code uint32 `protobuf:"varint,1,opt,name=code" json:"code,omitempty"`
	Info string `protobuf:"bytes,2,opt,name=info" json:"info,omitempty"`
}

func (m *ErrorResponse) Reset()                    { *m = ErrorResponse{} }
func (m *ErrorResponse) String() string            { return proto.CompactTextString(m) }
func (*ErrorResponse) ProtoMessage()               {}
func (*ErrorResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type OpenDBRequest struct {
	Name            string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Template        string `protobuf:"bytes,2,opt,name=template" json:"template,omitempty"`
	CreateIfMissing bool   `protobuf:"varint,3,opt,name=create_if_missing,json=createIfMissing" json:"create_if_missing,omitempty"`
	ErrorIfExists   bool   `protobuf:"varint,4,opt,name=error_if_exists,json=errorIfExists" json:"error_if_exists,omitempty"`
}

func (m *OpenDBRequest) Reset()                    { *m = OpenDBRequest{} }
func (m *OpenDBRequest) String() string            { return proto.CompactTextString(m) }
func (*OpenDBRequest) ProtoMessage()               {}
func (*OpenDBRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

type OpenDBResponse struct {
	Db      uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Version uint64 `protobuf:"varint,2,opt,name=version" json:"version,omitempty"`
}

func (m *OpenDBResponse) Reset()                    { *m = OpenDBResponse{} }
func (m *OpenDBResponse) String() string            { return proto.CompactTextString(m) }
func (*OpenDBResponse) ProtoMessage()               {}
func (*OpenDBResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type CloseDBRequest struct {
	Db uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
}

func (m *CloseDBRequest) Reset()                    { *m = CloseDBRequest{} }
func (m *CloseDBRequest) String() string            { return proto.CompactTextString(m) }
func (*CloseDBRequest) ProtoMessage()               {}
func (*CloseDBRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type ListDBRequest struct {
}

func (m *ListDBRequest) Reset()                    { *m = ListDBRequest{} }
func (m *ListDBRequest) String() string            { return proto.CompactTextString(m) }
func (*ListDBRequest) ProtoMessage()               {}
func (*ListDBRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type ListDBResponse struct {
	// name to version
	Databases map[string]uint32 `protobuf:"bytes,1,rep,name=databases" json:"databases,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"varint,2,opt,name=value"`
}

func (m *ListDBResponse) Reset()                    { *m = ListDBResponse{} }
func (m *ListDBResponse) String() string            { return proto.CompactTextString(m) }
func (*ListDBResponse) ProtoMessage()               {}
func (*ListDBResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *ListDBResponse) GetDatabases() map[string]uint32 {
	if m != nil {
		return m.Databases
	}
	return nil
}

type DeleteDBRequest struct {
	Name string `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
}

func (m *DeleteDBRequest) Reset()                    { *m = DeleteDBRequest{} }
func (m *DeleteDBRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteDBRequest) ProtoMessage()               {}
func (*DeleteDBRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

type GetRequest struct {
	Db   uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// protocol.FieldAny;
	// protocol.FieldTree;
	// protocol.FieldBinary.
	Type uint32 `protobuf:"varint,3,opt,name=type" json:"type,omitempty"`
}

func (m *GetRequest) Reset()                    { *m = GetRequest{} }
func (m *GetRequest) String() string            { return proto.CompactTextString(m) }
func (*GetRequest) ProtoMessage()               {}
func (*GetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

type GetResponse struct {
	// protocol.Marshaled
	Value []byte `protobuf:"bytes,1,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *GetResponse) Reset()                    { *m = GetResponse{} }
func (m *GetResponse) String() string            { return proto.CompactTextString(m) }
func (*GetResponse) ProtoMessage()               {}
func (*GetResponse) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

type SetRequest struct {
	Db   uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// protocol.Marshaled
	Value []byte `protobuf:"bytes,3,opt,name=value,proto3" json:"value,omitempty"`
}

func (m *SetRequest) Reset()                    { *m = SetRequest{} }
func (m *SetRequest) String() string            { return proto.CompactTextString(m) }
func (*SetRequest) ProtoMessage()               {}
func (*SetRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

type DeleteRequest struct {
	Db   uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *DeleteRequest) Reset()                    { *m = DeleteRequest{} }
func (m *DeleteRequest) String() string            { return proto.CompactTextString(m) }
func (*DeleteRequest) ProtoMessage()               {}
func (*DeleteRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

type TouchRequest struct {
	Db   uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
}

func (m *TouchRequest) Reset()                    { *m = TouchRequest{} }
func (m *TouchRequest) String() string            { return proto.CompactTextString(m) }
func (*TouchRequest) ProtoMessage()               {}
func (*TouchRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

type CacheRequest struct {
	Db   uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path" json:"path,omitempty"`
	// -2, never cached indirectly;
	// -1, deleted;
	//  0, default timeout duration;
	//  positive, timeout duration.
	Timeout int64 `protobuf:"zigzag64,3,opt,name=timeout" json:"timeout,omitempty"`
}

func (m *CacheRequest) Reset()                    { *m = CacheRequest{} }
func (m *CacheRequest) String() string            { return proto.CompactTextString(m) }
func (*CacheRequest) ProtoMessage()               {}
func (*CacheRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

type OptionRequest struct {
	Db    uint32 `protobuf:"varint,1,opt,name=db" json:"db,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Value string `protobuf:"bytes,3,opt,name=value" json:"value,omitempty"`
}

func (m *OptionRequest) Reset()                    { *m = OptionRequest{} }
func (m *OptionRequest) String() string            { return proto.CompactTextString(m) }
func (*OptionRequest) ProtoMessage()               {}
func (*OptionRequest) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func init() {
	proto.RegisterType((*OkResponse)(nil), "protocol.OkResponse")
	proto.RegisterType((*ErrorResponse)(nil), "protocol.ErrorResponse")
	proto.RegisterType((*OpenDBRequest)(nil), "protocol.OpenDBRequest")
	proto.RegisterType((*OpenDBResponse)(nil), "protocol.OpenDBResponse")
	proto.RegisterType((*CloseDBRequest)(nil), "protocol.CloseDBRequest")
	proto.RegisterType((*ListDBRequest)(nil), "protocol.ListDBRequest")
	proto.RegisterType((*ListDBResponse)(nil), "protocol.ListDBResponse")
	proto.RegisterType((*DeleteDBRequest)(nil), "protocol.DeleteDBRequest")
	proto.RegisterType((*GetRequest)(nil), "protocol.GetRequest")
	proto.RegisterType((*GetResponse)(nil), "protocol.GetResponse")
	proto.RegisterType((*SetRequest)(nil), "protocol.SetRequest")
	proto.RegisterType((*DeleteRequest)(nil), "protocol.DeleteRequest")
	proto.RegisterType((*TouchRequest)(nil), "protocol.TouchRequest")
	proto.RegisterType((*CacheRequest)(nil), "protocol.CacheRequest")
	proto.RegisterType((*OptionRequest)(nil), "protocol.OptionRequest")
}

var fileDescriptor0 = []byte{
	// 449 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x94, 0x53, 0xcf, 0x6f, 0xd3, 0x30,
	0x14, 0x56, 0xda, 0x02, 0xed, 0x5b, 0x93, 0x82, 0xc5, 0x21, 0xda, 0x69, 0x32, 0x02, 0x26, 0x0e,
	0x3d, 0x6c, 0x07, 0xd0, 0xc4, 0x89, 0xb5, 0xa0, 0x4a, 0x43, 0x93, 0x0c, 0xf7, 0xc9, 0x4d, 0x5f,
	0xa9, 0xb5, 0x24, 0x0e, 0xb1, 0x33, 0xd1, 0xbf, 0x03, 0xf1, 0xff, 0x62, 0x3b, 0x76, 0xb6, 0x22,
	0x31, 0x29, 0xa7, 0x7c, 0xef, 0xcb, 0xf7, 0x7e, 0x3f, 0xc3, 0x84, 0x57, 0x62, 0x5e, 0xd5, 0x52,
	0x4b, 0x32, 0x76, 0x9f, 0x4c, 0xe6, 0x74, 0x0a, 0x70, 0x7d, 0xcb, 0x50, 0x55, 0xb2, 0x54, 0x48,
	0xdf, 0x43, 0xbc, 0xac, 0x6b, 0x59, 0x07, 0x82, 0x10, 0x18, 0x65, 0x72, 0x83, 0x69, 0x74, 0x12,
	0x9d, 0xc6, 0xcc, 0x61, 0xcb, 0x89, 0x72, 0x2b, 0xd3, 0x81, 0xe1, 0x26, 0xcc, 0x61, 0xfa, 0x3b,
	0x82, 0xf8, 0xba, 0xc2, 0x72, 0xf1, 0x89, 0xe1, 0xcf, 0x06, 0x95, 0xb6, 0xaa, 0x92, 0x17, 0xad,
	0xa7, 0x51, 0x59, 0x4c, 0x8e, 0x61, 0xac, 0xb1, 0xa8, 0x72, 0xae, 0xd1, 0x7b, 0x77, 0x36, 0x79,
	0x07, 0x2f, 0xb2, 0x1a, 0x0d, 0xba, 0x11, 0xdb, 0x9b, 0x42, 0x28, 0x25, 0xca, 0x1f, 0xe9, 0xd0,
	0x88, 0xc6, 0x6c, 0xd6, 0xfe, 0x58, 0x6d, 0xbf, 0xb6, 0x34, 0x79, 0x03, 0x33, 0xb4, 0x65, 0x5a,
	0x29, 0xfe, 0x12, 0x4a, 0xab, 0x74, 0xe4, 0x94, 0xb1, 0xa3, 0x57, 0xdb, 0xa5, 0x23, 0xe9, 0x05,
	0x24, 0xa1, 0x28, 0xdf, 0x4f, 0x02, 0x83, 0xcd, 0xda, 0x77, 0x63, 0x10, 0x49, 0xe1, 0xd9, 0x1d,
	0xd6, 0x4a, 0xc8, 0xd2, 0x15, 0x34, 0x62, 0xc1, 0xa4, 0x27, 0x90, 0x5c, 0xe6, 0x52, 0xe1, 0x7d,
	0x47, 0xff, 0xf8, 0xd2, 0x19, 0xc4, 0x57, 0x26, 0x4d, 0x27, 0xa0, 0x7f, 0x22, 0x48, 0x02, 0xe3,
	0xf3, 0x2d, 0x61, 0xb2, 0xe1, 0x9a, 0xaf, 0xb9, 0x42, 0x65, 0x5c, 0x87, 0xa7, 0x47, 0x67, 0x6f,
	0xe7, 0x61, 0xf8, 0xf3, 0x43, 0xf1, 0x7c, 0x11, 0x94, 0xcb, 0x52, 0xd7, 0x7b, 0x76, 0xef, 0x79,
	0xfc, 0x11, 0x92, 0xc3, 0x9f, 0xe4, 0x39, 0x0c, 0x6f, 0x71, 0xef, 0xa7, 0x6b, 0x21, 0x79, 0x09,
	0x4f, 0xee, 0x78, 0xde, 0xb4, 0x93, 0x8d, 0x59, 0x6b, 0x5c, 0x0c, 0x3e, 0x44, 0xf4, 0x35, 0xcc,
	0x16, 0x98, 0xa3, 0xc6, 0x47, 0xb7, 0x43, 0x17, 0x00, 0x5f, 0x50, 0xff, 0xa7, 0x5b, 0xeb, 0x51,
	0x71, 0xbd, 0x0b, 0x5b, 0xb7, 0xd8, 0x72, 0x7a, 0x5f, 0xa1, 0x5b, 0x93, 0xb9, 0x0e, 0x8b, 0xe9,
	0x2b, 0x38, 0x72, 0x51, 0xfc, 0x00, 0xba, 0xaa, 0x6c, 0xa4, 0xa9, 0xaf, 0x8a, 0x7e, 0x06, 0xf8,
	0xd6, 0x2f, 0x55, 0x17, 0x67, 0xf8, 0x30, 0xce, 0x39, 0xc4, 0x6d, 0x67, 0x3d, 0x42, 0xd1, 0x33,
	0x98, 0x7e, 0x97, 0x4d, 0xb6, 0xeb, 0xe3, 0x73, 0x05, 0xd3, 0x4b, 0x9e, 0xed, 0xfa, 0xe4, 0xb1,
	0xb7, 0xa5, 0x45, 0x81, 0xb2, 0xd1, 0xae, 0x68, 0xc2, 0x82, 0x49, 0x57, 0xf6, 0xb1, 0x68, 0x73,
	0x65, 0x8f, 0x84, 0x73, 0xeb, 0x19, 0x3c, 0x78, 0x3c, 0x07, 0x13, 0x98, 0xf8, 0x09, 0xac, 0x9f,
	0xba, 0x63, 0x3a, 0xff, 0x1b, 0x00, 0x00, 0xff, 0xff, 0x35, 0x0f, 0x2e, 0x21, 0xdd, 0x03, 0x00,
	0x00,
}