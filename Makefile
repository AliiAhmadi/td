.DEFAUL_GOAL := run

run:
	@go run ./cmd/

test: clear_cache
	@go test ./...

clear_cache:
	@go clean -cache

.PHONY: run test clear_cache