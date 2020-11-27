FROM golang:buster AS builder

WORKDIR build/
COPY . .

RUN go generate delegates/*.go
RUN cat main.go
RUN go build main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder build/main app

CMD ["./app"]