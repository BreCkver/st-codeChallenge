FROM golang:1.17-buster

MAINTAINER Jaime Reyes Verea <breckver.dll@gmail.com>

WORKDIR /app
COPY . .

RUN go get -d -v ./...
RUN go build -v ./...

RUN go build -o /transactions

EXPOSE 8089

CMD ["/transactions"]