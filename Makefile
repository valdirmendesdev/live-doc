GO ?= go
TEST_RUN = ""
.PHONY: test

test:
	$(GO) test -v ./... -run $(TEST_RUN)

build-mocks:
	$(GO) get github.com/golang/mock/gomock
	$(GO) install github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=internal/core/usecases/customer/interface.go -destination=internal/core/usecases/customer/mock/customer.go -package=mock
