FROM golang:1.18-alpine AS build

WORKDIR /app
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY ./cmd ./cmd
COPY ./entities ./entities
COPY ./routes ./routes
COPY ./service ./service
COPY ./utils ./utils

RUN go build -o ./binary/keygen-service ./cmd/api/main.go

FROM alpine:3.16
WORKDIR /app
COPY --from=build /app/binary .
CMD ["./keygen-service"]