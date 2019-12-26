// Code generated by protoc-gen-go.
// source: manifest.proto
// DO NOT EDIT!

/*
Package manifest is a generated protocol buffer package.

It is generated from these files:
	manifest.proto

It has these top-level messages:
	Range
	Stream
	Metadata
	Content
	Entry
	EntrySet
	Page
*/
package manifest

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "github.com/golang/protobuf/ptypes/timestamp"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Metadata_Type int32

const (
	Metadata_UNKNOWN Metadata_Type = 0
	Metadata_FILE    Metadata_Type = 1
	Metadata_SYMLINK Metadata_Type = 2
)

var Metadata_Type_name = map[int32]string{
	0: "UNKNOWN",
	1: "FILE",
	2: "SYMLINK",
}
var Metadata_Type_value = map[string]int32{
	"UNKNOWN": 0,
	"FILE":    1,
	"SYMLINK": 2,
}

func (x Metadata_Type) String() string {
	return proto.EnumName(Metadata_Type_name, int32(x))
}
func (Metadata_Type) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{2, 0} }

type Range struct {
	Blob   string `protobuf:"bytes,1,opt,name=blob" json:"blob,omitempty"`
	Offset int64  `protobuf:"varint,2,opt,name=offset" json:"offset,omitempty"`
	Length int64  `protobuf:"varint,3,opt,name=length" json:"length,omitempty"`
}

func (m *Range) Reset()                    { *m = Range{} }
func (m *Range) String() string            { return proto.CompactTextString(m) }
func (*Range) ProtoMessage()               {}
func (*Range) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

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
	Ranges []*Range `protobuf:"bytes,1,rep,name=ranges" json:"ranges,omitempty"`
}

func (m *Stream) Reset()                    { *m = Stream{} }
func (m *Stream) String() string            { return proto.CompactTextString(m) }
func (*Stream) ProtoMessage()               {}
func (*Stream) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Stream) GetRanges() []*Range {
	if m != nil {
		return m.Ranges
	}
	return nil
}

type Metadata struct {
	Type       Metadata_Type              `protobuf:"varint,1,opt,name=type,enum=manifest.Metadata_Type" json:"type,omitempty"`
	Creation   *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=creation" json:"creation,omitempty"`
	Modified   *google_protobuf.Timestamp `protobuf:"bytes,3,opt,name=modified" json:"modified,omitempty"`
	Mode       uint32                     `protobuf:"varint,4,opt,name=mode" json:"mode,omitempty"`
	LinkTarget string                     `protobuf:"bytes,5,opt,name=link_target,json=linkTarget" json:"link_target,omitempty"`
}

func (m *Metadata) Reset()                    { *m = Metadata{} }
func (m *Metadata) String() string            { return proto.CompactTextString(m) }
func (*Metadata) ProtoMessage()               {}
func (*Metadata) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Metadata) GetType() Metadata_Type {
	if m != nil {
		return m.Type
	}
	return Metadata_UNKNOWN
}

func (m *Metadata) GetCreation() *google_protobuf.Timestamp {
	if m != nil {
		return m.Creation
	}
	return nil
}

func (m *Metadata) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

func (m *Metadata) GetMode() uint32 {
	if m != nil {
		return m.Mode
	}
	return 0
}

func (m *Metadata) GetLinkTarget() string {
	if m != nil {
		return m.LinkTarget
	}
	return ""
}

