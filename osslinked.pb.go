// Code generated by protoc-gen-go. DO NOT EDIT.
// source: osslinked.proto

package osslinked

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
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type Project struct {
	Name                 string        `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Author               string        `protobuf:"bytes,2,opt,name=author,proto3" json:"author,omitempty"`
	Email                string        `protobuf:"bytes,3,opt,name=email,proto3" json:"email,omitempty"`
	Version              string        `protobuf:"bytes,4,opt,name=version,proto3" json:"version,omitempty"`
	License              string        `protobuf:"bytes,5,opt,name=license,proto3" json:"license,omitempty"`
	Description          string        `protobuf:"bytes,6,opt,name=description,proto3" json:"description,omitempty"`
	Homepage             string        `protobuf:"bytes,7,opt,name=homepage,proto3" json:"homepage,omitempty"`
	Repository           *Repository   `protobuf:"bytes,8,opt,name=repository,proto3" json:"repository,omitempty"`
	Dependencies         []*Dependency `protobuf:"bytes,9,rep,name=dependencies,proto3" json:"dependencies,omitempty"`
	XXX_NoUnkeyedLiteral struct{}      `json:"-"`
	XXX_unrecognized     []byte        `json:"-"`
	XXX_sizecache        int32         `json:"-"`
}

func (m *Project) Reset()         { *m = Project{} }
func (m *Project) String() string { return proto.CompactTextString(m) }
func (*Project) ProtoMessage()    {}
func (*Project) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a93131606d3da39, []int{0}
}

func (m *Project) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Project.Unmarshal(m, b)
}
func (m *Project) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Project.Marshal(b, m, deterministic)
}
func (m *Project) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Project.Merge(m, src)
}
func (m *Project) XXX_Size() int {
	return xxx_messageInfo_Project.Size(m)
}
func (m *Project) XXX_DiscardUnknown() {
	xxx_messageInfo_Project.DiscardUnknown(m)
}

var xxx_messageInfo_Project proto.InternalMessageInfo

func (m *Project) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Project) GetAuthor() string {
	if m != nil {
		return m.Author
	}
	return ""
}

func (m *Project) GetEmail() string {
	if m != nil {
		return m.Email
	}
	return ""
}

func (m *Project) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func (m *Project) GetLicense() string {
	if m != nil {
		return m.License
	}
	return ""
}

func (m *Project) GetDescription() string {
	if m != nil {
		return m.Description
	}
	return ""
}

func (m *Project) GetHomepage() string {
	if m != nil {
		return m.Homepage
	}
	return ""
}

func (m *Project) GetRepository() *Repository {
	if m != nil {
		return m.Repository
	}
	return nil
}

func (m *Project) GetDependencies() []*Dependency {
	if m != nil {
		return m.Dependencies
	}
	return nil
}

type Repository struct {
	Type                 string   `protobuf:"bytes,1,opt,name=type,proto3" json:"type,omitempty"`
	Url                  string   `protobuf:"bytes,2,opt,name=url,proto3" json:"url,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Repository) Reset()         { *m = Repository{} }
func (m *Repository) String() string { return proto.CompactTextString(m) }
func (*Repository) ProtoMessage()    {}
func (*Repository) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a93131606d3da39, []int{1}
}

func (m *Repository) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Repository.Unmarshal(m, b)
}
func (m *Repository) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Repository.Marshal(b, m, deterministic)
}
func (m *Repository) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Repository.Merge(m, src)
}
func (m *Repository) XXX_Size() int {
	return xxx_messageInfo_Repository.Size(m)
}
func (m *Repository) XXX_DiscardUnknown() {
	xxx_messageInfo_Repository.DiscardUnknown(m)
}

var xxx_messageInfo_Repository proto.InternalMessageInfo

func (m *Repository) GetType() string {
	if m != nil {
		return m.Type
	}
	return ""
}

func (m *Repository) GetUrl() string {
	if m != nil {
		return m.Url
	}
	return ""
}

