# How to Build docker container
## 1. build binary
```
GOOS=linux GOARCH=amd64 go build jobnetes-manager.go
```

## 2. Build container
```
docker build . -t attsun/jobnetes-manager:latest
```
