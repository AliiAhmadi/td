.DEFAUL_GOAL := run

run:
	@go run ./cmd/todo

test: clear_cache
	@go test ./...

clear_cache:
	@go clean -cache

compile_check:
	@go build -o temp ./cmd/todo && rm -rf temp

.PHONY: run test clear_cache compile_check