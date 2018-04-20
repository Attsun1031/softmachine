# How to Build docker container
## 1. build binary
```
GOOS=linux GOARCH=amd64 go build dbmigration.go
```

## 2. Build container
```
docker build . -t attsun/jobnetes-migration:latest
```
