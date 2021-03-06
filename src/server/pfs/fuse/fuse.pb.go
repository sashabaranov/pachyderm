// Code generated by protoc-gen-go.
// source: server/pfs/fuse/fuse.proto
// DO NOT EDIT!

/*
Package fuse is a generated protocol buffer package.

It is generated from these files:
	server/pfs/fuse/fuse.proto

It has these top-level messages:
	CommitMount
	Filesystem
	Node
	Attr
	Dirent
	Root
	DirectoryAttr
	DirectoryLookup
	DirectoryReadDirAll
	DirectoryCreate
	DirectoryMkdir
	FileAttr
	FileSetAttr
	FileRead
	FileOpen
	FileWrite
	FileRemove
*/
package fuse

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import pfs "github.com/pachyderm/pachyderm/src/client/pfs"
import google_protobuf2 "go.pedge.io/pb/go/google/protobuf"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type CommitMount struct {
	Commit     *pfs.Commit     `protobuf:"bytes,1,opt,name=commit" json:"commit,omitempty"`
	DiffMethod *pfs.DiffMethod `protobuf:"bytes,2,opt,name=diff_method,json=diffMethod" json:"diff_method,omitempty"`
	FullFile   bool            `protobuf:"varint,3,opt,name=full_file,json=fullFile" json:"full_file,omitempty"`
	Alias      string          `protobuf:"bytes,4,opt,name=alias" json:"alias,omitempty"`
	Shard      *pfs.Shard      `protobuf:"bytes,5,opt,name=shard" json:"shard,omitempty"`
	Lazy       bool            `protobuf:"varint,6,opt,name=lazy" json:"lazy,omitempty"`
}

func (m *CommitMount) Reset()                    { *m = CommitMount{} }
func (m *CommitMount) String() string            { return proto.CompactTextString(m) }
func (*CommitMount) ProtoMessage()               {}
func (*CommitMount) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{0} }

func (m *CommitMount) GetCommit() *pfs.Commit {
	if m != nil {
		return m.Commit
	}
	return nil
}

func (m *CommitMount) GetDiffMethod() *pfs.DiffMethod {
	if m != nil {
		return m.DiffMethod
	}
	return nil
}

func (m *CommitMount) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

type Filesystem struct {
	Shard        *pfs.Shard     `protobuf:"bytes,1,opt,name=shard" json:"shard,omitempty"`
	CommitMounts []*CommitMount `protobuf:"bytes,2,rep,name=commit_mounts,json=commitMounts" json:"commit_mounts,omitempty"`
}

func (m *Filesystem) Reset()                    { *m = Filesystem{} }
func (m *Filesystem) String() string            { return proto.CompactTextString(m) }
func (*Filesystem) ProtoMessage()               {}
func (*Filesystem) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{1} }

func (m *Filesystem) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *Filesystem) GetCommitMounts() []*CommitMount {
	if m != nil {
		return m.CommitMounts
	}
	return nil
}

type Node struct {
	File      *pfs.File                   `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	RepoAlias string                      `protobuf:"bytes,2,opt,name=repo_alias,json=repoAlias" json:"repo_alias,omitempty"`
	Write     bool                        `protobuf:"varint,3,opt,name=write" json:"write,omitempty"`
	Shard     *pfs.Shard                  `protobuf:"bytes,4,opt,name=shard" json:"shard,omitempty"`
	Modified  *google_protobuf2.Timestamp `protobuf:"bytes,5,opt,name=modified" json:"modified,omitempty"`
}

func (m *Node) Reset()                    { *m = Node{} }
func (m *Node) String() string            { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()               {}
func (*Node) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{2} }

func (m *Node) GetFile() *pfs.File {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *Node) GetShard() *pfs.Shard {
	if m != nil {
		return m.Shard
	}
	return nil
}

func (m *Node) GetModified() *google_protobuf2.Timestamp {
	if m != nil {
		return m.Modified
	}
	return nil
}

type Attr struct {
	Mode uint32 `protobuf:"varint,1,opt,name=Mode,json=mode" json:"Mode,omitempty"`
}

func (m *Attr) Reset()                    { *m = Attr{} }
func (m *Attr) String() string            { return proto.CompactTextString(m) }
func (*Attr) ProtoMessage()               {}
func (*Attr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{3} }

type Dirent struct {
	Inode uint64 `protobuf:"varint,1,opt,name=inode" json:"inode,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
}

