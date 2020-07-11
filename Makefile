.PHONY: test
test:
	@go test -parallel 4 -v ./...

.PHONY: ci-test
cover:
	@go test -parallel 4 -race -coverprofile=coverage.txt -covermode=atomic -v ./...
	@go tool cover -func coverage.txt
	@rm coverage.txt

.PHONY: pack-contract-grpc
pack-contract-grpc:
	@rm -rf ./pb/sagagrpc && mkdir ./pb/sagagrpc
	@protoc -I pb --gofast_out=plugins=grpc:pb/sagagrpc pb/*.proto
	@go generate ./...
