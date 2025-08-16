package platform

//go:generate protoc --proto_path=../ --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/platform ../workflows/bridge_status/v1/bridge_status.proto
