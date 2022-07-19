FROM golang:1.18-alpine

RUN apk add --update && apk add --no-cache protoc && apk add protobuf-dev && apk add git &&\
    apk add --no-cache --upgrade bash && \
    GO111MODULE=on go install google.golang.org/protobuf/cmd/protoc-gen-go@latest && \
    GO111MODULE=on go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

RUN mkdir /input && mkdir /output

ENTRYPOINT ["protoc", "-I",  "/input", "--go_out=/output", "--go-grpc_out=require_unimplemented_servers=false:/output"]
