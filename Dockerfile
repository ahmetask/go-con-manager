FROM golang:1.13 AS builder

ADD . /go/src/go-con-manager

WORKDIR /go/src/go-con-manager
ENV GO111MODULE=on

RUN go mod download
RUN go build -o main

FROM debian:9.9-slim

ENV LANG C.UTF-8
ENV GOPATH /go

COPY --from=builder /go/src/go-con-manager/main /app/main
WORKDIR /app

RUN chmod +x main

ENTRYPOINT ["./main"]
EXPOSE 8088