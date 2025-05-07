.PHONY: preview-staging
preview-staging:
	iron-flask deploy-all -e staging -c ./iron-flask.yaml --preview

.PHONY: deploy-staging
deploy-staging:
	iron-flask deploy-all -e staging -c ./iron-flask.yaml

.PHONY: destroy-staging
destroy-staging:
	iron-flask deploy-all -e staging -c ./iron-flask.yaml --destroy

.PHONY: preview-prod
preview-prod:
	iron-flask deploy-all -e prod -c ./iron-flask.yaml --preview

.PHONY: deploy-prod
deploy-prod:
	iron-flask deploy-all -e prod -c ./iron-flask.yaml

.PHONY: destroy-prod
destroy-prod:
	iron-flask deploy-all -e prod -c ./iron-flask.yaml --destroy

.PHONY: clean
clean: ## Remove generated files
	find ./go -name '*.pb.go' -delete

.PHONY: generate
generate: clean ## Generate go-proto files
	cd go && go generate ./...
