.PHONY: preview-staging
preview-staging:
	iron-flask deploy-confluent -e staging -c ./iron-flask-data-feeds.yaml --preview

.PHONY: deploy-staging
deploy-staging:
	iron-flask deploy-confluent -e staging -c ./iron-flask-data-feeds.yaml

.PHONY: destroy-staging
destroy-staging:
	iron-flask deploy-confluent -e staging -c ./iron-flask-data-feeds.yaml --destroy

.PHONY: preview-prod
preview-prod:
	iron-flask deploy-confluent -e prod -c ./iron-flask-data-feeds.yaml --preview

.PHONY: deploy-prod
deploy-prod:
	iron-flask deploy-confluent -e prod -c ./iron-flask-data-feeds.yaml

.PHONY: destroy-prod
destroy-prod:
	iron-flask deploy-confluent -e prod -c ./iron-flask-data-feeds.yaml --destroy

.PHONY: clean
clean: ## Remove generated files
	find . -name '*.pb.go' -delete

.PHONY: generate
generate: clean ## Generate go-proto files
	go generate ./...
