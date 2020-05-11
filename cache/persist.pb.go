// Code generated by protoc-gen-go.
// source: persist.proto
// DO NOT EDIT!

/*
Package cache is a generated protocol buffer package.

It is generated from these files:
	persist.proto

It has these top-level messages:
	HitCount
	CacheState
*/
package cache

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type HitCount struct {
	Entry string `protobuf:"bytes,1,opt,name=entry" json:"entry,omitempty"`
	Count int64  `protobuf:"varint,2,opt,name=count" json:"count,omitempty"`
}

func (m *HitCount) Reset()                    { *m = HitCount{} }
func (m *HitCount) String() string            { return proto.CompactTextString(m) }
func (*HitCount) ProtoMessage()               {}
func (*HitCount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *HitCount) GetEntry() string {
	if m != nil {
		return m.Entry
	}
	return ""
}

func (m *HitCount) GetCount() int64 {
	if m != nil {
		return m.Count
	}
	return 0
}

type CacheState struct {
	Counts []*HitCount `protobuf:"bytes,1,rep,name=counts" json:"counts,omitempty"`
	Lru    []string    `protobuf:"bytes,2,rep,name=lru" json:"lru,omitempty"`
}

func (m *CacheState) Reset()                    { *m = CacheState{} }
func (m *CacheState) String() string            { return proto.CompactTextString(m) }
func (*CacheState) ProtoMessage()               {}
func (*CacheState) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *CacheState) GetCounts() []*HitCount {
	if m != nil {
		return m.Counts
	}
	return nil
}

func (m *CacheState) GetLru() []string {
	if m != nil {
		return m.Lru
	}
	return nil
}

func init() {
	proto.RegisterType((*HitCount)(nil), "cache.HitCount")
	proto.RegisterType((*CacheState)(nil), "cache.CacheState")
}

func init() { proto.RegisterFile("persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 152 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xe2, 0xe2, 0x2d, 0x48, 0x2d, 0x2a,
	0xce, 0x2c, 0x2e, 0xd1, 0x2b, 0x28, 0xca, 0x2f, 0xc9, 0x17, 0x62, 0x4d, 0x4e, 0x4c, 0xce, 0x48,
	0x55, 0x32, 0xe3, 0xe2, 0xf0, 0xc8, 0x2c, 0x71, 0xce, 0x2f, 0xcd, 0x2b, 0x11, 0x12, 0xe1, 0x62,
	0x4d, 0xcd, 0x2b, 0x29, 0xaa, 0x94, 0x60, 0x54, 0x60, 0xd4, 0xe0, 0x0c, 0x82, 0x70, 0x40, 0xa2,
	0xc9, 0x20, 0x69, 0x09, 0x26, 0x05, 0x46, 0x0d, 0xe6, 0x20, 0x08, 0x47, 0xc9, 0x9d, 0x8b, 0xcb,
	0x19, 0x64, 0x40, 0x70, 0x49, 0x62, 0x49, 0xaa, 0x90, 0x3a, 0x17, 0x1b, 0x58, 0xb8, 0x58, 0x82,
	0x51, 0x81, 0x59, 0x83, 0xdb, 0x88, 0x5f, 0x0f, 0x6c, 0xba, 0x1e, 0xcc, 0xe8, 0x20, 0xa8, 0xb4,
	0x90, 0x00, 0x17, 0x73, 0x4e, 0x51, 0xa9, 0x04, 0x93, 0x02, 0xb3, 0x06, 0x67, 0x10, 0x88, 0xe9,
	0xc4, 0x1e, 0x05, 0x71, 0x49, 0x12, 0x1b, 0xd8, 0x5d, 0xc6, 0x80, 0x00, 0x00, 0x00, 0xff, 0xff,
	0x17, 0x5a, 0xf4, 0xe6, 0xa8, 0x00, 0x00, 0x00,
}
