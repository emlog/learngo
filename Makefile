.PHONY: run
run: init swag
	@go run -tags=jsoniter main.go -run_env="testing"

.PHONY: init
init:
ifeq (, $(shell git config core.hooksPath))
	@echo "如果你的git版本低于2.9.0 请手动设置下面命令"
	@echo "find .git/hooks -type l -exec rm {} \;"
	@echo "find .githooks -type f -exec ln -sf ../../{} .git/hooks/ \;"
	@git config core.hooksPath .githooks
endif

.PHONY: swag
swag: check
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
lint: check
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
	@goimports -l -w -local github.com/emlog/goexample/judge $(shell find . -type f -name '*.go' -not -path './vendor/*' -not -path './docs/*')