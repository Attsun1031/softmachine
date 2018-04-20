# How to generate jobapi.bp.go
```
$ protoc --version
libprotoc 3.5.1
$ protoc -I proto/ proto/jobapi.proto --go_out=plugins=grpc:.
```

# How to Build docker container
## 1. build binary
```
GOOS=linux GOARCH=amd64 go build jobapi-server.go
```

## 2. Build container
```
docker build . -t attsun/jobapi-server:latest
```
