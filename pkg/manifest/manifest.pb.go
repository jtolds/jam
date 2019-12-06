// Code generated by protoc-gen-go. DO NOT EDIT.
// source: manifest.proto

package manifest

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	timestamp "github.com/golang/protobuf/ptypes/timestamp"
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

type Content_Type int32

const (
	Content_UNKNOWN Content_Type = 0
	Content_FILE    Content_Type = 1
	Content_SYMLINK Content_Type = 2
	Content_DIR     Content_Type = 3
)

var Content_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "FILE",
	2: "SYMLINK",
	3: "DIR",
}

var Content_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"FILE":    1,
	"SYMLINK": 2,
	"DIR":     3,
}

func (x Content_Type) String() string {
	return proto.EnumName(Content_Type_name, int32(x))
}

func (Content_Type) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{2, 0}
}

type Range struct {
	Blob                 string   `protobuf:"bytes,1,opt,name=blob,proto3" json:"blob,omitempty"`
	Offset               int64    `protobuf:"varint,2,opt,name=offset,proto3" json:"offset,omitempty"`
	Length               int64    `protobuf:"varint,3,opt,name=length,proto3" json:"length,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Range) Reset()         { *m = Range{} }
func (m *Range) String() string { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()    {}
func (*Range) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{0}
}

func (m *Range) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Range.Unmarshal(m, b)
}
func (m *Range) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Range.Marshal(b, m, deterministic)
}
func (m *Range) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Range.Merge(m, src)
}
func (m *Range) XXX_Size() int {
	return xxx_messageInfo_Range.Size(m)
}
func (m *Range) XXX_DiscardUnknown() {
	xxx_messageInfo_Range.DiscardUnknown(m)
}

var xxx_messageInfo_Range proto.InternalMessageInfo

func (m *Range) GetBlob() string {
	if m != nil {
		return m.Blob
	}
	return ""
}

func (m *Range) GetOffset() int64 {
	if m != nil {
		return m.Offset
	}
	return 0
}

func (m *Range) GetLength() int64 {
	if m != nil {
		return m.Length
	}
	return 0
}

type Stream struct {
	Ranges               []*Range `protobuf:"bytes,1,rep,name=ranges,proto3" json:"ranges,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Stream) Reset()         { *m = Stream{} }
func (m *Stream) String() string { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()    {}
func (*Stream) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{1}
}

func (m *Stream) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Stream.Unmarshal(m, b)
}
func (m *Stream) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Stream.Marshal(b, m, deterministic)
}
func (m *Stream) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Stream.Merge(m, src)
}
func (m *Stream) XXX_Size() int {
	return xxx_messageInfo_Stream.Size(m)
}
func (m *Stream) XXX_DiscardUnknown() {
	xxx_messageInfo_Stream.DiscardUnknown(m)
}

var xxx_messageInfo_Stream proto.InternalMessageInfo

