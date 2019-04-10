SOURCE_FILE := $(shell find . -name "*.go" -type f)

fmt:
	$(foreach var,$(SOURCE_FILE),$(shell go fmt $(var)))