type Dependency struct {
	Name                 string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Version              string   `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *Dependency) Reset()         { *m = Dependency{} }
func (m *Dependency) String() string { return proto.CompactTextString(m) }
func (*Dependency) ProtoMessage()    {}
func (*Dependency) Descriptor() ([]byte, []int) {
	return fileDescriptor_6a93131606d3da39, []int{2}
}

func (m *Dependency) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Dependency.Unmarshal(m, b)
}
func (m *Dependency) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Dependency.Marshal(b, m, deterministic)
}
func (m *Dependency) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Dependency.Merge(m, src)
}
func (m *Dependency) XXX_Size() int {
	return xxx_messageInfo_Dependency.Size(m)
}
func (m *Dependency) XXX_DiscardUnknown() {
	xxx_messageInfo_Dependency.DiscardUnknown(m)
}

var xxx_messageInfo_Dependency proto.InternalMessageInfo

func (m *Dependency) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Dependency) GetVersion() string {
	if m != nil {
		return m.Version
	}
	return ""
}

func init() {
	proto.RegisterType((*Project)(nil), "osslinked.Project")
	proto.RegisterType((*Repository)(nil), "osslinked.Repository")
	proto.RegisterType((*Dependency)(nil), "osslinked.Dependency")
}

func init() { proto.RegisterFile("osslinked.proto", fileDescriptor_6a93131606d3da39) }

var fileDescriptor_6a93131606d3da39 = []byte{
	// 279 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x6c, 0x91, 0xcd, 0x4e, 0x83, 0x40,
	0x10, 0xc7, 0x53, 0x68, 0xa1, 0x4c, 0x4d, 0x34, 0x13, 0x35, 0x1b, 0x0f, 0x86, 0x70, 0xea, 0x09,
	0x93, 0x1a, 0x0f, 0x7a, 0x34, 0x3e, 0x80, 0xe1, 0xe8, 0x8d, 0xc2, 0xa4, 0xac, 0xc2, 0xee, 0x66,
	0x77, 0x31, 0xe1, 0xa9, 0x7c, 0x45, 0xc3, 0x96, 0x02, 0x35, 0xde, 0xfe, 0x1f, 0x33, 0x13, 0xf8,
	0x2d, 0x5c, 0x4a, 0x63, 0x6a, 0x2e, 0xbe, 0xa8, 0x4c, 0x95, 0x96, 0x56, 0x62, 0x34, 0x06, 0xc9,
	0x8f, 0x07, 0xe1, 0xbb, 0x96, 0x9f, 0x54, 0x58, 0x44, 0x58, 0x8a, 0xbc, 0x21, 0xb6, 0x88, 0x17,
	0xdb, 0x28, 0x73, 0x1a, 0x6f, 0x21, 0xc8, 0x5b, 0x5b, 0x49, 0xcd, 0x3c, 0x97, 0x0e, 0x0e, 0xaf,
	0x61, 0x45, 0x4d, 0xce, 0x6b, 0xe6, 0xbb, 0xf8, 0x68, 0x90, 0x41, 0xf8, 0x4d, 0xda, 0x70, 0x29,
	0xd8, 0xd2, 0xe5, 0x27, 0xdb, 0x37, 0x35, 0x2f, 0x48, 0x18, 0x62, 0xab, 0x63, 0x33, 0x58, 0x8c,
	0x61, 0x53, 0x92, 0x29, 0x34, 0x57, 0xb6, 0xdf, 0x0b, 0x5c, 0x3b, 0x8f, 0xf0, 0x0e, 0xd6, 0x95,
	0x6c, 0x48, 0xe5, 0x07, 0x62, 0xa1, 0xab, 0x47, 0x8f, 0x4f, 0x00, 0x9a, 0x94, 0x34, 0xdc, 0x4a,
	0xdd, 0xb1, 0x75, 0xbc, 0xd8, 0x6e, 0x76, 0x37, 0xe9, 0xf4, 0xc3, 0xd9, 0x58, 0x66, 0xb3, 0x41,
	0x7c, 0x86, 0x8b, 0x92, 0x14, 0x89, 0x92, 0x44, 0xc1, 0xc9, 0xb0, 0x28, 0xf6, 0xff, 0x2c, 0xbe,
	0x9d, 0xea, 0x2e, 0x3b, 0x1b, 0x4d, 0x76, 0x00, 0xd3, 0xd1, 0x9e, 0x99, 0xed, 0xd4, 0xc8, 0xac,
	0xd7, 0x78, 0x05, 0x7e, 0xab, 0xeb, 0x01, 0x58, 0x2f, 0x93, 0x17, 0x80, 0xe9, 0xde, 0xbf, 0x9c,
	0x67, 0xe4, 0xbc, 0x33, 0x72, 0xaf, 0xf1, 0xc7, 0xfd, 0x81, 0xdb, 0xaa, 0xdd, 0xa7, 0x85, 0x6c,
	0x1e, 0xc6, 0x0f, 0x9c, 0xd4, 0x3e, 0x70, 0xaf, 0xfa, 0xf8, 0x1b, 0x00, 0x00, 0xff, 0xff, 0xd6,
	0x20, 0x36, 0x77, 0xe8, 0x01, 0x00, 0x00,
}
