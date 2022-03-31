FROM golang:1.14.9-alpine
RUN apk update && apk add git && go get gopkg.in/natefinch/lumberjack.v2
EXPOSE 8080

RUN mkdir /build
ADD go.mod go.sum main.go /build/
WORKDIR /build
RUN go build -o build
