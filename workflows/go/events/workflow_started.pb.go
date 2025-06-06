// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: workflows/v1/workflow_started.proto

package events

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type WorkflowExecutionStarted struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	M             *WorkflowMetadata      `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	Timestamp     string                 `protobuf:"bytes,2,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	TriggerID     string                 `protobuf:"bytes,3,opt,name=TriggerID,proto3" json:"TriggerID,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowExecutionStarted) Reset() {
	*x = WorkflowExecutionStarted{}
	mi := &file_workflows_v1_workflow_started_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowExecutionStarted) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowExecutionStarted) ProtoMessage() {}

func (x *WorkflowExecutionStarted) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_v1_workflow_started_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowExecutionStarted.ProtoReflect.Descriptor instead.
func (*WorkflowExecutionStarted) Descriptor() ([]byte, []int) {
	return file_workflows_v1_workflow_started_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowExecutionStarted) GetM() *WorkflowMetadata {
	if x != nil {
		return x.M
	}
	return nil
}

func (x *WorkflowExecutionStarted) GetTimestamp() string {
	if x != nil {
		return x.Timestamp
	}
	return ""
}

func (x *WorkflowExecutionStarted) GetTriggerID() string {
	if x != nil {
		return x.TriggerID
	}
	return ""
}

var File_workflows_v1_workflow_started_proto protoreflect.FileDescriptor

const file_workflows_v1_workflow_started_proto_rawDesc = "" +
	"\n" +
	"#workflows/v1/workflow_started.proto\x12\fworkflows.v1\x1a\x1bworkflows/v1/metadata.proto\"\x84\x01\n" +
	"\x18WorkflowExecutionStarted\x12,\n" +
	"\x01m\x18\x01 \x01(\v2\x1e.workflows.v1.WorkflowMetadataR\x01m\x12\x1c\n" +
	"\ttimestamp\x18\x02 \x01(\tR\ttimestamp\x12\x1c\n" +
	"\tTriggerID\x18\x03 \x01(\tR\tTriggerIDBBZ@github.com/smartcontractkit/chainlink-protos/workflows/go/eventsb\x06proto3"

var (
	file_workflows_v1_workflow_started_proto_rawDescOnce sync.Once
	file_workflows_v1_workflow_started_proto_rawDescData []byte
)

func file_workflows_v1_workflow_started_proto_rawDescGZIP() []byte {
	file_workflows_v1_workflow_started_proto_rawDescOnce.Do(func() {
		file_workflows_v1_workflow_started_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_workflows_v1_workflow_started_proto_rawDesc), len(file_workflows_v1_workflow_started_proto_rawDesc)))
	})
	return file_workflows_v1_workflow_started_proto_rawDescData
}

var file_workflows_v1_workflow_started_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_workflows_v1_workflow_started_proto_goTypes = []any{
	(*WorkflowExecutionStarted)(nil), // 0: workflows.v1.WorkflowExecutionStarted
	(*WorkflowMetadata)(nil),         // 1: workflows.v1.WorkflowMetadata
}
var file_workflows_v1_workflow_started_proto_depIdxs = []int32{
	1, // 0: workflows.v1.WorkflowExecutionStarted.m:type_name -> workflows.v1.WorkflowMetadata
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_workflows_v1_workflow_started_proto_init() }
func file_workflows_v1_workflow_started_proto_init() {
	if File_workflows_v1_workflow_started_proto != nil {
		return
	}
	file_workflows_v1_metadata_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_workflows_v1_workflow_started_proto_rawDesc), len(file_workflows_v1_workflow_started_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflows_v1_workflow_started_proto_goTypes,
		DependencyIndexes: file_workflows_v1_workflow_started_proto_depIdxs,
		MessageInfos:      file_workflows_v1_workflow_started_proto_msgTypes,
	}.Build()
	File_workflows_v1_workflow_started_proto = out.File
	file_workflows_v1_workflow_started_proto_goTypes = nil
	file_workflows_v1_workflow_started_proto_depIdxs = nil
}
