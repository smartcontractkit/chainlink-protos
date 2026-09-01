.PHONY: clean
clean: ## Remove generated files
	find ./go -name '*.pb.go' -delete

.PHONY: generate
generate: clean ## Generate go-proto files
	cd go && go generate ./...
