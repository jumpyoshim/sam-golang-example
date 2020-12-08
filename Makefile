build:
	cd cdk && \
	npm run build && \
	npx cdk synth --no-staging api > template.yml && \
	sam build

start: build
	sam local start-api -t cdk/template.yml

lint:
	golangci-lint run -v ./funcs/...

test:
	go test -p 1 ./funcs/...
