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
	docker run --network kittenserver_default -it --rm -v $(shell pwd):/src -w /src nicholasjackson/building-microservices-in-go bash -c  'ruby ./client/client.rb'
protoc:
	echo "Generating protobuf code"
	docker run -v $(shell pwd)/proto:/proto nicholasjackson/building-microservices-in-go bash -c 'protoc -I /proto --go_out=. /proto/kittens.proto'
	docker run -v $(shell pwd)/proto:/proto nicholasjackson/building-microservices-in-go bash -c 'protoc -I /proto --ruby_out=/proto /proto/kittens.proto'
