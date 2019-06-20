// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetImageCd

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

//web --> srv
type Request struct {
	Uuid                 string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
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

func (m *Request) GetUuid() string {
	if m != nil {
		return m.Uuid
	}
	return ""
}

//srv-->web
type Response struct {
	//错误码
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//图片 发送
	//pix  []uint8
	Pix []int32 `protobuf:"varint,3,rep,packed,name=pix,proto3" json:"pix,omitempty"`
	//Stride
	Stride               int64          `protobuf:"varint,4,opt,name=stride,proto3" json:"stride,omitempty"`
	Min                  *ResponsePoint `protobuf:"bytes,5,opt,name=min,proto3" json:"min,omitempty"`
	Max                  *ResponsePoint `protobuf:"bytes,6,opt,name=max,proto3" json:"max,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
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

func (m *Response) GetPix() []int32 {
	if m != nil {
		return m.Pix
	}
	return nil
}

func (m *Response) GetStride() int64 {
	if m != nil {
		return m.Stride
	}
	return 0
}

func (m *Response) GetMin() *ResponsePoint {
	if m != nil {
		return m.Min
	}
	return nil
}

func (m *Response) GetMax() *ResponsePoint {
	if m != nil {
		return m.Max
	}
	return nil
}

//point
type ResponsePoint struct {
	X                    int64    `protobuf:"varint,1,opt,name=x,proto3" json:"x,omitempty"`
	Y                    int64    `protobuf:"varint,2,opt,name=y,proto3" json:"y,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ResponsePoint) Reset()         { *m = ResponsePoint{} }
func (m *ResponsePoint) String() string { return proto.CompactTextString(m) }
func (*ResponsePoint) ProtoMessage()    {}
func (*ResponsePoint) Descriptor() ([]byte, []int) {
	return fileDescriptor_097b3f5db5cf5789, []int{1, 0}
}

func (m *ResponsePoint) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ResponsePoint.Unmarshal(m, b)
}
func (m *ResponsePoint) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ResponsePoint.Marshal(b, m, deterministic)
}
func (m *ResponsePoint) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ResponsePoint.Merge(m, src)
}
func (m *ResponsePoint) XXX_Size() int {
	return xxx_messageInfo_ResponsePoint.Size(m)
}
func (m *ResponsePoint) XXX_DiscardUnknown() {
	xxx_messageInfo_ResponsePoint.DiscardUnknown(m)
}

var xxx_messageInfo_ResponsePoint proto.InternalMessageInfo

func (m *ResponsePoint) GetX() int64 {
	if m != nil {
		return m.X
	}
	return 0
}

func (m *ResponsePoint) GetY() int64 {
	if m != nil {
		return m.Y
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetImageCd.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetImageCd.Response")
	proto.RegisterType((*ResponsePoint)(nil), "go.micro.srv.GetImageCd.Response.point")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 257 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x91, 0x41, 0x4b, 0xc4, 0x30,
	0x10, 0x85, 0x8d, 0xd9, 0xee, 0xea, 0xe8, 0x41, 0x06, 0xd1, 0xb0, 0x22, 0xd4, 0x7a, 0xb0, 0xa7,
	0x08, 0xeb, 0xc9, 0xb3, 0x88, 0x78, 0x8d, 0x77, 0xa1, 0xda, 0xa1, 0x04, 0x4c, 0x13, 0x93, 0x54,
	0xb2, 0x3f, 0x5e, 0x90, 0x66, 0xbb, 0x78, 0x52, 0xf1, 0x94, 0xf9, 0x86, 0x79, 0xc9, 0x9b, 0x17,
	0x38, 0x73, 0xde, 0x46, 0x7b, 0x4d, 0xa9, 0x31, 0xee, 0x8d, 0xb6, 0xa7, 0xcc, 0x5d, 0x3c, 0xed,
	0xac, 0x34, 0xfa, 0xd5, 0x5b, 0x19, 0xfc, 0x87, 0x7c, 0xa0, 0xf8, 0x68, 0x9a, 0x8e, 0xee, 0xda,
	0xea, 0x1c, 0x16, 0x8a, 0xde, 0x07, 0x0a, 0x11, 0x11, 0x66, 0xc3, 0xa0, 0x5b, 0xc1, 0x4a, 0x56,
	0xef, 0xab, 0x5c, 0x57, 0x9f, 0x0c, 0xf6, 0x14, 0x05, 0x67, 0xfb, 0x40, 0x78, 0x0c, 0x05, 0x79,
	0xdf, 0xdb, 0x69, 0x62, 0x03, 0x78, 0x02, 0x73, 0xf2, 0xde, 0x84, 0x4e, 0xec, 0xe6, 0xf6, 0x44,
	0x78, 0x04, 0xdc, 0xe9, 0x24, 0x78, 0xc9, 0xeb, 0x42, 0x8d, 0xe5, 0x38, 0x19, 0xa2, 0xd7, 0x2d,
	0x89, 0x59, 0xc9, 0x6a, 0xae, 0x26, 0xc2, 0x5b, 0xe0, 0x46, 0xf7, 0xa2, 0x28, 0x59, 0x7d, 0xb0,
	0xba, 0x92, 0x3f, 0x58, 0x95, 0x5b, 0x1f, 0xd2, 0x59, 0xdd, 0x47, 0x35, 0x6a, 0xb2, 0xb4, 0x49,
	0x62, 0xfe, 0x5f, 0x69, 0x93, 0x96, 0x97, 0x50, 0x64, 0xc2, 0x43, 0x60, 0x29, 0xaf, 0xc4, 0x15,
	0x4b, 0x23, 0xad, 0xf3, 0x26, 0x5c, 0xb1, 0xf5, 0xea, 0x19, 0x16, 0xf7, 0x9b, 0x20, 0xf1, 0x09,
	0xe0, 0xfb, 0x46, 0x2c, 0x7f, 0x79, 0x2b, 0xc7, 0xb9, 0xbc, 0xf8, 0xd3, 0x4d, 0xb5, 0xf3, 0x32,
	0xcf, 0xdf, 0x73, 0xf3, 0x15, 0x00, 0x00, 0xff, 0xff, 0x69, 0xf8, 0x1b, 0x75, 0xbd, 0x01, 0x00,
	0x00,
}