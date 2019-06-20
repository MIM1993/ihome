// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_PostAvatar

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
	//文件大小
	Filesize int64 `protobuf:"varint,2,opt,name=filesize,proto3" json:"filesize,omitempty"`
	//文件名
	Filename string `protobuf:"bytes,3,opt,name=filename,proto3" json:"filename,omitempty"`
	//二进制图片
	Buffer               []byte   `protobuf:"bytes,4,opt,name=buffer,proto3" json:"buffer,omitempty"`
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

func (m *Request) GetFilesize() int64 {
	if m != nil {
		return m.Filesize
	}
	return 0
}

func (m *Request) GetFilename() string {
	if m != nil {
		return m.Filename
	}
	return ""
}

func (m *Request) GetBuffer() []byte {
	if m != nil {
		return m.Buffer
	}
	return nil
}

//srv--->web
type Response struct {
	//错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//fileid
	Fileid               string   `protobuf:"bytes,3,opt,name=fileid,proto3" json:"fileid,omitempty"`
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

func (m *Response) GetFileid() string {
	if m != nil {
		return m.Fileid
	}
	return ""
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.PostAvatar.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.PostAvatar.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 230 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xb1, 0x4e, 0xc3, 0x30,
	0x10, 0x86, 0x09, 0x85, 0xb6, 0x39, 0x31, 0x59, 0x08, 0xa2, 0xc2, 0x10, 0x32, 0x65, 0x32, 0x12,
	0x3c, 0x01, 0x03, 0x7b, 0x65, 0x76, 0x24, 0x97, 0x5e, 0x2a, 0x4b, 0xb5, 0x2f, 0xdc, 0xb9, 0x05,
	0xf1, 0xf4, 0x28, 0x8e, 0xdb, 0x4e, 0x30, 0x59, 0xdf, 0x6f, 0x9f, 0x7f, 0x7f, 0x86, 0xbb, 0x9e,
	0x29, 0xd2, 0x23, 0x7e, 0x5b, 0xdf, 0x6f, 0xf1, 0xb0, 0xea, 0x94, 0xaa, 0xdb, 0x0d, 0x69, 0xef,
	0x3e, 0x98, 0xb4, 0xf0, 0x5e, 0x2f, 0x49, 0xe2, 0xcb, 0xde, 0x46, 0xcb, 0xcd, 0x17, 0xcc, 0x0c,
	0x7e, 0xee, 0x50, 0xa2, 0xba, 0x87, 0x52, 0x50, 0xc4, 0x51, 0x70, 0xeb, 0xaa, 0xa8, 0x8b, 0xb6,
	0x34, 0xa7, 0x40, 0x2d, 0x60, 0xde, 0xb9, 0x2d, 0x8a, 0xfb, 0xc1, 0xea, 0xbc, 0x2e, 0xda, 0x89,
	0x39, 0xf2, 0x61, 0x2f, 0x58, 0x8f, 0xd5, 0x24, 0x0d, 0x1e, 0x59, 0xdd, 0xc0, 0x74, 0xb5, 0xeb,
	0x3a, 0xe4, 0xea, 0xa2, 0x2e, 0xda, 0x2b, 0x93, 0xa9, 0x59, 0xc2, 0xdc, 0xa0, 0xf4, 0x14, 0x04,
	0xd5, 0x35, 0x5c, 0x22, 0x73, 0xa0, 0xdc, 0x3a, 0xc2, 0x30, 0x89, 0xcc, 0x5e, 0x36, 0xa9, 0xaf,
	0x34, 0x99, 0x86, 0x7c, 0xb8, 0xdd, 0xad, 0x73, 0x57, 0xa6, 0xa7, 0x77, 0x98, 0xbd, 0x8e, 0xd2,
	0xea, 0x0d, 0xe0, 0xe4, 0xa8, 0x6a, 0xfd, 0x87, 0xbd, 0xce, 0xea, 0x8b, 0x87, 0x7f, 0x4e, 0x8c,
	0x6f, 0x6c, 0xce, 0x56, 0xd3, 0xf4, 0x95, 0xcf, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0x7e, 0x37,
	0xcf, 0xaa, 0x69, 0x01, 0x00, 0x00,
}