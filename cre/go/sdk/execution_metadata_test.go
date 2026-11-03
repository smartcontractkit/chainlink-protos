package sdk_test

import (
	"slices"
	"testing"

	"github.com/smartcontractkit/chainlink-protos/cre/go/sdk"
	"github.com/stretchr/testify/require"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protodesc"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/dynamicpb"
)

// Synthetic identifiers use production encodings; execution IDs follow the
// engine hash for the workflow ID, a cron event timestamp, and trigger index 0.
const (
	metadataWorkflowIDFixture  = "0020350fb4f89856a031268826d6e9b9340996637aa27859baa7b07db58bb020"
	metadataExecutionIDFixture = "9f38bdece8892b3452c401a7a2d5a170f078ab9b2c98c591e24028ce94b33ed2"
	metadataOwnerFixture       = "5366890c26e8737b89d9084b4c632f7de73d1674"
)

func TestExecuteRequest_MetadataRoundTrip(t *testing.T) {
	for _, tc := range []struct {
		name     string
		metadata *sdk.WorkflowExecutionMetadata
	}{
		{name: "absent"},
		{name: "empty", metadata: &sdk.WorkflowExecutionMetadata{}},
		{name: "present", metadata: &sdk.WorkflowExecutionMetadata{WorkflowId: metadataWorkflowIDFixture, ExecutionId: metadataExecutionIDFixture, WorkflowOwner: metadataOwnerFixture, DonId: 17, DonF: 3}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			request := &sdk.ExecuteRequest{Config: []byte(`{"schedule":"0 */5 * * * *"}`), MaxResponseSize: 2048, Request: &sdk.ExecuteRequest_Trigger{Trigger: &sdk.Trigger{Id: 0}}, ExecutionMetadata: tc.metadata}
			encoded, err := proto.Marshal(request)
			require.NoError(t, err)
			var decoded sdk.ExecuteRequest
			require.NoError(t, proto.Unmarshal(encoded, &decoded))
			require.True(t, proto.Equal(request, &decoded))
		})
	}
}

func TestExecuteRequest_LegacyReaderPreservesMetadata(t *testing.T) {
	request := &sdk.ExecuteRequest{
		Config: []byte(`{"schedule":"0 */5 * * * *"}`), MaxResponseSize: 2048,
		Request:           &sdk.ExecuteRequest_Trigger{Trigger: &sdk.Trigger{Id: 0}},
		ExecutionMetadata: &sdk.WorkflowExecutionMetadata{WorkflowId: metadataWorkflowIDFixture, ExecutionId: metadataExecutionIDFixture, WorkflowOwner: metadataOwnerFixture, DonId: 17, DonF: 3},
	}
	encoded, err := proto.Marshal(request)
	require.NoError(t, err)

	// Remove metadata from the descriptor to model an actual legacy reader.
	descriptor := protodesc.ToFileDescriptorProto(sdk.File_sdk_v1alpha_sdk_proto)
	message := descriptor.MessageType[request.ProtoReflect().Descriptor().Index()]
	message.Field = slices.DeleteFunc(message.Field, func(field *descriptorpb.FieldDescriptorProto) bool {
		return field.GetNumber() == 7
	})
	file, err := protodesc.NewFile(descriptor, protoregistry.GlobalFiles)
	require.NoError(t, err)
	legacy := dynamicpb.NewMessage(file.Messages().ByName("ExecuteRequest"))
	require.NoError(t, proto.Unmarshal(encoded, legacy))
	require.NotEmpty(t, legacy.GetUnknown())

	reencoded, err := proto.Marshal(legacy)
	require.NoError(t, err)
	var restored sdk.ExecuteRequest
	require.NoError(t, proto.Unmarshal(reencoded, &restored))
	require.True(t, proto.Equal(request, &restored))
}
