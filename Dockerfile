FROM golang:1.18-alpine AS builder
RUN apk update
RUN apk add git
RUN mkdir /build
ADD . /build
WORKDIR /build
RUN go mod download
RUN go build -o main .

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/gateway /app/
COPY gateway/ /app/gateway
WORKDIR /app

CMD ["sh", "./gateway"]
