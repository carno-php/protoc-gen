FROM golang:alpine

RUN apk --no-cache add git gcc musl-dev binutils

RUN mkdir /build

ADD . /build/

WORKDIR /build

RUN go get github.com/gobuffalo/packr/packr \
    && CGO_ENABLED=0 GOOS=linux packr build -a -installsuffix cgo -ldflags '-extldflags "-static"' -o protoc-gen . \
    && strip protoc-gen
