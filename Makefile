PACKAGES = $(wildcard funcs/*)
BUILD_PACKAGES = $(PACKAGES:%=build_%)

.PHONY: build $(BUILD_PACKAGES) lint test

build: $(BUILD_PACKAGES)
$(BUILD_PACKAGES): build_%:
	GOARCH=amd64 GOOS=linux go build -o build/$*/main ./$*


synth:
	cd cdk && \
	npm run build && \
	npx cdk synth --no-staging api-private > template.yml

start: build
	sam local start-api -t cdk/template.yml \
		--docker-network sam-golang-example_default \
		--port 8000 \
		--host 0.0.0.0 \
		--env-vars env.json

lint:
	golangci-lint run -v ./funcs/...

test:
	go test -v ./funcs/...
