package workflows

//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/capability_finished.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/capability_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/metadata.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/metering.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/transmit_schedule_event.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/transmissions_scheduled_event.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/user_logs.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/workflow_finished.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/workflow_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/workflow_status_changed.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../common/v1/base_message.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/operational_events.proto
// do not include deprecated workflows/pb/base_message_legacy.proto
