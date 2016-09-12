get:
	go get ./...

buildserver:
	echo "Build API"
	CGO_ENABLED=0 GOOS=linux go build -o ./server/kittenserver ./server/main.go
runserver: buildserver
	echo "Run API"
	docker-compose -p kittenserver -f ./server/docker-compose.yml up --build
runclient:
	echo "Running client code"
	docker run --network kittenserver_default -it --rm -v $(shell pwd):/src -w /src protodocker ruby ./client/client.rb
protoc: protodocker
	echo "Generating protobuf code"
	docker run -v $(shell pwd)/proto:/proto protodocker bash -c 'protoc -I /proto --go_out=. /proto/kittens.proto'
	docker run -v $(shell pwd)/proto:/proto protodocker bash -c 'protoc -I /proto --ruby_out=/proto --grpc_out=/proto --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` /proto/kittens.proto'
protodocker:
	echo "Building protoc for Ruby"
	docker build -t protodocker .