func (m *Dirent) Reset()                    { *m = Dirent{} }
func (m *Dirent) String() string            { return proto.CompactTextString(m) }
func (*Dirent) ProtoMessage()               {}
func (*Dirent) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{4} }

type Root struct {
	Filesystem *Filesystem `protobuf:"bytes,1,opt,name=filesystem" json:"filesystem,omitempty"`
	Result     *Node       `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error      string      `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *Root) Reset()                    { *m = Root{} }
func (m *Root) String() string            { return proto.CompactTextString(m) }
func (*Root) ProtoMessage()               {}
func (*Root) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{5} }

func (m *Root) GetFilesystem() *Filesystem {
	if m != nil {
		return m.Filesystem
	}
	return nil
}

func (m *Root) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryAttr struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryAttr) Reset()                    { *m = DirectoryAttr{} }
func (m *DirectoryAttr) String() string            { return proto.CompactTextString(m) }
func (*DirectoryAttr) ProtoMessage()               {}
func (*DirectoryAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{6} }

func (m *DirectoryAttr) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryLookup struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Result    *Node  `protobuf:"bytes,3,opt,name=result" json:"result,omitempty"`
	Err       string `protobuf:"bytes,4,opt,name=err" json:"err,omitempty"`
}

func (m *DirectoryLookup) Reset()                    { *m = DirectoryLookup{} }
func (m *DirectoryLookup) String() string            { return proto.CompactTextString(m) }
func (*DirectoryLookup) ProtoMessage()               {}
func (*DirectoryLookup) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{7} }

func (m *DirectoryLookup) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryLookup) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryReadDirAll struct {
	Directory *Node     `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    []*Dirent `protobuf:"bytes,2,rep,name=result" json:"result,omitempty"`
	Error     string    `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryReadDirAll) Reset()                    { *m = DirectoryReadDirAll{} }
func (m *DirectoryReadDirAll) String() string            { return proto.CompactTextString(m) }
func (*DirectoryReadDirAll) ProtoMessage()               {}
func (*DirectoryReadDirAll) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{8} }

func (m *DirectoryReadDirAll) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryReadDirAll) GetResult() []*Dirent {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryCreate struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryCreate) Reset()                    { *m = DirectoryCreate{} }
func (m *DirectoryCreate) String() string            { return proto.CompactTextString(m) }
func (*DirectoryCreate) ProtoMessage()               {}
func (*DirectoryCreate) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{9} }

func (m *DirectoryCreate) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryCreate) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type DirectoryMkdir struct {
	Directory *Node  `protobuf:"bytes,1,opt,name=directory" json:"directory,omitempty"`
	Result    *Node  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error     string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *DirectoryMkdir) Reset()                    { *m = DirectoryMkdir{} }
func (m *DirectoryMkdir) String() string            { return proto.CompactTextString(m) }
func (*DirectoryMkdir) ProtoMessage()               {}
func (*DirectoryMkdir) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{10} }

func (m *DirectoryMkdir) GetDirectory() *Node {
	if m != nil {
		return m.Directory
	}
	return nil
}

func (m *DirectoryMkdir) GetResult() *Node {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileAttr struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Result *Attr  `protobuf:"bytes,2,opt,name=result" json:"result,omitempty"`
	Error  string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *FileAttr) Reset()                    { *m = FileAttr{} }
func (m *FileAttr) String() string            { return proto.CompactTextString(m) }
func (*FileAttr) ProtoMessage()               {}
func (*FileAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{11} }

func (m *FileAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func (m *FileAttr) GetResult() *Attr {
	if m != nil {
		return m.Result
	}
	return nil
}

type FileSetAttr struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileSetAttr) Reset()                    { *m = FileSetAttr{} }
func (m *FileSetAttr) String() string            { return proto.CompactTextString(m) }
func (*FileSetAttr) ProtoMessage()               {}
func (*FileSetAttr) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{12} }

func (m *FileSetAttr) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileRead struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Data  string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Error string `protobuf:"bytes,3,opt,name=error" json:"error,omitempty"`
}

