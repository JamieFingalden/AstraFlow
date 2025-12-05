# ---------- build ----------
FROM golang:1.24-alpine AS builder
ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
ENV GOPROXY=https://goproxy.cn,direct GO111MODULE=on
WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
# 确认源码目录存在
RUN ls -l cmd/

# 分别编译
RUN go build -ldflags="-s -w" -o migrate ./cmd/migrate \
 && go build -ldflags="-s -w" -o server  ./cmd/server

# ---------- run ----------
FROM alpine:latest
RUN apk add --no-cache ca-certificates tzdata
WORKDIR /app
COPY --from=builder /build/migrate /build/server ./
COPY --from=builder /build/pkg/config /app/pkg/config
# 复制配置文件
COPY --from=builder /build/pkg/config /app/pkg/config
ENTRYPOINT ["/app/server"]
