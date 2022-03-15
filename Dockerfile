FROM golang:1.16-alpine AS build
WORKDIR /app/
COPY . .
RUN go install
RUN go test ./pool -cover .
