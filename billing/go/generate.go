package billing

//go:generate protoc --proto_path=../ --proto_path=../../workflows --go_out=./ --go_opt=Mevents/metering.proto=github.com/smartcontractkit/chainlink-protos/workflows/go/events --go_opt=module=github.com/smartcontractkit/chainlink-protos/billing/go --go-grpc_out=./ --go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/billing/go ../billing/v1/credit_reservation_message_types.proto ../billing/v1/credit_reservation_service.proto
