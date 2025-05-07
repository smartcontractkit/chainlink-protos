module github.com/smartcontractkit/chainlink-protos/billing/go

go 1.24.2

require (
	github.com/smartcontractkit/chainlink-protos/workflows/go v0.0.0-20250430163438-97d324ef9061
	google.golang.org/grpc v1.72.0
	google.golang.org/protobuf v1.36.6
)

require (
	golang.org/x/net v0.35.0 // indirect
	golang.org/x/sys v0.30.0 // indirect
	golang.org/x/text v0.22.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20250218202821-56aae31c358a // indirect
)

replace github.com/smartcontractkit/chainlink-protos/workflows/go => ../../workflows/go
