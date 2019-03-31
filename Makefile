BUILD_CONTEXT := ./build

.PHONY: go-test
go-test:
	@echo "Run all project tests..."
	go test -v -p 1 ./...

.PHONY: bin
bin: clean
	@echo "Build project binaries..."
	GOOS=linux GOARCH=386 go build -v -o $(BUILD_CONTEXT)/respawn_linux_386
	GOOS=darwin GOARCH=386 go build -v -o $(BUILD_CONTEXT)/respawn_darwin_386
	GOOS=linux GOARCH=386 go build -v -o $(BUILD_CONTEXT)/respawn_windows_386

.PHONY: clean
clean:
	rm -rf $(BUILD_CONTEXT)/

.PHONY: go-get
go-get:
	@echo "Fetch project dependencies..."
	GO111MODULE=on go get -u ./...
	GO111MODULE=on go mod vendor

.PHONY: setup
setup: clean go-get go-build

.PHONY: test
test: go-get go-test

.PHONY: run
run:
	GO111MODULE=on go run respawn.go
