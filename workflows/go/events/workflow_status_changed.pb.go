// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: workflows/v1/workflow_status_changed.proto

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

type WorkflowStatusChanged struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	M             *WorkflowMetadata      `protobuf:"bytes,1,opt,name=m,proto3" json:"m,omitempty"`
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	Details       string                 `protobuf:"bytes,3,opt,name=details,proto3" json:"details,omitempty"`
	TxInfo        *TransactionInfo       `protobuf:"bytes,4,opt,name=txInfo,proto3" json:"txInfo,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *WorkflowStatusChanged) Reset() {
	*x = WorkflowStatusChanged{}
	mi := &file_workflows_v1_workflow_status_changed_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *WorkflowStatusChanged) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*WorkflowStatusChanged) ProtoMessage() {}

func (x *WorkflowStatusChanged) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_v1_workflow_status_changed_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use WorkflowStatusChanged.ProtoReflect.Descriptor instead.
func (*WorkflowStatusChanged) Descriptor() ([]byte, []int) {
	return file_workflows_v1_workflow_status_changed_proto_rawDescGZIP(), []int{0}
}

func (x *WorkflowStatusChanged) GetM() *WorkflowMetadata {
	if x != nil {
		return x.M
	}
	return nil
}

func (x *WorkflowStatusChanged) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *WorkflowStatusChanged) GetDetails() string {
	if x != nil {
		return x.Details
	}
	return ""
}

func (x *WorkflowStatusChanged) GetTxInfo() *TransactionInfo {
	if x != nil {
		return x.TxInfo
	}
	return nil
}

type TransactionInfo struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ChainSelector string                 `protobuf:"bytes,1,opt,name=chainSelector,proto3" json:"chainSelector,omitempty"`
	TxHash        string                 `protobuf:"bytes,2,opt,name=txHash,proto3" json:"txHash,omitempty"`
	GasCost       string                 `protobuf:"bytes,3,opt,name=gasCost,proto3" json:"gasCost,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransactionInfo) Reset() {
	*x = TransactionInfo{}
	mi := &file_workflows_v1_workflow_status_changed_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransactionInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionInfo) ProtoMessage() {}

func (x *TransactionInfo) ProtoReflect() protoreflect.Message {
	mi := &file_workflows_v1_workflow_status_changed_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionInfo.ProtoReflect.Descriptor instead.
func (*TransactionInfo) Descriptor() ([]byte, []int) {
	return file_workflows_v1_workflow_status_changed_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionInfo) GetChainSelector() string {
	if x != nil {
		return x.ChainSelector
	}
	return ""
}

func (x *TransactionInfo) GetTxHash() string {
	if x != nil {
		return x.TxHash
	}
	return ""
}

func (x *TransactionInfo) GetGasCost() string {
	if x != nil {
		return x.GasCost
	}
	return ""
}

var File_workflows_v1_workflow_status_changed_proto protoreflect.FileDescriptor

const file_workflows_v1_workflow_status_changed_proto_rawDesc = "" +
	"\n" +
	"*workflows/v1/workflow_status_changed.proto\x12\fworkflows.v1\x1a\x1bworkflows/v1/metadata.proto\"\xae\x01\n" +
	"\x15WorkflowStatusChanged\x12,\n" +
	"\x01m\x18\x01 \x01(\v2\x1e.workflows.v1.WorkflowMetadataR\x01m\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x12\x18\n" +
	"\adetails\x18\x03 \x01(\tR\adetails\x125\n" +
	"\x06txInfo\x18\x04 \x01(\v2\x1d.workflows.v1.TransactionInfoR\x06txInfo\"i\n" +
	"\x0fTransactionInfo\x12$\n" +
	"\rchainSelector\x18\x01 \x01(\tR\rchainSelector\x12\x16\n" +
	"\x06txHash\x18\x02 \x01(\tR\x06txHash\x12\x18\n" +
	"\agasCost\x18\x03 \x01(\tR\agasCostBBZ@github.com/smartcontractkit/chainlink-protos/workflows/go/eventsb\x06proto3"

var (
	file_workflows_v1_workflow_status_changed_proto_rawDescOnce sync.Once
	file_workflows_v1_workflow_status_changed_proto_rawDescData []byte
)

func file_workflows_v1_workflow_status_changed_proto_rawDescGZIP() []byte {
	file_workflows_v1_workflow_status_changed_proto_rawDescOnce.Do(func() {
		file_workflows_v1_workflow_status_changed_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_workflows_v1_workflow_status_changed_proto_rawDesc), len(file_workflows_v1_workflow_status_changed_proto_rawDesc)))
	})
	return file_workflows_v1_workflow_status_changed_proto_rawDescData
}

var file_workflows_v1_workflow_status_changed_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_workflows_v1_workflow_status_changed_proto_goTypes = []any{
	(*WorkflowStatusChanged)(nil), // 0: workflows.v1.WorkflowStatusChanged
	(*TransactionInfo)(nil),       // 1: workflows.v1.TransactionInfo
	(*WorkflowMetadata)(nil),      // 2: workflows.v1.WorkflowMetadata
}
var file_workflows_v1_workflow_status_changed_proto_depIdxs = []int32{
	2, // 0: workflows.v1.WorkflowStatusChanged.m:type_name -> workflows.v1.WorkflowMetadata
	1, // 1: workflows.v1.WorkflowStatusChanged.txInfo:type_name -> workflows.v1.TransactionInfo
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_workflows_v1_workflow_status_changed_proto_init() }
func file_workflows_v1_workflow_status_changed_proto_init() {
	if File_workflows_v1_workflow_status_changed_proto != nil {
		return
	}
	file_workflows_v1_metadata_proto_init()
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_workflows_v1_workflow_status_changed_proto_rawDesc), len(file_workflows_v1_workflow_status_changed_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_workflows_v1_workflow_status_changed_proto_goTypes,
		DependencyIndexes: file_workflows_v1_workflow_status_changed_proto_depIdxs,
		MessageInfos:      file_workflows_v1_workflow_status_changed_proto_msgTypes,
	}.Build()
	File_workflows_v1_workflow_status_changed_proto = out.File
	file_workflows_v1_workflow_status_changed_proto_goTypes = nil
	file_workflows_v1_workflow_status_changed_proto_depIdxs = nil
}
