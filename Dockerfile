FROM golang:1.14.6-alpine3.12

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build ./*.go

CMD ["./main"]
