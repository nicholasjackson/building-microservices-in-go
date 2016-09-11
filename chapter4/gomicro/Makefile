get:
	go get ./...

buildserver:
	go build -o kittenserver ./server/main.go

protoc:
	protoc --go_out=. ./proto/kittens.proto