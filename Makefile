NAME=textanalysis
PORT=7777
VERSION=$(shell git describe --tags --always --dirty)

export GOBIN=$(shell pwd)/bin

build: neat
	mkdir -p bin
	$(RM) bin/$(NAME)
	find src/cmd -type f -name '*.go' | xargs -I xxx go install -ldflags "-X pkg/version.VERSION=$(VERSION)" xxx

run: build
	./bin/$(NAME)

neat:
	find src/pkg src/cmd -type f -name '*.go' | xargs -I xxx go fmt xxx
