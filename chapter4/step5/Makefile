cucumber:
	docker-compose up -d
	cd features && godog ./
	docker-compose stop

run:
	docker-compose up -d
	go run main.go
	docker-compose stop

unit:
	go test -race -v ./...

test: unit cucumber
