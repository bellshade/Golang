# EXCLUDE_FOLDER = "basic\|intermediate\|utilities"
# TEST_PACKAGE := $(shell go list ./... | grep -v $(EXCLUDE_FOLDER))
TEST_PACKAGE := $(shell go list ./tests/...)

.DEFAULT_GOAL := test

test:
	@echo "testing file golang"
	go test -v $(TEST_PACKAGE)

test-race:
	@echo "testing dengan race detector"
	go test -race -v $(TEST_PACKAGE)

format:
	@echo "format kode"
	go fmt ./...

help:
	@echo "command param tersedia"
	@echo " make test       - jalankan unittesting"
	@echo " make test-race  - jalankan testing dengan race detector"
	@echo " format          - format kode"
