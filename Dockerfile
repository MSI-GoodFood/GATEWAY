FROM golang:alpine as builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

LABEL maintainer="GALLIER Thomas"

RUN apk update && apk add --no-cache git

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go mod tidy

RUN go build -a -installsuffix cgo -o main .

WORKDIR /dist

RUN cp /build/main .

FROM alpine:latest

COPY --from=builder /dist/main /

EXPOSE 8080 8080

CMD ["/bin/sh"]

ENTRYPOINT ["/main"]
