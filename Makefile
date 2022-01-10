build:
	go mod tidy
	go build cmd/http-service.go
run:
	./http-service
test:
	go mod tidy
	go test ./tests
