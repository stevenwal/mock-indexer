FROM golang:1.19-bullseye

ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
ENV APP_HOME /mock-indexer

RUN mkdir -p "$APP_HOME"

EXPOSE 4321

WORKDIR "$APP_HOME"
COPY ./ .

RUN go mod download
RUN go mod vendor
RUN go mod verify
RUN go build -o bin/main main/main.go
