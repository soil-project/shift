# This Makefile is meant to be used by people that do not usually work
# with Go source code. If you know what GOPATH is then you probably
# don't need to bother with make.

.PHONY: shift evm mist all test travis-test-with-coverage clean
GOBIN = build/bin

shift:
	build/env.sh go install -v $(shell build/ldflags.sh) ./cmd/shift
	@echo "Done building."
	@echo "Run \"$(GOBIN)/shift\" to launch shift."

evm:
	build/env.sh $(GOROOT)/bin/go install -v $(shell build/ldflags.sh) ./cmd/evm
	@echo "Done building."
	@echo "Run \"$(GOBIN)/evm to start the evm."
mist:
	build/env.sh go install -v $(shell build/ldflags.sh) ./cmd/mist
	@echo "Done building."
	@echo "Run \"$(GOBIN)/mist --asset_path=cmd/mist/assets\" to launch mist."

all:
	build/env.sh go install -v $(shell build/ldflags.sh) ./...

test: all
	build/env.sh go test ./...

travis-test-with-coverage: all
	build/env.sh build/test-global-coverage.sh

clean:
	rm -fr build/_workspace/pkg/ Godeps/_workspace/pkg $(GOBIN)/*
