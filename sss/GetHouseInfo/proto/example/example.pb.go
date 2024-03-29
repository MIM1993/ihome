// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/example/example.proto

package go_micro_srv_GetHouseInfo

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
	//房屋id
	HouseId              string   `protobuf:"bytes,2,opt,name=house_id,json=houseId,proto3" json:"house_id,omitempty"`
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

func (m *Request) GetHouseId() string {
	if m != nil {
		return m.HouseId
	}
	return ""
}

//srv-->web
type Response struct {
	//错误吗
	Errno string `protobuf:"bytes,1,opt,name=errno,proto3" json:"errno,omitempty"`
	//错误信息
	Errmsg string `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	//房屋信息
	Housedata []byte `protobuf:"bytes,3,opt,name=housedata,proto3" json:"housedata,omitempty"`
	//用户id
	UserId               int64    `protobuf:"varint,4,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
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

func (m *Response) GetHousedata() []byte {
	if m != nil {
		return m.Housedata
	}
	return nil
}

func (m *Response) GetUserId() int64 {
	if m != nil {
		return m.UserId
	}
	return 0
}

func init() {
	proto.RegisterType((*Request)(nil), "go.micro.srv.GetHouseInfo.Request")
	proto.RegisterType((*Response)(nil), "go.micro.srv.GetHouseInfo.Response")
}

func init() { proto.RegisterFile("proto/example/example.proto", fileDescriptor_097b3f5db5cf5789) }

var fileDescriptor_097b3f5db5cf5789 = []byte{
	// 229 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x90, 0xc1, 0x4a, 0xc4, 0x30,
	0x10, 0x86, 0xad, 0xab, 0xed, 0xee, 0xb0, 0xa7, 0x41, 0xb4, 0xab, 0x1e, 0x4a, 0xbd, 0xf4, 0x14,
	0x41, 0xdf, 0x40, 0x10, 0xed, 0x35, 0x37, 0x4f, 0x52, 0xcd, 0xb8, 0x16, 0x6c, 0xa7, 0xce, 0xb4,
	0xe2, 0xe3, 0x4b, 0xd2, 0xc8, 0x7a, 0x71, 0x4f, 0xe1, 0xfb, 0x93, 0x99, 0x9f, 0x2f, 0x70, 0x31,
	0x08, 0x8f, 0x7c, 0x4d, 0xdf, 0x4d, 0x37, 0x7c, 0xd0, 0xef, 0x69, 0x42, 0x8a, 0x9b, 0x2d, 0x9b,
	0xae, 0x7d, 0x15, 0x36, 0x2a, 0x5f, 0xe6, 0x81, 0xc6, 0x47, 0x9e, 0x94, 0xea, 0xfe, 0x8d, 0xcb,
	0x3b, 0xc8, 0x2c, 0x7d, 0x4e, 0xa4, 0x23, 0x5e, 0xc2, 0x4a, 0x49, 0xb5, 0xe5, 0xbe, 0x75, 0x79,
	0x52, 0x24, 0xd5, 0xca, 0xee, 0x02, 0xdc, 0xc0, 0xf2, 0xdd, 0x4f, 0x3d, 0xb7, 0x2e, 0x3f, 0x0c,
	0x97, 0x59, 0xe0, 0xda, 0x95, 0x0c, 0x4b, 0x4b, 0x3a, 0x70, 0xaf, 0x84, 0x27, 0x70, 0x4c, 0x22,
	0x3d, 0xc7, 0x05, 0x33, 0xe0, 0x29, 0xa4, 0x24, 0xd2, 0xe9, 0x36, 0x8e, 0x46, 0xf2, 0x95, 0x61,
	0x89, 0x6b, 0xc6, 0x26, 0x5f, 0x14, 0x49, 0xb5, 0xb6, 0xbb, 0x00, 0xcf, 0x20, 0x9b, 0x94, 0xc4,
	0x37, 0x1e, 0x15, 0x49, 0xb5, 0xb0, 0xa9, 0xc7, 0xda, 0xdd, 0x38, 0xc8, 0xee, 0x67, 0x41, 0x7c,
	0x82, 0xf5, 0x5f, 0x1f, 0x2c, 0xcd, 0xbf, 0xae, 0x26, 0x8a, 0x9e, 0x5f, 0xed, 0x7d, 0x33, 0x8b,
	0x94, 0x07, 0x2f, 0x69, 0xf8, 0xbc, 0xdb, 0x9f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x1d, 0x15, 0xac,
	0x5f, 0x5b, 0x01, 0x00, 0x00,
}
