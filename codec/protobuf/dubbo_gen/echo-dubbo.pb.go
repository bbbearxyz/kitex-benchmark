// Code generated by protoc-gen-go. DO NOT EDIT.
// source: echo-dubbo.proto

package dubbo

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
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type Request struct {
	Action               string   `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Field1               string   `protobuf:"bytes,2,opt,name=field1,proto3" json:"field1,omitempty"`
	Field2               string   `protobuf:"bytes,3,opt,name=field2,proto3" json:"field2,omitempty"`
	Field3               string   `protobuf:"bytes,4,opt,name=field3,proto3" json:"field3,omitempty"`
	Field4               string   `protobuf:"bytes,5,opt,name=field4,proto3" json:"field4,omitempty"`
	Field5               string   `protobuf:"bytes,6,opt,name=field5,proto3" json:"field5,omitempty"`
	Field6               string   `protobuf:"bytes,7,opt,name=field6,proto3" json:"field6,omitempty"`
	Field7               string   `protobuf:"bytes,8,opt,name=field7,proto3" json:"field7,omitempty"`
	Field8               string   `protobuf:"bytes,9,opt,name=field8,proto3" json:"field8,omitempty"`
	Field9               string   `protobuf:"bytes,10,opt,name=field9,proto3" json:"field9,omitempty"`
	Field10              string   `protobuf:"bytes,11,opt,name=field10,proto3" json:"field10,omitempty"`
	Time                 int64    `protobuf:"varint,12,opt,name=time,proto3" json:"time,omitempty"`
	Length               int64    `protobuf:"varint,13,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb415af32831e763, []int{0}
}

func (m *Request) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Request.Unmarshal(m, b)
}
func (m *Request) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Request.Marshal(b, m, deterministic)
}
func (m *Request) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Request.Merge(m, src)
}
func (m *Request) XXX_Size() int {
	return xxx_messageInfo_Request.Size(m)
}
func (m *Request) XXX_DiscardUnknown() {
	xxx_messageInfo_Request.DiscardUnknown(m)
}

var xxx_messageInfo_Request proto.InternalMessageInfo

func (m *Request) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Request) GetField1() string {
	if m != nil {
		return m.Field1
	}
	return ""
}

func (m *Request) GetField2() string {
	if m != nil {
		return m.Field2
	}
	return ""
}

func (m *Request) GetField3() string {
	if m != nil {
		return m.Field3
	}
	return ""
}

func (m *Request) GetField4() string {
	if m != nil {
		return m.Field4
	}
	return ""
}

func (m *Request) GetField5() string {
	if m != nil {
		return m.Field5
	}
	return ""
}

func (m *Request) GetField6() string {
	if m != nil {
		return m.Field6
	}
	return ""
}

func (m *Request) GetField7() string {
	if m != nil {
		return m.Field7
	}
	return ""
}

func (m *Request) GetField8() string {
	if m != nil {
		return m.Field8
	}
	return ""
}

func (m *Request) GetField9() string {
	if m != nil {
		return m.Field9
	}
	return ""
}

func (m *Request) GetField10() string {
	if m != nil {
		return m.Field10
	}
	return ""
}

func (m *Request) GetTime() int64 {
	if m != nil {
		return m.Time
	}
	return 0
}

