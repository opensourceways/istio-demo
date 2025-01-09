# 使用 Go 1.19 镜像
FROM golang:latest

# 设置工作目录
WORKDIR /app

# 复制 go 文件到容器中
COPY . .

# 编译 Go 服务
RUN go mod tidy
RUN go build -o go-istio-client .

# 暴露端口
EXPOSE 8081

# 启动 Go 服务
CMD ["./go-istio-client"]
