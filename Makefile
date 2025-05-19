run: build
	@./cmd/bin

build: | cmd
	@go build -o cmd/bin main.go 

cmd: 
	@mkdir -p cmd

test:
	go test ./...
