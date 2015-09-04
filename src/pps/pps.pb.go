// Code generated by protoc-gen-go.
// source: pps/pps.proto
// DO NOT EDIT!

/*
Package pps is a generated protocol buffer package.

It is generated from these files:
	pps/pps.proto

It has these top-level messages:
	Input
	Output
	Node
	DockerService
	Pipeline
	GithubPipelineSource
	PipelineSource
	PipelineRun
	PipelineRunStatus
	PipelineRunContainer
	PipelineRunLog
	PfsCommitMapping
	CreatePipelineSourceRequest
	CreatePipelineSourceResponse
	GetPipelineSourceRequest
	GetPipelineSourceResponse
	UpdatePipelineSourceRequest
	UpdatePipelineSourceResponse
	DeletePipelineSourceRequest
	ListPipelineSourcesRequest
	ListPipelineSourcesResponse
	GetPipelineRequest
	GetPipelineResponse
	CreatePipelineRunRequest
	CreatePipelineRunResponse
	StartPipelineRunRequest
	ListPipelineRunsRequest
	ListPipelineRunsResponse
	GetPipelineRunStatusRequest
	GetPipelineRunStatusResponse
	GetPipelineRunLogsRequest
	GetPipelineRunLogsResponse
*/
package pps

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"
import google_protobuf "go.pedge.io/google-protobuf"
import google_protobuf1 "go.pedge.io/google-protobuf"

