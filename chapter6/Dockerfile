FROM ruby:latest
MAINTAINER jackson.nic@gmail.com

RUN apt-get update && \
    apt-get -y install git unzip build-essential autoconf libtool wget

RUN mkdir /go

# Install Go
ENV GOLANG_VERSION 1.7.1
ENV GOLANG_DOWNLOAD_URL https://golang.org/dl/go$GOLANG_VERSION.linux-amd64.tar.gz
ENV GOLANG_DOWNLOAD_SHA256 43ad621c9b014cde8db17393dc108378d37bc853aa351a6c74bf6432c1bbd182

RUN curl -fsSL "$GOLANG_DOWNLOAD_URL" -o golang.tar.gz \
    && echo "$GOLANG_DOWNLOAD_SHA256  golang.tar.gz" | sha256sum -c - \
    && tar -C /usr/local -xzf golang.tar.gz \
    && rm golang.tar.gz

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
RUN go get google.golang.org/grpc && \
    go get github.com/golang/protobuf/protoc-gen-go

# Install protobuffers
RUN git clone -b v1.0.x https://github.com/grpc/grpc && \
    cd grpc && \
    git submodule update --init && \
    make && \
    make install && \
    ldconfig && \
    cp /grpc/third_party/protobuf/src/protoc /usr/local/bin/protoc

RUN gem install grpc -v 1.0.0 && \
    gem install json-rpc-objects && \
    gem install json-rpc-objects-json && \
    gem install protobuf
