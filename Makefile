lint:
	./bin/lint

test:
	mkdir -p $(COVERAGE_DIR)
	go test -p 1 $(TEST_OPTIONS) ./funcs/... ./libs/...
