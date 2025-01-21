# Amis Base

使用 GoFiber 和 amis 构建你的数据面板~

### 运行

首先, 你需要复制一个配置文件到 `config/config.yaml`, 并配置数据库信息

```shell
cp config/config.example.yaml config/config.yaml
```

#### 使用 Docker Compose

```shell
docker-compose up
```

#### 使用 Docker

```shell
# 构建
docker build -t amis-base .
# 运行
docker run -d \
    --name amis-base \
    -p 8080:8080 \
    -v $(pwd)/config/config.yaml:/app/config.yaml \
    -v $(pwd)/assets:/app/assets \
    amis-base
```

#### 手动运行

```shell
# install
go mod download
go mod tidy

# vendor (可选)
# go mod vendor

# 运行
go run cmd/amis-base/main.go server
```