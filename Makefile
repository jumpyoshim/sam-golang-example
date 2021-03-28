PACKAGES = $(wildcard funcs/*)
BUILD_PACKAGES = $(PACKAGES:%=build_%)

.PHONY: build $(BUILD_PACKAGES) lint test synth

synth:
	cd cdk && \
	npm run build && \
	npx cdk synth --no-staging api-private > template.yml

setup: synth
	docker-compose up -d && \
	DYNAMODB_LOCAL_ENDPOINT=http://localhost:8800 go run scripts/reset-tables/main.go

build: $(BUILD_PACKAGES)
$(BUILD_PACKAGES): build_%:
	GOARCH=amd64 GOOS=linux go build -o build/$*/main ./$*

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
