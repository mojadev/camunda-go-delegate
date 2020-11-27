FROM golang:buster AS builder

WORKDIR build/
COPY . .
RUN go generate ./...
RUN go build main.go - app

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder build/app app

CMD ["./app"]