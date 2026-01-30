package ring

//go:generate protoc --proto_path=../pb --go_out=./ --go_opt=module=github.com/smartcontractkit/chainlink-protos/ring/go --go-grpc_out=./ --go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/ring/go ../pb/shared.proto ../pb/arbiter.proto ../pb/consensus.proto ../pb/shard_orchestrator.proto
