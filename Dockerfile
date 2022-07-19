FROM golang:1.17-alpine as builder

WORKDIR /app

COPY . .

RUN GO111MODULE=on CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go mod vendor && go build ./cmd/main.go

FROM alpine:latest as target

WORKDIR /app

COPY --from=builder /app/main .
RUN ls /app

EXPOSE 50051

CMD ["./main"]
