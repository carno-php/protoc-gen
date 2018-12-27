FROM carno/protoc-assets as builder

FROM alpine

ARG PROTOC_VER=3.6.1

RUN apk --no-cache add libc6-compat
RUN wget -qO- https://github.com/protocolbuffers/protobuf/releases/download/v${PROTOC_VER}/protoc-${PROTOC_VER}-linux-x86_64.zip | \
    unzip -d /usr/local - && \
    chmod a+x /usr/local/bin/protoc

COPY --from=builder /build/protoc-gen /usr/local/bin

ADD gen-kit /usr/local/bin

WORKDIR /protobufs

ENTRYPOINT ["gen-kit"]
