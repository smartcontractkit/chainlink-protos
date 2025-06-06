// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: serialization/rmn_offchain.proto

package serialization

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

type LaneSource struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	SourceChainSelector uint64                 `protobuf:"varint,1,opt,name=source_chain_selector,json=sourceChainSelector,proto3" json:"source_chain_selector,omitempty"`
	OnrampAddress       []byte                 `protobuf:"bytes,2,opt,name=onramp_address,json=onrampAddress,proto3" json:"onramp_address,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *LaneSource) Reset() {
	*x = LaneSource{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LaneSource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneSource) ProtoMessage() {}

func (x *LaneSource) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneSource.ProtoReflect.Descriptor instead.
func (*LaneSource) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{0}
}

func (x *LaneSource) GetSourceChainSelector() uint64 {
	if x != nil {
		return x.SourceChainSelector
	}
	return 0
}

func (x *LaneSource) GetOnrampAddress() []byte {
	if x != nil {
		return x.OnrampAddress
	}
	return nil
}

type LaneDest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	DestChainSelector uint64                 `protobuf:"varint,1,opt,name=dest_chain_selector,json=destChainSelector,proto3" json:"dest_chain_selector,omitempty"`
	OfframpAddress    []byte                 `protobuf:"bytes,2,opt,name=offramp_address,json=offrampAddress,proto3" json:"offramp_address,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *LaneDest) Reset() {
	*x = LaneDest{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *LaneDest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LaneDest) ProtoMessage() {}

func (x *LaneDest) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LaneDest.ProtoReflect.Descriptor instead.
func (*LaneDest) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{1}
}

func (x *LaneDest) GetDestChainSelector() uint64 {
	if x != nil {
		return x.DestChainSelector
	}
	return 0
}

func (x *LaneDest) GetOfframpAddress() []byte {
	if x != nil {
		return x.OfframpAddress
	}
	return nil
}

type ClosedInterval struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	MinMsgNr      uint64                 `protobuf:"varint,1,opt,name=min_msg_nr,json=minMsgNr,proto3" json:"min_msg_nr,omitempty"`
	MaxMsgNr      uint64                 `protobuf:"varint,2,opt,name=max_msg_nr,json=maxMsgNr,proto3" json:"max_msg_nr,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ClosedInterval) Reset() {
	*x = ClosedInterval{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ClosedInterval) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ClosedInterval) ProtoMessage() {}

func (x *ClosedInterval) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ClosedInterval.ProtoReflect.Descriptor instead.
func (*ClosedInterval) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{2}
}

func (x *ClosedInterval) GetMinMsgNr() uint64 {
	if x != nil {
		return x.MinMsgNr
	}
	return 0
}

func (x *ClosedInterval) GetMaxMsgNr() uint64 {
	if x != nil {
		return x.MaxMsgNr
	}
	return 0
}

type FixedDestLaneUpdateRequest struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	LaneSource     *LaneSource            `protobuf:"bytes,1,opt,name=lane_source,json=laneSource,proto3" json:"lane_source,omitempty"`
	ClosedInterval *ClosedInterval        `protobuf:"bytes,2,opt,name=closed_interval,json=closedInterval,proto3" json:"closed_interval,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FixedDestLaneUpdateRequest) Reset() {
	*x = FixedDestLaneUpdateRequest{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FixedDestLaneUpdateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedDestLaneUpdateRequest) ProtoMessage() {}

func (x *FixedDestLaneUpdateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedDestLaneUpdateRequest.ProtoReflect.Descriptor instead.
func (*FixedDestLaneUpdateRequest) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{3}
}

func (x *FixedDestLaneUpdateRequest) GetLaneSource() *LaneSource {
	if x != nil {
		return x.LaneSource
	}
	return nil
}

func (x *FixedDestLaneUpdateRequest) GetClosedInterval() *ClosedInterval {
	if x != nil {
		return x.ClosedInterval
	}
	return nil
}

