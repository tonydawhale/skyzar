# syntax=docker/dockerfile:1

FROM golang:1.20-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 23570

RUN CGO_ENABLED=0 GOOD=linux go build -o /backend

ENTRYPOINT [ "/backend" ]