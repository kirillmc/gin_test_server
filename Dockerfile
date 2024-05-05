FROM golang:1.21.9-alpine3.18 AS builder

COPY . /github.com/kirillmc/HTTP_TEST_SERVER_WITH_GIN/source/
WORKDIR /github.com/kirillmc/HTTP_TEST_SERVER_WITH_GIN/source/

RUN go mod download
RUN go build -o ./bin/HTTP_TEST_SERVER_WITH_GIN cmd/http_server/main.go

FROM alpine:latest

WORKDIR /root/
COPY --from=builder /github.com/kirillmc/HTTP_TEST_SERVER_WITH_GIN/source/bin/HTTP_TEST_SERVER_WITH_GIN .
COPY .env .
CMD ["./HTTP_TEST_SERVER_WITH_GIN"]



