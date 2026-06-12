package metering

//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/metering/go ../metering/v1/meter_record.proto
