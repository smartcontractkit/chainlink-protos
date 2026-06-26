package metering

//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/metering/go ../metering/v1/identity.proto ../metering/v1/utilization.proto ../metering/v1/meter_record.proto ../metering/v1/snapshot.proto
