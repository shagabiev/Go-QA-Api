FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN CGO_ENABLED=0 GOOS=linux go build -o /qa-api ./cmd/server/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /qa-api .

COPY --from=builder /go/bin/goose /usr/local/bin/goose

COPY migrations ./migrations

EXPOSE 8080
CMD ["./go-qa-api"]