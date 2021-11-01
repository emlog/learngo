
all:
	go build github.com/emlog/goexample/cmd/bookstore

.PHONY: fmt
fmt:
	@goimports -l -w -local github.com/emlog/goexample $(shell find . -type f -name '*.go' -not -path './vendor/*' -not -path './docs/*')
