GO ?= go
TEST_RUN = ""
.PHONY: test

test:
	$(GO) test -v ./... -run $(TEST_RUN)

build-test-coverage:
	$(GO) test ./... -coverprofile=coverage/profile.out

show-test-coverage:	build-test-coverage
	$(GO) tool cover -html=coverage/profile.out

build-mocks:
	$(GO) get github.com/golang/mock/gomock
	$(GO) install github.com/golang/mock/mockgen@v1.6.0
	@~/go/bin/mockgen -source=internal/live-docs/core/customer/repository.go -destination=internal/live-docs/core/mocks/customer.go -package=mocks -mock_names=Repository=MockCustomer