import (
	context "golang.org/x/net/context"
	grpc "google.golang.org/grpc"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type PipelineRunStatusType int32

const (
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_NONE    PipelineRunStatusType = 0
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_ADDED   PipelineRunStatusType = 1
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_STARTED PipelineRunStatusType = 2
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_ERROR   PipelineRunStatusType = 3
	PipelineRunStatusType_PIPELINE_RUN_STATUS_TYPE_SUCCESS PipelineRunStatusType = 4
)

var PipelineRunStatusType_name = map[int32]string{
	0: "PIPELINE_RUN_STATUS_TYPE_NONE",
	1: "PIPELINE_RUN_STATUS_TYPE_ADDED",
	2: "PIPELINE_RUN_STATUS_TYPE_STARTED",
	3: "PIPELINE_RUN_STATUS_TYPE_ERROR",
	4: "PIPELINE_RUN_STATUS_TYPE_SUCCESS",
}
var PipelineRunStatusType_value = map[string]int32{
	"PIPELINE_RUN_STATUS_TYPE_NONE":    0,
	"PIPELINE_RUN_STATUS_TYPE_ADDED":   1,
	"PIPELINE_RUN_STATUS_TYPE_STARTED": 2,
	"PIPELINE_RUN_STATUS_TYPE_ERROR":   3,
	"PIPELINE_RUN_STATUS_TYPE_SUCCESS": 4,
}

func (x PipelineRunStatusType) String() string {
	return proto.EnumName(PipelineRunStatusType_name, int32(x))
}

type OutputStream int32

const (
	OutputStream_OUTPUT_STREAM_NONE   OutputStream = 0
	OutputStream_OUTPUT_STREAM_STDOUT OutputStream = 1
	OutputStream_OUTPUT_STREAM_STDERR OutputStream = 2
)

var OutputStream_name = map[int32]string{
	0: "OUTPUT_STREAM_NONE",
	1: "OUTPUT_STREAM_STDOUT",
	2: "OUTPUT_STREAM_STDERR",
}
var OutputStream_value = map[string]int32{
	"OUTPUT_STREAM_NONE":   0,
	"OUTPUT_STREAM_STDOUT": 1,
	"OUTPUT_STREAM_STDERR": 2,
}

func (x OutputStream) String() string {
	return proto.EnumName(OutputStream_name, int32(x))
}

type Input struct {
	Node []string          `protobuf:"bytes,1,rep,name=node" json:"node,omitempty"`
	Host map[string]string `protobuf:"bytes,2,rep,name=host" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pfs  map[string]string `protobuf:"bytes,3,rep,name=pfs" json:"pfs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Input) Reset()         { *m = Input{} }
func (m *Input) String() string { return proto.CompactTextString(m) }
func (*Input) ProtoMessage()    {}

func (m *Input) GetHost() map[string]string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Input) GetPfs() map[string]string {
	if m != nil {
		return m.Pfs
	}
	return nil
}

type Output struct {
	Host map[string]string `protobuf:"bytes,1,rep,name=host" json:"host,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	Pfs  map[string]string `protobuf:"bytes,2,rep,name=pfs" json:"pfs,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Output) Reset()         { *m = Output{} }
func (m *Output) String() string { return proto.CompactTextString(m) }
func (*Output) ProtoMessage()    {}

func (m *Output) GetHost() map[string]string {
	if m != nil {
		return m.Host
	}
	return nil
}

func (m *Output) GetPfs() map[string]string {
	if m != nil {
		return m.Pfs
	}
	return nil
}

type Node struct {
	Service string   `protobuf:"bytes,1,opt,name=service" json:"service,omitempty"`
	Input   *Input   `protobuf:"bytes,2,opt,name=input" json:"input,omitempty"`
	Output  *Output  `protobuf:"bytes,3,opt,name=output" json:"output,omitempty"`
	Run     []string `protobuf:"bytes,4,rep,name=run" json:"run,omitempty"`
}

func (m *Node) Reset()         { *m = Node{} }
func (m *Node) String() string { return proto.CompactTextString(m) }
func (*Node) ProtoMessage()    {}

func (m *Node) GetInput() *Input {
	if m != nil {
		return m.Input
	}
	return nil
}

func (m *Node) GetOutput() *Output {
	if m != nil {
		return m.Output
	}
	return nil
}

type DockerService struct {
	Image      string `protobuf:"bytes,1,opt,name=image" json:"image,omitempty"`
	Build      string `protobuf:"bytes,2,opt,name=build" json:"build,omitempty"`
	Dockerfile string `protobuf:"bytes,3,opt,name=dockerfile" json:"dockerfile,omitempty"`
}

func (m *DockerService) Reset()         { *m = DockerService{} }
func (m *DockerService) String() string { return proto.CompactTextString(m) }
func (*DockerService) ProtoMessage()    {}

type Pipeline struct {
	NameToNode          map[string]*Node          `protobuf:"bytes,1,rep,name=name_to_node" json:"name_to_node,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	NameToDockerService map[string]*DockerService `protobuf:"bytes,2,rep,name=name_to_docker_service" json:"name_to_docker_service,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *Pipeline) Reset()         { *m = Pipeline{} }
func (m *Pipeline) String() string { return proto.CompactTextString(m) }
func (*Pipeline) ProtoMessage()    {}

func (m *Pipeline) GetNameToNode() map[string]*Node {
	if m != nil {
		return m.NameToNode
	}
	return nil
}

func (m *Pipeline) GetNameToDockerService() map[string]*DockerService {
	if m != nil {
		return m.NameToDockerService
	}
	return nil
}

type GithubPipelineSource struct {
	ContextDir  string `protobuf:"bytes,1,opt,name=context_dir" json:"context_dir,omitempty"`
	User        string `protobuf:"bytes,2,opt,name=user" json:"user,omitempty"`
	Repository  string `protobuf:"bytes,3,opt,name=repository" json:"repository,omitempty"`
	Branch      string `protobuf:"bytes,4,opt,name=branch" json:"branch,omitempty"`
	CommitId    string `protobuf:"bytes,5,opt,name=commit_id" json:"commit_id,omitempty"`
	AccessToken string `protobuf:"bytes,6,opt,name=access_token" json:"access_token,omitempty"`
}

func (m *GithubPipelineSource) Reset()         { *m = GithubPipelineSource{} }
func (m *GithubPipelineSource) String() string { return proto.CompactTextString(m) }
func (*GithubPipelineSource) ProtoMessage()    {}

type PipelineSource struct {
	Id     string            `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	Active bool              `protobuf:"varint,2,opt,name=active" json:"active,omitempty"`
	Tags   map[string]string `protobuf:"bytes,3,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	// Types that are valid to be assigned to TypedPipelineSource:
	//	*PipelineSource_GithubPipelineSource
	TypedPipelineSource isPipelineSource_TypedPipelineSource `protobuf_oneof:"typed_pipeline_source"`
}

func (m *PipelineSource) Reset()         { *m = PipelineSource{} }
func (m *PipelineSource) String() string { return proto.CompactTextString(m) }
func (*PipelineSource) ProtoMessage()    {}

type isPipelineSource_TypedPipelineSource interface {
	isPipelineSource_TypedPipelineSource()
}

type PipelineSource_GithubPipelineSource struct {
	GithubPipelineSource *GithubPipelineSource `protobuf:"bytes,4,opt,name=github_pipeline_source"`
}

func (*PipelineSource_GithubPipelineSource) isPipelineSource_TypedPipelineSource() {}

func (m *PipelineSource) GetTypedPipelineSource() isPipelineSource_TypedPipelineSource {
	if m != nil {
		return m.TypedPipelineSource
	}
	return nil
}

func (m *PipelineSource) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

func (m *PipelineSource) GetGithubPipelineSource() *GithubPipelineSource {
	if x, ok := m.GetTypedPipelineSource().(*PipelineSource_GithubPipelineSource); ok {
		return x.GithubPipelineSource
	}
	return nil
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*PipelineSource) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), []interface{}) {
	return _PipelineSource_OneofMarshaler, _PipelineSource_OneofUnmarshaler, []interface{}{
		(*PipelineSource_GithubPipelineSource)(nil),
	}
}

func _PipelineSource_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*PipelineSource)
	// typed_pipeline_source
	switch x := m.TypedPipelineSource.(type) {
	case *PipelineSource_GithubPipelineSource:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.GithubPipelineSource); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("PipelineSource.TypedPipelineSource has unexpected type %T", x)
	}
	return nil
}

