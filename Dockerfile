FROM golang:1.24-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY .env ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/go-api .

# The run stage
FROM alpine:latest
WORKDIR /app
COPY .env .

COPY --from=builder /app/go-api .
RUN chmod +x /app/go-api
EXPOSE 8080
CMD ["/app/go-api"]

