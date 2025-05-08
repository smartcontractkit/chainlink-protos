package workflows

//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/capability_finished.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/capability_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/metadata.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/metering.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/transmit_schedule_event.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/workflow_finished.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/workflow_started.proto
//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflows/go ../workflows/v1/workflow_status_changed.proto
