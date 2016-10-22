get:
	go get ./...

buildserver:
	echo "Build API"
	CGO_ENABLED=0 GOOS=linux go build -o ./server/kittenserver ./server/main.go
runserver: buildserver
	echo "Run API"
	docker-compose -p kittenserver -f ./docker-compose.yml up --build
runclient:
	echo "Running client code"
	go run ./client/client.go
runrubyclient:
	echo "Running ruby code"
	docker run -v $(shell pwd):/src -w /src --network=kittenserver_default -link=kittenserver_kittenserver_1 nicholasjackson/building-microservices-in-go /bin/bash -c 'ruby ./client/client.rb'
protoc:
	echo "Generating protobuf code"
	docker run -v $(shell pwd)/proto:/proto nicholasjackson/building-microservices-in-go bash -c 'protoc -I /proto /proto/kittens.proto --go_out=plugins=grpc:/proto'
	docker run -v $(shell pwd)/proto:/proto nicholasjackson/building-microservices-in-go bash -c 'protoc -I /proto --ruby_out=/proto --grpc_out=/proto --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` /proto/kittens.proto'
