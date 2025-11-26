package chainlink_ccv

//go:generate sh -c "protoc -I../.. -I\"$(go list -m -f '{{.Dir}}' github.com/googleapis/googleapis)\" --go_out=./v1/ --go_opt=module=github.com/smartcontractkit/chainlink-protos/chainlink-ccv/go/v1 --go-grpc_out=./v1/ --go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/chainlink-ccv/go/v1 chainlink-ccv/v1/common.proto chainlink-ccv/v1/committee-verifier.proto chainlink-ccv/v1/message-discovery.proto chainlink-ccv/v1/verifier-result-api.proto"
