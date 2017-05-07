build_read:
		echo "Build Products read server"
		CGO_ENABLED=0 GOOS=linux go build -o ./product-read/server ./product-read/. 

build_write:
	  echo "Build Products write server"
		CGO_ENABLED=0 GOOS=linux go build -o ./product-write/server ./product-write/. 

run: build_read build_write
	docker-compose up
