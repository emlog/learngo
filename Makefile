.PHONY: run
run: swag
	@go run main.go -c ./deployments/config.testing.toml

.PHONY: swag
swag:
	@swag i -g ./main.go

.PHONY: mod
mod:
	@go mod tidy

.PHONY: test
test:
	@go test -gcflags "all=-l" -race -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out

.PHONY: clean
clean:
	-@rm -rf logs/* bin/* coverage.out

.PHONY: lint
lint:
	@golangci-lint config path && golangci-lint run

.PHONY: check
check:
ifeq (, $(shell which golangci-lint))
	$(error golangci-lint not found in PATH)
endif
ifeq (, $(shell which goimports))
	$(error goimports not found in PATH)
endif
ifeq (, $(shell which swag))
	$(error swag not found in PATH)
endif

.PHONY: fmt
fmt: check
	@goimports -l -w -local gitlab.luojilab.com/igetserver/ledgers $(shell find . -type f -name '*.go' -not -path './vendor/*' -not -path './docs/*' -not -path './models/models*')
