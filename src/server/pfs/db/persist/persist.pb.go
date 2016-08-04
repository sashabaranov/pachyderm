// Code generated by protoc-gen-go.
// source: server/pfs/db/persist/persist.proto
// DO NOT EDIT!

/*
Package persist is a generated protocol buffer package.

It is generated from these files:
	server/pfs/db/persist/persist.proto

It has these top-level messages:
	Clock
	ClockID
	BranchClock
	Repo
	Branch
	BlockRef
	Diff
	Commit
*/
package persist

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/pb/go/google/protobuf"
import _ "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type FileType int32

const (
	FileType_NONE FileType = 0
	FileType_FILE FileType = 1
	FileType_DIR  FileType = 2
)

var FileType_name = map[int32]string{
	0: "NONE",
	1: "FILE",
	2: "DIR",
}
var FileType_value = map[string]int32{
	"NONE": 0,
	"FILE": 1,
	"DIR":  2,
}

func (x FileType) String() string {
	return proto.EnumName(FileType_name, int32(x))
}
func (FileType) EnumDescriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type Clock struct {
	// a document either has these two fields
	Branch string `protobuf:"bytes,1,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,2,opt,name=clock" json:"clock,omitempty"`
}

func (m *Clock) Reset()                    { *m = Clock{} }
func (m *Clock) String() string            { return proto.CompactTextString(m) }
func (*Clock) ProtoMessage()               {}
func (*Clock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

type ClockID struct {
	ID     string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo   string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Branch string `protobuf:"bytes,3,opt,name=branch" json:"branch,omitempty"`
	Clock  uint64 `protobuf:"varint,4,opt,name=clock" json:"clock,omitempty"`
}

func (m *ClockID) Reset()                    { *m = ClockID{} }
func (m *ClockID) String() string            { return proto.CompactTextString(m) }
func (*ClockID) ProtoMessage()               {}
func (*ClockID) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

type BranchClock struct {
	Clocks []*Clock `protobuf:"bytes,1,rep,name=clocks" json:"clocks,omitempty"`
}

func (m *BranchClock) Reset()                    { *m = BranchClock{} }
func (m *BranchClock) String() string            { return proto.CompactTextString(m) }
func (*BranchClock) ProtoMessage()               {}
func (*BranchClock) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *BranchClock) GetClocks() []*Clock {
	if m != nil {
		return m.Clocks
	}
	return nil
}

type Repo struct {
	Name    string                     `protobuf:"bytes,1,opt,name=name" json:"name,omitempty"`
	Created *google_protobuf.Timestamp `protobuf:"bytes,2,opt,name=created" json:"created,omitempty"`
}

func (m *Repo) Reset()                    { *m = Repo{} }
func (m *Repo) String() string            { return proto.CompactTextString(m) }
func (*Repo) ProtoMessage()               {}
func (*Repo) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

func (m *Repo) GetCreated() *google_protobuf.Timestamp {
	if m != nil {
		return m.Created
	}
	return nil
}

type Branch struct {
	ID   string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	Name string `protobuf:"bytes,3,opt,name=name" json:"name,omitempty"`
}

func (m *Branch) Reset()                    { *m = Branch{} }
func (m *Branch) String() string            { return proto.CompactTextString(m) }
func (*Branch) ProtoMessage()               {}
func (*Branch) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type BlockRef struct {
	Hash  string `protobuf:"bytes,1,opt,name=hash" json:"hash,omitempty"`
	Lower uint64 `protobuf:"varint,2,opt,name=lower" json:"lower,omitempty"`
	Upper uint64 `protobuf:"varint,3,opt,name=upper" json:"upper,omitempty"`
}

func (m *BlockRef) Reset()                    { *m = BlockRef{} }
func (m *BlockRef) String() string            { return proto.CompactTextString(m) }
func (*BlockRef) ProtoMessage()               {}
func (*BlockRef) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

