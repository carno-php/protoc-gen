FROM carno/protoc-assets as builder

FROM alpine

RUN apk --no-cache add protobuf

COPY --from=builder /build/protoc-gen /usr/bin

ADD gen-kit /usr/bin

WORKDIR /protobufs

ENTRYPOINT ["gen-kit"]
