// Code generated by protoc-gen-go. DO NOT EDIT.
// source: base/v1/base.proto

package base // import "github.com/didiyun/didiyun-go-sdk/base/v1"

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import _ "google.golang.org/genproto/googleapis/api/annotations"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

// 公共的错误信息返回
type Error struct {
	Errno                int32    `protobuf:"varint,1,opt,name=errno,proto3" json:"errno,omitempty"`
	Errmsg               string   `protobuf:"bytes,2,opt,name=errmsg,proto3" json:"errmsg,omitempty"`
	RequestId            string   `protobuf:"bytes,3,opt,name=requestId,proto3" json:"requestId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Error) Reset()         { *m = Error{} }
func (m *Error) String() string { return proto.CompactTextString(m) }
func (*Error) ProtoMessage()    {}
func (*Error) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_146c5b6eaf525713, []int{0}
}
func (m *Error) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Error.Unmarshal(m, b)
}
func (m *Error) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Error.Marshal(b, m, deterministic)
}
func (dst *Error) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Error.Merge(dst, src)
}
func (m *Error) XXX_Size() int {
	return xxx_messageInfo_Error.Size(m)
}
func (m *Error) XXX_DiscardUnknown() {
	xxx_messageInfo_Error.DiscardUnknown(m)
}

var xxx_messageInfo_Error proto.InternalMessageInfo

func (m *Error) GetErrno() int32 {
	if m != nil {
		return m.Errno
	}
	return 0
}

func (m *Error) GetErrmsg() string {
	if m != nil {
		return m.Errmsg
	}
	return ""
}

func (m *Error) GetRequestId() string {
	if m != nil {
		return m.RequestId
	}
	return ""
}

// 公共的请求头
type Header struct {
	RegionId             string   `protobuf:"bytes,1,opt,name=regionId,proto3" json:"regionId,omitempty"`
	ZoneId               string   `protobuf:"bytes,2,opt,name=zoneId,proto3" json:"zoneId,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Header) Reset()         { *m = Header{} }
func (m *Header) String() string { return proto.CompactTextString(m) }
func (*Header) ProtoMessage()    {}
func (*Header) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_146c5b6eaf525713, []int{1}
}
func (m *Header) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Header.Unmarshal(m, b)
}
func (m *Header) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Header.Marshal(b, m, deterministic)
}
func (dst *Header) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Header.Merge(dst, src)
}
func (m *Header) XXX_Size() int {
	return xxx_messageInfo_Header.Size(m)
}
func (m *Header) XXX_DiscardUnknown() {
	xxx_messageInfo_Header.DiscardUnknown(m)
}

var xxx_messageInfo_Header proto.InternalMessageInfo

func (m *Header) GetRegionId() string {
	if m != nil {
		return m.RegionId
	}
	return ""
}

func (m *Header) GetZoneId() string {
	if m != nil {
		return m.ZoneId
	}
	return ""
}