func _PipelineSource_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*PipelineSource)
	switch tag {
	case 4: // typed_pipeline_source.github_pipeline_source
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GithubPipelineSource)
		err := b.DecodeMessage(msg)
		m.TypedPipelineSource = &PipelineSource_GithubPipelineSource{msg}
		return true, err
	default:
		return false, nil
	}
}

type PipelineRun struct {
	Id             string          `protobuf:"bytes,1,opt,name=id" json:"id,omitempty"`
	PipelineSource *PipelineSource `protobuf:"bytes,2,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
	Pipeline       *Pipeline       `protobuf:"bytes,3,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *PipelineRun) Reset()         { *m = PipelineRun{} }
func (m *PipelineRun) String() string { return proto.CompactTextString(m) }
func (*PipelineRun) ProtoMessage()    {}

func (m *PipelineRun) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

func (m *PipelineRun) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type PipelineRunStatus struct {
	PipelineRunId         string                      `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	PipelineRunStatusType PipelineRunStatusType       `protobuf:"varint,2,opt,name=pipeline_run_status_type,enum=pps.PipelineRunStatusType" json:"pipeline_run_status_type,omitempty"`
	Timestamp             *google_protobuf1.Timestamp `protobuf:"bytes,3,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PipelineRunStatus) Reset()         { *m = PipelineRunStatus{} }
func (m *PipelineRunStatus) String() string { return proto.CompactTextString(m) }
func (*PipelineRunStatus) ProtoMessage()    {}

func (m *PipelineRunStatus) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PipelineRunContainer struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	ContainerId   string `protobuf:"bytes,2,opt,name=container_id" json:"container_id,omitempty"`
	Node          string `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
}

func (m *PipelineRunContainer) Reset()         { *m = PipelineRunContainer{} }
func (m *PipelineRunContainer) String() string { return proto.CompactTextString(m) }
func (*PipelineRunContainer) ProtoMessage()    {}

type PipelineRunLog struct {
	PipelineRunId string                      `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	ContainerId   string                      `protobuf:"bytes,2,opt,name=container_id" json:"container_id,omitempty"`
	Node          string                      `protobuf:"bytes,3,opt,name=node" json:"node,omitempty"`
	Timestamp     *google_protobuf1.Timestamp `protobuf:"bytes,4,opt,name=timestamp" json:"timestamp,omitempty"`
	OutputStream  OutputStream                `protobuf:"varint,5,opt,name=output_stream,enum=pps.OutputStream" json:"output_stream,omitempty"`
	Data          []byte                      `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
}

func (m *PipelineRunLog) Reset()         { *m = PipelineRunLog{} }
func (m *PipelineRunLog) String() string { return proto.CompactTextString(m) }
func (*PipelineRunLog) ProtoMessage()    {}

func (m *PipelineRunLog) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type PfsCommitMapping struct {
	InputRepository  string                      `protobuf:"bytes,1,opt,name=input_repository" json:"input_repository,omitempty"`
	InputCommitId    string                      `protobuf:"bytes,2,opt,name=input_commit_id" json:"input_commit_id,omitempty"`
	OutputRepository string                      `protobuf:"bytes,3,opt,name=output_repository" json:"output_repository,omitempty"`
	OutputCommitId   string                      `protobuf:"bytes,4,opt,name=output_commit_id" json:"output_commit_id,omitempty"`
	Timestamp        *google_protobuf1.Timestamp `protobuf:"bytes,5,opt,name=timestamp" json:"timestamp,omitempty"`
}

func (m *PfsCommitMapping) Reset()         { *m = PfsCommitMapping{} }
func (m *PfsCommitMapping) String() string { return proto.CompactTextString(m) }
func (*PfsCommitMapping) ProtoMessage()    {}

func (m *PfsCommitMapping) GetTimestamp() *google_protobuf1.Timestamp {
	if m != nil {
		return m.Timestamp
	}
	return nil
}

type CreatePipelineSourceRequest struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *CreatePipelineSourceRequest) Reset()         { *m = CreatePipelineSourceRequest{} }
func (m *CreatePipelineSourceRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineSourceRequest) ProtoMessage()    {}

func (m *CreatePipelineSourceRequest) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type CreatePipelineSourceResponse struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *CreatePipelineSourceResponse) Reset()         { *m = CreatePipelineSourceResponse{} }
func (m *CreatePipelineSourceResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineSourceResponse) ProtoMessage()    {}

func (m *CreatePipelineSourceResponse) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type GetPipelineSourceRequest struct {
	PipelineSourceId string `protobuf:"bytes,1,opt,name=pipeline_source_id" json:"pipeline_source_id,omitempty"`
}

func (m *GetPipelineSourceRequest) Reset()         { *m = GetPipelineSourceRequest{} }
func (m *GetPipelineSourceRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineSourceRequest) ProtoMessage()    {}

type GetPipelineSourceResponse struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *GetPipelineSourceResponse) Reset()         { *m = GetPipelineSourceResponse{} }
func (m *GetPipelineSourceResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineSourceResponse) ProtoMessage()    {}

func (m *GetPipelineSourceResponse) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type UpdatePipelineSourceRequest struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *UpdatePipelineSourceRequest) Reset()         { *m = UpdatePipelineSourceRequest{} }
func (m *UpdatePipelineSourceRequest) String() string { return proto.CompactTextString(m) }
func (*UpdatePipelineSourceRequest) ProtoMessage()    {}

func (m *UpdatePipelineSourceRequest) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type UpdatePipelineSourceResponse struct {
	PipelineSource *PipelineSource `protobuf:"bytes,1,opt,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *UpdatePipelineSourceResponse) Reset()         { *m = UpdatePipelineSourceResponse{} }
func (m *UpdatePipelineSourceResponse) String() string { return proto.CompactTextString(m) }
func (*UpdatePipelineSourceResponse) ProtoMessage()    {}

func (m *UpdatePipelineSourceResponse) GetPipelineSource() *PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type DeletePipelineSourceRequest struct {
	PipelineSourceId string `protobuf:"bytes,1,opt,name=pipeline_source_id" json:"pipeline_source_id,omitempty"`
}

func (m *DeletePipelineSourceRequest) Reset()         { *m = DeletePipelineSourceRequest{} }
func (m *DeletePipelineSourceRequest) String() string { return proto.CompactTextString(m) }
func (*DeletePipelineSourceRequest) ProtoMessage()    {}

type ListPipelineSourcesRequest struct {
	Tags map[string]string `protobuf:"bytes,1,rep,name=tags" json:"tags,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
}

func (m *ListPipelineSourcesRequest) Reset()         { *m = ListPipelineSourcesRequest{} }
func (m *ListPipelineSourcesRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelineSourcesRequest) ProtoMessage()    {}

func (m *ListPipelineSourcesRequest) GetTags() map[string]string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type ListPipelineSourcesResponse struct {
	PipelineSource []*PipelineSource `protobuf:"bytes,1,rep,name=pipeline_source" json:"pipeline_source,omitempty"`
}

func (m *ListPipelineSourcesResponse) Reset()         { *m = ListPipelineSourcesResponse{} }
func (m *ListPipelineSourcesResponse) String() string { return proto.CompactTextString(m) }
func (*ListPipelineSourcesResponse) ProtoMessage()    {}

func (m *ListPipelineSourcesResponse) GetPipelineSource() []*PipelineSource {
	if m != nil {
		return m.PipelineSource
	}
	return nil
}

type GetPipelineRequest struct {
	PipelineSourceId string `protobuf:"bytes,1,opt,name=pipeline_source_id" json:"pipeline_source_id,omitempty"`
}

func (m *GetPipelineRequest) Reset()         { *m = GetPipelineRequest{} }
func (m *GetPipelineRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRequest) ProtoMessage()    {}

type GetPipelineResponse struct {
	Pipeline *Pipeline `protobuf:"bytes,1,opt,name=pipeline" json:"pipeline,omitempty"`
}

func (m *GetPipelineResponse) Reset()         { *m = GetPipelineResponse{} }
func (m *GetPipelineResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineResponse) ProtoMessage()    {}

func (m *GetPipelineResponse) GetPipeline() *Pipeline {
	if m != nil {
		return m.Pipeline
	}
	return nil
}

type CreatePipelineRunRequest struct {
	PipelineSourceId string `protobuf:"bytes,1,opt,name=pipeline_source_id" json:"pipeline_source_id,omitempty"`
}

func (m *CreatePipelineRunRequest) Reset()         { *m = CreatePipelineRunRequest{} }
func (m *CreatePipelineRunRequest) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineRunRequest) ProtoMessage()    {}

type CreatePipelineRunResponse struct {
	PipelineRun *PipelineRun `protobuf:"bytes,1,opt,name=pipeline_run" json:"pipeline_run,omitempty"`
}

func (m *CreatePipelineRunResponse) Reset()         { *m = CreatePipelineRunResponse{} }
func (m *CreatePipelineRunResponse) String() string { return proto.CompactTextString(m) }
func (*CreatePipelineRunResponse) ProtoMessage()    {}

func (m *CreatePipelineRunResponse) GetPipelineRun() *PipelineRun {
	if m != nil {
		return m.PipelineRun
	}
	return nil
}

type StartPipelineRunRequest struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
}

func (m *StartPipelineRunRequest) Reset()         { *m = StartPipelineRunRequest{} }
func (m *StartPipelineRunRequest) String() string { return proto.CompactTextString(m) }
func (*StartPipelineRunRequest) ProtoMessage()    {}

type ListPipelineRunsRequest struct {
	PipelineSourceId string `protobuf:"bytes,1,opt,name=pipeline_source_id" json:"pipeline_source_id,omitempty"`
}

func (m *ListPipelineRunsRequest) Reset()         { *m = ListPipelineRunsRequest{} }
func (m *ListPipelineRunsRequest) String() string { return proto.CompactTextString(m) }
func (*ListPipelineRunsRequest) ProtoMessage()    {}

type ListPipelineRunsResponse struct {
	PipelineRun []*PipelineRun `protobuf:"bytes,1,rep,name=pipeline_run" json:"pipeline_run,omitempty"`
}

func (m *ListPipelineRunsResponse) Reset()         { *m = ListPipelineRunsResponse{} }
func (m *ListPipelineRunsResponse) String() string { return proto.CompactTextString(m) }
func (*ListPipelineRunsResponse) ProtoMessage()    {}

func (m *ListPipelineRunsResponse) GetPipelineRun() []*PipelineRun {
	if m != nil {
		return m.PipelineRun
	}
	return nil
}

type GetPipelineRunStatusRequest struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	All           bool   `protobuf:"varint,2,opt,name=all" json:"all,omitempty"`
}

func (m *GetPipelineRunStatusRequest) Reset()         { *m = GetPipelineRunStatusRequest{} }
func (m *GetPipelineRunStatusRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunStatusRequest) ProtoMessage()    {}

type GetPipelineRunStatusResponse struct {
	PipelineRunStatus []*PipelineRunStatus `protobuf:"bytes,1,rep,name=pipeline_run_status" json:"pipeline_run_status,omitempty"`
}

func (m *GetPipelineRunStatusResponse) Reset()         { *m = GetPipelineRunStatusResponse{} }
func (m *GetPipelineRunStatusResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunStatusResponse) ProtoMessage()    {}

func (m *GetPipelineRunStatusResponse) GetPipelineRunStatus() []*PipelineRunStatus {
	if m != nil {
		return m.PipelineRunStatus
	}
	return nil
}

type GetPipelineRunLogsRequest struct {
	PipelineRunId string `protobuf:"bytes,1,opt,name=pipeline_run_id" json:"pipeline_run_id,omitempty"`
	Node          string `protobuf:"bytes,2,opt,name=node" json:"node,omitempty"`
}

func (m *GetPipelineRunLogsRequest) Reset()         { *m = GetPipelineRunLogsRequest{} }
func (m *GetPipelineRunLogsRequest) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunLogsRequest) ProtoMessage()    {}

type GetPipelineRunLogsResponse struct {
	PipelineRunLog []*PipelineRunLog `protobuf:"bytes,1,rep,name=pipeline_run_log" json:"pipeline_run_log,omitempty"`
}

func (m *GetPipelineRunLogsResponse) Reset()         { *m = GetPipelineRunLogsResponse{} }
func (m *GetPipelineRunLogsResponse) String() string { return proto.CompactTextString(m) }
func (*GetPipelineRunLogsResponse) ProtoMessage()    {}

func (m *GetPipelineRunLogsResponse) GetPipelineRunLog() []*PipelineRunLog {
	if m != nil {
		return m.PipelineRunLog
	}
	return nil
}

func init() {
	proto.RegisterEnum("pps.PipelineRunStatusType", PipelineRunStatusType_name, PipelineRunStatusType_value)
	proto.RegisterEnum("pps.OutputStream", OutputStream_name, OutputStream_value)
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConn

// Client API for Api service

type ApiClient interface {
	CreatePipelineSource(ctx context.Context, in *CreatePipelineSourceRequest, opts ...grpc.CallOption) (*CreatePipelineSourceResponse, error)
	GetPipelineSource(ctx context.Context, in *GetPipelineSourceRequest, opts ...grpc.CallOption) (*GetPipelineSourceResponse, error)
	UpdatePipelineSource(ctx context.Context, in *UpdatePipelineSourceRequest, opts ...grpc.CallOption) (*UpdatePipelineSourceResponse, error)
	DeletePipelineSource(ctx context.Context, in *DeletePipelineSourceRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ListPipelineSources(ctx context.Context, in *ListPipelineSourcesRequest, opts ...grpc.CallOption) (*ListPipelineSourcesResponse, error)
	GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error)
	CreatePipelineRun(ctx context.Context, in *CreatePipelineRunRequest, opts ...grpc.CallOption) (*CreatePipelineRunResponse, error)
	StartPipelineRun(ctx context.Context, in *StartPipelineRunRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error)
	ListPipelineRuns(ctx context.Context, in *ListPipelineRunsRequest, opts ...grpc.CallOption) (*ListPipelineRunsResponse, error)
	GetPipelineRunStatus(ctx context.Context, in *GetPipelineRunStatusRequest, opts ...grpc.CallOption) (*GetPipelineRunStatusResponse, error)
	GetPipelineRunLogs(ctx context.Context, in *GetPipelineRunLogsRequest, opts ...grpc.CallOption) (*GetPipelineRunLogsResponse, error)
}

type apiClient struct {
	cc *grpc.ClientConn
}

func NewApiClient(cc *grpc.ClientConn) ApiClient {
	return &apiClient{cc}
}

func (c *apiClient) CreatePipelineSource(ctx context.Context, in *CreatePipelineSourceRequest, opts ...grpc.CallOption) (*CreatePipelineSourceResponse, error) {
	out := new(CreatePipelineSourceResponse)
	err := grpc.Invoke(ctx, "/pps.Api/CreatePipelineSource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPipelineSource(ctx context.Context, in *GetPipelineSourceRequest, opts ...grpc.CallOption) (*GetPipelineSourceResponse, error) {
	out := new(GetPipelineSourceResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipelineSource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) UpdatePipelineSource(ctx context.Context, in *UpdatePipelineSourceRequest, opts ...grpc.CallOption) (*UpdatePipelineSourceResponse, error) {
	out := new(UpdatePipelineSourceResponse)
	err := grpc.Invoke(ctx, "/pps.Api/UpdatePipelineSource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) DeletePipelineSource(ctx context.Context, in *DeletePipelineSourceRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pps.Api/DeletePipelineSource", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListPipelineSources(ctx context.Context, in *ListPipelineSourcesRequest, opts ...grpc.CallOption) (*ListPipelineSourcesResponse, error) {
	out := new(ListPipelineSourcesResponse)
	err := grpc.Invoke(ctx, "/pps.Api/ListPipelineSources", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPipeline(ctx context.Context, in *GetPipelineRequest, opts ...grpc.CallOption) (*GetPipelineResponse, error) {
	out := new(GetPipelineResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipeline", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) CreatePipelineRun(ctx context.Context, in *CreatePipelineRunRequest, opts ...grpc.CallOption) (*CreatePipelineRunResponse, error) {
	out := new(CreatePipelineRunResponse)
	err := grpc.Invoke(ctx, "/pps.Api/CreatePipelineRun", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) StartPipelineRun(ctx context.Context, in *StartPipelineRunRequest, opts ...grpc.CallOption) (*google_protobuf.Empty, error) {
	out := new(google_protobuf.Empty)
	err := grpc.Invoke(ctx, "/pps.Api/StartPipelineRun", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) ListPipelineRuns(ctx context.Context, in *ListPipelineRunsRequest, opts ...grpc.CallOption) (*ListPipelineRunsResponse, error) {
	out := new(ListPipelineRunsResponse)
	err := grpc.Invoke(ctx, "/pps.Api/ListPipelineRuns", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPipelineRunStatus(ctx context.Context, in *GetPipelineRunStatusRequest, opts ...grpc.CallOption) (*GetPipelineRunStatusResponse, error) {
	out := new(GetPipelineRunStatusResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipelineRunStatus", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *apiClient) GetPipelineRunLogs(ctx context.Context, in *GetPipelineRunLogsRequest, opts ...grpc.CallOption) (*GetPipelineRunLogsResponse, error) {
	out := new(GetPipelineRunLogsResponse)
	err := grpc.Invoke(ctx, "/pps.Api/GetPipelineRunLogs", in, out, c.cc, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Server API for Api service

type ApiServer interface {
	CreatePipelineSource(context.Context, *CreatePipelineSourceRequest) (*CreatePipelineSourceResponse, error)
	GetPipelineSource(context.Context, *GetPipelineSourceRequest) (*GetPipelineSourceResponse, error)
	UpdatePipelineSource(context.Context, *UpdatePipelineSourceRequest) (*UpdatePipelineSourceResponse, error)
	DeletePipelineSource(context.Context, *DeletePipelineSourceRequest) (*google_protobuf.Empty, error)
	ListPipelineSources(context.Context, *ListPipelineSourcesRequest) (*ListPipelineSourcesResponse, error)
	GetPipeline(context.Context, *GetPipelineRequest) (*GetPipelineResponse, error)
	CreatePipelineRun(context.Context, *CreatePipelineRunRequest) (*CreatePipelineRunResponse, error)
	StartPipelineRun(context.Context, *StartPipelineRunRequest) (*google_protobuf.Empty, error)
	ListPipelineRuns(context.Context, *ListPipelineRunsRequest) (*ListPipelineRunsResponse, error)
	GetPipelineRunStatus(context.Context, *GetPipelineRunStatusRequest) (*GetPipelineRunStatusResponse, error)
	GetPipelineRunLogs(context.Context, *GetPipelineRunLogsRequest) (*GetPipelineRunLogsResponse, error)
}

func RegisterApiServer(s *grpc.Server, srv ApiServer) {
	s.RegisterService(&_Api_serviceDesc, srv)
}

func _Api_CreatePipelineSource_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(CreatePipelineSourceRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).CreatePipelineSource(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetPipelineSource_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineSourceRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipelineSource(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_UpdatePipelineSource_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(UpdatePipelineSourceRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).UpdatePipelineSource(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_DeletePipelineSource_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(DeletePipelineSourceRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).DeletePipelineSource(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_ListPipelineSources_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ListPipelineSourcesRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).ListPipelineSources(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetPipeline_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipeline(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_CreatePipelineRun_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(CreatePipelineRunRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).CreatePipelineRun(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_StartPipelineRun_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(StartPipelineRunRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).StartPipelineRun(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_ListPipelineRuns_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(ListPipelineRunsRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).ListPipelineRuns(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetPipelineRunStatus_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineRunStatusRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipelineRunStatus(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func _Api_GetPipelineRunLogs_Handler(srv interface{}, ctx context.Context, codec grpc.Codec, buf []byte) (interface{}, error) {
	in := new(GetPipelineRunLogsRequest)
	if err := codec.Unmarshal(buf, in); err != nil {
		return nil, err
	}
	out, err := srv.(ApiServer).GetPipelineRunLogs(ctx, in)
	if err != nil {
		return nil, err
	}
	return out, nil
}

var _Api_serviceDesc = grpc.ServiceDesc{
	ServiceName: "pps.Api",
	HandlerType: (*ApiServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "CreatePipelineSource",
			Handler:    _Api_CreatePipelineSource_Handler,
		},
		{
			MethodName: "GetPipelineSource",
			Handler:    _Api_GetPipelineSource_Handler,
		},
		{
			MethodName: "UpdatePipelineSource",
			Handler:    _Api_UpdatePipelineSource_Handler,
		},
		{
			MethodName: "DeletePipelineSource",
			Handler:    _Api_DeletePipelineSource_Handler,
		},
		{
			MethodName: "ListPipelineSources",
			Handler:    _Api_ListPipelineSources_Handler,
		},
		{
			MethodName: "GetPipeline",
			Handler:    _Api_GetPipeline_Handler,
		},
		{
			MethodName: "CreatePipelineRun",
			Handler:    _Api_CreatePipelineRun_Handler,
		},
		{
			MethodName: "StartPipelineRun",
			Handler:    _Api_StartPipelineRun_Handler,
		},
		{
			MethodName: "ListPipelineRuns",
			Handler:    _Api_ListPipelineRuns_Handler,
		},
		{
			MethodName: "GetPipelineRunStatus",
			Handler:    _Api_GetPipelineRunStatus_Handler,
		},
		{
			MethodName: "GetPipelineRunLogs",
			Handler:    _Api_GetPipelineRunLogs_Handler,
		},
	},
	Streams: []grpc.StreamDesc{},
}