func (m *Stream) GetRanges() []*Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type Content struct {
	Type                 Content_Type         `protobuf:"varint,1,opt,name=type,proto3,enum=manifest.Content_Type" json:"type,omitempty"`
	Creation             *timestamp.Timestamp `protobuf:"bytes,2,opt,name=creation,proto3" json:"creation,omitempty"`
	Modified             *timestamp.Timestamp `protobuf:"bytes,3,opt,name=modified,proto3" json:"modified,omitempty"`
	Mode                 uint32               `protobuf:"varint,4,opt,name=mode,proto3" json:"mode,omitempty"`
	LinkTarget           string               `protobuf:"bytes,5,opt,name=link_target,json=linkTarget,proto3" json:"link_target,omitempty"`
	Data                 *Stream              `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Content) Reset()         { *m = Content{} }
func (m *Content) String() string { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()    {}
func (*Content) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{2}
}

func (m *Content) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Content.Unmarshal(m, b)
}
func (m *Content) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Content.Marshal(b, m, deterministic)
}
func (m *Content) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Content.Merge(m, src)
}
func (m *Content) XXX_Size() int {
	return xxx_messageInfo_Content.Size(m)
}
func (m *Content) XXX_DiscardUnknown() {
	xxx_messageInfo_Content.DiscardUnknown(m)
}

var xxx_messageInfo_Content proto.InternalMessageInfo

func (m *Content) GetType() Content_Type {
	if m != nil {
		return m.Type
	}
	return Content_UNKNOWN
}

func (m *Content) GetCreation() *timestamp.Timestamp {
	if m != nil {
		return m.Creation
	}
	return nil
}

func (m *Content) GetModified() *timestamp.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

func (m *Content) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Content) GetLinkTarget() string {
	if m != nil {
		return m.LinkTarget
	}
	return ""
}

func (m *Content) GetData() *Stream {
	if m != nil {
		return m.Data
	}
	return nil
}

type Entry struct {
	Path                 string   `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	Content              *Content `protobuf:"bytes,2,opt,name=content,proto3" json:"content,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Entry) Reset()         { *m = Entry{} }
func (m *Entry) String() string { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()    {}
func (*Entry) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{3}
}

func (m *Entry) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Entry.Unmarshal(m, b)
}
func (m *Entry) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Entry.Marshal(b, m, deterministic)
}
func (m *Entry) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Entry.Merge(m, src)
}
func (m *Entry) XXX_Size() int {
	return xxx_messageInfo_Entry.Size(m)
}
func (m *Entry) XXX_DiscardUnknown() {
	xxx_messageInfo_Entry.DiscardUnknown(m)
}

var xxx_messageInfo_Entry proto.InternalMessageInfo

func (m *Entry) GetPath() string {
	if m != nil {
		return m.Path
	}
	return ""
}

func (m *Entry) GetContent() *Content {
	if m != nil {
		return m.Content
	}
	return nil
}

type EntrySet struct {
	Entries              []*Entry `protobuf:"bytes,1,rep,name=entries,proto3" json:"entries,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EntrySet) Reset()         { *m = EntrySet{} }
func (m *EntrySet) String() string { return proto.CompactTextString(m) }
func (*EntrySet) ProtoMessage()    {}
func (*EntrySet) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{4}
}

func (m *EntrySet) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EntrySet.Unmarshal(m, b)
}
func (m *EntrySet) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EntrySet.Marshal(b, m, deterministic)
}
func (m *EntrySet) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EntrySet.Merge(m, src)
}
func (m *EntrySet) XXX_Size() int {
	return xxx_messageInfo_EntrySet.Size(m)
}
func (m *EntrySet) XXX_DiscardUnknown() {
	xxx_messageInfo_EntrySet.DiscardUnknown(m)
}

var xxx_messageInfo_EntrySet proto.InternalMessageInfo

func (m *EntrySet) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Page struct {
	Prefix string `protobuf:"bytes,1,opt,name=prefix,proto3" json:"prefix,omitempty"`
	// Types that are valid to be assigned to Descendents:
	//	*Page_Branch
	//	*Page_Entries
	Descendents          isPage_Descendents `protobuf_oneof:"descendents"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *Page) Reset()         { *m = Page{} }
func (m *Page) String() string { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()    {}
func (*Page) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{5}
}

func (m *Page) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Page.Unmarshal(m, b)
}
func (m *Page) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Page.Marshal(b, m, deterministic)
}
func (m *Page) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Page.Merge(m, src)
}
func (m *Page) XXX_Size() int {
	return xxx_messageInfo_Page.Size(m)
}
func (m *Page) XXX_DiscardUnknown() {
	xxx_messageInfo_Page.DiscardUnknown(m)
}

var xxx_messageInfo_Page proto.InternalMessageInfo

func (m *Page) GetPrefix() string {
	if m != nil {
		return m.Prefix
	}
	return ""
}

type isPage_Descendents interface {
	isPage_Descendents()
}

type Page_Branch struct {
	Branch *Stream `protobuf:"bytes,2,opt,name=branch,proto3,oneof"`
}

type Page_Entries struct {
	Entries *EntrySet `protobuf:"bytes,3,opt,name=entries,proto3,oneof"`
}

func (*Page_Branch) isPage_Descendents() {}

func (*Page_Entries) isPage_Descendents() {}

func (m *Page) GetDescendents() isPage_Descendents {
	if m != nil {
		return m.Descendents
	}
	return nil
}

func (m *Page) GetBranch() *Stream {
	if x, ok := m.GetDescendents().(*Page_Branch); ok {
		return x.Branch
	}
	return nil
}

func (m *Page) GetEntries() *EntrySet {
	if x, ok := m.GetDescendents().(*Page_Entries); ok {
		return x.Entries
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Page) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Page_Branch)(nil),
		(*Page_Entries)(nil),
	}
}

type Manifest struct {
	Pages                []*Page  `protobuf:"bytes,1,rep,name=pages,proto3" json:"pages,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Manifest) Reset()         { *m = Manifest{} }
func (m *Manifest) String() string { return proto.CompactTextString(m) }
func (*Manifest) ProtoMessage()    {}
func (*Manifest) Descriptor() ([]byte, []int) {
	return fileDescriptor_0bb23f43f7afb4c1, []int{6}
}

func (m *Manifest) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Manifest.Unmarshal(m, b)
}
func (m *Manifest) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Manifest.Marshal(b, m, deterministic)
}
func (m *Manifest) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Manifest.Merge(m, src)
}
func (m *Manifest) XXX_Size() int {
	return xxx_messageInfo_Manifest.Size(m)
}
func (m *Manifest) XXX_DiscardUnknown() {
	xxx_messageInfo_Manifest.DiscardUnknown(m)
}