type Content struct {
	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata" json:"metadata,omitempty"`
	Data     *Stream   `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
}

func (m *Content) Reset()                    { *m = Content{} }
func (m *Content) String() string            { return proto.CompactTextString(m) }
func (*Content) ProtoMessage()               {}
func (*Content) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Content) GetMetadata() *Metadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Content) GetData() *Stream {
	if m != nil {
		return m.Data
	}
	return nil
}

type Entry struct {
	Path    string   `protobuf:"bytes,1,opt,name=path" json:"path,omitempty"`
	Content *Content `protobuf:"bytes,2,opt,name=content" json:"content,omitempty"`
}

func (m *Entry) Reset()                    { *m = Entry{} }
func (m *Entry) String() string            { return proto.CompactTextString(m) }
func (*Entry) ProtoMessage()               {}
func (*Entry) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

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
	Entries []*Entry `protobuf:"bytes,1,rep,name=entries" json:"entries,omitempty"`
}

func (m *EntrySet) Reset()                    { *m = EntrySet{} }
func (m *EntrySet) String() string            { return proto.CompactTextString(m) }
func (*EntrySet) ProtoMessage()               {}
func (*EntrySet) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *EntrySet) GetEntries() []*Entry {
	if m != nil {
		return m.Entries
	}
	return nil
}

type Page struct {
	SortKey string `protobuf:"bytes,1,opt,name=sort_key,json=sortKey" json:"sort_key,omitempty"`
	// Types that are valid to be assigned to Descendents:
	//	*Page_Branch
	//	*Page_Entries
	Descendents isPage_Descendents `protobuf_oneof:"descendents"`
}

func (m *Page) Reset()                    { *m = Page{} }
func (m *Page) String() string            { return proto.CompactTextString(m) }
func (*Page) ProtoMessage()               {}
func (*Page) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

type isPage_Descendents interface {
	isPage_Descendents()
}

type Page_Branch struct {
	Branch *Stream `protobuf:"bytes,2,opt,name=branch,oneof"`
}
type Page_Entries struct {
	Entries *EntrySet `protobuf:"bytes,3,opt,name=entries,oneof"`
}

func (*Page_Branch) isPage_Descendents()  {}
func (*Page_Entries) isPage_Descendents() {}

func (m *Page) GetDescendents() isPage_Descendents {
	if m != nil {
		return m.Descendents
	}
	return nil
}

func (m *Page) GetSortKey() string {
	if m != nil {
		return m.SortKey
	}
	return ""
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

// XXX_OneofFuncs is for the internal use of the proto package.
func (*Page) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _Page_OneofMarshaler, _Page_OneofUnmarshaler, _Page_OneofSizer, []interface{}{
		(*Page_Branch)(nil),
		(*Page_Entries)(nil),
	}
}

func _Page_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*Page)
	// descendents
	switch x := m.Descendents.(type) {
	case *Page_Branch:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Branch); err != nil {
			return err
		}
	case *Page_Entries:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Entries); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("Page.Descendents has unexpected type %T", x)
	}
	return nil
}

func _Page_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*Page)
	switch tag {
	case 2: // descendents.branch
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(Stream)
		err := b.DecodeMessage(msg)
		m.Descendents = &Page_Branch{msg}
		return true, err
	case 3: // descendents.entries
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(EntrySet)
		err := b.DecodeMessage(msg)
		m.Descendents = &Page_Entries{msg}
		return true, err
	default:
		return false, nil
	}
}

func _Page_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*Page)
	// descendents
	switch x := m.Descendents.(type) {
	case *Page_Branch:
		s := proto.Size(x.Branch)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *Page_Entries:
		s := proto.Size(x.Entries)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

func init() {
	proto.RegisterType((*Range)(nil), "manifest.Range")
	proto.RegisterType((*Stream)(nil), "manifest.Stream")
	proto.RegisterType((*Metadata)(nil), "manifest.Metadata")
	proto.RegisterType((*Content)(nil), "manifest.Content")
	proto.RegisterType((*Entry)(nil), "manifest.Entry")
	proto.RegisterType((*EntrySet)(nil), "manifest.EntrySet")
	proto.RegisterType((*Page)(nil), "manifest.Page")
	proto.RegisterEnum("manifest.Metadata_Type", Metadata_Type_name, Metadata_Type_value)
}

func init() { proto.RegisterFile("manifest.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 476 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x84, 0x53, 0x4d, 0x6f, 0xd3, 0x40,
	0x10, 0x8d, 0x13, 0x27, 0x31, 0x63, 0xb5, 0x84, 0x3d, 0x80, 0xe9, 0xa5, 0x91, 0x85, 0x44, 0x68,
	0x25, 0x57, 0x04, 0xc1, 0x0f, 0x28, 0x2a, 0x4a, 0x95, 0x36, 0xa0, 0x4d, 0x10, 0x82, 0x4b, 0xb4,
	0x8e, 0xc7, 0x8e, 0xd5, 0x78, 0xd7, 0xb2, 0x87, 0x83, 0xff, 0x01, 0x07, 0x7e, 0x34, 0xda, 0xf5,
	0x47, 0x24, 0x3e, 0xc4, 0x6d, 0xe7, 0xed, 0xdb, 0x99, 0xf7, 0xde, 0x68, 0xe1, 0x34, 0x13, 0x32,
	0x8d, 0xb1, 0xa4, 0x20, 0x2f, 0x14, 0x29, 0xe6, 0xb4, 0xf5, 0xd9, 0x79, 0xa2, 0x54, 0x72, 0xc0,
	0x2b, 0x83, 0x87, 0xdf, 0xe3, 0x2b, 0x4a, 0x33, 0x2c, 0x49, 0x64, 0x79, 0x4d, 0xf5, 0x97, 0x30,
	0xe4, 0x42, 0x26, 0xc8, 0x18, 0xd8, 0xe1, 0x41, 0x85, 0x9e, 0x35, 0xb5, 0x66, 0x8f, 0xb8, 0x39,
	0xb3, 0xa7, 0x30, 0x52, 0x71, 0x5c, 0x22, 0x79, 0xfd, 0xa9, 0x35, 0x1b, 0xf0, 0xa6, 0xd2, 0xf8,
	0x01, 0x65, 0x42, 0x7b, 0x6f, 0x50, 0xe3, 0x75, 0xe5, 0xbf, 0x86, 0xd1, 0x9a, 0x0a, 0x14, 0x19,
	0x7b, 0x09, 0xa3, 0x42, 0xb7, 0x2d, 0x3d, 0x6b, 0x3a, 0x98, 0xb9, 0xf3, 0xc7, 0x41, 0x27, 0xd1,
	0x8c, 0xe3, 0xcd, 0xb5, 0xff, 0xa3, 0x0f, 0xce, 0x3d, 0x92, 0x88, 0x04, 0x09, 0x76, 0x09, 0x36,
	0x55, 0x39, 0x1a, 0x0d, 0xa7, 0xf3, 0x67, 0xc7, 0x37, 0x2d, 0x23, 0xd8, 0x54, 0x39, 0x72, 0x43,
	0x62, 0xef, 0xc0, 0xd9, 0x15, 0x28, 0x28, 0x55, 0xd2, 0xc8, 0x73, 0xe7, 0x67, 0x41, 0xed, 0x36,
	0x68, 0xdd, 0x06, 0x9b, 0xd6, 0x2d, 0xef, 0xb8, 0xfa, 0x5d, 0xa6, 0xa2, 0x34, 0x4e, 0x31, 0x32,
	0xf2, 0xff, 0xf3, 0xae, 0xe5, 0xea, 0x80, 0x32, 0x15, 0xa1, 0x67, 0x4f, 0xad, 0xd9, 0x09, 0x37,
	0x67, 0x76, 0x0e, 0xee, 0x21, 0x95, 0x0f, 0x5b, 0x12, 0x45, 0x82, 0xe4, 0x0d, 0x4d, 0x76, 0xa0,
	0xa1, 0x8d, 0x41, 0xfc, 0x0b, 0xb0, 0xb5, 0x64, 0xe6, 0xc2, 0xf8, 0xf3, 0x6a, 0xb9, 0xfa, 0xf8,
	0x65, 0x35, 0xe9, 0x31, 0x07, 0xec, 0x0f, 0xb7, 0x77, 0x37, 0x13, 0x4b, 0xc3, 0xeb, 0xaf, 0xf7,
	0x77, 0xb7, 0xab, 0xe5, 0xa4, 0xef, 0x6f, 0x61, 0xfc, 0x5e, 0x49, 0x42, 0x49, 0x2c, 0x00, 0x27,
	0x6b, 0x2c, 0x9b, 0x30, 0xdc, 0x39, 0xfb, 0x33, 0x0c, 0xde, 0x71, 0xd8, 0x0b, 0xb0, 0x0d, 0xb7,
	0xce, 0x61, 0x72, 0xe4, 0xd6, 0xeb, 0xe0, 0xe6, 0xd6, 0x5f, 0xc0, 0xf0, 0x46, 0x52, 0x51, 0x69,
	0x2b, 0xb9, 0xa0, 0x7d, 0xbb, 0x6b, 0x7d, 0x66, 0x97, 0x30, 0xde, 0xd5, 0xd3, 0x9b, 0x2e, 0x4f,
	0x8e, 0x5d, 0x1a, 0x59, 0xbc, 0x65, 0xf8, 0x6f, 0xc1, 0x31, 0x9d, 0xd6, 0x48, 0xec, 0x15, 0x8c,
	0x51, 0x52, 0x91, 0xfe, 0x6d, 0xd7, 0x86, 0xc4, 0xdb, 0x7b, 0xff, 0xa7, 0x05, 0xf6, 0x27, 0x91,
	0x20, 0x7b, 0x0e, 0x4e, 0xa9, 0x0a, 0xda, 0x3e, 0x60, 0xd5, 0x88, 0x18, 0xeb, 0x7a, 0x89, 0x15,
	0xbb, 0x80, 0x51, 0x58, 0x08, 0xb9, 0xdb, 0xff, 0xcb, 0xcc, 0xa2, 0xc7, 0x1b, 0x06, 0x0b, 0x8e,
	0xa3, 0x07, 0xbf, 0xa7, 0xd4, 0xea, 0x5b, 0xf4, 0xba, 0xf9, 0xd7, 0x27, 0xe0, 0x46, 0x58, 0xee,
	0x50, 0x46, 0x28, 0xa9, 0xbc, 0x86, 0x6f, 0xdd, 0x47, 0x09, 0x47, 0x66, 0xf7, 0x6f, 0x7e, 0x05,
	0x00, 0x00, 0xff, 0xff, 0x4b, 0x57, 0x21, 0xa8, 0x4b, 0x03, 0x00, 0x00,
}
