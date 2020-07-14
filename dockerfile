FROM golang:1.14.1

RUN mkdir -p /app

WORKDIR /app

ADD . /app

RUN go build

CMD ["./ip-service"]