type FixedDestLaneUpdate struct {
	state          protoimpl.MessageState `protogen:"open.v1"`
	LaneSource     *LaneSource            `protobuf:"bytes,1,opt,name=lane_source,json=laneSource,proto3" json:"lane_source,omitempty"`
	ClosedInterval *ClosedInterval        `protobuf:"bytes,2,opt,name=closed_interval,json=closedInterval,proto3" json:"closed_interval,omitempty"`
	Root           []byte                 `protobuf:"bytes,3,opt,name=root,proto3" json:"root,omitempty"`
	unknownFields  protoimpl.UnknownFields
	sizeCache      protoimpl.SizeCache
}

func (x *FixedDestLaneUpdate) Reset() {
	*x = FixedDestLaneUpdate{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *FixedDestLaneUpdate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*FixedDestLaneUpdate) ProtoMessage() {}

func (x *FixedDestLaneUpdate) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use FixedDestLaneUpdate.ProtoReflect.Descriptor instead.
func (*FixedDestLaneUpdate) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{4}
}

func (x *FixedDestLaneUpdate) GetLaneSource() *LaneSource {
	if x != nil {
		return x.LaneSource
	}
	return nil
}

func (x *FixedDestLaneUpdate) GetClosedInterval() *ClosedInterval {
	if x != nil {
		return x.ClosedInterval
	}
	return nil
}

func (x *FixedDestLaneUpdate) GetRoot() []byte {
	if x != nil {
		return x.Root
	}
	return nil
}

type ObservationRequest struct {
	state                       protoimpl.MessageState        `protogen:"open.v1"`
	LaneDest                    *LaneDest                     `protobuf:"bytes,1,opt,name=lane_dest,json=laneDest,proto3" json:"lane_dest,omitempty"` // could be implied
	FixedDestLaneUpdateRequests []*FixedDestLaneUpdateRequest `protobuf:"bytes,2,rep,name=fixed_dest_lane_update_requests,json=fixedDestLaneUpdateRequests,proto3" json:"fixed_dest_lane_update_requests,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *ObservationRequest) Reset() {
	*x = ObservationRequest{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ObservationRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ObservationRequest) ProtoMessage() {}

func (x *ObservationRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ObservationRequest.ProtoReflect.Descriptor instead.
func (*ObservationRequest) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{5}
}

func (x *ObservationRequest) GetLaneDest() *LaneDest {
	if x != nil {
		return x.LaneDest
	}
	return nil
}

func (x *ObservationRequest) GetFixedDestLaneUpdateRequests() []*FixedDestLaneUpdateRequest {
	if x != nil {
		return x.FixedDestLaneUpdateRequests
	}
	return nil
}

// TODO: For terseness, we might want to split this into two messages down the line:
// An observation containing only the things that cannot be inferred
// An observation representing the exact message that is signed by the RMN node
type Observation struct {
	state                       protoimpl.MessageState `protogen:"open.v1"`
	RmnHomeContractConfigDigest []byte                 `protobuf:"bytes,1,opt,name=rmn_home_contract_config_digest,json=rmnHomeContractConfigDigest,proto3" json:"rmn_home_contract_config_digest,omitempty"` // could be implied
	LaneDest                    *LaneDest              `protobuf:"bytes,2,opt,name=lane_dest,json=laneDest,proto3" json:"lane_dest,omitempty"`                                                                // could be implied
	FixedDestLaneUpdates        []*FixedDestLaneUpdate `protobuf:"bytes,3,rep,name=fixed_dest_lane_updates,json=fixedDestLaneUpdates,proto3" json:"fixed_dest_lane_updates,omitempty"`
	Timestamp                   uint64                 `protobuf:"varint,4,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *Observation) Reset() {
	*x = Observation{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Observation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Observation) ProtoMessage() {}

func (x *Observation) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Observation.ProtoReflect.Descriptor instead.
func (*Observation) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{6}
}

func (x *Observation) GetRmnHomeContractConfigDigest() []byte {
	if x != nil {
		return x.RmnHomeContractConfigDigest
	}
	return nil
}

func (x *Observation) GetLaneDest() *LaneDest {
	if x != nil {
		return x.LaneDest
	}
	return nil
}

