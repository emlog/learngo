.PHONY: all test vendor docker golint build clean swag

SHELL = /bin/bash

# preset constants and variables
PWD = $(shell pwd)
DEPS := $(wildcard *.go)

# color printing
CE = "\033[0m"
CRED = "\033[31m"
CGREEN = "\033[32m"
CYELLOW = "\033[33m"

# version set
GOCILINT_VERSION = "v1.44.2"

# program detection
SWAG_EXISTS := $(shell which swag 2>/dev/null | grep -c "/swag" )
GOCILINT_EXISTS := $(shell which golangci-lint 2>/dev/null | grep -c "/golangci-lint" )
DOCKER_CHECK := $(shell sysctl net.inet.tcp.sack 2>/dev/null | grep -c '0')

# commands
all: swag build

# swag gofmt
build: clean
	@-echo -e ${CGREEN} "make build........................................................."
	@-go mod tidy
	@-go build -o application
	@-echo -e ${CGREEN} "build succeeds with application as server executable..............."

run: swag build r

r:
	@-echo -e ${CGREEN} "make run.........................................................."
	RUN_ENV=development ./application

clean:
	@-echo -e ${CGREEN} "make clean........................................................."
	@-rm -rf logs
	@-rm -f application

test:
	@-echo -e ${CGREEN} "make test........................................................."
	@go test ./...

gofmt:
	@-echo -e ${CGREEN} "make gofmt........................................................."
	@-echo 'gofmt -l -w  $$(find . -type f -name "*.go" -not -path "./vendor/*")  '
	@-gofmt -l -w  $$(find . -type f -name '*.go' -not -path './vendor/*') >/dev/null

swag:
	@-echo -e ${CGREEN} "make swag........................................................."
	@-if [ ${SWAG_EXISTS} -eq 0 ]; then  \
		GO111MODULE="on" go get github.com/swaggo/gin-swagger; \
		GO111MODULE="on" go install github.com/swaggo/swag/cmd/swag@latest; \
		if [ ${SWAG_EXISTS} -eq 0 ]; then  \
		   export PATH=$$PATH:$$GOPATH/bin;\
		   echo -e ${CYELLOW} "请确保你的 ~/.bashrc 或 ~/.zshrc 中 PATH 包括 $$GOPATH/bin";\
		fi; \
	fi
	@if [ $$(swag i -g main.go 2>&1 | tee /dev/stderr | grep -c "cannot find type definition") -gt 0 ]; then \
    	swag i --pd --parseInternal --parseDepth 1 -g main.go; \
    fi \


vendor:
	@-echo -e ${CGREEN} "make vendor........................................................."
	-GO111MODULE="on" go mod tidy
	-GO111MODULE="on" go mod vendor

lint:
	@-echo -e ${CGREEN} "make lint........................................................."
	@-if [ ${GOCILINT_EXISTS} -eq 0 ]; then  \
		echo "未检测到golangci-lint，开始安装.........................................................";\
		echo "> 不使用go get方式安装的原因，参见：https://golangci-lint.run/usage/install/#install-from-source";\
		echo "> 不使用brew install，原因暂时无法指定版本安装";\
		echo "> 关于golangci-lint更多信息，请查阅 https://golangci-lint.run/usage/quick-start/";\
		echo "> 若提示无法找到golangci-lint，尝试 source ~/.bashrc 或 ~/.zshrc";\
	fi
	@golangci-lint run
