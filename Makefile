VERBOSE_FLAG = $(if $(VERBOSE),-v)

VERSION = $$(git describe --tags --always --dirty) ($$(git name-rev --name-only HEAD))

BUILD_FLAGS = -ldflags "-X autil.Version \"$(VERSION)\""

build: deps
		go build $(VERBOSE_FLAG) $(BUILD_FLAGS)

test: testdeps
		go test $(VERBOSE_FLAG) ./...

deps:
		go get -d $(VERBOSE_FLAG)

testdeps:
		go get -d -t $(VERBOSE_FLAG)

install: deps
		go install $(VERBOSE_FLAG) $(BUILD_FLAGS)
		cp -p bin/sshec2 $(GOROOT)/bin

.PHONY: build test deps testdeps install