func (x *Observation) GetFixedDestLaneUpdates() []*FixedDestLaneUpdate {
	if x != nil {
		return x.FixedDestLaneUpdates
	}
	return nil
}

func (x *Observation) GetTimestamp() uint64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

type SignedObservation struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Observation   *Observation           `protobuf:"bytes,1,opt,name=observation,proto3" json:"observation,omitempty"`
	Signature     []byte                 `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"` // sign(sha256("chainlink ccip 1.6 rmn observation"|sha256(observation)))
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *SignedObservation) Reset() {
	*x = SignedObservation{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *SignedObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SignedObservation) ProtoMessage() {}

func (x *SignedObservation) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SignedObservation.ProtoReflect.Descriptor instead.
func (*SignedObservation) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{7}
}

func (x *SignedObservation) GetObservation() *Observation {
	if x != nil {
		return x.Observation
	}
	return nil
}

func (x *SignedObservation) GetSignature() []byte {
	if x != nil {
		return x.Signature
	}
	return nil
}

type AttributedSignedObservation struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	SignedObservation *SignedObservation     `protobuf:"bytes,1,opt,name=signed_observation,json=signedObservation,proto3" json:"signed_observation,omitempty"`
	SignerNodeIndex   uint32                 `protobuf:"varint,2,opt,name=signer_node_index,json=signerNodeIndex,proto3" json:"signer_node_index,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *AttributedSignedObservation) Reset() {
	*x = AttributedSignedObservation{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AttributedSignedObservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AttributedSignedObservation) ProtoMessage() {}

func (x *AttributedSignedObservation) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AttributedSignedObservation.ProtoReflect.Descriptor instead.
func (*AttributedSignedObservation) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{8}
}

func (x *AttributedSignedObservation) GetSignedObservation() *SignedObservation {
	if x != nil {
		return x.SignedObservation
	}
	return nil
}

func (x *AttributedSignedObservation) GetSignerNodeIndex() uint32 {
	if x != nil {
		return x.SignerNodeIndex
	}
	return 0
}

// Signed along with the report
type ReportContext struct {
	state                       protoimpl.MessageState `protogen:"open.v1"`
	EvmDestChainId              uint64                 `protobuf:"varint,1,opt,name=evm_dest_chain_id,json=evmDestChainId,proto3" json:"evm_dest_chain_id,omitempty"`
	RmnRemoteContractAddress    []byte                 `protobuf:"bytes,2,opt,name=rmn_remote_contract_address,json=rmnRemoteContractAddress,proto3" json:"rmn_remote_contract_address,omitempty"`
	RmnHomeContractConfigDigest []byte                 `protobuf:"bytes,3,opt,name=rmn_home_contract_config_digest,json=rmnHomeContractConfigDigest,proto3" json:"rmn_home_contract_config_digest,omitempty"` // can lag behind home chain to support blue/green
	LaneDest                    *LaneDest              `protobuf:"bytes,4,opt,name=lane_dest,json=laneDest,proto3" json:"lane_dest,omitempty"`
	unknownFields               protoimpl.UnknownFields
	sizeCache                   protoimpl.SizeCache
}

func (x *ReportContext) Reset() {
	*x = ReportContext{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportContext) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportContext) ProtoMessage() {}

func (x *ReportContext) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportContext.ProtoReflect.Descriptor instead.
func (*ReportContext) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{9}
}

func (x *ReportContext) GetEvmDestChainId() uint64 {
	if x != nil {
		return x.EvmDestChainId
	}
	return 0
}

func (x *ReportContext) GetRmnRemoteContractAddress() []byte {
	if x != nil {
		return x.RmnRemoteContractAddress
	}
	return nil
}

func (x *ReportContext) GetRmnHomeContractConfigDigest() []byte {
	if x != nil {
		return x.RmnHomeContractConfigDigest
	}
	return nil
}

func (x *ReportContext) GetLaneDest() *LaneDest {
	if x != nil {
		return x.LaneDest
	}
	return nil
}

type ReportSignatureRequest struct {
	state                        protoimpl.MessageState         `protogen:"open.v1"`
	Context                      *ReportContext                 `protobuf:"bytes,1,opt,name=context,proto3" json:"context,omitempty"`
	AttributedSignedObservations []*AttributedSignedObservation `protobuf:"bytes,2,rep,name=attributed_signed_observations,json=attributedSignedObservations,proto3" json:"attributed_signed_observations,omitempty"`
	unknownFields                protoimpl.UnknownFields
	sizeCache                    protoimpl.SizeCache
}

func (x *ReportSignatureRequest) Reset() {
	*x = ReportSignatureRequest{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportSignatureRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSignatureRequest) ProtoMessage() {}

func (x *ReportSignatureRequest) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSignatureRequest.ProtoReflect.Descriptor instead.
func (*ReportSignatureRequest) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{10}
}

func (x *ReportSignatureRequest) GetContext() *ReportContext {
	if x != nil {
		return x.Context
	}
	return nil
}

func (x *ReportSignatureRequest) GetAttributedSignedObservations() []*AttributedSignedObservation {
	if x != nil {
		return x.AttributedSignedObservations
	}
	return nil
}

type ReportSignature struct {
	state protoimpl.MessageState `protogen:"open.v1"`
	// bytes signed_hash = 1; // needless since we have a request_id
	Signature     *EcdsaSignature `protobuf:"bytes,2,opt,name=signature,proto3" json:"signature,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ReportSignature) Reset() {
	*x = ReportSignature{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportSignature) ProtoMessage() {}

