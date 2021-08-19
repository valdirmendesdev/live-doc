GO ?= go
TEST_RUN = ""
.PHONY: test

test:
	$(GO) test -v ./... -run $(TEST_RUN)

build-mocks:
	$(GO) get github.com/golang/mock/gomock
	$(GO) install github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=internal/core/repositories/customer.go -destination=internal/core/repositories/mocks/customer.go -package=mocks
