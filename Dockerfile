# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
# Agregamos bash
RUN apk update && apk add bash
RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' ntp.go

# Ejecutar: /app/ntp pool.ntp.rg
#CMD [ "/app/ntp pool.ntp.org" ]
