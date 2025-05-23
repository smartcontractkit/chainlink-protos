// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.29.3
// source: serialization/keystore.proto

package serialization

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Keystore struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	OffchainSecretKey   []byte `protobuf:"bytes,1,opt,name=offchain_secret_key,json=offchainSecretKey,proto3" json:"offchain_secret_key,omitempty"`
	EvmOnchainSecretKey []byte `protobuf:"bytes,2,opt,name=evm_onchain_secret_key,json=evmOnchainSecretKey,proto3" json:"evm_onchain_secret_key,omitempty"`
}

func (x *Keystore) Reset() {
	*x = Keystore{}
	if protoimpl.UnsafeEnabled {
		mi := &file_serialization_keystore_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Keystore) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Keystore) ProtoMessage() {}

func (x *Keystore) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_keystore_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Keystore.ProtoReflect.Descriptor instead.
func (*Keystore) Descriptor() ([]byte, []int) {
	return file_serialization_keystore_proto_rawDescGZIP(), []int{0}
}

func (x *Keystore) GetOffchainSecretKey() []byte {
	if x != nil {
		return x.OffchainSecretKey
	}
	return nil
}

func (x *Keystore) GetEvmOnchainSecretKey() []byte {
	if x != nil {
		return x.EvmOnchainSecretKey
	}
	return nil
}

var File_serialization_keystore_proto protoreflect.FileDescriptor

var file_serialization_keystore_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2f,
	0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x08,
	0x6b, 0x65, 0x79, 0x73, 0x74, 0x6f, 0x72, 0x65, 0x22, 0x6f, 0x0a, 0x08, 0x4b, 0x65, 0x79, 0x73,
	0x74, 0x6f, 0x72, 0x65, 0x12, 0x2e, 0x0a, 0x13, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x11, 0x6f, 0x66, 0x66, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x53, 0x65, 0x63, 0x72, 0x65,
	0x74, 0x4b, 0x65, 0x79, 0x12, 0x33, 0x0a, 0x16, 0x65, 0x76, 0x6d, 0x5f, 0x6f, 0x6e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x5f, 0x73, 0x65, 0x63, 0x72, 0x65, 0x74, 0x5f, 0x6b, 0x65, 0x79, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x0c, 0x52, 0x13, 0x65, 0x76, 0x6d, 0x4f, 0x6e, 0x63, 0x68, 0x61, 0x69, 0x6e,
	0x53, 0x65, 0x63, 0x72, 0x65, 0x74, 0x4b, 0x65, 0x79, 0x42, 0x12, 0x5a, 0x10, 0x2e, 0x2f, 0x3b,
	0x73, 0x65, 0x72, 0x69, 0x61, 0x6c, 0x69, 0x7a, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_serialization_keystore_proto_rawDescOnce sync.Once
	file_serialization_keystore_proto_rawDescData = file_serialization_keystore_proto_rawDesc
)

func file_serialization_keystore_proto_rawDescGZIP() []byte {
	file_serialization_keystore_proto_rawDescOnce.Do(func() {
		file_serialization_keystore_proto_rawDescData = protoimpl.X.CompressGZIP(file_serialization_keystore_proto_rawDescData)
	})
	return file_serialization_keystore_proto_rawDescData
}

var file_serialization_keystore_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_serialization_keystore_proto_goTypes = []any{
	(*Keystore)(nil), // 0: keystore.Keystore
}
var file_serialization_keystore_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_serialization_keystore_proto_init() }
func file_serialization_keystore_proto_init() {
	if File_serialization_keystore_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_serialization_keystore_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Keystore); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_serialization_keystore_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serialization_keystore_proto_goTypes,
		DependencyIndexes: file_serialization_keystore_proto_depIdxs,
		MessageInfos:      file_serialization_keystore_proto_msgTypes,
	}.Build()
	File_serialization_keystore_proto = out.File
	file_serialization_keystore_proto_rawDesc = nil
	file_serialization_keystore_proto_goTypes = nil
	file_serialization_keystore_proto_depIdxs = nil
}
