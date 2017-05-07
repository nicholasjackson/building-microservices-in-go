FROM alpine:latest

RUN mkdir /service
COPY ./server /service/server

CMD /service/server -nats nats:4222
