module github.com/smartcontractkit/chainlink-protos/billing/go

go 1.24.2

require (
	github.com/smartcontractkit/chainlink-protos/workflows/go v0.0.0-20250430163438-97d324ef9061
	google.golang.org/protobuf v1.36.6
)

replace github.com/smartcontractkit/chainlink-protos/workflows/go => ../../workflows/go
