package workflow_ownership_service

//go:generate protoc --proto_path=../ --go_out=./v1/ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflow-ownership-service/go/v1 --go-grpc_out=./v1/ --go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/workflow-ownership-service/go/v1 ../v1/workflow-ownership-service.proto
//go:generate protoc --proto_path=../ --go_out=./v1/ --go_opt=module=github.com/smartcontractkit/chainlink-protos/workflow-ownership-service/go/v1 ../v1/workflow-ownership-types.proto
