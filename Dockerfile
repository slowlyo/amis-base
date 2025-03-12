# 构建
FROM golang:1.23.5 AS builder

# 设置环境变量 (禁用CGO, 才能在 alpine 中运行)
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOPROXY=https://goproxy.cn,direct

# 设置工作目录
WORKDIR /app

# 复制代码到容器中
COPY . .

# 下载依赖并编译
RUN go mod download && \
    go mod tidy && \
    go mod vendor && \
    go build -v -o amis-base cmd/amis-base/main.go

# 运行
FROM alpine:latest

# 设置工作目录
WORKDIR /app

# 从构建阶段复制编译好的二进制文件
COPY --from=builder /app/amis-base .
# 复制前端
COPY --from=builder /app/web/dist ./web/dist

# 设置时区为上海, 给文件添加权限
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && \
    chmod +x ./amis-base

# 暴露端口
EXPOSE 8080

# 设置启动命令
ENTRYPOINT ["./amis-base", "server"]
