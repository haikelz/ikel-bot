FROM golang:1.25.1-alpine AS builder

RUN apk add --no-cache \
    tesseract-ocr \
    tesseract-ocr-dev \
    gcc \
    musl-dev \
    build-base

WORKDIR /app

COPY go.mod go.sum ./

COPY . ./

RUN go mod download && go mod verify && go mod tidy

RUN go build -o main cmd/app/main.go 
