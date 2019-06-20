// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PutUserInfo

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
	Sessionid string `protobuf:"bytes,1,opt,name=sessionid,proto3" json:"sessionid,omitempty"`
	//更换的用户名
	Username             string   `protobuf:"bytes,2,opt,name=username,proto3" json:"username,omitempty"`
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

func (m *Request) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

//srv-->web
type Response struct {
	//错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//用户名
	Username             string   `protobuf:"bytes,3,opt,name=username,proto3" json:"username,omitempty"`
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

func (m *Response) GetUsername() string {
	if m != nil {
		return m.Username
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PutUserInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PutUserInfo.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 200 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xe2, 0x92, 0x2e, 0x28, 0xca, 0x2f,
	0xc9, 0xd7, 0x4f, 0xad, 0x48, 0xcc, 0x2d, 0xc8, 0x49, 0x85, 0xd1, 0x7a, 0x60, 0x51, 0x21, 0x89,
	0xf4, 0x7c, 0xbd, 0xdc, 0xcc, 0xe4, 0xa2, 0x7c, 0xbd, 0xe2, 0xa2, 0x32, 0xbd, 0x80, 0xd2, 0x92,
	0xd0, 0xe2, 0xd4, 0x22, 0xcf, 0xbc, 0xb4, 0x7c, 0x25, 0x67, 0x2e, 0xf6, 0xa0, 0xd4, 0xc2, 0xd2,
	0xd4, 0xe2, 0x12, 0x21, 0x19, 0x2e, 0xce, 0xe2, 0xd4, 0xe2, 0xe2, 0xcc, 0xfc, 0xbc, 0xcc, 0x14,
	0x09, 0x46, 0x05, 0x46, 0x0d, 0xce, 0x20, 0x84, 0x80, 0x90, 0x14, 0x17, 0x47, 0x69, 0x71, 0x6a,
	0x51, 0x5e, 0x62, 0x6e, 0xaa, 0x04, 0x13, 0x58, 0x12, 0xce, 0x57, 0x0a, 0xe1, 0xe2, 0x08, 0x4a,
	0x2d, 0x2e, 0xc8, 0xcf, 0x2b, 0x4e, 0x15, 0x12, 0xe1, 0x62, 0x4d, 0x2d, 0x2a, 0xca, 0xcb, 0x87,
	0x9a, 0x00, 0xe1, 0x08, 0x89, 0x71, 0xb1, 0xa5, 0x16, 0x15, 0xe5, 0x16, 0xa7, 0x43, 0xf5, 0x42,
	0x79, 0x28, 0xa6, 0x32, 0xa3, 0x9a, 0x6a, 0x94, 0xc8, 0xc5, 0xee, 0x0a, 0xf1, 0x85, 0x50, 0x18,
	0x17, 0x37, 0x92, 0xa3, 0x85, 0x14, 0xf5, 0x70, 0xf9, 0x47, 0x0f, 0xea, 0x19, 0x29, 0x25, 0x7c,
	0x4a, 0x20, 0x4e, 0x55, 0x62, 0x48, 0x62, 0x03, 0x07, 0x8f, 0x31, 0x20, 0x00, 0x00, 0xff, 0xff,
	0x17, 0xea, 0x7b, 0x34, 0x3d, 0x01, 0x00, 0x00,
}