var xxx_messageInfo_Manifest proto.InternalMessageInfo

func (m *Manifest) GetPages() []*Page {
	if m != nil {
		return m.Pages
	}
	return nil
}

func init() {
	proto.RegisterEnum("manifest.Content_Type", Content_Type_name, Content_Type_value)
	proto.RegisterType((*Range)(nil), "manifest.Range")
	proto.RegisterType((*Stream)(nil), "manifest.Stream")
	proto.RegisterType((*Content)(nil), "manifest.Content")
	proto.RegisterType((*Entry)(nil), "manifest.Entry")
	proto.RegisterType((*EntrySet)(nil), "manifest.EntrySet")
	proto.RegisterType((*Page)(nil), "manifest.Page")
	proto.RegisterType((*Manifest)(nil), "manifest.Manifest")
}

func init() { proto.RegisterFile("manifest.proto", fileDescriptor_0bb23f43f7afb4c1) }

var fileDescriptor_0bb23f43f7afb4c1 = []byte{
	// 481 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x84, 0x53, 0xd1, 0x6e, 0xd3, 0x40,
	0x10, 0x8c, 0x13, 0xc7, 0x09, 0x1b, 0x35, 0x98, 0x7b, 0xa8, 0xac, 0xbe, 0x34, 0xb2, 0x2a, 0x11,
	0x8a, 0xe4, 0x42, 0x2a, 0xf8, 0x80, 0x42, 0x51, 0xa2, 0xb4, 0x01, 0x5d, 0x82, 0x10, 0xbc, 0xa0,
	0x73, 0xbc, 0x76, 0x2c, 0xe2, 0x3b, 0xcb, 0x5e, 0x24, 0xf2, 0x09, 0x7c, 0x10, 0xff, 0x87, 0xee,
	0x6c, 0xc7, 0x12, 0x45, 0xea, 0xdb, 0xed, 0xec, 0xdc, 0xed, 0xcc, 0xac, 0x0d, 0xe3, 0x4c, 0xc8,
	0x34, 0xc6, 0x92, 0x82, 0xbc, 0x50, 0xa4, 0xd8, 0xb0, 0xa9, 0xcf, 0xce, 0x13, 0xa5, 0x92, 0x3d,
	0x5e, 0x19, 0x3c, 0xfc, 0x19, 0x5f, 0x51, 0x9a, 0x61, 0x49, 0x22, 0xcb, 0x2b, 0xaa, 0xbf, 0x84,
	0x3e, 0x17, 0x32, 0x41, 0xc6, 0xc0, 0x0e, 0xf7, 0x2a, 0xf4, 0xac, 0x89, 0x35, 0x7d, 0xc2, 0xcd,
	0x99, 0x9d, 0x82, 0xa3, 0xe2, 0xb8, 0x44, 0xf2, 0xba, 0x13, 0x6b, 0xda, 0xe3, 0x75, 0xa5, 0xf1,
	0x3d, 0xca, 0x84, 0x76, 0x5e, 0xaf, 0xc2, 0xab, 0xca, 0x7f, 0x0d, 0xce, 0x9a, 0x0a, 0x14, 0x19,
	0x7b, 0x0e, 0x4e, 0xa1, 0x9f, 0x2d, 0x3d, 0x6b, 0xd2, 0x9b, 0x8e, 0x66, 0x4f, 0x83, 0xa3, 0x44,
	0x33, 0x8e, 0xd7, 0x6d, 0xff, 0x4f, 0x17, 0x06, 0xef, 0x94, 0x24, 0x94, 0xc4, 0x2e, 0xc1, 0xa6,
	0x43, 0x8e, 0x46, 0xc2, 0x78, 0x76, 0xda, 0x5e, 0xa9, 0x09, 0xc1, 0xe6, 0x90, 0x23, 0x37, 0x1c,
	0xf6, 0x16, 0x86, 0xdb, 0x02, 0x05, 0xa5, 0x4a, 0x1a, 0x71, 0xa3, 0xd9, 0x59, 0x50, 0x79, 0x0d,
	0x1a, 0xaf, 0xc1, 0xa6, 0xf1, 0xca, 0x8f, 0x5c, 0x7d, 0x2f, 0x53, 0x51, 0x1a, 0xa7, 0x18, 0x19,
	0xf1, 0x8f, 0xdc, 0x6b, 0xb8, 0x3a, 0x9e, 0x4c, 0x45, 0xe8, 0xd9, 0x13, 0x6b, 0x7a, 0xc2, 0xcd,
	0x99, 0x9d, 0xc3, 0x68, 0x9f, 0xca, 0x1f, 0xdf, 0x49, 0x14, 0x09, 0x92, 0xd7, 0x37, 0xc9, 0x81,
	0x86, 0x36, 0x06, 0x61, 0x17, 0x60, 0x47, 0x82, 0x84, 0xe7, 0x98, 0x41, 0x6e, 0x6b, 0xa8, 0x4a,
	0x89, 0x9b, 0xae, 0x7f, 0x0d, 0xb6, 0x36, 0xc6, 0x46, 0x30, 0xf8, 0xbc, 0x5a, 0xae, 0x3e, 0x7e,
	0x59, 0xb9, 0x1d, 0x36, 0x04, 0xfb, 0xc3, 0xe2, 0xee, 0xd6, 0xb5, 0x34, 0xbc, 0xfe, 0x7a, 0x7f,
	0xb7, 0x58, 0x2d, 0xdd, 0x2e, 0x1b, 0x40, 0xef, 0xfd, 0x82, 0xbb, 0x3d, 0x7f, 0x0e, 0xfd, 0x5b,
	0x49, 0xc5, 0x41, 0x0b, 0xcb, 0x05, 0xed, 0x9a, 0xbd, 0xe9, 0x33, 0x7b, 0x09, 0x83, 0x6d, 0x15,
	0x59, 0x9d, 0xcd, 0xb3, 0x07, 0x59, 0xf2, 0x86, 0xe1, 0xbf, 0x81, 0xa1, 0x79, 0x69, 0x8d, 0xc4,
	0x5e, 0xc0, 0x00, 0x25, 0x15, 0xe9, 0xff, 0xf6, 0x66, 0x48, 0xbc, 0xe9, 0xfb, 0xbf, 0x2d, 0xb0,
	0x3f, 0x89, 0x04, 0xf5, 0xc7, 0x90, 0x17, 0x18, 0xa7, 0xbf, 0x6a, 0x09, 0x75, 0xc5, 0x2e, 0xc1,
	0x09, 0x0b, 0x21, 0xb7, 0xbb, 0x5a, 0xc3, 0x03, 0xfb, 0xf3, 0x0e, 0xaf, 0x19, 0x2c, 0x68, 0xe7,
	0x56, 0x4b, 0x61, 0xff, 0xcc, 0x5d, 0x23, 0xcd, 0x3b, 0xc7, 0xe1, 0x37, 0x27, 0x30, 0x8a, 0xb0,
	0xdc, 0xa2, 0x8c, 0x50, 0x52, 0xe9, 0xbf, 0x82, 0xe1, 0x7d, 0x4d, 0x67, 0x17, 0xd0, 0xcf, 0x45,
	0xfb, 0xe1, 0x8d, 0xdb, 0x87, 0xb4, 0x5a, 0x5e, 0x35, 0x6f, 0xe0, 0xdb, 0xf1, 0x1f, 0x09, 0x1d,
	0xb3, 0xf8, 0xeb, 0xbf, 0x01, 0x00, 0x00, 0xff, 0xff, 0xd4, 0x5b, 0xee, 0xbd, 0x46, 0x03, 0x00,
	0x00,
}
