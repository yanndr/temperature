all:
	go build
	
test: 
	go test -race ./...

run-example:
	go run ./example/main.go