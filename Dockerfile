FROM golang:1.18-alpine as buider

RUN apk --no-cache add make git gcc libtool musl-dev ca-certificates dumb-init

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . /app
RUN go install github.com/onsi/ginkgo/v2/ginkgo@latest
