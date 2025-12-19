# -------- BUILD STAGE --------
FROM golang:1.25.5-alpine AS builder

RUN apk add --no-cache git
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o app

# -------- RUNTIME STAGE --------
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /root

COPY --from=builder /app/app .
EXPOSE 8090
ENTRYPOINT ["./app"]
