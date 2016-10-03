get:
	go get ./...

buildserver:
	echo "Build API"
	CGO_ENABLED=0 GOOS=linux go build -o ./server/kittenserver ./server/main.go
buildkontrol:
	echo "Build Kontrol"
	go get github.com/koding/kite/kontrol/kontrol
	CGO_ENABLED=0 GOOS=linux go build -o ./kontrol/kontrol $(GOPATH)/src/github.com/koding/kite/kontrol/kontrol/main.go
runserver: buildserver buildkontrol
	echo "Run API"
	docker-compose -p kittenserver -f ./docker-compose.yml up --build
runclient:
	echo "Running client code"
	go run ./client/client.go
