package storage_service

//go:generate protoc --proto_path=../ --proto_path=../../workflows --go_out=./ --go_opt=Mevents/metering.proto=github.com/smartcontractkit/chainlink-protos/storage-service/go/events --go_opt=module=github.com/smartcontractkit/chainlink-protos/storage-service/go --go-grpc_out=./ --go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/storage-service/go ../node_service/v1/node_service.proto
