cucumber:
	docker-compose up -d
	cd features && godog ./
	docker-compose stop

buildserver:
	CGO_ENABLED=0 GOOS=linux go build -o ./kittenserver ./main.go

runserver: buildserver
	docker-compose -p kittenserver -f ./docker-compose.yml up --build

unit:
	go test -race -v ./...

test: unit cucumber
