FROM golang:alpine AS builder

ENV GO111MODULE=on \
    GOOS=linux \
    GOARCH=amd64 \
    CGO_ENABLED=1

ADD config.yaml /go/bin/

ADD . /go/src/github.com/szemskov/herald
WORKDIR /go/src/github.com/szemskov/herald

RUN go mod init
RUN go get ./...
RUN go install -ldflags="-w -s" ./...

FROM alpine:latest

ADD config.yaml /usr/local/bin/

COPY --from=builder /go/bin/herald-server /usr/local/bin/herald-server

WORKDIR /usr/local/bin/

ENTRYPOINT ["herald-server", "--host", "0.0.0.0", "--port", "8080"]

EXPOSE 8080