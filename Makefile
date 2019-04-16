GOFMT_FILES?=$$(find . -name '*.go' | grep -v vendor)

fmt:
	gofmt -l -w $(GOFMT_FILES)

