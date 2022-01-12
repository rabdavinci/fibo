build:
	go mod tidy
	go build cmd/*.go
run:
	./http-service
test:
	go mod tidy
	go test ./tests
gen-proto:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	go mod tidy
clean-proto:
	rm gen/proto/*.go