func (x *ReportSignature) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportSignature.ProtoReflect.Descriptor instead.
func (*ReportSignature) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{11}
}

func (x *ReportSignature) GetSignature() *EcdsaSignature {
	if x != nil {
		return x.Signature
	}
	return nil
}

// CCIP -> RMN
type Request struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	RequestId uint64                 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Request:
	//
	//	*Request_ObservationRequest
	//	*Request_ReportSignatureRequest
	Request       isRequest_Request `protobuf_oneof:"request"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Request) Reset() {
	*x = Request{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[12]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[12]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{12}
}

func (x *Request) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Request) GetRequest() isRequest_Request {
	if x != nil {
		return x.Request
	}
	return nil
}

func (x *Request) GetObservationRequest() *ObservationRequest {
	if x != nil {
		if x, ok := x.Request.(*Request_ObservationRequest); ok {
			return x.ObservationRequest
		}
	}
	return nil
}

func (x *Request) GetReportSignatureRequest() *ReportSignatureRequest {
	if x != nil {
		if x, ok := x.Request.(*Request_ReportSignatureRequest); ok {
			return x.ReportSignatureRequest
		}
	}
	return nil
}

type isRequest_Request interface {
	isRequest_Request()
}

type Request_ObservationRequest struct {
	ObservationRequest *ObservationRequest `protobuf:"bytes,2,opt,name=observation_request,json=observationRequest,proto3,oneof"`
}

type Request_ReportSignatureRequest struct {
	ReportSignatureRequest *ReportSignatureRequest `protobuf:"bytes,3,opt,name=report_signature_request,json=reportSignatureRequest,proto3,oneof"`
}

func (*Request_ObservationRequest) isRequest_Request() {}

func (*Request_ReportSignatureRequest) isRequest_Request() {}

// RMN -> CCIP
type Response struct {
	state     protoimpl.MessageState `protogen:"open.v1"`
	RequestId uint64                 `protobuf:"varint,1,opt,name=request_id,json=requestId,proto3" json:"request_id,omitempty"`
	// Types that are valid to be assigned to Response:
	//
	//	*Response_SignedObservation
	//	*Response_ReportSignature
	Response      isResponse_Response `protobuf_oneof:"response"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Response) Reset() {
	*x = Response{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[13]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[13]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{13}
}

func (x *Response) GetRequestId() uint64 {
	if x != nil {
		return x.RequestId
	}
	return 0
}

func (x *Response) GetResponse() isResponse_Response {
	if x != nil {
		return x.Response
	}
	return nil
}

func (x *Response) GetSignedObservation() *SignedObservation {
	if x != nil {
		if x, ok := x.Response.(*Response_SignedObservation); ok {
			return x.SignedObservation
		}
	}
	return nil
}

func (x *Response) GetReportSignature() *ReportSignature {
	if x != nil {
		if x, ok := x.Response.(*Response_ReportSignature); ok {
			return x.ReportSignature
		}
	}
	return nil
}

type isResponse_Response interface {
	isResponse_Response()
}

type Response_SignedObservation struct {
	SignedObservation *SignedObservation `protobuf:"bytes,2,opt,name=signed_observation,json=signedObservation,proto3,oneof"`
}

type Response_ReportSignature struct {
	ReportSignature *ReportSignature `protobuf:"bytes,3,opt,name=report_signature,json=reportSignature,proto3,oneof"`
}

func (*Response_SignedObservation) isResponse_Response() {}

func (*Response_ReportSignature) isResponse_Response() {}

type EcdsaSignature struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	R             []byte                 `protobuf:"bytes,1,opt,name=r,proto3" json:"r,omitempty"`
	S             []byte                 `protobuf:"bytes,2,opt,name=s,proto3" json:"s,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *EcdsaSignature) Reset() {
	*x = EcdsaSignature{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[14]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *EcdsaSignature) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EcdsaSignature) ProtoMessage() {}

func (x *EcdsaSignature) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[14]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EcdsaSignature.ProtoReflect.Descriptor instead.
func (*EcdsaSignature) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{14}
}

func (x *EcdsaSignature) GetR() []byte {
	if x != nil {
		return x.R
	}
	return nil
}

func (x *EcdsaSignature) GetS() []byte {
	if x != nil {
		return x.S
	}
	return nil
}

// To be sent by the OCR leader in a "BuildingReports" round.
type ReportingPluginQuery struct {
	state                protoimpl.MessageState `protogen:"open.v1"`
	FixedDestLaneUpdates []*FixedDestLaneUpdate `protobuf:"bytes,1,rep,name=fixed_dest_lane_updates,json=fixedDestLaneUpdates,proto3" json:"fixed_dest_lane_updates,omitempty"`
	EcdsaSignatures      []*EcdsaSignature      `protobuf:"bytes,2,rep,name=ecdsa_signatures,json=ecdsaSignatures,proto3" json:"ecdsa_signatures,omitempty"`
	unknownFields        protoimpl.UnknownFields
	sizeCache            protoimpl.SizeCache
}

func (x *ReportingPluginQuery) Reset() {
	*x = ReportingPluginQuery{}
	mi := &file_serialization_rmn_offchain_proto_msgTypes[15]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ReportingPluginQuery) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReportingPluginQuery) ProtoMessage() {}

func (x *ReportingPluginQuery) ProtoReflect() protoreflect.Message {
	mi := &file_serialization_rmn_offchain_proto_msgTypes[15]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReportingPluginQuery.ProtoReflect.Descriptor instead.
func (*ReportingPluginQuery) Descriptor() ([]byte, []int) {
	return file_serialization_rmn_offchain_proto_rawDescGZIP(), []int{15}
}

func (x *ReportingPluginQuery) GetFixedDestLaneUpdates() []*FixedDestLaneUpdate {
	if x != nil {
		return x.FixedDestLaneUpdates
	}
	return nil
}

func (x *ReportingPluginQuery) GetEcdsaSignatures() []*EcdsaSignature {
	if x != nil {
		return x.EcdsaSignatures
	}
	return nil
}

var File_serialization_rmn_offchain_proto protoreflect.FileDescriptor

const file_serialization_rmn_offchain_proto_rawDesc = "" +
	"\n" +
	" serialization/rmn_offchain.proto\x12\frmn_offchain\"g\n" +
	"\n" +
	"LaneSource\x122\n" +
	"\x15source_chain_selector\x18\x01 \x01(\x04R\x13sourceChainSelector\x12%\n" +
	"\x0eonramp_address\x18\x02 \x01(\fR\ronrampAddress\"c\n" +
	"\bLaneDest\x12.\n" +
	"\x13dest_chain_selector\x18\x01 \x01(\x04R\x11destChainSelector\x12'\n" +
	"\x0fofframp_address\x18\x02 \x01(\fR\x0eofframpAddress\"L\n" +
	"\x0eClosedInterval\x12\x1c\n" +
	"\n" +
	"min_msg_nr\x18\x01 \x01(\x04R\bminMsgNr\x12\x1c\n" +
	"\n" +
	"max_msg_nr\x18\x02 \x01(\x04R\bmaxMsgNr\"\x9e\x01\n" +
	"\x1aFixedDestLaneUpdateRequest\x129\n" +
	"\vlane_source\x18\x01 \x01(\v2\x18.rmn_offchain.LaneSourceR\n" +
	"laneSource\x12E\n" +
	"\x0fclosed_interval\x18\x02 \x01(\v2\x1c.rmn_offchain.ClosedIntervalR\x0eclosedInterval\"\xab\x01\n" +
	"\x13FixedDestLaneUpdate\x129\n" +
	"\vlane_source\x18\x01 \x01(\v2\x18.rmn_offchain.LaneSourceR\n" +
	"laneSource\x12E\n" +
	"\x0fclosed_interval\x18\x02 \x01(\v2\x1c.rmn_offchain.ClosedIntervalR\x0eclosedInterval\x12\x12\n" +
	"\x04root\x18\x03 \x01(\fR\x04root\"\xb9\x01\n" +
	"\x12ObservationRequest\x123\n" +
	"\tlane_dest\x18\x01 \x01(\v2\x16.rmn_offchain.LaneDestR\blaneDest\x12n\n" +
	"\x1ffixed_dest_lane_update_requests\x18\x02 \x03(\v2(.rmn_offchain.FixedDestLaneUpdateRequestR\x1bfixedDestLaneUpdateRequests\"\x80\x02\n" +
	"\vObservation\x12D\n" +
	"\x1frmn_home_contract_config_digest\x18\x01 \x01(\fR\x1brmnHomeContractConfigDigest\x123\n" +
	"\tlane_dest\x18\x02 \x01(\v2\x16.rmn_offchain.LaneDestR\blaneDest\x12X\n" +
	"\x17fixed_dest_lane_updates\x18\x03 \x03(\v2!.rmn_offchain.FixedDestLaneUpdateR\x14fixedDestLaneUpdates\x12\x1c\n" +
	"\ttimestamp\x18\x04 \x01(\x04R\ttimestamp\"n\n" +
	"\x11SignedObservation\x12;\n" +
	"\vobservation\x18\x01 \x01(\v2\x19.rmn_offchain.ObservationR\vobservation\x12\x1c\n" +
	"\tsignature\x18\x02 \x01(\fR\tsignature\"\x99\x01\n" +
	"\x1bAttributedSignedObservation\x12N\n" +
	"\x12signed_observation\x18\x01 \x01(\v2\x1f.rmn_offchain.SignedObservationR\x11signedObservation\x12*\n" +
	"\x11signer_node_index\x18\x02 \x01(\rR\x0fsignerNodeIndex\"\xf4\x01\n" +
	"\rReportContext\x12)\n" +
	"\x11evm_dest_chain_id\x18\x01 \x01(\x04R\x0eevmDestChainId\x12=\n" +
	"\x1brmn_remote_contract_address\x18\x02 \x01(\fR\x18rmnRemoteContractAddress\x12D\n" +
	"\x1frmn_home_contract_config_digest\x18\x03 \x01(\fR\x1brmnHomeContractConfigDigest\x123\n" +
	"\tlane_dest\x18\x04 \x01(\v2\x16.rmn_offchain.LaneDestR\blaneDest\"\xc0\x01\n" +
	"\x16ReportSignatureRequest\x125\n" +
	"\acontext\x18\x01 \x01(\v2\x1b.rmn_offchain.ReportContextR\acontext\x12o\n" +
	"\x1eattributed_signed_observations\x18\x02 \x03(\v2).rmn_offchain.AttributedSignedObservationR\x1cattributedSignedObservations\"M\n" +
	"\x0fReportSignature\x12:\n" +
	"\tsignature\x18\x02 \x01(\v2\x1c.rmn_offchain.EcdsaSignatureR\tsignature\"\xea\x01\n" +
	"\aRequest\x12\x1d\n" +
	"\n" +
	"request_id\x18\x01 \x01(\x04R\trequestId\x12S\n" +
	"\x13observation_request\x18\x02 \x01(\v2 .rmn_offchain.ObservationRequestH\x00R\x12observationRequest\x12`\n" +
	"\x18report_signature_request\x18\x03 \x01(\v2$.rmn_offchain.ReportSignatureRequestH\x00R\x16reportSignatureRequestB\t\n" +
	"\arequest\"\xd3\x01\n" +
	"\bResponse\x12\x1d\n" +
	"\n" +
	"request_id\x18\x01 \x01(\x04R\trequestId\x12P\n" +
	"\x12signed_observation\x18\x02 \x01(\v2\x1f.rmn_offchain.SignedObservationH\x00R\x11signedObservation\x12J\n" +
	"\x10report_signature\x18\x03 \x01(\v2\x1d.rmn_offchain.ReportSignatureH\x00R\x0freportSignatureB\n" +
	"\n" +
	"\bresponse\",\n" +
	"\x0eEcdsaSignature\x12\f\n" +
	"\x01r\x18\x01 \x01(\fR\x01r\x12\f\n" +
	"\x01s\x18\x02 \x01(\fR\x01s\"\xb9\x01\n" +
	"\x14ReportingPluginQuery\x12X\n" +
	"\x17fixed_dest_lane_updates\x18\x01 \x03(\v2!.rmn_offchain.FixedDestLaneUpdateR\x14fixedDestLaneUpdates\x12G\n" +
	"\x10ecdsa_signatures\x18\x02 \x03(\v2\x1c.rmn_offchain.EcdsaSignatureR\x0fecdsaSignaturesB\x12Z\x10./;serializationb\x06proto3"

var (
	file_serialization_rmn_offchain_proto_rawDescOnce sync.Once
	file_serialization_rmn_offchain_proto_rawDescData []byte
)

func file_serialization_rmn_offchain_proto_rawDescGZIP() []byte {
	file_serialization_rmn_offchain_proto_rawDescOnce.Do(func() {
		file_serialization_rmn_offchain_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_serialization_rmn_offchain_proto_rawDesc), len(file_serialization_rmn_offchain_proto_rawDesc)))
	})
	return file_serialization_rmn_offchain_proto_rawDescData
}

