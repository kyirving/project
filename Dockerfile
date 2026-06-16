# Stage 1: Build
FROM golang:1.26-alpine AS builder

ENV GOPROXY=https://goproxy.cn,direct

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 go build -ldflags="-s -w" -o /src/bin/app main.go

# Stage 2: Runtime
FROM alpine:3.21

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app
COPY --from=builder /src/bin/app .

EXPOSE 8080

ENTRYPOINT ["./app"]
CMD ["--config_dir", "/app/config", "--config_file", "config.yaml"]

# 构建镜像
 # docker build -t app:latest .

 # 推送镜像
 # docker push app:latest