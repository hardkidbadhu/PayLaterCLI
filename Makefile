install:
		go install github.com/PayLaterCLI

test:
		go test ./...

deamond:
		go run ./deamon/main.go 4444

format_check:
	go fmt ./...
	go vet ./...