func (m *FileRead) Reset()                    { *m = FileRead{} }
func (m *FileRead) String() string            { return proto.CompactTextString(m) }
func (*FileRead) ProtoMessage()               {}
func (*FileRead) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{13} }

func (m *FileRead) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileOpen struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Error string `protobuf:"bytes,2,opt,name=error" json:"error,omitempty"`
}

func (m *FileOpen) Reset()                    { *m = FileOpen{} }
func (m *FileOpen) String() string            { return proto.CompactTextString(m) }
func (*FileOpen) ProtoMessage()               {}
func (*FileOpen) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{14} }

func (m *FileOpen) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileWrite struct {
	File   *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Data   string `protobuf:"bytes,2,opt,name=data" json:"data,omitempty"`
	Offset int64  `protobuf:"varint,3,opt,name=offset" json:"offset,omitempty"`
	Error  string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *FileWrite) Reset()                    { *m = FileWrite{} }
func (m *FileWrite) String() string            { return proto.CompactTextString(m) }
func (*FileWrite) ProtoMessage()               {}
func (*FileWrite) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{15} }

func (m *FileWrite) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

type FileRemove struct {
	File  *Node  `protobuf:"bytes,1,opt,name=file" json:"file,omitempty"`
	Name  string `protobuf:"bytes,2,opt,name=name" json:"name,omitempty"`
	Dir   bool   `protobuf:"varint,3,opt,name=dir" json:"dir,omitempty"`
	Error string `protobuf:"bytes,4,opt,name=error" json:"error,omitempty"`
}

func (m *FileRemove) Reset()                    { *m = FileRemove{} }
func (m *FileRemove) String() string            { return proto.CompactTextString(m) }
func (*FileRemove) ProtoMessage()               {}
func (*FileRemove) Descriptor() ([]byte, []int) { return fileDescriptor0, []int{16} }

func (m *FileRemove) GetFile() *Node {
	if m != nil {
		return m.File
	}
	return nil
}

func init() {
	proto.RegisterType((*CommitMount)(nil), "fuse.CommitMount")
	proto.RegisterType((*Filesystem)(nil), "fuse.Filesystem")
	proto.RegisterType((*Node)(nil), "fuse.Node")
	proto.RegisterType((*Attr)(nil), "fuse.Attr")
	proto.RegisterType((*Dirent)(nil), "fuse.Dirent")
	proto.RegisterType((*Root)(nil), "fuse.Root")
	proto.RegisterType((*DirectoryAttr)(nil), "fuse.DirectoryAttr")
	proto.RegisterType((*DirectoryLookup)(nil), "fuse.DirectoryLookup")
	proto.RegisterType((*DirectoryReadDirAll)(nil), "fuse.DirectoryReadDirAll")
	proto.RegisterType((*DirectoryCreate)(nil), "fuse.DirectoryCreate")
	proto.RegisterType((*DirectoryMkdir)(nil), "fuse.DirectoryMkdir")
	proto.RegisterType((*FileAttr)(nil), "fuse.FileAttr")
	proto.RegisterType((*FileSetAttr)(nil), "fuse.FileSetAttr")
	proto.RegisterType((*FileRead)(nil), "fuse.FileRead")
	proto.RegisterType((*FileOpen)(nil), "fuse.FileOpen")
	proto.RegisterType((*FileWrite)(nil), "fuse.FileWrite")
	proto.RegisterType((*FileRemove)(nil), "fuse.FileRemove")
}

func init() { proto.RegisterFile("server/pfs/fuse/fuse.proto", fileDescriptor0) }

