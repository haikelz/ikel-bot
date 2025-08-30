FROM golang:1.23.4-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

RUN go mod download && go mod verify && go mod tidy

RUN go build -o main cmd/app/main.go 
