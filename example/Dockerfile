FROM golang:alpine as builder

WORKDIR /build/
COPY . .

RUN go generate steps/*.go
RUN cat main.go
RUN go build -o app main.go

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=builder /build/app app

CMD ["./app"]