var fileDescriptor0 = []byte{
	// 669 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xbc, 0x54, 0x4f, 0x6f, 0xd3, 0x4e,
	0x10, 0x95, 0x13, 0x27, 0x8a, 0x27, 0xed, 0xaf, 0xfd, 0x99, 0x0a, 0x59, 0x41, 0x85, 0xc8, 0x70,
	0xc8, 0x29, 0x41, 0x45, 0xe2, 0x4c, 0xd5, 0x8a, 0x13, 0x01, 0x69, 0x5b, 0x89, 0x63, 0xe4, 0x66,
	0x67, 0xdb, 0x55, 0x6d, 0xaf, 0xb5, 0xbb, 0x2e, 0x2a, 0x9c, 0xf9, 0x46, 0x7c, 0x0d, 0xbe, 0x13,
	0xda, 0x5d, 0xff, 0x43, 0x34, 0x6a, 0x1b, 0x24, 0x2e, 0xd1, 0xcc, 0xee, 0xe4, 0xbd, 0x37, 0x6f,
	0x66, 0x0d, 0x13, 0x85, 0xf2, 0x06, 0xe5, 0xa2, 0x60, 0x6a, 0xc1, 0x4a, 0x85, 0xf6, 0x67, 0x5e,
	0x48, 0xa1, 0x45, 0xe8, 0x9b, 0x78, 0x72, 0xb0, 0x4e, 0x39, 0xe6, 0xda, 0x56, 0x14, 0x4c, 0xb9,
	0xbb, 0xc9, 0x8b, 0x4b, 0x21, 0x2e, 0x53, 0x5c, 0xd8, 0xec, 0xa2, 0x64, 0x0b, 0xcd, 0x33, 0x54,
	0x3a, 0xc9, 0x0a, 0x57, 0x10, 0xff, 0xf4, 0x60, 0x7c, 0x22, 0xb2, 0x8c, 0xeb, 0xa5, 0x28, 0x73,
	0x1d, 0xbe, 0x84, 0xe1, 0xda, 0xa6, 0x91, 0x37, 0xf5, 0x66, 0xe3, 0xa3, 0xf1, 0xdc, 0x80, 0xb9,
	0x0a, 0x52, 0x5d, 0x85, 0xaf, 0x61, 0x4c, 0x39, 0x63, 0xab, 0x0c, 0xf5, 0x95, 0xa0, 0x51, 0xcf,
	0x56, 0xee, 0xd9, 0xca, 0x53, 0xce, 0xd8, 0xd2, 0x1e, 0x13, 0xa0, 0x4d, 0x1c, 0x3e, 0x83, 0x80,
	0x95, 0x69, 0xba, 0x62, 0x3c, 0xc5, 0xa8, 0x3f, 0xf5, 0x66, 0x23, 0x32, 0x32, 0x07, 0xef, 0x79,
	0x8a, 0xe1, 0x01, 0x0c, 0x92, 0x94, 0x27, 0x2a, 0xf2, 0xa7, 0xde, 0x2c, 0x20, 0x2e, 0x09, 0xa7,
	0x30, 0x50, 0x57, 0x89, 0xa4, 0xd1, 0xc0, 0xc2, 0x83, 0x85, 0x3f, 0x33, 0x27, 0xc4, 0x5d, 0x84,
	0x21, 0xf8, 0x69, 0xf2, 0xf5, 0x36, 0x1a, 0x5a, 0x3c, 0x1b, 0xc7, 0x0c, 0xc0, 0x60, 0xaa, 0x5b,
	0xa5, 0x31, 0x6b, 0x31, 0xbc, 0x4d, 0x18, 0x6f, 0x61, 0xd7, 0x35, 0xb5, 0xca, 0x4c, 0xff, 0x2a,
	0xea, 0x4d, 0xfb, 0xb3, 0xf1, 0xd1, 0xff, 0x73, 0x6b, 0x70, 0xc7, 0x19, 0xb2, 0xb3, 0x6e, 0x13,
	0x15, 0xff, 0xf0, 0xc0, 0xff, 0x28, 0x28, 0x86, 0x87, 0xe0, 0xdb, 0xa6, 0x1c, 0x43, 0x60, 0x19,
	0x8c, 0x02, 0x62, 0x8f, 0xc3, 0x43, 0x00, 0x89, 0x85, 0x58, 0xb9, 0x06, 0x7b, 0xb6, 0xc1, 0xc0,
	0x9c, 0x1c, 0xdb, 0x26, 0x0f, 0x60, 0xf0, 0x45, 0x72, 0x5d, 0x7b, 0xe2, 0x92, 0x56, 0xb6, 0xbf,
	0x59, 0xf6, 0x28, 0x13, 0x94, 0x33, 0x8e, 0xb5, 0x3f, 0x93, 0xb9, 0x1b, 0xf5, 0xbc, 0x1e, 0xf5,
	0xfc, 0xbc, 0x1e, 0x35, 0x69, 0x6a, 0xe3, 0x09, 0xf8, 0xc7, 0x5a, 0x4b, 0x63, 0xdd, 0x52, 0x50,
	0xa7, 0x7a, 0x97, 0xf8, 0x99, 0xa0, 0x18, 0x1f, 0xc1, 0xf0, 0x94, 0x4b, 0xcc, 0xb5, 0x51, 0xc5,
	0xf3, 0xfa, 0xda, 0x27, 0x2e, 0x31, 0xff, 0xc9, 0x93, 0x0c, 0xab, 0x26, 0x6c, 0x1c, 0x4b, 0xf0,
	0x89, 0x10, 0x66, 0x23, 0x80, 0x35, 0xb6, 0x57, 0x5e, 0xec, 0x3b, 0x0f, 0xdb, 0x71, 0x90, 0x4e,
	0x4d, 0x18, 0xc3, 0x50, 0xa2, 0x2a, 0x53, 0x5d, 0xad, 0x0f, 0xb8, 0x6a, 0xe3, 0x29, 0xa9, 0x6e,
	0x8c, 0x0e, 0x94, 0x52, 0x48, 0xeb, 0x4e, 0x40, 0x5c, 0x12, 0x2b, 0xd8, 0x35, 0x3a, 0xd7, 0x5a,
	0xc8, 0x5b, 0xdb, 0xcc, 0x0c, 0x02, 0x5a, 0x1f, 0x34, 0x93, 0x6e, 0xd1, 0xda, 0xcb, 0x4d, 0xa4,
	0x06, 0xe5, 0x1e, 0xd2, 0xef, 0x1e, 0xec, 0x35, 0xac, 0x1f, 0x84, 0xb8, 0x2e, 0x8b, 0x47, 0xf0,
	0xde, 0x61, 0x5d, 0x47, 0x4b, 0x7f, 0xa3, 0x01, 0xfb, 0xd0, 0x47, 0x29, 0xab, 0x77, 0x61, 0xc2,
	0xf8, 0x1b, 0x3c, 0x69, 0x64, 0x10, 0x4c, 0xe8, 0x29, 0x97, 0xc7, 0x69, 0xfa, 0x08, 0x29, 0xaf,
	0x3a, 0x16, 0x98, 0x4d, 0xdf, 0x71, 0x65, 0x6e, 0xf2, 0xf7, 0x98, 0x50, 0x76, 0x3c, 0x38, 0x91,
	0x98, 0x68, 0xfc, 0x7b, 0xef, 0x1f, 0x30, 0x70, 0x0d, 0xff, 0x35, 0xb4, 0xcb, 0x6b, 0xca, 0xe5,
	0x3f, 0x61, 0xa5, 0x30, 0x32, 0xab, 0x6b, 0x37, 0xec, 0xf9, 0x6f, 0x8f, 0xbc, 0x8b, 0xe1, 0x5e,
	0xf9, 0xf6, 0x7b, 0x75, 0x02, 0x63, 0xc3, 0x72, 0x86, 0xfa, 0x41, 0x44, 0x0d, 0x48, 0xaf, 0x0b,
	0x72, 0xee, 0xa4, 0x9a, 0x7d, 0xb8, 0x17, 0x21, 0x04, 0x9f, 0x26, 0x3a, 0xa9, 0x57, 0xd1, 0xc4,
	0x1b, 0xa4, 0xbd, 0x73, 0xa8, 0x9f, 0x0a, 0xcc, 0xb7, 0xd4, 0x95, 0x41, 0x60, 0x10, 0x3e, 0xdb,
	0x8f, 0xda, 0x36, 0xc2, 0x9e, 0xc2, 0x50, 0x30, 0xa6, 0xd0, 0xbd, 0x91, 0x3e, 0xa9, 0xb2, 0x96,
	0xce, 0xef, 0xd2, 0x5d, 0xb9, 0x6f, 0x3f, 0xc1, 0x4c, 0xdc, 0x3c, 0x88, 0xef, 0x8f, 0x37, 0xb9,
	0x0f, 0x7d, 0xca, 0x65, 0xf5, 0x31, 0x36, 0xe1, 0xdd, 0x4c, 0x17, 0x43, 0xfb, 0x91, 0x7d, 0xf3,
	0x2b, 0x00, 0x00, 0xff, 0xff, 0x16, 0x22, 0x70, 0x3b, 0x97, 0x07, 0x00, 0x00,
}
