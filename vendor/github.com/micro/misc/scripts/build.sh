#!/bin/bash

set -e
set -x

REGISTRY=microhq

# Used to rebuild all the things

find * -type d -maxdepth 0 -print | while read dir; do
	if [ ! -f $dir/Dockerfile ]; then
		continue
	fi

	if [ -f $dir/.skip ]; then
		continue
	fi

	pushd $dir >/dev/null

	# test
	go test -v ./...

	# build static binary
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-w' -o $dir ./main.go

	# build docker image
	docker build -t $REGISTRY/$dir .

	# push docker image
	docker push $REGISTRY/$dir

	# remove binary
	rm $dir

	popd >/dev/null
done
