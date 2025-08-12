FROM golang:1.23-alpine

WORKDIR /app

COPY . /app

RUN go build -o gomorse-cli

ENTRYPOINT ["./gomorse-cli"]
