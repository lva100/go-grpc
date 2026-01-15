FROM golang:1.24-alpine AS builder

COPY . /github.com/lva100/go-grpc
WORKDIR /github.com/lva100/go-grpc

RUN go mod download

RUN go build -o ./bin/crud_server cmd/grpc_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/lva100/go-grpc/bin/crud_server .

CMD ["./crud_server"]