type Diff struct {
	ID       string `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo     string `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	CommitID string `protobuf:"bytes,3,opt,name=commit_id,json=commitId" json:"commit_id,omitempty"`
	Path     string `protobuf:"bytes,4,opt,name=path" json:"path,omitempty"`
	// block_refs and delete cannot both be set
	BlockRefs    []*BlockRef                `protobuf:"bytes,5,rep,name=block_refs,json=blockRefs" json:"block_refs,omitempty"`
	Delete       bool                       `protobuf:"varint,6,opt,name=delete" json:"delete,omitempty"`
	Size         uint64                     `protobuf:"varint,7,opt,name=size" json:"size,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,8,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	FileType     FileType                   `protobuf:"varint,9,opt,name=file_type,json=fileType,enum=FileType" json:"file_type,omitempty"`
	Modified     *google_protobuf.Timestamp `protobuf:"bytes,10,opt,name=modified" json:"modified,omitempty"`
}

func (m *Diff) Reset()                    { *m = Diff{} }
func (m *Diff) String() string            { return proto.CompactTextString(m) }
func (*Diff) ProtoMessage()               {}
func (*Diff) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *Diff) GetBlockRefs() []*BlockRef {
	if m != nil {
		return m.BlockRefs
	}
	return nil
}

func (m *Diff) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Diff) GetModified() *google_protobuf.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Commit struct {
	ID           string                     `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Repo         string                     `protobuf:"bytes,2,opt,name=repo" json:"repo,omitempty"`
	BranchClocks []*BranchClock             `protobuf:"bytes,3,rep,name=branch_clocks,json=branchClocks" json:"branch_clocks,omitempty"`
	Started      *google_protobuf.Timestamp `protobuf:"bytes,4,opt,name=started" json:"started,omitempty"`
	Finished     *google_protobuf.Timestamp `protobuf:"bytes,5,opt,name=finished" json:"finished,omitempty"`
	Cancelled    bool                       `protobuf:"varint,6,opt,name=cancelled" json:"cancelled,omitempty"`
	Provenance   []string                   `protobuf:"bytes,7,rep,name=provenance" json:"provenance,omitempty"`
}

func (m *Commit) Reset()                    { *m = Commit{} }
func (m *Commit) String() string            { return proto.CompactTextString(m) }
func (*Commit) ProtoMessage()               {}
func (*Commit) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *Commit) GetBranchClocks() []*BranchClock {
	if m != nil {
		return m.BranchClocks
	}
	return nil
}

func (m *Commit) GetStarted() *google_protobuf.Timestamp {
	if m != nil {
		return m.Started
	}
	return nil
}

func (m *Commit) GetFinished() *google_protobuf.Timestamp {
	if m != nil {
		return m.Finished
	}
	return nil
}

func init() {
	proto.RegisterType((*Clock)(nil), "Clock")
	proto.RegisterType((*ClockID)(nil), "ClockID")
	proto.RegisterType((*BranchClock)(nil), "BranchClock")
	proto.RegisterType((*Repo)(nil), "Repo")
	proto.RegisterType((*Branch)(nil), "Branch")
	proto.RegisterType((*BlockRef)(nil), "BlockRef")
	proto.RegisterType((*Diff)(nil), "Diff")
	proto.RegisterType((*Commit)(nil), "Commit")
	proto.RegisterEnum("FileType", FileType_name, FileType_value)
}

