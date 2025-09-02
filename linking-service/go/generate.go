package linking_service

//go:generate protoc --proto_path=../ --go_out=./v1/ --go_opt=module=github.com/smartcontractkit/chainlink-protos/linking-service/go/v1 --go-grpc_out=./v1/ --go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/linking-service/go/v1 ../v1/linking-service.proto
//go:generate protoc --proto_path=../ --go_out=./v1/ --go_opt=module=github.com/smartcontractkit/chainlink-protos/linking-service/go/v1 ../v1/linking-types.proto
