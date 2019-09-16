build:
	go build -ldflags '-s -w' -o public/main.cgi cmd/cgi/main.go

fmt:
	go fmt ./...
