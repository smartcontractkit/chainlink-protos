package workflows

//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/capability_finished.proto ../workflows/v1/capability_started.proto ../workflows/v1/metadata.proto ../workflows/v1/metering_detail.proto ../workflows/v1/metering_step.proto ../workflows/v1/metering.proto ../workflows/v1/transmit_schedule_event.proto ../workflows/v1/workflow_finished.proto ../workflows/v1/workflow_started.proto ../workflows/v1/workflow_status_changed.proto
