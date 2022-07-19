# GRPC TIMEOUT TESTING

This project created by research related to timeout.

## How to generate protobufs

```shell
$ docker build -t protoc .
$ docker run --rm -v "$(pwd)"/protobuf:/input \
  -v "$(pwd)"/pb:/output protoc \
  /input/example.proto
```


