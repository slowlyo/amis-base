# Amis Base

### Run 

#### Docker
```shell
# copy config
cp config/config.example.yaml config/config.yaml

# use docker-compose
docker-compose up

# use docker
docker build -t amis-base .
docker run -d \
    --name amis-base \
    -p 8080:8080 \
    -v $(pwd)/config/config.yaml:/app/config.yaml \
    -v $(pwd)/assets:/app/assets \
    amis-base
```

#### Manual

```shell
# install
go mod tidy

# vendor (optional)
# go mod vendor

# copy config
cp config/config.example.yaml config/config.yaml

# run
go run cmd/amis-base/main.go server
```