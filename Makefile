PACKAGES = $(wildcard funcs/*)
BUILD_PACKAGES = $(PACKAGES:%=build_%)

.PHONY: build $(BUILD_PACKAGES) lint test

build: $(BUILD_PACKAGES)
$(BUILD_PACKAGES): build_%:
	GOARCH=amd64 GOOS=linux go build -o build/$*/main ./$*

lint:
	golangci-lint run -v ./funcs/...

test:
	go test -p 1 ./funcs/...
