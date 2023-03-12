help:
	@echo "Available commands:"
	@echo "* build   - Build the Pandaroll binary"
	@echo "* test    - Run tests"
	@echo "* test-db - Spin up the test DB containers"

build:
	$(eval VERSION=$(shell git describe --exact-match --tags $(git log -n1 --pretty='%h') || echo 'development'))
	$(eval COMMIT_HASH=$(shell git rev-parse HEAD))

	go build \
		-o bin/pandaroll \
		-ldflags="-X 'blobdev.com/pandaroll/internal/build.Version=$(VERSION)'" \
		-ldflags="-X 'blobdev.com/pandaroll/internal/build.Commit=$(COMMIT_HASH)'"

test:
	go test ./... -v

test-db:
	docker-compose up