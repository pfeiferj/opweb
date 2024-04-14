VERSION 0.8
FROM golang:1.22-bookworm
WORKDIR /opweb

deps:
    COPY go.mod go.sum ./
    RUN apt update
    RUN apt install capnproto libcapnp-dev -y
    RUN go mod download
    RUN go install github.com/a-h/templ/cmd/templ@latest
    SAVE ARTIFACT go.mod AS LOCAL go.mod
    SAVE ARTIFACT go.sum AS LOCAL go.sum

templ:
    FROM +deps
    COPY ./components/*.templ components/
    RUN templ generate .
    SAVE ARTIFACT components/*.go AS LOCAL components/

build-deps:
    FROM +templ
    COPY static/* static/
    COPY templates/* templates/
    COPY proto/*.go proto/
    COPY *.go .


build:
    FROM +build-deps
    RUN go build -ldflags="-extldflags=-static -s -w" -o build/opwebd
    SAVE ARTIFACT build/opwebd /opwebd AS LOCAL build/opwebd

build-release:
    BUILD --platform=linux/arm64 +build
