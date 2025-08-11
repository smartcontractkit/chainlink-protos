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

// v2 operational events
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/execution_status.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/cre_info.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow_key.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/transaction_info.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow_deployed.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow_updated.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow_deleted.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/trigger_execution_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow_execution_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/workflow_execution_finished.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/capability_execution_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v2/capability_execution_finished.proto

// do not include deprecated workflows/pb/base_message_legacy.proto
