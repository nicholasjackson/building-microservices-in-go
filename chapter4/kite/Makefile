get:
	go get ./...

buildserver:
	echo "Build API"
	docker run --rm -v $(shell echo $$GOPATH):$(shell echo $$GOPATH) -e GOPATH=$(shell echo $$GOPATH) -v $(shell pwd)/server:/out -w $(shell pwd)/server golang:alpine /bin/sh -c 'go build -o /out/kittenserver ./main.go'
buildkontrol:
	# Cannot read kite.key: user: Current not implemented on linux/amd64
	# this is because os/user needs native code which is compiled in with cgo
	echo "Build Kontrol"
	go get github.com/koding/kite/kontrol/kontrol
	docker run --rm -v $(shell echo $$GOPATH):$(shell echo $$GOPATH) -e GOPATH=$(shell echo $$GOPATH) -v $(shell pwd)/kontrol:/out -w $(shell echo $$GOPATH)/src/github.com/koding/kite/kontrol/kontrol golang:alpine /bin/sh -c 'go build -o /out/kontrol ./main.go'
runserver: buildserver buildkontrol
	echo "Run API"
	docker-compose -p kittenserver -f ./docker-compose.yml up --build
runclient:
	echo "Running client code"
	go run ./client/client.go
