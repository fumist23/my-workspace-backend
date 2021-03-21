FROM golang:1.15 AS builder

ENV GOOS=linux
ENV GOARCH=amd64
ENV CGP_ENABLED=0
ENV GOMODULE111 on

WORKDIR /app

COPY . . 

RUN go mod download && \
    go get -u github.com/cosmtrek/air
CMD ["air"]