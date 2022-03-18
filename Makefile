# Makefile

.PHONY: all
all: proto test

.PHONY: test
test: ## Run unit tests
	@go test ./...

.PHONY: vet
vet: ## Run go vet
	@go vet ./...

.PHONY: coverage
coverage: ## Generate global code coverage report
	@go test -cover ./...

.PHONY: deps
deps: ## Get the dependencies
	@go get ./...

.PHONY: proto
proto: ## build proto definitions
	@./script/generate.sh

.PHONY: proto_check
proto_check: ## proto: Check committed files match just-generated files
	@make proto_clean 1>/dev/null
	@make proto 1>/dev/null
	@files="$$(git diff --name-only generated/ vega/ swagger/ data-node/)" ; \
	if test -n "$$files" ; then \
		echo "Committed files do not match just-generated files: " $$files ; \
		test -n "$(CI)" && git diff vega/ ; \
		exit 1 ; \
	fi

.PHONY: proto_clean
proto_clean:
	@find vega swagger data-node -name '*.pb.go' -o -name '*.pb.gw.go' -o -name '*.validator.pb.go' -o -name '*.swagger.json' \
		| xargs -r rm

# Misc Targets

.PHONY: buflint
buflint: ## Run buf lint
	@buf lint


.PHONY: help
help: ## Display this help screen
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONE: clean
clean:
	rm -rf ./**/*-re
