.PHONY: clean
clean: ## Remove generated files
	find ./go -name '*.pb.go' -delete

.PHONY: generate
generate: clean ## Generate go-proto files
	protoc \
		--proto_path=./ \
		--proto_path=../workflows \
		--go_out=./go \
		--go_opt=Mevents/metering.proto=github.com/smartcontractkit/chainlink-protos/workflows/go/events \
		--go_opt=module=github.com/smartcontractkit/chainlink-protos/billing/go \
		--go-grpc_out=./go \
		--go-grpc_opt=module=github.com/smartcontractkit/chainlink-protos/billing/go \
		$(shell find ./billing/v1 -name "*.proto")