GO              := GO111MODULE=on CGO_ENABLED=0 go
GOBUILD         := $(GO) build $(BUILD_FLAG) -tags codes

test:
	$(GO) test $(BUILD_FLAG) -tags codes ./...

.PHONY: test