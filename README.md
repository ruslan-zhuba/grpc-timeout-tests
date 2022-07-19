# GRPC TIMEOUT TESTING

This project created by research related to timeout.

## How to generate protobufs

```shell
$ docker build -t protoc -f protoc.Dockerfile .
$ docker run --rm -v "$(pwd)"/protobuf:/input \
  -v "$(pwd)"/pb:/output protoc \
  /input/example.proto
```

## How to start locally

```shell
$ docker build -t example-grpc .
$ docker run -p 50051:50051 --rm --name grpc example-grpc
```
or
```shell
$ docker run -p 50051:50051 --rm \
  --name grpc afrofunkylover/grpc-timeout-tests
```
