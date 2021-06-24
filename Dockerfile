ARG GO_VERSION=1.16

FROM golang:${GO_VERSION}-alpine AS builder

RUN apk update && apk add alpine-sdk git && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download
RUN go install github.com/swaggo/swag/cmd/swag@latest

COPY . .
RUN swag init
RUN go build -o ./api ./main.go

FROM alpine:latest

RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir -p /app
WORKDIR /app
COPY --from=builder /app/api .

EXPOSE 5000

ENTRYPOINT ["./api"]