// 任务返回定义
type JobInfo struct {
	JobUuid              string   `protobuf:"bytes,1,opt,name=jobUuid,proto3" json:"jobUuid,omitempty"`
	ResourceUuid         string   `protobuf:"bytes,2,opt,name=resourceUuid,proto3" json:"resourceUuid,omitempty"`
	Progress             float64  `protobuf:"fixed64,3,opt,name=progress,proto3" json:"progress,omitempty"`
	Type                 string   `protobuf:"bytes,4,opt,name=type,proto3" json:"type,omitempty"`
	Done                 bool     `protobuf:"varint,5,opt,name=done,proto3" json:"done,omitempty"`
	Success              bool     `protobuf:"varint,6,opt,name=success,proto3" json:"success,omitempty"`
	Result               string   `protobuf:"bytes,7,opt,name=result,proto3" json:"result,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JobInfo) Reset()         { *m = JobInfo{} }
func (m *JobInfo) String() string { return proto.CompactTextString(m) }
func (*JobInfo) ProtoMessage()    {}
func (*JobInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_146c5b6eaf525713, []int{2}
}
func (m *JobInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JobInfo.Unmarshal(m, b)
}
func (m *JobInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JobInfo.Marshal(b, m, deterministic)
}
func (dst *JobInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JobInfo.Merge(dst, src)
}
func (m *JobInfo) XXX_Size() int {
	return xxx_messageInfo_JobInfo.Size(m)
}
func (m *JobInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JobInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JobInfo proto.InternalMessageInfo

func (m *JobInfo) GetJobUuid() string {
	if m != nil {
		return m.JobUuid
	}
	return ""
}

func (m *JobInfo) GetResourceUuid() string {
	if m != nil {
		return m.ResourceUuid
	}
	return ""
}

func (m *JobInfo) GetProgress() float64 {
	if m != nil {
		return m.Progress
	}
	return 0
}

func (m *JobInfo) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *JobInfo) GetDone() bool {
	if m != nil {
		return m.Done
	}
	return false
}

func (m *JobInfo) GetSuccess() bool {
	if m != nil {
		return m.Success
	}
	return false
}

func (m *JobInfo) GetResult() string {
	if m != nil {
		return m.Result
	}
	return ""
}

type RegionInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AreaName             string   `protobuf:"bytes,3,opt,name=areaName,proto3" json:"areaName,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RegionInfo) Reset()         { *m = RegionInfo{} }
func (m *RegionInfo) String() string { return proto.CompactTextString(m) }
func (*RegionInfo) ProtoMessage()    {}
func (*RegionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_146c5b6eaf525713, []int{3}
}
func (m *RegionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionInfo.Unmarshal(m, b)
}
func (m *RegionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionInfo.Marshal(b, m, deterministic)
}
func (dst *RegionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionInfo.Merge(dst, src)
}
func (m *RegionInfo) XXX_Size() int {
	return xxx_messageInfo_RegionInfo.Size(m)
}
func (m *RegionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RegionInfo proto.InternalMessageInfo

func (m *RegionInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegionInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegionInfo) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

type ZoneInfo struct {
	Id                   string   `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string   `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ZoneInfo) Reset()         { *m = ZoneInfo{} }
func (m *ZoneInfo) String() string { return proto.CompactTextString(m) }
func (*ZoneInfo) ProtoMessage()    {}
func (*ZoneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_146c5b6eaf525713, []int{4}
}
func (m *ZoneInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ZoneInfo.Unmarshal(m, b)
}
func (m *ZoneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ZoneInfo.Marshal(b, m, deterministic)
}
func (dst *ZoneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ZoneInfo.Merge(dst, src)
}
func (m *ZoneInfo) XXX_Size() int {
	return xxx_messageInfo_ZoneInfo.Size(m)
}
func (m *ZoneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ZoneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ZoneInfo proto.InternalMessageInfo

func (m *ZoneInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *ZoneInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

type RegionAndZoneInfo struct {
	Id                   string    `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name                 string    `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	AreaName             string    `protobuf:"bytes,3,opt,name=areaName,proto3" json:"areaName,omitempty"`
	Zone                 *ZoneInfo `protobuf:"bytes,5,opt,name=zone,proto3" json:"zone,omitempty"`
	XXX_NoUnkeyedLiteral struct{}  `json:"-"`
	XXX_unrecognized     []byte    `json:"-"`
	XXX_sizecache        int32     `json:"-"`
}

func (m *RegionAndZoneInfo) Reset()         { *m = RegionAndZoneInfo{} }
func (m *RegionAndZoneInfo) String() string { return proto.CompactTextString(m) }
func (*RegionAndZoneInfo) ProtoMessage()    {}
func (*RegionAndZoneInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_base_146c5b6eaf525713, []int{5}
}
func (m *RegionAndZoneInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RegionAndZoneInfo.Unmarshal(m, b)
}
func (m *RegionAndZoneInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RegionAndZoneInfo.Marshal(b, m, deterministic)
}
func (dst *RegionAndZoneInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RegionAndZoneInfo.Merge(dst, src)
}
func (m *RegionAndZoneInfo) XXX_Size() int {
	return xxx_messageInfo_RegionAndZoneInfo.Size(m)
}
func (m *RegionAndZoneInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_RegionAndZoneInfo.DiscardUnknown(m)
}

var xxx_messageInfo_RegionAndZoneInfo proto.InternalMessageInfo

func (m *RegionAndZoneInfo) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *RegionAndZoneInfo) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *RegionAndZoneInfo) GetAreaName() string {
	if m != nil {
		return m.AreaName
	}
	return ""
}

func (m *RegionAndZoneInfo) GetZone() *ZoneInfo {
	if m != nil {
		return m.Zone
	}
	return nil
}

func init() {
	proto.RegisterType((*Error)(nil), "didi.cloud.base.v1.Error")
	proto.RegisterType((*Header)(nil), "didi.cloud.base.v1.Header")
	proto.RegisterType((*JobInfo)(nil), "didi.cloud.base.v1.JobInfo")
	proto.RegisterType((*RegionInfo)(nil), "didi.cloud.base.v1.RegionInfo")
	proto.RegisterType((*ZoneInfo)(nil), "didi.cloud.base.v1.ZoneInfo")
	proto.RegisterType((*RegionAndZoneInfo)(nil), "didi.cloud.base.v1.RegionAndZoneInfo")
}

func init() { proto.RegisterFile("base/v1/base.proto", fileDescriptor_base_146c5b6eaf525713) }

var fileDescriptor_base_146c5b6eaf525713 = []byte{
	// 438 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x94, 0x52, 0x4d, 0xab, 0x13, 0x31,
	0x14, 0x25, 0xb5, 0x9f, 0x57, 0x11, 0x8c, 0x22, 0xc3, 0xa3, 0x8b, 0x32, 0xab, 0x6e, 0x5e, 0xc6,
	0xea, 0x52, 0x37, 0x56, 0x05, 0x2b, 0x22, 0x32, 0xe2, 0x5b, 0x3c, 0xba, 0xc9, 0x4c, 0xae, 0xe3,
	0x68, 0x9b, 0x5b, 0x6f, 0x66, 0x1e, 0xb4, 0x7b, 0xff, 0x8c, 0xff, 0xc1, 0x3f, 0xe0, 0x2f, 0x72,
	0x29, 0x49, 0x3a, 0x15, 0x11, 0xc4, 0xb7, 0x4a, 0xce, 0xc9, 0xbd, 0xe7, 0x9c, 0x5c, 0x2e, 0xc8,
	0x42, 0x3b, 0xcc, 0xae, 0x16, 0x99, 0x3f, 0xd5, 0x8e, 0xa9, 0x21, 0x29, 0x4d, 0x6d, 0x6a, 0x55,
	0x6e, 0xa8, 0x35, 0x2a, 0xd0, 0x57, 0x8b, 0xb3, 0x69, 0x45, 0x54, 0x6d, 0x30, 0xd3, 0xbb, 0x3a,
	0xd3, 0xd6, 0x52, 0xa3, 0x9b, 0x9a, 0xac, 0x8b, 0x1d, 0xe9, 0x3b, 0x18, 0xbc, 0x60, 0x26, 0x96,
	0xf7, 0x60, 0x80, 0xcc, 0x96, 0x12, 0x31, 0x13, 0xf3, 0x41, 0x1e, 0x81, 0xbc, 0x0f, 0x43, 0x64,
	0xde, 0xba, 0x2a, 0xe9, 0xcd, 0xc4, 0x7c, 0x92, 0x1f, 0x91, 0x9c, 0xc2, 0x84, 0xf1, 0x4b, 0x8b,
	0xae, 0x59, 0x99, 0xe4, 0x46, 0x78, 0xfa, 0x4d, 0xa4, 0x4f, 0x60, 0xf8, 0x12, 0xb5, 0x41, 0x96,
	0x67, 0x30, 0x66, 0xac, 0x6a, 0xb2, 0x2b, 0x13, 0x84, 0x27, 0xf9, 0x09, 0x7b, 0xed, 0x03, 0x59,
	0x5c, 0x99, 0x4e, 0x3b, 0xa2, 0xf4, 0xbb, 0x80, 0xd1, 0x2b, 0x2a, 0x56, 0xf6, 0x03, 0xc9, 0x04,
	0x46, 0x9f, 0xa8, 0x78, 0xdf, 0xd6, 0x5d, 0x7b, 0x07, 0x65, 0x0a, 0xb7, 0x18, 0x1d, 0xb5, 0x5c,
	0x62, 0x78, 0x8e, 0x1a, 0x7f, 0x70, 0xde, 0x7d, 0xc7, 0x54, 0x31, 0x3a, 0x17, 0x42, 0x8a, 0xfc,
	0x84, 0xa5, 0x84, 0x7e, 0xb3, 0xdf, 0x61, 0xd2, 0x0f, 0x7d, 0xe1, 0xee, 0x39, 0x43, 0x16, 0x93,
	0xc1, 0x4c, 0xcc, 0xc7, 0x79, 0xb8, 0xfb, 0x04, 0xae, 0x2d, 0x4b, 0x2f, 0x31, 0x0c, 0x74, 0x07,
	0x7d, 0x7e, 0x46, 0xd7, 0x6e, 0x9a, 0x64, 0x14, 0xf3, 0x47, 0x94, 0xbe, 0x06, 0xc8, 0xe3, 0x1f,
	0xfd, 0x0f, 0x6e, 0x43, 0xef, 0x14, 0xbe, 0x57, 0x1b, 0xef, 0x61, 0xf5, 0x16, 0x8f, 0x79, 0xc3,
	0xdd, 0xe7, 0xd4, 0x8c, 0xfa, 0x8d, 0xe7, 0xe3, 0x30, 0x4f, 0x38, 0x55, 0x30, 0xbe, 0xf4, 0x73,
	0xf9, 0x4f, 0xad, 0xf4, 0xab, 0x80, 0x3b, 0xd1, 0xfe, 0xa9, 0x35, 0xd7, 0xe9, 0xfc, 0x57, 0x0a,
	0xf9, 0x00, 0xfa, 0x87, 0x6e, 0x32, 0x37, 0x1f, 0x4e, 0xd5, 0xdf, 0x7b, 0xa6, 0x3a, 0xaf, 0x3c,
	0x54, 0x2e, 0x0f, 0x70, 0xb7, 0xa4, 0x6d, 0x28, 0xdc, 0xb7, 0xb6, 0xab, 0x5a, 0xf6, 0x97, 0xda,
	0xe1, 0x5b, 0x71, 0xa9, 0xaa, 0xba, 0xf9, 0xd8, 0x16, 0xaa, 0xa4, 0x6d, 0x76, 0xac, 0xe9, 0xce,
	0xf3, 0x8a, 0xce, 0x9d, 0xf9, 0x9c, 0x1d, 0xf7, 0xfb, 0xb1, 0x3f, 0x7f, 0x0a, 0xf1, 0xad, 0x27,
	0x9f, 0x7b, 0xd7, 0x67, 0xc1, 0xd5, 0xcb, 0xa8, 0x8b, 0xc5, 0x8f, 0x48, 0xae, 0x03, 0xb9, 0xf6,
	0xe4, 0xfa, 0x62, 0x51, 0x0c, 0xc3, 0x6e, 0x3f, 0xfa, 0x15, 0x00, 0x00, 0xff, 0xff, 0x06, 0x6e,
	0xfc, 0x18, 0x23, 0x03, 0x00, 0x00,
}