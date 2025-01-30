# Build stage
FROM golang:1.23-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /jardin-app ./internal/main.go

# Final stage
FROM alpine:latest
WORKDIR /app
RUN apk add --no-cache tzdata
COPY --from=builder /jardin-app .
COPY resources /resources
ENV TZ=Europe/Paris
EXPOSE 8001
ENTRYPOINT ["./jardin-app"]


