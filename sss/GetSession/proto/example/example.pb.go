// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetSession

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

//web-->srv
type Request struct {
	//sessionid
	Sessionid            string   `protobuf:"bytes,1,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Request) Reset()         { *m = Request{} }
func (m *Request) String() string { return proto.CompactTextString(m) }
func (*Request) ProtoMessage()    {}
func (*Request) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{0}
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

func (m *Request) GetSessionid() string {
	if m != nil {
		return m.Sessionid
	}
	return ""
}

//srv--web
type Response struct {
	//错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//用户名
	Name                 string   `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Response) Reset()         { *m = Response{} }
func (m *Response) String() string { return proto.CompactTextString(m) }
func (*Response) ProtoMessage()    {}
func (*Response) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1}
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

func (m *Response) GetErrno() string {
	if m != nil {
		return m.Errno
	}
	return ""
}

func (m *Response) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Response) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetSession.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetSession.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 186 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x8f, 0x31, 0x8b, 0xc2, 0x40,
	0x10, 0x85, 0x2f, 0x77, 0x67, 0x62, 0xa6, 0x1c, 0x44, 0x83, 0x5a, 0xc4, 0x34, 0x5a, 0xad, 0xa0,
	0xbf, 0x41, 0x6c, 0xac, 0x92, 0x5e, 0x88, 0x3a, 0x84, 0x80, 0xbb, 0x1b, 0x67, 0xa2, 0xf8, 0xf3,
	0x85, 0xdd, 0x48, 0x2a, 0xad, 0x66, 0xde, 0xc7, 0x2b, 0xbe, 0x07, 0xb3, 0x86, 0x6d, 0x6b, 0xd7,
	0xf4, 0x2c, 0x75, 0x73, 0xa5, 0xf7, 0x55, 0x8e, 0xe2, 0xa4, 0xb2, 0x4a, 0xd7, 0x67, 0xb6, 0x4a,
	0xf8, 0xa1, 0xf6, 0xd4, 0x16, 0x24, 0x52, 0x5b, 0x93, 0x2d, 0x21, 0xca, 0xe9, 0x76, 0x27, 0x69,
	0x71, 0x0e, 0xb1, 0x78, 0x5a, 0x5f, 0x92, 0x20, 0x0d, 0x56, 0x71, 0xde, 0x83, 0xec, 0x00, 0xc3,
	0x9c, 0xa4, 0xb1, 0x46, 0x08, 0x47, 0x30, 0x20, 0x66, 0x63, 0xbb, 0x96, 0x0f, 0x38, 0x86, 0x90,
	0x98, 0xb5, 0x54, 0xc9, 0xaf, 0xc3, 0x5d, 0x42, 0x84, 0x7f, 0x53, 0x6a, 0x4a, 0xfe, 0x1c, 0x75,
	0xff, 0xe6, 0x08, 0xd1, 0xce, 0x0b, 0x62, 0x01, 0xd0, 0xfb, 0x60, 0xaa, 0x3e, 0x98, 0xaa, 0x4e,
	0x73, 0xba, 0xf8, 0xd2, 0xf0, 0x7e, 0xd9, 0xcf, 0x29, 0x74, 0xb3, 0xb7, 0xaf, 0x00, 0x00, 0x00,
	0xff, 0xff, 0x18, 0x1d, 0x91, 0xca, 0x15, 0x01, 0x00, 0x00,
}
