CURRENT_DIR=$(shell pwd)

APP=$(shell basename ${CURRENT_DIR})

APP_CMD_DIR=${CURRENT_DIR}/cmd

build:
	CGO_ENABLED=0 GOOS=linux go mod vendor && go build -mod=vendor -a -installsuffix cgo -o ${CURRENT_DIR}/bin/${APP} ${APP_CMD_DIR}/main.go
run:
	./http-service
test:
	go mod vendor
	go test ./tests/*
gen-proto:
	protoc --proto_path=proto proto/*.proto --go_out=gen/
	protoc --proto_path=proto proto/*.proto --go-grpc_out=gen/
	go mod vendor
clean-proto:
	rm gen/proto/*.go