func init() { proto.RegisterFile("server/pfs/db/persist/persist.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 540 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x8c, 0x53, 0x5d, 0x8b, 0xd3, 0x40,
	0x14, 0xb5, 0x4d, 0x9a, 0x26, 0xb7, 0xeb, 0x52, 0x06, 0x91, 0xb0, 0xca, 0x2a, 0x11, 0xb4, 0x08,
	0xa6, 0xb8, 0x7e, 0x3c, 0xcb, 0x6e, 0x77, 0xa1, 0x22, 0xab, 0x0c, 0xfb, 0xe6, 0x43, 0x49, 0x9a,
	0xc9, 0x76, 0x30, 0x69, 0xc2, 0x4c, 0x76, 0x17, 0xfd, 0x0b, 0xfe, 0x2d, 0x7f, 0x98, 0x77, 0x6e,
	0x32, 0xb6, 0xa8, 0xd0, 0x3e, 0xf5, 0xde, 0x93, 0xfb, 0x71, 0xee, 0x39, 0x53, 0x78, 0xa6, 0x85,
	0xba, 0x15, 0x6a, 0x5a, 0xe7, 0x7a, 0x9a, 0xa5, 0xd3, 0x5a, 0x28, 0x2d, 0x75, 0x63, 0x7f, 0xe3,
	0x5a, 0x55, 0x4d, 0x75, 0xf4, 0xe4, 0xba, 0xaa, 0xae, 0x0b, 0x31, 0xa5, 0x2c, 0xbd, 0xc9, 0xa7,
	0x8d, 0x2c, 0x85, 0x6e, 0x92, 0xb2, 0xee, 0x0a, 0x8e, 0xff, 0x2e, 0xb8, 0x53, 0x49, 0x6d, 0x66,
	0xb4, 0xdf, 0xa3, 0x77, 0x30, 0x38, 0x2b, 0xaa, 0xe5, 0x37, 0xf6, 0x10, 0xbc, 0x54, 0x25, 0xeb,
	0xe5, 0x2a, 0xec, 0x3d, 0xed, 0x4d, 0x02, 0xde, 0x65, 0xec, 0x01, 0x0c, 0x96, 0xa6, 0x20, 0xec,
	0x23, 0xec, 0xf2, 0x36, 0x89, 0xbe, 0xc2, 0x90, 0xda, 0xe6, 0x33, 0x76, 0x08, 0x7d, 0x99, 0x75,
	0x4d, 0x18, 0x31, 0x06, 0xae, 0x12, 0x75, 0x45, 0xf5, 0x01, 0xa7, 0x78, 0x6b, 0xb8, 0xf3, 0xff,
	0xe1, 0xee, 0xf6, 0xf0, 0x57, 0x30, 0x3a, 0xa5, 0xef, 0x2d, 0xb3, 0x63, 0xf0, 0x08, 0xd7, 0xb8,
	0xc4, 0x99, 0x8c, 0x4e, 0xbc, 0x98, 0x70, 0xde, 0xa1, 0xd1, 0x17, 0x70, 0xb9, 0x59, 0x82, 0x8b,
	0xd7, 0x49, 0x29, 0x3a, 0x2a, 0x14, 0xb3, 0xb7, 0x30, 0x5c, 0x2a, 0x91, 0x34, 0x22, 0x23, 0x3e,
	0xa3, 0x93, 0xa3, 0xb8, 0x15, 0x24, 0xb6, 0x82, 0xc4, 0x57, 0x56, 0x31, 0x6e, 0x4b, 0xa3, 0x0f,
	0xe0, 0xb5, 0x04, 0xf6, 0x3a, 0xce, 0xee, 0x75, 0x36, 0x7b, 0xa3, 0x8f, 0xe0, 0x9f, 0x12, 0x49,
	0x91, 0x9b, 0xef, 0xab, 0x44, 0x5b, 0x5d, 0x29, 0x36, 0x87, 0x17, 0xd5, 0x9d, 0x50, 0x56, 0x55,
	0x4a, 0x0c, 0x7a, 0x63, 0xcc, 0xa1, 0x51, 0x88, 0x52, 0x12, 0xfd, 0xea, 0x83, 0x3b, 0x93, 0x79,
	0xbe, 0x17, 0x99, 0x47, 0x10, 0x2c, 0xab, 0xb2, 0x94, 0xcd, 0x02, 0x4b, 0x5b, 0x46, 0x7e, 0x0b,
	0xcc, 0xa9, 0xa1, 0x4e, 0x9a, 0x15, 0xa9, 0x8d, 0x0d, 0x26, 0x66, 0x13, 0x80, 0xd4, 0x30, 0x5d,
	0x28, 0x91, 0xeb, 0x70, 0x40, 0x0a, 0x07, 0xb1, 0x25, 0xcf, 0x83, 0xb4, 0x8b, 0xb4, 0x31, 0x31,
	0x13, 0x85, 0x68, 0x44, 0xe8, 0x61, 0xbf, 0xcf, 0xbb, 0xcc, 0x4c, 0xd5, 0xf2, 0x87, 0x08, 0x87,
	0x44, 0x9a, 0x62, 0xf6, 0x1a, 0xee, 0xb7, 0x16, 0x2f, 0x3a, 0xeb, 0x7c, 0x1a, 0x7c, 0x10, 0x6f,
	0x19, 0xcb, 0x0f, 0xd2, 0x4d, 0xa2, 0xd9, 0x73, 0x08, 0x72, 0x59, 0x88, 0x45, 0xf3, 0xbd, 0x16,
	0x61, 0x80, 0xb3, 0x0e, 0x91, 0xc7, 0x05, 0x22, 0x57, 0x08, 0x70, 0x3f, 0xef, 0x22, 0xf6, 0x1e,
	0xfc, 0xb2, 0xca, 0x64, 0x2e, 0xd1, 0x53, 0xd8, 0xe9, 0xe9, 0x9f, 0xda, 0xe8, 0x67, 0x1f, 0xbc,
	0x33, 0x52, 0x62, 0x2f, 0x21, 0xff, 0xb9, 0xc0, 0xd9, 0x79, 0x01, 0x3e, 0x36, 0x5c, 0xaa, 0xcc,
	0x63, 0x73, 0x77, 0x3f, 0xb6, 0xae, 0xd4, 0xdc, 0x93, 0xcb, 0xb5, 0xd4, 0x2b, 0x6c, 0x1b, 0xec,
	0xbe, 0xc7, 0xd6, 0xb2, 0xc7, 0xe8, 0x34, 0x2e, 0x17, 0x45, 0x81, 0x8d, 0xad, 0x23, 0x1b, 0x00,
	0xff, 0x34, 0x80, 0xdd, 0xb7, 0x62, 0x6d, 0x10, 0xb4, 0xc6, 0xc1, 0xc3, 0xb6, 0x90, 0x97, 0x2f,
	0xc0, 0xb7, 0xda, 0x32, 0x1f, 0xdc, 0xcb, 0xcf, 0x97, 0xe7, 0xe3, 0x7b, 0x26, 0xba, 0x98, 0x7f,
	0x3a, 0x1f, 0xf7, 0xd8, 0x10, 0x9c, 0xd9, 0x9c, 0x8f, 0xfb, 0xa9, 0x47, 0x24, 0xde, 0xfc, 0x0e,
	0x00, 0x00, 0xff, 0xff, 0xdb, 0x8d, 0x30, 0x52, 0x8f, 0x04, 0x00, 0x00,
}