var file_serialization_rmn_offchain_proto_msgTypes = make([]protoimpl.MessageInfo, 16)
var file_serialization_rmn_offchain_proto_goTypes = []any{
	(*LaneSource)(nil),                  // 0: rmn_offchain.LaneSource
	(*LaneDest)(nil),                    // 1: rmn_offchain.LaneDest
	(*ClosedInterval)(nil),              // 2: rmn_offchain.ClosedInterval
	(*FixedDestLaneUpdateRequest)(nil),  // 3: rmn_offchain.FixedDestLaneUpdateRequest
	(*FixedDestLaneUpdate)(nil),         // 4: rmn_offchain.FixedDestLaneUpdate
	(*ObservationRequest)(nil),          // 5: rmn_offchain.ObservationRequest
	(*Observation)(nil),                 // 6: rmn_offchain.Observation
	(*SignedObservation)(nil),           // 7: rmn_offchain.SignedObservation
	(*AttributedSignedObservation)(nil), // 8: rmn_offchain.AttributedSignedObservation
	(*ReportContext)(nil),               // 9: rmn_offchain.ReportContext
	(*ReportSignatureRequest)(nil),      // 10: rmn_offchain.ReportSignatureRequest
	(*ReportSignature)(nil),             // 11: rmn_offchain.ReportSignature
	(*Request)(nil),                     // 12: rmn_offchain.Request
	(*Response)(nil),                    // 13: rmn_offchain.Response
	(*EcdsaSignature)(nil),              // 14: rmn_offchain.EcdsaSignature
	(*ReportingPluginQuery)(nil),        // 15: rmn_offchain.ReportingPluginQuery
}
var file_serialization_rmn_offchain_proto_depIdxs = []int32{
	0,  // 0: rmn_offchain.FixedDestLaneUpdateRequest.lane_source:type_name -> rmn_offchain.LaneSource
	2,  // 1: rmn_offchain.FixedDestLaneUpdateRequest.closed_interval:type_name -> rmn_offchain.ClosedInterval
	0,  // 2: rmn_offchain.FixedDestLaneUpdate.lane_source:type_name -> rmn_offchain.LaneSource
	2,  // 3: rmn_offchain.FixedDestLaneUpdate.closed_interval:type_name -> rmn_offchain.ClosedInterval
	1,  // 4: rmn_offchain.ObservationRequest.lane_dest:type_name -> rmn_offchain.LaneDest
	3,  // 5: rmn_offchain.ObservationRequest.fixed_dest_lane_update_requests:type_name -> rmn_offchain.FixedDestLaneUpdateRequest
	1,  // 6: rmn_offchain.Observation.lane_dest:type_name -> rmn_offchain.LaneDest
	4,  // 7: rmn_offchain.Observation.fixed_dest_lane_updates:type_name -> rmn_offchain.FixedDestLaneUpdate
	6,  // 8: rmn_offchain.SignedObservation.observation:type_name -> rmn_offchain.Observation
	7,  // 9: rmn_offchain.AttributedSignedObservation.signed_observation:type_name -> rmn_offchain.SignedObservation
	1,  // 10: rmn_offchain.ReportContext.lane_dest:type_name -> rmn_offchain.LaneDest
	9,  // 11: rmn_offchain.ReportSignatureRequest.context:type_name -> rmn_offchain.ReportContext
	8,  // 12: rmn_offchain.ReportSignatureRequest.attributed_signed_observations:type_name -> rmn_offchain.AttributedSignedObservation
	14, // 13: rmn_offchain.ReportSignature.signature:type_name -> rmn_offchain.EcdsaSignature
	5,  // 14: rmn_offchain.Request.observation_request:type_name -> rmn_offchain.ObservationRequest
	10, // 15: rmn_offchain.Request.report_signature_request:type_name -> rmn_offchain.ReportSignatureRequest
	7,  // 16: rmn_offchain.Response.signed_observation:type_name -> rmn_offchain.SignedObservation
	11, // 17: rmn_offchain.Response.report_signature:type_name -> rmn_offchain.ReportSignature
	4,  // 18: rmn_offchain.ReportingPluginQuery.fixed_dest_lane_updates:type_name -> rmn_offchain.FixedDestLaneUpdate
	14, // 19: rmn_offchain.ReportingPluginQuery.ecdsa_signatures:type_name -> rmn_offchain.EcdsaSignature
	20, // [20:20] is the sub-list for method output_type
	20, // [20:20] is the sub-list for method input_type
	20, // [20:20] is the sub-list for extension type_name
	20, // [20:20] is the sub-list for extension extendee
	0,  // [0:20] is the sub-list for field type_name
}

func init() { file_serialization_rmn_offchain_proto_init() }
func file_serialization_rmn_offchain_proto_init() {
	if File_serialization_rmn_offchain_proto != nil {
		return
	}
	file_serialization_rmn_offchain_proto_msgTypes[12].OneofWrappers = []any{
		(*Request_ObservationRequest)(nil),
		(*Request_ReportSignatureRequest)(nil),
	}
	file_serialization_rmn_offchain_proto_msgTypes[13].OneofWrappers = []any{
		(*Response_SignedObservation)(nil),
		(*Response_ReportSignature)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_serialization_rmn_offchain_proto_rawDesc), len(file_serialization_rmn_offchain_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   16,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_serialization_rmn_offchain_proto_goTypes,
		DependencyIndexes: file_serialization_rmn_offchain_proto_depIdxs,
		MessageInfos:      file_serialization_rmn_offchain_proto_msgTypes,
	}.Build()
	File_serialization_rmn_offchain_proto = out.File
	file_serialization_rmn_offchain_proto_goTypes = nil
	file_serialization_rmn_offchain_proto_depIdxs = nil
}
