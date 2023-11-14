FROM golang:1.21.3-alpine AS builder 

WORKDIR /app

COPY . .

# Dowload deps
RUN go mod download 

# Build the application
RUN go build -o /app/backend /app/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/backend /app/backend  

ENTRYPOINT ["./backend"]
