# ---------- build ----------
FROM golang:1.22-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
WORKDIR /build
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# 同时编译两个命令
RUN go build -ldflags="-s -w" -o migrate ./cmd/migrate && \
    go build -ldflags="-s -w" -o server  ./cmd/server

# ---------- run ----------
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /build/migrate .
COPY --from=builder /build/server  .
# 默认入口留空，由 compose 覆盖
ENTRYPOINT ["/app/server"]