func (m *Request) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Response struct {
	Action               string   `protobuf:"bytes,1,opt,name=Action,proto3" json:"Action,omitempty"`
	Msg                  string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	IsEnd                bool     `protobuf:"varint,3,opt,name=isEnd,proto3" json:"isEnd,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_eb415af32831e763, []int{1}
}

func (m *Response) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Response.Unmarshal(m, b)
}
func (m *Response) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Response.Marshal(b, m, deterministic)
}
func (m *Response) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Response.Merge(m, src)
}
func (m *Response) XXX_Size() int {
	return xxx_messageInfo_Response.Size(m)
}
func (m *Response) XXX_DiscardUnknown() {
	xxx_messageInfo_Response.DiscardUnknown(m)
}

var xxx_messageInfo_Response proto.InternalMessageInfo

func (m *Response) GetAction() string {
	if m != nil {
		return m.Action
	}
	return ""
}

func (m *Response) GetMsg() string {
	if m != nil {
		return m.Msg
	}
	return ""
}

func (m *Response) GetIsEnd() bool {
	if m != nil {
		return m.IsEnd
	}
	return false
}

func init() {
	proto.RegisterType((*Request)(nil), "protobuf.Request")
	proto.RegisterType((*Response)(nil), "protobuf.Response")
}

func init() {
	proto.RegisterFile("echo-dubbo.proto", fileDescriptor_eb415af32831e763)
}

var fileDescriptor_eb415af32831e763 = []byte{
	// 291 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0xd2, 0xcf, 0x4f, 0x83, 0x30,
	0x14, 0x07, 0x70, 0x19, 0x0c, 0xd8, 0x53, 0x93, 0xd9, 0x18, 0xf3, 0xe2, 0x69, 0xe1, 0xc4, 0x45,
	0x36, 0x61, 0xbf, 0x88, 0x27, 0x4d, 0x76, 0xf1, 0xc8, 0x3c, 0x79, 0x13, 0xe8, 0x06, 0xc9, 0xa0,
	0x13, 0xca, 0xdf, 0xe7, 0xbf, 0x66, 0x68, 0x59, 0xd3, 0x8b, 0x89, 0x27, 0xde, 0xf7, 0x03, 0xb4,
	0xe9, 0x7b, 0x85, 0x29, 0xcd, 0x0a, 0xf6, 0x94, 0x77, 0x69, 0xca, 0x82, 0x73, 0xc3, 0x38, 0x23,
	0xae, 0x78, 0xa4, 0xdd, 0xc1, 0xfb, 0x19, 0x81, 0x93, 0xd0, 0xef, 0x8e, 0xb6, 0x9c, 0x3c, 0x80,
	0xfd, 0x9a, 0xf1, 0x92, 0xd5, 0x68, 0xcc, 0x0c, 0x7f, 0x92, 0x0c, 0xa9, 0xf7, 0x43, 0x49, 0x4f,
	0xf9, 0x33, 0x8e, 0xa4, 0xcb, 0xa4, 0x3c, 0x44, 0x53, 0xf3, 0x50, 0x79, 0x84, 0x96, 0xe6, 0x91,
	0xf2, 0x25, 0x8e, 0x35, 0x5f, 0x2a, 0x5f, 0xa1, 0xad, 0xf9, 0x4a, 0xf9, 0x1a, 0x1d, 0xcd, 0xd7,
	0xca, 0x37, 0xe8, 0x6a, 0xbe, 0x51, 0xbe, 0xc5, 0x89, 0xe6, 0x5b, 0xe5, 0x31, 0x82, 0xe6, 0x31,
	0x41, 0x70, 0xe4, 0x49, 0x16, 0x78, 0x2d, 0x5e, 0x5c, 0x22, 0x21, 0x60, 0xf1, 0xb2, 0xa2, 0x78,
	0x33, 0x33, 0x7c, 0x33, 0x11, 0x75, 0xbf, 0xca, 0x89, 0xd6, 0x47, 0x5e, 0xe0, 0xad, 0xd0, 0x21,
	0x79, 0xef, 0xe0, 0x26, 0xb4, 0x3d, 0xb3, 0xba, 0xa5, 0x7f, 0x76, 0x70, 0x0a, 0x66, 0xd5, 0x1e,
	0x87, 0xf6, 0xf5, 0x25, 0xb9, 0x87, 0x71, 0xd9, 0xee, 0xea, 0x5c, 0xb4, 0xce, 0x4d, 0x64, 0x08,
	0x1b, 0xb0, 0x76, 0x59, 0xc1, 0xc8, 0x1c, 0xac, 0x3d, 0xad, 0x73, 0x72, 0x17, 0x5c, 0x06, 0x15,
	0x0c, 0x43, 0x7a, 0x24, 0x3a, 0xc9, 0x6d, 0xbd, 0x2b, 0x12, 0x03, 0xec, 0x79, 0x43, 0xbf, 0xaa,
	0x8f, 0x7e, 0x90, 0xff, 0xfd, 0xcd, 0x37, 0x16, 0xc6, 0x1b, 0x7c, 0xba, 0xc1, 0xfc, 0x45, 0xdc,
	0x8e, 0xd4, 0x16, 0x1f, 0x45, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd9, 0x29, 0x49, 0xdd, 0x32,
	0x02, 0x00, 0x00,
}
